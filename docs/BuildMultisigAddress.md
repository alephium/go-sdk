# BuildMultisigAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **[]string** |  | 
**KeyTypes** | Pointer to **[]string** |  | [optional] 
**Mrequired** | **int32** |  | 
**MultiSigType** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildMultisigAddress

`func NewBuildMultisigAddress(keys []string, mrequired int32, ) *BuildMultisigAddress`

NewBuildMultisigAddress instantiates a new BuildMultisigAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildMultisigAddressWithDefaults

`func NewBuildMultisigAddressWithDefaults() *BuildMultisigAddress`

NewBuildMultisigAddressWithDefaults instantiates a new BuildMultisigAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *BuildMultisigAddress) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *BuildMultisigAddress) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *BuildMultisigAddress) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### GetKeyTypes

`func (o *BuildMultisigAddress) GetKeyTypes() []string`

GetKeyTypes returns the KeyTypes field if non-nil, zero value otherwise.

### GetKeyTypesOk

`func (o *BuildMultisigAddress) GetKeyTypesOk() (*[]string, bool)`

GetKeyTypesOk returns a tuple with the KeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypes

`func (o *BuildMultisigAddress) SetKeyTypes(v []string)`

SetKeyTypes sets KeyTypes field to given value.

### HasKeyTypes

`func (o *BuildMultisigAddress) HasKeyTypes() bool`

HasKeyTypes returns a boolean if a field has been set.

### GetMrequired

`func (o *BuildMultisigAddress) GetMrequired() int32`

GetMrequired returns the Mrequired field if non-nil, zero value otherwise.

### GetMrequiredOk

`func (o *BuildMultisigAddress) GetMrequiredOk() (*int32, bool)`

GetMrequiredOk returns a tuple with the Mrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrequired

`func (o *BuildMultisigAddress) SetMrequired(v int32)`

SetMrequired sets Mrequired field to given value.


### GetMultiSigType

`func (o *BuildMultisigAddress) GetMultiSigType() string`

GetMultiSigType returns the MultiSigType field if non-nil, zero value otherwise.

### GetMultiSigTypeOk

`func (o *BuildMultisigAddress) GetMultiSigTypeOk() (*string, bool)`

GetMultiSigTypeOk returns a tuple with the MultiSigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSigType

`func (o *BuildMultisigAddress) SetMultiSigType(v string)`

SetMultiSigType sets MultiSigType field to given value.

### HasMultiSigType

`func (o *BuildMultisigAddress) HasMultiSigType() bool`

HasMultiSigType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


