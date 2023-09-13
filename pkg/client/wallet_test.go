package client

import (
	"crypto/ecdsa"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestWallet(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	w := &Wallet{
		PrivateKey:  privateKey,
		PublicKey:   publicKeyECDSA,
		FromAddress: fromAddress,
	}

	// Assert that the private key is not nil
	if w.PrivateKey == nil {
		t.Fatal("Expected private key to be set, got nil")
	}

	// Assert that the public key is not nil
	if w.PublicKey == nil {
		t.Fatal("Expected public key to be set, got nil")
	}

	// Assert that the from address is not the zero address
	if w.FromAddress == (common.Address{}) {
		t.Fatal("Expected from address to be set, got zero address")
	}

	// Assert that converting the public key to an address and back results in the original public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	msg := crypto.Keccak256([]byte("foo"))
	sig, err := crypto.Sign(msg, privateKey)
	if err != nil {
		t.Errorf("Sign error: %s", err)
	}
	recoveredPubKey, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPubKey)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if address != recoveredAddr {
		t.Errorf("Address mismatch: want: %x have: %x", address, recoveredAddr)
	}

	// should be equal to SigToPub
	recoveredPub2, err := crypto.SigToPub(msg, sig)
	if err != nil {
		t.Errorf("ECRecover error: %s", err)
	}
	recoveredAddr2 := crypto.PubkeyToAddress(*recoveredPub2)
	if address != recoveredAddr2 {
		t.Errorf("Address mismatch: want: %x have: %x", address, recoveredAddr2)
	}
}
