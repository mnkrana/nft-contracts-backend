package handler

import (
	"log"

	"github.com/mnkrana/nft-contracts-backend/internal/adapters/commands"
	"github.com/mnkrana/nft-contracts-backend/internal/ports"
	"github.com/mnkrana/nft-contracts-backend/internal/utils"
)

type Router struct {
	RpcPort  ports.RpcPort
	Commands map[string]ports.Command
}

func NewRouter(
	rpc ports.RpcPort,
	ranaAdapter ports.NFTPort,
) *Router {
	return &Router{
		RpcPort:  rpc,
		Commands: initializeCommands(rpc, ranaAdapter),
	}
}

func initializeCommands(
	rpc ports.RpcPort,
	ranaAdapter ports.NFTPort,
) map[string]ports.Command {

	commands := map[string]ports.Command{
		"mintbatch": commands.NewMintBatchCommand(ranaAdapter),
		"transfer":  commands.NewTransferCommand(rpc, ranaAdapter),
	}
	return commands
}

func (r *Router) HandleRequest(action string, payload any) (string, error) {
	cmd, exists := r.Commands[action]
	if !exists {
		return utils.HandleError("Unknown command: "+action, nil)
	}

	log.Printf("Executing command: %s\n", action)
	return cmd.ExecuteRequest(action, payload)
}
