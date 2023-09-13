package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lmittmann/w3"
)

func EthClient(providerURL string) *ethclient.Client {
	client, err := ethclient.Dial(providerURL)
	if err != nil {
		log.Panic("Failed client", err)
	}
	return client
}

func W3Client(providerURL string) *w3.Client {
	client := w3.MustDial(providerURL)
	return client
}
