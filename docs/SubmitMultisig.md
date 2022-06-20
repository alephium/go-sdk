# SubmitMultisig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnsignedTx** | **string** |  | 
**Signatures** | **[]string** |  | 

## Methods

### NewSubmitMultisig

`func NewSubmitMultisig(unsignedTx string, signatures []string, ) *SubmitMultisig`

NewSubmitMultisig instantiates a new SubmitMultisig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitMultisigWithDefaults

`func NewSubmitMultisigWithDefaults() *SubmitMultisig`

NewSubmitMultisigWithDefaults instantiates a new SubmitMultisig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsignedTx

`func (o *SubmitMultisig) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *SubmitMultisig) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *SubmitMultisig) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetSignatures

`func (o *SubmitMultisig) GetSignatures() []string`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *SubmitMultisig) GetSignaturesOk() (*[]string, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *SubmitMultisig) SetSignatures(v []string)`

SetSignatures sets Signatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


