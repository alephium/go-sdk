# BuildInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseVersion** | **string** |  | 
**Commit** | **string** |  | 

## Methods

### NewBuildInfo

`func NewBuildInfo(releaseVersion string, commit string, ) *BuildInfo`

NewBuildInfo instantiates a new BuildInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildInfoWithDefaults

`func NewBuildInfoWithDefaults() *BuildInfo`

NewBuildInfoWithDefaults instantiates a new BuildInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseVersion

`func (o *BuildInfo) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *BuildInfo) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *BuildInfo) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.


### GetCommit

`func (o *BuildInfo) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *BuildInfo) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *BuildInfo) SetCommit(v string)`

SetCommit sets Commit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


