package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalValSucceed(t *testing.T) {
	tests := []struct {
		data   []byte
		result Val
	}{
		{
			data:   []byte(`{"type":"Bool","value":false}`),
			result: Val{ValBool: &ValBool{false, "Bool"}},
		},
		{
			data:   []byte(`{"type":"Bool","value":true}`),
			result: Val{ValBool: &ValBool{true, "Bool"}},
		},
		{
			data:   []byte(`{"type":"I256","value":"123"}`),
			result: Val{ValI256: &ValI256{"123", "I256"}},
		},
		{
			data:   []byte(`{"type":"U256","value":"123"}`),
			result: Val{ValU256: &ValU256{"123", "U256"}},
		},
		{
			data:   []byte(`{"type":"ByteVec","value":"0011223344"}`),
			result: Val{ValByteVec: &ValByteVec{"0011223344", "ByteVec"}},
		},
		{
			data:   []byte(`{"type":"Address","value":"aabbccddeezz"}`),
			result: Val{ValAddress: &ValAddress{"aabbccddeezz", "Address"}},
		},
		{
			data:   []byte(`{"type":"Address","value":"aabbccddeezz"}`),
			result: Val{ValAddress: &ValAddress{"aabbccddeezz", "Address"}},
		},
		{
			data: []byte(`{
				"type":"Array",
				"value":[
					{
						"type":"Array",
						"value":[
							{
								"type":"Bool",
								"value":true
							},
							{
								"type":"ByteVec",
								"value":"0011"
							}
						]
					},
					{
						"type":"Bool",
						"value":true
					},
					{
						"type":"ByteVec",
						"value":"0011"
					}
				]
			}`),
			result: Val{
				ValArray: &ValArray{
					Type: "Array",
					Value: []Val{
						{
							ValArray: &ValArray{
								Type: "Array",
								Value: []Val{
									{ValBool: &ValBool{true, "Bool"}},
									{ValByteVec: &ValByteVec{"0011", "ByteVec"}},
								},
							},
						},
						{ValBool: &ValBool{true, "Bool"}},
						{ValByteVec: &ValByteVec{"0011", "ByteVec"}},
					},
				},
			},
		},
	}

	for _, c := range tests {
		var v Val
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}

func TestUnmarshalValFailed(t *testing.T) {
	tests := []struct {
		data []byte
	}{
		{
			data: []byte(`{"type":"Bool","value":1}`),
		},
		{
			data: []byte(`{"type":"I256","value":123}`),
		},
		{
			data: []byte(`{"type":"U256","value":123}`),
		},
		{
			data: []byte(`{"type":"ByteVec","value":false}`),
		},
		{
			data: []byte(`{"type":"Address","value":{}}`),
		},
		{
			data: []byte(`{"type":"Array","value":[{""}]}`),
		},
	}

	for _, c := range tests {
		var v Val
		err := json.Unmarshal(c.data, &v)
		assert.NotNil(t, err)
	}
}
