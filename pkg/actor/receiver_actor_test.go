package actor

import (
	"encoding/json"
	"fmt"
	"oracle-actor/config"
	"oracle-actor/pkg/client"
	"oracle-actor/pkg/generated/contracts/optimistic_oracle"
	"testing"
)

func TestOracleEventReceivedActor_Receive(t *testing.T) {
	clientProvider := client.EthClient(config.LocalRPCWS)
	txOpts, err := client.CreateTransactor(1200000)
	if err != nil {
		t.Fatal(err)
	}
	contractAddress, txs, oracle, err := optimistic_oracle.DeployOptimisticOracle(txOpts, clientProvider)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("contractAddress:", contractAddress)
	dataTxs, _ := json.MarshalIndent(txs, "", " ")
	fmt.Println("txs:", string(dataTxs))

	dataOracle, _ := json.MarshalIndent(*oracle, "", " ")
	fmt.Println("dataOracle:", string(dataOracle))

}
