package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalContractEventWithoutTimestamp(t *testing.T) {
	data := []byte(`{
		"events": [{
			"blockHash": "block-hash",
			"txId": "tx-id",
			"timestamp": 0,
			"contractAddress": "contract-address",
			"eventIndex": 0,
			"fields": [{
				"type": "U256",
				"value": "1"
			}]
		}],
		"nextStart": 1
	}`)

	var events ContractEvents
	err := json.Unmarshal(data, &events)
	assert.Nil(t, err)
	assert.Len(t, events.Events, 1)
	assert.Equal(t, int64(0), events.Events[0].Timestamp)
	assert.Equal(t, int32(1), events.NextStart)
	assert.Len(t, events.Events[0].Fields, 1)
	assert.NotNil(t, events.Events[0].Fields[0].ValU256)
	assert.Equal(t, "1", events.Events[0].Fields[0].ValU256.Value)
}
