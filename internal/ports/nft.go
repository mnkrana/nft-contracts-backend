package ports

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type NFTPort interface {
	GetBalance(address common.Address) (*big.Int, error)
	GetBalanceWithTokenID(address common.Address, tokenId *big.Int) (*big.Int, error)
	Mint(address common.Address, id *big.Int, amount *big.Int, uri string) (string, error)
	MintBatch(address common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (string, error)
	GetUris(address common.Address, balance int64) (any, error)
	SafeTransferNFT(from common.Address, to common.Address, id *big.Int, amount *big.Int) (string, error)
}
