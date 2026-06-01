package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBuildTransferTxResult(t *testing.T) {
	tests := []struct {
		data   []byte
		result BuildTransferTxResult
	}{
		{
			data: []byte(`{
				"unsignedTx": "unsigned",
				"gasAmount": 1000,
				"gasPrice": "100",
				"txId": "tx",
				"fromGroup": 0,
				"toGroup": 1
			}`),
			result: BuildTransferTxResult{BuildSimpleTransferTxResult: &BuildSimpleTransferTxResult{
				UnsignedTx: "unsigned",
				GasAmount:  1000,
				GasPrice:   "100",
				TxId:       "tx",
				FromGroup:  0,
				ToGroup:    1,
			}},
		},
		{
			data: []byte(`{
				"unsignedTx": "unsigned",
				"gasAmount": 1000,
				"gasPrice": "100",
				"txId": "tx",
				"fromGroup": 0,
				"toGroup": 1,
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
			result: BuildTransferTxResult{BuildGrouplessTransferTxResult: &BuildGrouplessTransferTxResult{
				UnsignedTx: "unsigned",
				GasAmount:  1000,
				GasPrice:   "100",
				TxId:       "tx",
				FromGroup:  0,
				ToGroup:    1,
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
		var v BuildTransferTxResult
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
