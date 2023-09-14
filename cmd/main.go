package main

import (
	"fmt"
	"github.com/anthdm/hollywood/log"
	oracleActor "oracle-actor/pkg/actor"
	"oracle-actor/pkg/adapter"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Let's gooo")

	log.SetLevel(log.LevelInfo)
	e := oracleActor.NewEngine()
	done := make(chan struct{})
	go oracleActor.WatchEventMultiActor(e, done)

	go adapter.WatchOptimisticOracle(e)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)
	<-signalCh
	close(done)
}
