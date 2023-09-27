# HtmlFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguage** | Pointer to **string** | Language of original file | [optional] [default to "en"]
**TargetLanguages** | Pointer to **[]string** | List of target languages | [optional] 
**File** | Pointer to **NullableString** | File as byte array | [optional] 
**OriginalFileName** | Pointer to **NullableString** | Type in the file name. If null will be as request ID. | [optional] 
**Url** | Pointer to **NullableString** | Link to file for translation. Ignore, if \&quot;file\&quot; property not null | [optional] 
**Origin** | Pointer to **NullableString** | Url or name of application using this SDK. Not required. | [optional] 
**SavingMode** | Pointer to **string** | Toggle file saving mode for storage.  Is Files by default. | [optional] 
**OutputFormat** | Pointer to **string** | output file format | [optional] 

## Methods

### NewHtmlFileRequest

`func NewHtmlFileRequest() *HtmlFileRequest`

NewHtmlFileRequest instantiates a new HtmlFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHtmlFileRequestWithDefaults

`func NewHtmlFileRequestWithDefaults() *HtmlFileRequest`

NewHtmlFileRequestWithDefaults instantiates a new HtmlFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *HtmlFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *HtmlFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *HtmlFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.

### HasSourceLanguage

`func (o *HtmlFileRequest) HasSourceLanguage() bool`

HasSourceLanguage returns a boolean if a field has been set.

### GetTargetLanguages

`func (o *HtmlFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *HtmlFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *HtmlFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.

### HasTargetLanguages

`func (o *HtmlFileRequest) HasTargetLanguages() bool`

HasTargetLanguages returns a boolean if a field has been set.

### GetFile

`func (o *HtmlFileRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *HtmlFileRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *HtmlFileRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *HtmlFileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *HtmlFileRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *HtmlFileRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetOriginalFileName

`func (o *HtmlFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *HtmlFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *HtmlFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *HtmlFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *HtmlFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *HtmlFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *HtmlFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HtmlFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HtmlFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HtmlFileRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *HtmlFileRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *HtmlFileRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetOrigin

`func (o *HtmlFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *HtmlFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *HtmlFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *HtmlFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *HtmlFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *HtmlFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetSavingMode

`func (o *HtmlFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *HtmlFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *HtmlFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *HtmlFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetOutputFormat

`func (o *HtmlFileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *HtmlFileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *HtmlFileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *HtmlFileRequest) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


