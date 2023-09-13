package main

import (
	"fmt"
	"oracle-actor/config"
)

func main() {
	fmt.Println("Let's gooo")
	fmt.Println(config.LocalRPCWS)
	fmt.Println(config.PrivateKey)
}
