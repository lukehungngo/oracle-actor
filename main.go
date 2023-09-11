package main

import (
	"fmt"
	"github.com/anthdm/hollywood/actor"
	"time"
)

type message struct {
	data string
}

type foo struct{}

func newFoo() actor.Receiver {
	return &foo{}
}

func (f *foo) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("foo has started")
	case *message:
		fmt.Println("foo has received", msg.data)
	}
}

func main() {
	engine := actor.NewEngine()
	pid := engine.Spawn(newFoo, "foo")

	ticker := time.Tick(100 * time.Millisecond)
	for _ = range ticker {
		engine.Send(pid, &message{data: "hello world!"})
	}

	// Stop the actor, but let it process its messages first.
	engine.Poison(pid).Wait()
}
