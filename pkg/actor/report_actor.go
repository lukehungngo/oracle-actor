package actor

import (
	"encoding/json"
	"fmt"
	"github.com/anthdm/hollywood/actor"
	"oracle-actor/pkg/generated/contracts/optimistic_oracle"
	"time"
)

type OracleReportCreatedActor struct {
}

func NewOracleReportCreatedActor() actor.Receiver {
	return &OracleReportCreatedActor{}
}

func (f *OracleReportCreatedActor) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Printf("OracleEventReceivedActor has started")
	case *optimistic_oracle.OptimisticOracleReportCreated:
		data, _ := json.MarshalIndent(msg, "", " ")
		fmt.Println("OracleEventReceivedActor has received", string(data))
		time.Sleep(3 * time.Second)
	}
}
