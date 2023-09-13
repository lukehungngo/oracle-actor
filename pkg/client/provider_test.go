package client

import (
	"reflect"
	"testing"
)

func TestEthClient(t *testing.T) {
	type args struct {
		providerURL string
	}
	tests := []struct {
		name string
		args args
		want *ethclient.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EthClient(tt.args.providerURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EthClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
