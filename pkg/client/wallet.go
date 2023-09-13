// Copyright 2023 ConcaveFi
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"crypto/ecdsa"
	"oracle-actor/config"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey  *ecdsa.PrivateKey
	PublicKey   *ecdsa.PublicKey
	FromAddress common.Address
}

func wallet() (*Wallet, error) {
	PK := config.PrivateKey
	if strings.HasPrefix(PK, "0x") {
		PK, _ = strings.CutPrefix(PK, "0x")
	}
	privateKey, err := crypto.HexToECDSA(PK)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	w := &Wallet{
		PrivateKey:  privateKey,
		PublicKey:   publicKeyECDSA,
		FromAddress: fromAddress,
	}

	return w, nil
}
