package commands

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/mnkrana/nft-contracts-backend/internal/ports"
	"github.com/mnkrana/nft-contracts-backend/internal/utils"
)

type MintBatchRequest struct {
	Address string   `json:"address,omitempty"`
	Ids     []int    `json:"ids,omitempty"`
	Amounts []int    `json:"amounts,omitempty"`
	Uris    []string `json:"uris,omitempty"`
}

type MintBatchCommand struct {
	Contract ports.NFTPort
}

func NewMintBatchCommand(port ports.NFTPort) *MintBatchCommand {
	return &MintBatchCommand{Contract: port}
}

func (m *MintBatchCommand) ExecuteRequest(action string, request any) (string, error) {
	req, ok := request.(*MintBatchRequest)
	log.Println("execute req", req)
	if !ok {
		return utils.HandleError("invalid request format", nil)
	}

	if req.Address == "" {
		return utils.HandleError("address is required", nil)
	}

	if len(req.Ids) == 0 {
		return utils.HandleError("ids are required", nil)
	}

	ids := []*big.Int{}
	for _, id := range req.Ids {
		ids = append(ids, big.NewInt((int64(id))))
	}

	if len(req.Amounts) == 0 {
		return utils.HandleError("amounts are required", nil)
	}

	amounts := []*big.Int{}
	for _, amount := range req.Amounts {
		amounts = append(amounts, big.NewInt((int64(amount))))
	}

	if len(req.Uris) == 0 {
		return utils.HandleError("uris are required", nil)
	}

	for _, value := range req.Uris {
		log.Println("Uris: ", value)
	}

	playerAddress, err := utils.GetAddressFromRaw(req.Address)
	if err != nil {
		return "", err
	}

	log.Println("Token ids", ids)
	tx, err := m.Contract.MintBatch(playerAddress, ids, amounts, req.Uris)
	if err != nil {
		log.Println("error in minting from contract ", err)
		return "", err
	}

	txHash := Transaction{
		Hash: tx,
	}

	return txHash.getTxDetails(), nil
}

type Transaction struct {
	Hash string `json:"txHash"`
}

func (tx *Transaction) getTxDetails() string {
	e, err := json.Marshal(tx)
	if err != nil {
		return err.Error()
	}
	return string(e)
}
