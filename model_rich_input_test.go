package alephium

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalRichInputsWithType(t *testing.T) {
	assetInputData := []byte(`{
		"type": "AssetInput",
		"hint": 1,
		"key": "asset-input-key",
		"unlockScript": "00",
		"attoAlphAmount": "1000",
		"address": "asset-input-address",
		"tokens": [],
		"outputRefTxId": "asset-input-output-ref"
	}`)
	var assetInput RichAssetInput
	err := json.Unmarshal(assetInputData, &assetInput)
	assert.Nil(t, err)
	assert.Equal(t, RichAssetInput{
		Hint:           1,
		Key:            "asset-input-key",
		UnlockScript:   "00",
		AttoAlphAmount: "1000",
		Address:        "asset-input-address",
		Tokens:         []Token{},
		OutputRefTxId:  "asset-input-output-ref",
	}, assetInput)

	contractInputData := []byte(`{
		"type": "ContractInput",
		"hint": 2,
		"key": "contract-input-key",
		"attoAlphAmount": "2000",
		"address": "contract-input-address",
		"tokens": [],
		"outputRefTxId": "contract-input-output-ref"
	}`)
	var contractInput RichContractInput
	err = json.Unmarshal(contractInputData, &contractInput)
	assert.Nil(t, err)
	assert.Equal(t, RichContractInput{
		Hint:           2,
		Key:            "contract-input-key",
		AttoAlphAmount: "2000",
		Address:        "contract-input-address",
		Tokens:         []Token{},
		OutputRefTxId:  "contract-input-output-ref",
	}, contractInput)
}

func TestUnmarshalRichInputsInvalidType(t *testing.T) {
	var assetInput RichAssetInput
	err := json.Unmarshal([]byte(`{
		"type": "ContractInput",
		"hint": 1,
		"key": "asset-input-key",
		"unlockScript": "00",
		"attoAlphAmount": "1000",
		"address": "asset-input-address",
		"tokens": [],
		"outputRefTxId": "asset-input-output-ref"
	}`), &assetInput)
	assert.EqualError(t, err, "invalid rich asset input type ContractInput")

	var contractInput RichContractInput
	err = json.Unmarshal([]byte(`{
		"type": "AssetInput",
		"hint": 2,
		"key": "contract-input-key",
		"attoAlphAmount": "2000",
		"address": "contract-input-address",
		"tokens": [],
		"outputRefTxId": "contract-input-output-ref"
	}`), &contractInput)
	assert.EqualError(t, err, "invalid rich contract input type AssetInput")
}

func TestUnmarshalRichTransactionWithTypedInputs(t *testing.T) {
	data := []byte(`{
		"unsigned": {
			"txId": "tx-id",
			"version": 1,
			"networkId": 0,
			"gasAmount": 20000,
			"gasPrice": "100000000000",
			"inputs": [{
				"type": "AssetInput",
				"hint": 1,
				"key": "asset-input-key",
				"unlockScript": "00",
				"attoAlphAmount": "1000",
				"address": "asset-input-address",
				"tokens": [],
				"outputRefTxId": "asset-input-output-ref"
			}],
			"fixedOutputs": [{
				"hint": 2,
				"key": "fixed-output-key",
				"attoAlphAmount": "900",
				"address": "fixed-output-address",
				"tokens": [],
				"lockTime": 0,
				"message": ""
			}]
		},
		"scriptExecutionOk": true,
		"contractInputs": [{
			"type": "ContractInput",
			"hint": 3,
			"key": "contract-input-key",
			"attoAlphAmount": "2000",
			"address": "contract-input-address",
			"tokens": [],
			"outputRefTxId": "contract-input-output-ref"
		}],
		"generatedOutputs": [{
			"type": "ContractOutput",
			"hint": 4,
			"key": "contract-output-key",
			"attoAlphAmount": "3000",
			"address": "contract-output-address",
			"tokens": []
		}],
		"inputSignatures": [],
		"scriptSignatures": []
	}`)

	var tx RichTransaction
	err := json.Unmarshal(data, &tx)
	assert.Nil(t, err)
	assert.Equal(t, "tx-id", tx.Unsigned.TxId)
	assert.Len(t, tx.Unsigned.Inputs, 1)
	assert.Equal(t, "asset-input-key", tx.Unsigned.Inputs[0].Key)
	assert.Len(t, tx.ContractInputs, 1)
	assert.Equal(t, "contract-input-key", tx.ContractInputs[0].Key)
	assert.Len(t, tx.GeneratedOutputs, 1)
	assert.NotNil(t, tx.GeneratedOutputs[0].ContractOutput)
	assert.Equal(t, "contract-output-key", tx.GeneratedOutputs[0].ContractOutput.Key)
}
