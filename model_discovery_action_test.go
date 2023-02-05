package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalDiscoveryAction(t *testing.T) {

	tests := []struct {
		data   []byte
		result DiscoveryAction
	}{
		{
			data: []byte(`{
			   "type": "Reachable",
			   "peers": ["1.2.3.4", "2.3.4.5"]
			 }`),
			result: DiscoveryAction{Reachable: &Reachable{
				Type:  "Reachable",
				Peers: []string{"1.2.3.4", "2.3.4.5"},
			}},
		},
		{
			data: []byte(`{
				   "type": "Unreachable",
				   "peers": ["1.2.3.4", "2.3.4.5"]
				 }`),
			result: DiscoveryAction{Unreachable: &Unreachable{
				Type:  "Unreachable",
				Peers: []string{"1.2.3.4", "2.3.4.5"},
			}},
		},
	}

	for _, c := range tests {
		var v DiscoveryAction
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
