package actor

import (
	"encoding/json"
	"fmt"
	"oracle-actor/config"
	"oracle-actor/pkg/client"
	"oracle-actor/pkg/generated/contracts/optimistic_oracle"
	"sync"
	"testing"
	"time"
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

func TestEventStream(t *testing.T) {
	e := NewEngine()
	wg := new(sync.WaitGroup)
	eventSub := e.EventStream.Subscribe(func(event any) {
		fmt.Println("received event", event)
	})

	eventSub2 := e.EventStream.Subscribe(func(event any) {
		fmt.Println("received event2", event)
	})
	go func() {
		wg.Add(1)
		e.EventStream.Publish("this is a message")
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		e.EventStream.Publish("this is a second message")
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	wg.Wait()

	// Cleanup subscription
	defer e.EventStream.Unsubscribe(eventSub)
	defer e.EventStream.Unsubscribe(eventSub2)
}
