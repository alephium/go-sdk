package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBuildExecuteScriptTxResult(t *testing.T) {
	simulationResult := SimulationResult{
		ContractInputs:   []AddressAssetState{},
		GeneratedOutputs: []AddressAssetState{},
	}
	tests := []struct {
		data   []byte
		result BuildExecuteScriptTxResult
	}{
		{
			data: []byte(`{
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
			}`),
			result: BuildExecuteScriptTxResult{BuildSimpleExecuteScriptTxResult: &BuildSimpleExecuteScriptTxResult{
				FromGroup:        0,
				ToGroup:          1,
				UnsignedTx:       "unsigned",
				GasAmount:        1000,
				GasPrice:         "100",
				TxId:             "tx",
				SimulationResult: simulationResult,
			}},
		},
		{
			data: []byte(`{
				"fromGroup": 0,
				"toGroup": 1,
				"unsignedTx": "unsigned",
				"gasAmount": 1000,
				"gasPrice": "100",
				"txId": "tx",
				"simulationResult": {
					"contractInputs": [],
					"generatedOutputs": []
				},
				"fundingTxs": [
					{
						"unsignedTx": "funding",
						"gasAmount": 10,
						"gasPrice": "1",
						"txId": "funding-tx",
						"fromGroup": 2,
						"toGroup": 3
					}
				]
			}`),
			result: BuildExecuteScriptTxResult{BuildGrouplessExecuteScriptTxResult: &BuildGrouplessExecuteScriptTxResult{
				FromGroup:        0,
				ToGroup:          1,
				UnsignedTx:       "unsigned",
				GasAmount:        1000,
				GasPrice:         "100",
				TxId:             "tx",
				SimulationResult: simulationResult,
				FundingTxs: []BuildSimpleTransferTxResult{{
					UnsignedTx: "funding",
					GasAmount:  10,
					GasPrice:   "1",
					TxId:       "funding-tx",
					FromGroup:  2,
					ToGroup:    3,
				}},
			}},
		},
	}

	for _, c := range tests {
		var v BuildExecuteScriptTxResult
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
