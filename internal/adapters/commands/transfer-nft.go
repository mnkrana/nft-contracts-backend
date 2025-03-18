package commands

import (
	"log"
	"math/big"

	"github.com/mnkrana/nft-contracts-backend/internal/ports"
	"github.com/mnkrana/nft-contracts-backend/internal/utils"
)

type TransferRequest struct {
	Address string `json:"address,omitempty"`
	Id      int    `json:"id,omitempty"`
	Amount  int    `json:"amount,omitempty"`
}

type TransferCommand struct {
	Rpc      ports.RpcPort
	Contract ports.NFTPort
}

func NewTransferCommand(rpc ports.RpcPort, port ports.NFTPort) *TransferCommand {
	return &TransferCommand{Rpc: rpc, Contract: port}
}

func (m *TransferCommand) ExecuteRequest(action string, request any) (string, error) {
	req, ok := request.(*TransferRequest)
	log.Println("execute req", req)
	if !ok {
		return utils.HandleError("invalid request format", nil)
	}

	if req.Address == "" {
		return utils.HandleError("address is required", nil)
	}

	if req.Id == 0 {
		return utils.HandleError("id is required", nil)
	}

	if req.Amount == 0 {
		return utils.HandleError("amount is required", nil)
	}

	playerAddress, err := utils.GetAddressFromRaw(req.Address)
	if err != nil {
		return "", err
	}

	log.Println("Token id", req.Id)
	tx, err := m.Contract.SafeTransferNFT(
		m.Rpc.GetUserAddress(),
		playerAddress,
		big.NewInt(int64(req.Id)),
		big.NewInt(int64(req.Amount)))

	if err != nil {
		log.Println("error in transfer  ", err)
		return "", err
	}

	txHash := Transaction{
		Hash: tx,
	}

	return txHash.getTxDetails(), nil
}
