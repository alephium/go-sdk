package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalTxStatus(t *testing.T) {
	tests := []struct {
		data   []byte
		result TxStatus
	}{
		{
			data:   []byte(`{"type":"MemPooled"}`),
			result: TxStatus{MemPooled: &MemPooled{"MemPooled"}},
		},
		{
			data:   []byte(`{"type":"TxNotFound"}`),
			result: TxStatus{TxNotFound: &TxNotFound{"TxNotFound"}},
		},
		{
			data: []byte(`{
				"type":"Confirmed",
				"blockHash":"00",
				"txIndex":0,
				"chainConfirmations":0,
				"fromGroupConfirmations":1,
				"toGroupConfirmations":2
			}`),
			result: TxStatus{Confirmed: &Confirmed{
				Type:                   "Confirmed",
				BlockHash:              "00",
				TxIndex:                0,
				ChainConfirmations:     0,
				FromGroupConfirmations: 1,
				ToGroupConfirmations:   2,
			}},
		},
	}

	for _, c := range tests {
		var v TxStatus
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
