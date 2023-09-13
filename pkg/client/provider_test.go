package client

import (
	"oracle-actor/config"
	"testing"
)

func TestEthClient(t *testing.T) {
	client := EthClient(config.LocalRPCHTTP)

	if client == nil {
		t.Errorf("Failed to create Ethereum client.")
	}
}
