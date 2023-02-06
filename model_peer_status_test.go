package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalPeerStatus(t *testing.T) {

	tests := []struct {
		data   []byte
		result PeerStatus
	}{
		{
			data: []byte(`{
			   "type": "Banned",
			   "until": 1675580963000
			 }`),
			result: PeerStatus{Banned: &Banned{
				Until: 1675580963000,
				Type:  "Banned",
			}},
		},
		{
			data: []byte(`{
				   "type": "Penalty",
				   "value": 1234
				 }`),
			result: PeerStatus{Penalty: &Penalty{
				Type:  "Penalty",
				Value: 1234,
			}},
		},
	}

	for _, c := range tests {
		var v PeerStatus
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
