package actor

import (
	"encoding/json"
	"fmt"
	"github.com/anthdm/hollywood/actor"
	"oracle-actor/pkg/generated/contracts/optimistic_oracle"
)

type OracleEventReceivedActor struct {
}

func (f *OracleEventReceivedActor) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("OracleEventReceivedActor has started")
	case *optimistic_oracle.OptimisticOracleReportCreated:
		data, _ := json.MarshalIndent(msg, "", " ")
		fmt.Println("OracleEventReceivedActor has received", string(data))
	}
}
