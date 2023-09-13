package actor

import "testing"

func TestOracleEventReceivedActor_Receive(t *testing.T) {
	type args struct {
		ctx *actor.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &OracleEventReceivedActor{}
			f.Receive(tt.args.ctx)
		})
	}
}
