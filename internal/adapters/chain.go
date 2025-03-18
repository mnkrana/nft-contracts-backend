package adapters

import (
	"os"

	"github.com/mnkrana/nft-contracts-backend/internal/domain"
	"github.com/mnkrana/nft-contracts-backend/internal/utils"
)

func NewChain() *domain.Chain {
	network := os.Getenv(utils.NetworkKey)
	chain := &domain.Chain{
		Network: os.Getenv(network + utils.ChainNetworkKey),
		RpcUrl:  os.Getenv(network + utils.ChainRPCUrlKey),
		PKey:    os.Getenv(network + utils.ChainPKey),
	}
	return chain
}
