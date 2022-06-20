# SubmitTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnsignedTx** | **string** |  | 
**Signature** | **string** |  | 

## Methods

### NewSubmitTransaction

`func NewSubmitTransaction(unsignedTx string, signature string, ) *SubmitTransaction`

NewSubmitTransaction instantiates a new SubmitTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTransactionWithDefaults

`func NewSubmitTransactionWithDefaults() *SubmitTransaction`

NewSubmitTransactionWithDefaults instantiates a new SubmitTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsignedTx

`func (o *SubmitTransaction) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *SubmitTransaction) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *SubmitTransaction) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetSignature

`func (o *SubmitTransaction) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SubmitTransaction) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SubmitTransaction) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


