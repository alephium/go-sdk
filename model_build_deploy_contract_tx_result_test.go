package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBuildDeployContractTxResult(t *testing.T) {
	tests := []struct {
		data   []byte
		result BuildDeployContractTxResult
	}{
		{
			data: []byte(`{
				"fromGroup": 0,
				"toGroup": 1,
				"unsignedTx": "unsigned",
				"gasAmount": 1000,
				"gasPrice": "100",
				"txId": "tx",
				"contractAddress": "contract"
			}`),
			result: BuildDeployContractTxResult{BuildSimpleDeployContractTxResult: &BuildSimpleDeployContractTxResult{
				FromGroup:       0,
				ToGroup:         1,
				UnsignedTx:      "unsigned",
				GasAmount:       1000,
				GasPrice:        "100",
				TxId:            "tx",
				ContractAddress: "contract",
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
				"contractAddress": "contract",
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
			result: BuildDeployContractTxResult{BuildGrouplessDeployContractTxResult: &BuildGrouplessDeployContractTxResult{
				FromGroup:       0,
				ToGroup:         1,
				UnsignedTx:      "unsigned",
				GasAmount:       1000,
				GasPrice:        "100",
				TxId:            "tx",
				ContractAddress: "contract",
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
		var v BuildDeployContractTxResult
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
