package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBuildChainedTx(t *testing.T) {
	tests := []struct {
		data   []byte
		result BuildChainedTx
	}{
		{
			data: []byte(`{
				"type": "DeployContract",
				"value": {
					"fromPublicKey": "public-key",
					"bytecode": "bytecode"
				}
			}`),
			result: BuildChainedTx{BuildChainedDeployContractTx: &BuildChainedDeployContractTx{
				Type: "DeployContract",
				Value: BuildDeployContractTx{
					FromPublicKey: "public-key",
					Bytecode:      "bytecode",
				},
			}},
		},
		{
			data: []byte(`{
				"type": "ExecuteScript",
				"value": {
					"fromPublicKey": "public-key",
					"bytecode": "bytecode"
				}
			}`),
			result: BuildChainedTx{BuildChainedExecuteScriptTx: &BuildChainedExecuteScriptTx{
				Type: "ExecuteScript",
				Value: BuildExecuteScriptTx{
					FromPublicKey: "public-key",
					Bytecode:      "bytecode",
				},
			}},
		},
		{
			data: []byte(`{
				"type": "Transfer",
				"value": {
					"fromPublicKey": "public-key",
					"destinations": [{"address": "address"}]
				}
			}`),
			result: BuildChainedTx{BuildChainedTransferTx: &BuildChainedTransferTx{
				Type: "Transfer",
				Value: BuildTransferTx{
					FromPublicKey: "public-key",
					Destinations:  []Destination{{Address: "address"}},
				},
			}},
		},
	}

	for _, c := range tests {
		var v BuildChainedTx
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
