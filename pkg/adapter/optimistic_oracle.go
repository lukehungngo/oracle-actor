package adapter

import (
	"github.com/anthdm/hollywood/actor"
	"log"
	"oracle-actor/pkg/client"
	"oracle-actor/pkg/generated/contracts/optimistic_oracle"
)

func WatchOptimisticOracle(e *actor.Engine) {
	i, c := client.InstanceOptimisticOracle()

	ch := make(chan *optimistic_oracle.OptimisticOracleReportCreated, 10)
	_, err := i.WatchReportCreated(nil, ch, nil, nil, nil)
	if err != nil {
		log.Printf("error watching report created event: %v\n", err)
		c.Close()
	}

	go func() {
		for reportEvent := range ch {
			log.Println("Receive Report: id=", reportEvent.ReportId)
			e.EventStream.Publish(reportEvent)
		}
	}()
}
