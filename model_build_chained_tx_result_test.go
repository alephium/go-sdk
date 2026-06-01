package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBuildChainedTxResult(t *testing.T) {
	simulationResult := SimulationResult{
		ContractInputs:   []AddressAssetState{},
		GeneratedOutputs: []AddressAssetState{},
	}
	tests := []struct {
		data   []byte
		result BuildChainedTxResult
	}{
		{
			data: []byte(`{
				"type": "DeployContract",
				"value": {
					"fromGroup": 0,
					"toGroup": 1,
					"unsignedTx": "unsigned",
					"gasAmount": 1000,
					"gasPrice": "100",
					"txId": "tx",
					"contractAddress": "contract"
				}
			}`),
			result: BuildChainedTxResult{BuildChainedDeployContractTxResult: &BuildChainedDeployContractTxResult{
				Type: "DeployContract",
				Value: BuildSimpleDeployContractTxResult{
					FromGroup:       0,
					ToGroup:         1,
					UnsignedTx:      "unsigned",
					GasAmount:       1000,
					GasPrice:        "100",
					TxId:            "tx",
					ContractAddress: "contract",
				},
			}},
		},
		{
			data: []byte(`{
				"type": "ExecuteScript",
				"value": {
					"fromGroup": 0,
					"toGroup": 1,
					"unsignedTx": "unsigned",
					"gasAmount": 1000,
					"gasPrice": "100",
					"txId": "tx",
					"simulationResult": {
						"contractInputs": [],
						"generatedOutputs": []
					}
				}
			}`),
			result: BuildChainedTxResult{BuildChainedExecuteScriptTxResult: &BuildChainedExecuteScriptTxResult{
				Type: "ExecuteScript",
				Value: BuildSimpleExecuteScriptTxResult{
					FromGroup:        0,
					ToGroup:          1,
					UnsignedTx:       "unsigned",
					GasAmount:        1000,
					GasPrice:         "100",
					TxId:             "tx",
					SimulationResult: simulationResult,
				},
			}},
		},
		{
			data: []byte(`{
				"type": "Transfer",
				"value": {
					"unsignedTx": "unsigned",
					"gasAmount": 1000,
					"gasPrice": "100",
					"txId": "tx",
					"fromGroup": 0,
					"toGroup": 1
				}
			}`),
			result: BuildChainedTxResult{BuildChainedTransferTxResult: &BuildChainedTransferTxResult{
				Type: "Transfer",
				Value: BuildSimpleTransferTxResult{
					UnsignedTx: "unsigned",
					GasAmount:  1000,
					GasPrice:   "100",
					TxId:       "tx",
					FromGroup:  0,
					ToGroup:    1,
				},
			}},
		},
	}

	for _, c := range tests {
		var v BuildChainedTxResult
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
