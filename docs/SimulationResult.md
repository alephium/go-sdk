# SimulationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractInputs** | [**[]AddressAssetState**](AddressAssetState.md) |  | 
**GeneratedOutputs** | [**[]AddressAssetState**](AddressAssetState.md) |  | 

## Methods

### NewSimulationResult

`func NewSimulationResult(contractInputs []AddressAssetState, generatedOutputs []AddressAssetState, ) *SimulationResult`

NewSimulationResult instantiates a new SimulationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationResultWithDefaults

`func NewSimulationResultWithDefaults() *SimulationResult`

NewSimulationResultWithDefaults instantiates a new SimulationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractInputs

`func (o *SimulationResult) GetContractInputs() []AddressAssetState`

GetContractInputs returns the ContractInputs field if non-nil, zero value otherwise.

### GetContractInputsOk

`func (o *SimulationResult) GetContractInputsOk() (*[]AddressAssetState, bool)`

GetContractInputsOk returns a tuple with the ContractInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractInputs

`func (o *SimulationResult) SetContractInputs(v []AddressAssetState)`

SetContractInputs sets ContractInputs field to given value.


### GetGeneratedOutputs

`func (o *SimulationResult) GetGeneratedOutputs() []AddressAssetState`

GetGeneratedOutputs returns the GeneratedOutputs field if non-nil, zero value otherwise.

### GetGeneratedOutputsOk

`func (o *SimulationResult) GetGeneratedOutputsOk() (*[]AddressAssetState, bool)`

GetGeneratedOutputsOk returns a tuple with the GeneratedOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedOutputs

`func (o *SimulationResult) SetGeneratedOutputs(v []AddressAssetState)`

SetGeneratedOutputs sets GeneratedOutputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


