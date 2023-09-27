# HugoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **NullableString** | File as byte array | [optional] 
**Url** | Pointer to **NullableString** | Link to file for translation | [optional] 

## Methods

### NewHugoRequest

`func NewHugoRequest() *HugoRequest`

NewHugoRequest instantiates a new HugoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHugoRequestWithDefaults

`func NewHugoRequestWithDefaults() *HugoRequest`

NewHugoRequestWithDefaults instantiates a new HugoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *HugoRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *HugoRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *HugoRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *HugoRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *HugoRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *HugoRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetUrl

`func (o *HugoRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HugoRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HugoRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HugoRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *HugoRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *HugoRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


