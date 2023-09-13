package config

import (
	"os"
)

var (
	PrivateKey   = os.Getenv("LOCAL_PRIVATE_KEY")
	LocalRPCWS   = "ws://" + os.Getenv("LOCAL_RPC")
	LocalRPCHTTP = "http://" + os.Getenv("LOCAL_RPC")
)
