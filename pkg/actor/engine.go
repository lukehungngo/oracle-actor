package actor

import (
	"encoding/json"
	"github.com/anthdm/hollywood/actor"
	"log"
	"oracle-actor/pkg/generated/contracts/optimistic_oracle"
	"time"
)

func NewEngine() *actor.Engine {
	var e = actor.NewEngine()
	return e
}

func WatchEventSpawnOnAction(e *actor.Engine, done chan struct{}) {

	log.Println("Begin Watch Actor Event")
	accumulateTime := int64(0)
	eventSub := e.EventStream.Subscribe(func(event any) {
		switch event.(type) {
		case *actor.TerminationEvent:
			log.Println("process terminated:", event.(*actor.TerminationEvent).PID)
		case *optimistic_oracle.OptimisticOracleReportCreated:
			start := time.Now()
			pid := e.Spawn(NewOracleReportCreatedActor, "OracleReportCreatedActor")
			e.Send(pid, event)
			e.Poison(pid)
			accumulateTime += time.Now().Sub(start).Nanoseconds()
			log.Println("AccumulationTime=", accumulateTime)
		default:
			data, _ := json.MarshalIndent(event, "", " ")
			log.Println("received event", string(data))
		}

	})

	<-done

	e.EventStream.Unsubscribe(eventSub)
}

func WatchEventNoSpawn(e *actor.Engine, done chan struct{}) {

	log.Println("Begin Watch Actor Event")
	pid := e.Spawn(NewOracleReportCreatedActor, "OracleReportCreatedActor")

	accumulateTime := int64(0)
	eventSub := e.EventStream.Subscribe(func(event any) {
		switch event.(type) {
		case *actor.TerminationEvent:
			log.Println("process terminated:", event.(*actor.TerminationEvent).PID)
		case *optimistic_oracle.OptimisticOracleReportCreated:
			start := time.Now()
			e.Send(pid, event)
			accumulateTime += time.Now().Sub(start).Nanoseconds()
			log.Println("AccumulationTime=", accumulateTime)
		default:
			data, _ := json.MarshalIndent(event, "", " ")
			log.Println("received event", string(data))
		}

	})

	<-done

	e.EventStream.Unsubscribe(eventSub)
}

func WatchEventMultiActor(e *actor.Engine, done chan struct{}) {

	log.Println("Begin Watch Actor Event")
	maxPid := 5
	pids := []*actor.PID{}
	for i := 0; i < maxPid; i++ {
		pid := e.Spawn(NewOracleReportCreatedActor, "OracleReportCreatedActor_"+string(i))
		pids = append(pids, pid)
	}

	accumulateTime := int64(0)
	counter := 0
	eventSub := e.EventStream.Subscribe(func(event any) {
		switch event.(type) {
		case *actor.TerminationEvent:
			log.Println("process terminated:", event.(*actor.TerminationEvent).PID)
		case *optimistic_oracle.OptimisticOracleReportCreated:
			start := time.Now()
			e.Send(pids[counter%maxPid], event)
			accumulateTime += time.Now().Sub(start).Nanoseconds()
			log.Println("AccumulationTime=", accumulateTime)
			counter++
		default:
			data, _ := json.MarshalIndent(event, "", " ")
			log.Println("received event", string(data))
		}

	})

	<-done

	e.EventStream.Unsubscribe(eventSub)
}
