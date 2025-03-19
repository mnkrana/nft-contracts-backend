package commands

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mnkrana/nft-contracts-backend/internal/ports"
	"github.com/mnkrana/nft-contracts-backend/internal/utils"
)

type TransactionRequest struct {
	Hash string `json:"tx,omitempty"`
}

type TransactionCommand struct {
	Rpc      ports.RpcPort
	Contract ports.NFTPort
}

func NewTransactionCommand(rpc ports.RpcPort, port ports.NFTPort) *TransactionCommand {
	return &TransactionCommand{Rpc: rpc, Contract: port}
}

func (m *TransactionCommand) ExecuteRequest(action string, request any) (string, error) {
	req, ok := request.(*TransactionRequest)
	log.Println("execute req", req)
	if !ok {
		return utils.HandleError("invalid request format", nil)
	}
	return m.CheckTransaction(req.Hash)
}

func (t *TransactionCommand) CheckTransaction(hash string) (string, error) {
	hash = strings.Trim(hash, "\"")
	log.Printf("hash: %s", hash)

	for {
		time.Sleep(1 * time.Second)
		success, err := t.Transfer(hash)
		if err != nil {
			log.Printf("error occurred during  %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if success {
			return "success", nil
		} else {
			log.Printf("failed tx")
			return "failed", nil
		}
	}
}

func (t *TransactionCommand) Transfer(hash string) (bool, error) {
	receipt, err := t.Rpc.GetEthClient().TransactionReceipt(
		context.Background(),
		common.HexToHash(hash))

	if err != nil {
		return false, err
	}

	if receipt.Status == 0 {
		return false, err
	}

	logEntry := receipt.Logs[1]
	user := common.BytesToAddress(logEntry.Topics[1].Bytes())
	tokenId := new(big.Int).SetBytes(logEntry.Data[:32]).Uint64()
	log.Println("User:", user.Hex(), "TokenId:", tokenId)

	bal, err := t.Contract.GetBalanceWithTokenID(t.Rpc.GetUserAddress(), big.NewInt(int64(tokenId)))
	if err != nil {
		log.Println("error in balance ", err)
		return false, err
	}

	//in case of deployer doesn't have requested NFT
	//this will never happen, just adding it for fail safe
	if bal.Int64() < 1 {
		log.Println("balance is low ", bal)
		return false, err
	}

	tx, err := t.Contract.SafeTransferNFT(
		t.Rpc.GetUserAddress(),
		user,
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(1)))

	if err != nil {
		log.Println("error in transfer  ", err)
		return false, err
	}

	log.Println("tx transfer", tx)
	return true, nil
}
