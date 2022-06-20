# SweepAddressTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** |  | 
**UnsignedTx** | **string** |  | 
**GasAmount** | **int32** |  | 
**GasPrice** | **string** |  | 

## Methods

### NewSweepAddressTransaction

`func NewSweepAddressTransaction(txId string, unsignedTx string, gasAmount int32, gasPrice string, ) *SweepAddressTransaction`

NewSweepAddressTransaction instantiates a new SweepAddressTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepAddressTransactionWithDefaults

`func NewSweepAddressTransactionWithDefaults() *SweepAddressTransaction`

NewSweepAddressTransactionWithDefaults instantiates a new SweepAddressTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *SweepAddressTransaction) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *SweepAddressTransaction) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *SweepAddressTransaction) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetUnsignedTx

`func (o *SweepAddressTransaction) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *SweepAddressTransaction) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *SweepAddressTransaction) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetGasAmount

`func (o *SweepAddressTransaction) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *SweepAddressTransaction) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *SweepAddressTransaction) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *SweepAddressTransaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *SweepAddressTransaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *SweepAddressTransaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


