package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalOutput(t *testing.T) {

	// Tx https://explorer-v112.testnet.alephium.org/transactions/d78e6b405f1e7239f660e11f08e9b20775dab396f617673e341ef1c41596167f
	// is causing some unmarshalling error: "data matches more than one schema in oneOf(Output)"
	tests := []struct {
		data   []byte
		result Output
	}{
		{
			data: []byte(`{
			   "type": "ContractOutput",
			   "hint": 905276618,
			   "key": "e68624d324aba3cfbb84c1edab5167ba167244c2bc40bae33b1941645a75a93e",
			   "attoAlphAmount": "1000000000000000000",
			   "address": "2ACpX9pgDSSVAbBZJrNXcDUW7woxfjsvwYk6EKmEmswy9",
			   "tokens": []
			 }`),
			result: Output{ContractOutput: &ContractOutput{
				Type:           "ContractOutput",
				Hint:           905276618,
				Key:            "e68624d324aba3cfbb84c1edab5167ba167244c2bc40bae33b1941645a75a93e",
				AttoAlphAmount: "1000000000000000000",
				Address:        "2ACpX9pgDSSVAbBZJrNXcDUW7woxfjsvwYk6EKmEmswy9",
				Tokens:         []Token{},
			}},
		},
		{
			data: []byte(`{
				   "type": "AssetOutput",
				   "hint": 933512263,
				   "key": "68f87e99a2396a31e814f926f11c75cfbd5a44515c0ee496ca48c85c31d7d476",
				   "attoAlphAmount": "1992242400000000000",
				   "address": "1DrDyTr9RpRsQnDnXo2YRiPzPW4ooHX5LLoqXrqfMrpQH",
				   "tokens": [],
				   "lockTime": 0,
				   "message": ""
				 }`),
			result: Output{AssetOutput: &AssetOutput{
				Type:           "AssetOutput",
				Hint:           933512263,
				Key:            "68f87e99a2396a31e814f926f11c75cfbd5a44515c0ee496ca48c85c31d7d476",
				AttoAlphAmount: "1992242400000000000",
				Address:        "1DrDyTr9RpRsQnDnXo2YRiPzPW4ooHX5LLoqXrqfMrpQH",
				Tokens:         []Token{},
				LockTime:       0,
				Message:        "",
			}},
		},
	}

	for _, c := range tests {
		var v Output
		err := json.Unmarshal(c.data, &v)
		assert.Nil(t, err)
		assert.Equal(t, v, c.result)
	}
}
