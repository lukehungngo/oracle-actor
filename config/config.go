package config

import (
	"os"
)

var (
	LocalRPCWS    = "ws://" + os.Getenv("LOCAL_RPC")
	LocalRPCHTTP  = "http://" + os.Getenv("LOCAL_RPC")
	PrivateKey    = os.Getenv("LOCAL_PRIVATE_KEY")
	OracleAddress = os.Getenv("ORACLE_ADDRESS")
)
