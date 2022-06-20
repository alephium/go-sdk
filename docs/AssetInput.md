# AssetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputRef** | [**OutputRef**](OutputRef.md) |  | 
**UnlockScript** | **string** |  | 

## Methods

### NewAssetInput

`func NewAssetInput(outputRef OutputRef, unlockScript string, ) *AssetInput`

NewAssetInput instantiates a new AssetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetInputWithDefaults

`func NewAssetInputWithDefaults() *AssetInput`

NewAssetInputWithDefaults instantiates a new AssetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputRef

`func (o *AssetInput) GetOutputRef() OutputRef`

GetOutputRef returns the OutputRef field if non-nil, zero value otherwise.

### GetOutputRefOk

`func (o *AssetInput) GetOutputRefOk() (*OutputRef, bool)`

GetOutputRefOk returns a tuple with the OutputRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRef

`func (o *AssetInput) SetOutputRef(v OutputRef)`

SetOutputRef sets OutputRef field to given value.


### GetUnlockScript

`func (o *AssetInput) GetUnlockScript() string`

GetUnlockScript returns the UnlockScript field if non-nil, zero value otherwise.

### GetUnlockScriptOk

`func (o *AssetInput) GetUnlockScriptOk() (*string, bool)`

GetUnlockScriptOk returns a tuple with the UnlockScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockScript

`func (o *AssetInput) SetUnlockScript(v string)`

SetUnlockScript sets UnlockScript field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


