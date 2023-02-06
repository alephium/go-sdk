package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMisbehaviorAction(t *testing.T) {

	tests := []struct {
		data   []byte
		result MisbehaviorAction
	}{
		{
			data: []byte(`{
			   "type": "Ban",
			   "peers": ["1.2.3.4", "2.3.4.5"]
			 }`),
			result: MisbehaviorAction{Ban: &Ban{
				Type:  "Ban",
				Peers: []string{"1.2.3.4", "2.3.4.5"},
			}},
		},
		{
			data: []byte(`{
				   "type": "Unban",
				   "peers": ["1.2.3.4", "2.3.4.5"]
				 }`),
			result: MisbehaviorAction{Unban: &Unban{
				Type:  "Unban",
				Peers: []string{"1.2.3.4", "2.3.4.5"},
			}},
		},
	}

	for _, c := range tests {
		var v MisbehaviorAction
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
