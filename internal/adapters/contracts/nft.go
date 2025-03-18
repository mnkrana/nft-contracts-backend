package contracts

import (
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mnkrana/nft-contracts-backend/bindings/contracts/rana"
	"github.com/mnkrana/nft-contracts-backend/internal/ports"
	"github.com/mnkrana/nft-contracts-backend/internal/utils"
)

type RanaAdapter struct {
	RPC      ports.RpcPort
	Instance *rana.Rana
}

func InitNFT(rpcPort ports.RpcPort) *RanaAdapter {
	address, _ := utils.GetAddressFromRaw(os.Getenv(utils.NFTKey))
	instance, err := rana.NewRana(address, rpcPort.GetEthClient())
	if err != nil {
		log.Fatalln("error in creating rana contract instance", err)
	}

	contract := &RanaAdapter{
		RPC:      rpcPort,
		Instance: instance,
	}
	return contract
}

func (c *RanaAdapter) GetPlayerBalance(address common.Address) (*big.Int, error) {
	log.Println("Balance command is not implemented")
	return big.NewInt(0), errors.New("balance command is not implemented")
}

func (c *RanaAdapter) GetPlayerBalanceWithCardID(address common.Address, cardId *big.Int) (*big.Int, error) {
	balance, err := c.Instance.BalanceOf(&bind.CallOpts{From: address}, address, cardId)
	if err != nil {
		return big.NewInt(0), err
	}
	return balance, nil
}

func (c *RanaAdapter) Mint(address common.Address, id *big.Int, amount *big.Int, uri string) (string, error) {
	authValue := big.NewInt(0)
	auth, err := c.RPC.GetAuth(authValue)
	if err != nil {
		log.Println("error in auth, can't mint", err)
		return "", err
	}
	tx, err := c.Instance.Mint(auth, address, id, amount, uri)
	if err != nil {
		log.Println("error in minting", err)
		return "", err
	}
	log.Println("mint tx", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func (c *RanaAdapter) MintBatch(address common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (string, error) {
	authValue := big.NewInt(0)
	auth, err := c.RPC.GetAuth(authValue)
	if err != nil {
		log.Println("error in auth, can't mint", err)
		return "", err
	}
	tx, err := c.Instance.MintBatch(auth, address, ids, amounts, uris)
	if err != nil {
		log.Println("error in minting", err)
		return "", err
	}
	log.Println("mint tx", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

func (c *RanaAdapter) GetUris(address common.Address, balance int64) (interface{}, error) {
	return nil, nil
}

func (c *RanaAdapter) SafeTransferNFT(from common.Address, to common.Address, id *big.Int, amount *big.Int) (string, error) {
	authValue := big.NewInt(0)
	auth, err := c.RPC.GetAuth(authValue)
	if err != nil {
		log.Println("error in auth, can't transfer NFT", err)
		return "", err
	}

	tx, err := c.Instance.SafeTransferFrom(auth, from, to, id, amount, []byte{})
	if err != nil {
		log.Println("error in transferring NFT", err)
		return "", err
	}

	log.Println("transfer tx", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}
