# EventSig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Signature** | **string** |  | 
**FieldNames** | **[]string** |  | 
**FieldTypes** | **[]string** |  | 

## Methods

### NewEventSig

`func NewEventSig(name string, signature string, fieldNames []string, fieldTypes []string, ) *EventSig`

NewEventSig instantiates a new EventSig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSigWithDefaults

`func NewEventSigWithDefaults() *EventSig`

NewEventSigWithDefaults instantiates a new EventSig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventSig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventSig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventSig) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *EventSig) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *EventSig) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *EventSig) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetFieldNames

`func (o *EventSig) GetFieldNames() []string`

GetFieldNames returns the FieldNames field if non-nil, zero value otherwise.

### GetFieldNamesOk

`func (o *EventSig) GetFieldNamesOk() (*[]string, bool)`

GetFieldNamesOk returns a tuple with the FieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNames

`func (o *EventSig) SetFieldNames(v []string)`

SetFieldNames sets FieldNames field to given value.


### GetFieldTypes

`func (o *EventSig) GetFieldTypes() []string`

GetFieldTypes returns the FieldTypes field if non-nil, zero value otherwise.

### GetFieldTypesOk

`func (o *EventSig) GetFieldTypesOk() (*[]string, bool)`

GetFieldTypesOk returns a tuple with the FieldTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypes

`func (o *EventSig) SetFieldTypes(v []string)`

SetFieldTypes sets FieldTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


