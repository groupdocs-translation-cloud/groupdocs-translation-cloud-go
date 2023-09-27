# ImageToTextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | Originnal file format | [optional] [default to "Unknown"]
**Source** | Pointer to **string** | Language of original file | [optional] [default to "en"]
**Targets** | Pointer to **[]string** | List of target languages | [optional] 
**File** | Pointer to **NullableString** | File for translation | [optional] 
**Url** | Pointer to **NullableString** | Link to file for translation | [optional] 
**Rotate** | Pointer to **int32** | Left to write angle to rotate scanned image / pdf | [optional] 
**Ishandwritten** | Pointer to **bool** | is handwritten text | [optional] 
**Origin** | Pointer to **NullableString** | for analysis only | [optional] 
**Route** | Pointer to **NullableString** | endpoints route | [optional] 

## Methods

### NewImageToTextRequest

`func NewImageToTextRequest() *ImageToTextRequest`

NewImageToTextRequest instantiates a new ImageToTextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageToTextRequestWithDefaults

`func NewImageToTextRequestWithDefaults() *ImageToTextRequest`

NewImageToTextRequestWithDefaults instantiates a new ImageToTextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ImageToTextRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ImageToTextRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ImageToTextRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ImageToTextRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSource

`func (o *ImageToTextRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImageToTextRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImageToTextRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ImageToTextRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTargets

`func (o *ImageToTextRequest) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ImageToTextRequest) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ImageToTextRequest) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ImageToTextRequest) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetFile

`func (o *ImageToTextRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ImageToTextRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ImageToTextRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *ImageToTextRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *ImageToTextRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *ImageToTextRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetUrl

`func (o *ImageToTextRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageToTextRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageToTextRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImageToTextRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ImageToTextRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ImageToTextRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRotate

`func (o *ImageToTextRequest) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *ImageToTextRequest) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *ImageToTextRequest) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *ImageToTextRequest) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetIshandwritten

`func (o *ImageToTextRequest) GetIshandwritten() bool`

GetIshandwritten returns the Ishandwritten field if non-nil, zero value otherwise.

### GetIshandwrittenOk

`func (o *ImageToTextRequest) GetIshandwrittenOk() (*bool, bool)`

GetIshandwrittenOk returns a tuple with the Ishandwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIshandwritten

`func (o *ImageToTextRequest) SetIshandwritten(v bool)`

SetIshandwritten sets Ishandwritten field to given value.

### HasIshandwritten

`func (o *ImageToTextRequest) HasIshandwritten() bool`

HasIshandwritten returns a boolean if a field has been set.

### GetOrigin

`func (o *ImageToTextRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ImageToTextRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ImageToTextRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ImageToTextRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ImageToTextRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ImageToTextRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetRoute

`func (o *ImageToTextRequest) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *ImageToTextRequest) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *ImageToTextRequest) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *ImageToTextRequest) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### SetRouteNil

`func (o *ImageToTextRequest) SetRouteNil(b bool)`

 SetRouteNil sets the value for Route to be an explicit nil

### UnsetRoute
`func (o *ImageToTextRequest) UnsetRoute()`

UnsetRoute ensures that no value is present for Route, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


