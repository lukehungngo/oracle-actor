package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"oracle-actor/config"
	"oracle-actor/pkg/generated/contracts/optimistic_oracle"
)

func InstanceOptimisticOracle() (oracle *optimistic_oracle.OptimisticOracle, c *ethclient.Client) {
	clientProvider := EthClient(config.LocalRPCWS)
	oracleAddress := common.HexToAddress(config.OracleAddress)
	optimisticOracleInstance, err := optimistic_oracle.NewOptimisticOracle(oracleAddress, clientProvider)
	if err != nil {
		log.Println(err)
		defer clientProvider.Close()
	}
	return optimisticOracleInstance, clientProvider
}
