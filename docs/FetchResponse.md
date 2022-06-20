# FetchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | [**[][]BlockEntry**]([]BlockEntry.md) |  | 

## Methods

### NewFetchResponse

`func NewFetchResponse(blocks [][]BlockEntry, ) *FetchResponse`

NewFetchResponse instantiates a new FetchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchResponseWithDefaults

`func NewFetchResponseWithDefaults() *FetchResponse`

NewFetchResponseWithDefaults instantiates a new FetchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *FetchResponse) GetBlocks() [][]BlockEntry`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *FetchResponse) GetBlocksOk() (*[][]BlockEntry, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *FetchResponse) SetBlocks(v [][]BlockEntry)`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


