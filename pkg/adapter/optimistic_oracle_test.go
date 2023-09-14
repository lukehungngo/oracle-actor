package adapter

import (
	"encoding/json"
	"math/big"
	"oracle-actor/pkg/client"
	"testing"
)

func TestWatchOptimisticOracle(t *testing.T) {
	i, _ := client.InstanceOptimisticOracle()
	for counter := 0; counter < 100; counter++ {
		txOpts, err := client.CreateTransactor(520000)
		if err != nil {
			t.Fatal(err)
		}
		txs, err := i.Report(txOpts, big.NewInt(int64(counter)), 1, []byte("triggerType"), []byte("proof"))
		if err != nil {
			t.Fatal(err)
		}
		data, _ := json.MarshalIndent(txs, "", " ")
		t.Log(string(data))
	}
}
