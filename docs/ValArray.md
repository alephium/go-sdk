# ValArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**[]Val**](Val.md) |  | 
**Type** | **string** |  | 

## Methods

### NewValArray

`func NewValArray(value []Val, type_ string, ) *ValArray`

NewValArray instantiates a new ValArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValArrayWithDefaults

`func NewValArrayWithDefaults() *ValArray`

NewValArrayWithDefaults instantiates a new ValArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ValArray) GetValue() []Val`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValArray) GetValueOk() (*[]Val, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValArray) SetValue(v []Val)`

SetValue sets Value field to given value.


### GetType

`func (o *ValArray) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValArray) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValArray) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


