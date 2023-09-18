package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalCallContractResult(t *testing.T) {
	tests := []struct {
		data   []byte
		result CallContractResult
	}{
		{
			data: []byte(`{
				"type":"CallContractSucceeded",
				"returns": [{"type": "U256","value": "1"}],
				"gasUsed":1000,
				"contracts": [
					{
						"address": "uomjgUz6D4tLejTkQtbNJMY8apAjTm1bgQf7em1wDV7S",
						"bytecode": "00010700000000000118",
						"codeHash": "0000000000000000000000000000000000000000000000000000000000000000",
						"immFields": [],
						"mutFields": [],
						"asset": {"attoAlphAmount": "1000000000000000000","tokens": []}
					}
				],
				"txInputs": [],
				"txOutputs": [],
				"events": []
			}`),
			result: CallContractResult{CallContractSucceeded: &CallContractSucceeded{
				Returns: []Val{{ValU256: &ValU256{Value: "1", Type: "U256"}}},
				GasUsed: 1000,
				Contracts: []ContractState{{
					Address:   "uomjgUz6D4tLejTkQtbNJMY8apAjTm1bgQf7em1wDV7S",
					Bytecode:  "00010700000000000118",
					CodeHash:  "0000000000000000000000000000000000000000000000000000000000000000",
					ImmFields: []Val{},
					MutFields: []Val{},
					Asset:     AssetState{"1000000000000000000", []Token{}},
				}},
				TxInputs:  []string{},
				TxOutputs: []Output{},
				Events:    []ContractEventByTxId{},
				Type:      "CallContractSucceeded",
			}},
		},
		{
			data: []byte(`{"type":"CallContractFailed","Error":"invalid contract id"}`),
			result: CallContractResult{CallContractFailed: &CallContractFailed{
				Error: "invalid contract id",
				Type:  "CallContractFailed",
			}},
		},
	}

	for _, c := range tests {
		var v CallContractResult
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
