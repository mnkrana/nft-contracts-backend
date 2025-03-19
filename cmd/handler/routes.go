package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mnkrana/nft-contracts-backend/internal/adapters/commands"
)

// RegisterRoutes registers all routes to the Echo instance
func RegisterRoutes(e *echo.Echo, router *Router) {
	// POST routes with structured payloads
	e.POST("/mintbatch", handleStructuredRequest(router, "mintbatch", &commands.MintBatchRequest{}))
	e.POST("/transfer", handleStructuredRequest(router, "transfer", &commands.TransferRequest{}))
	e.POST("/tx", handleStructuredRequest(router, "tx", &commands.TransactionRequest{}))
}

// handleStructuredRequest handles routes with structured payloads
func handleStructuredRequest(router *Router, action string, requestPrototype interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read and log the raw request body
		bodyBytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}
		log.Printf("Raw Request Body: %s", string(bodyBytes))

		// Create a new instance of the request struct
		reqType := reflect.TypeOf(requestPrototype).Elem() // Get the element type
		reqValue := reflect.New(reqType).Interface()       // Create a new pointer to the struct

		// Reassign the request body so it can be read by c.Bind
		c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if err := c.Bind(reqValue); err != nil {
			log.Printf("Bind Error: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}
		log.Printf("Parsed Request (%s): %+v", action, reqValue)

		// Dispatch the request
		response, err := router.HandleRequest(action, reqValue)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Ensure response is properly formatted JSON
		var formattedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(response), &formattedResponse); err != nil {
			// If unmarshalling fails, return as a plain string
			return c.JSON(http.StatusOK, response)
		}

		return c.JSON(http.StatusOK, formattedResponse)
	}
}
