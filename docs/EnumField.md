# EnumField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | [**Val**](Val.md) |  | 

## Methods

### NewEnumField

`func NewEnumField(name string, value Val, ) *EnumField`

NewEnumField instantiates a new EnumField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumFieldWithDefaults

`func NewEnumFieldWithDefaults() *EnumField`

NewEnumFieldWithDefaults instantiates a new EnumField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnumField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumField) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *EnumField) GetValue() Val`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnumField) GetValueOk() (*Val, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnumField) SetValue(v Val)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


