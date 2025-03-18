package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mnkrana/nft-contracts-backend/cmd/handler"
	"github.com/mnkrana/nft-contracts-backend/internal/adapters"
	"github.com/mnkrana/nft-contracts-backend/internal/adapters/contracts"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: .env file not loaded: %v", err)
	}
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	return e
}

func initDependencies() (*adapters.RPC, *contracts.RanaAdapter) {
	chainConfig := adapters.NewChain()
	rpcAdapter := adapters.NewRPCAdapter(chainConfig)
	rpcAdapter.Print()

	ranaAdapter := contracts.InitNFT(rpcAdapter)
	return rpcAdapter, ranaAdapter
}

func main() {
	loadEnv()

	e := setupEcho()

	rpcAdapter, ranaAdapter := initDependencies()
	router := handler.NewRouter(rpcAdapter, ranaAdapter)
	handler.RegisterRoutes(e, router)
	echoLambda := echoadapter.New(e)

	env := os.Getenv("APP_ENV")
	if env == "local" {
		log.Println("Starting server on :8080...")
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		log.Println("Starting Lambda handler...")
		lambda.Start(echoLambda.ProxyWithContext)
	}
}
