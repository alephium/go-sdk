package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalContractEventByBlockHashWithoutTimestamp(t *testing.T) {
	data := []byte(`{
		"events": [{
			"txId": "tx-id",
			"contractAddress": "contract-address",
			"eventIndex": 0,
			"fields": [{
				"type": "U256",
				"value": "1"
			}]
		}]
	}`)

	var events ContractEventsByBlockHash
	err := json.Unmarshal(data, &events)
	assert.Nil(t, err)
	assert.Len(t, events.Events, 1)
	assert.Equal(t, int64(0), events.Events[0].Timestamp)
	assert.Len(t, events.Events[0].Fields, 1)
	assert.NotNil(t, events.Events[0].Fields[0].ValU256)
	assert.Equal(t, "1", events.Events[0].Fields[0].ValU256.Value)
}
