package client

import (
	"log"
	"math/big"
	"oracle-actor/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/module/eth"
)

func CreateTransactor(gasLimit int) (*bind.TransactOpts, error) {
	var (
		w, _        = wallet()
		chainID     uint64
		nonce       uint64
		gasPrice    big.Int
		blockHeader types.Header
		cw3         = W3Client(config.LocalRPCWS)
	)
	if err := cw3.Call(
		eth.ChainID().Returns(&chainID),
		eth.Nonce(w.FromAddress, nil).Returns(&nonce),
		eth.GasPrice().Returns(&gasPrice),
		eth.HeaderByNumber(nil).Returns(&blockHeader),
	); err != nil {
		log.Println("get blockchain Infos failed: ", err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(w.PrivateKey, new(big.Int).SetUint64(chainID))
	if err != nil {
		log.Println("error with the transactor with chain ID: ", err)
		return nil, err
	}
	auth.From = w.FromAddress
	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.GasPrice = &gasPrice
	auth.GasLimit = uint64(float64(gasLimit) * 1.25)

	return auth, nil
}
