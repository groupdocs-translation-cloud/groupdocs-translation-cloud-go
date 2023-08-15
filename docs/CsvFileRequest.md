# CsvFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguage** | **string** | Language of original file | [default to "en"]
**TargetLanguages** | **[]string** | List of target languages | 
**File** | Pointer to **NullableString** | File as byte array | [optional] 
**OriginalFileName** | Pointer to **NullableString** | Type in the file name. If null will be as request ID. | [optional] 
**Url** | Pointer to **NullableString** | Link to file for translation. Ignore, if \&quot;file\&quot; property not null | [optional] 
**Origin** | Pointer to **NullableString** | Url or name of application using this SDK. Not required. | [optional] 
**SavingMode** | Pointer to **string** | Toggle file saving mode for storage.  Is Files by default. | [optional] 
**Format** | Pointer to **string** | Input file format | [optional] [default to "Csv"]
**OutFormat** | **string** | output file format | 
**Separator** | Pointer to **NullableString** | Separator in files | [optional] 

## Methods

### NewCsvFileRequest

`func NewCsvFileRequest(sourceLanguage string, targetLanguages []string, outFormat string, ) *CsvFileRequest`

NewCsvFileRequest instantiates a new CsvFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsvFileRequestWithDefaults

`func NewCsvFileRequestWithDefaults() *CsvFileRequest`

NewCsvFileRequestWithDefaults instantiates a new CsvFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *CsvFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *CsvFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *CsvFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.


### GetTargetLanguages

`func (o *CsvFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *CsvFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *CsvFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.


### GetFile

`func (o *CsvFileRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CsvFileRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CsvFileRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CsvFileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *CsvFileRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *CsvFileRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetOriginalFileName

`func (o *CsvFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *CsvFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *CsvFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *CsvFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *CsvFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *CsvFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *CsvFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CsvFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CsvFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CsvFileRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CsvFileRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CsvFileRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetOrigin

`func (o *CsvFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CsvFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CsvFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CsvFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *CsvFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *CsvFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetSavingMode

`func (o *CsvFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *CsvFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *CsvFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *CsvFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetFormat

`func (o *CsvFileRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CsvFileRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CsvFileRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CsvFileRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOutFormat

`func (o *CsvFileRequest) GetOutFormat() string`

GetOutFormat returns the OutFormat field if non-nil, zero value otherwise.

### GetOutFormatOk

`func (o *CsvFileRequest) GetOutFormatOk() (*string, bool)`

GetOutFormatOk returns a tuple with the OutFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutFormat

`func (o *CsvFileRequest) SetOutFormat(v string)`

SetOutFormat sets OutFormat field to given value.


### GetSeparator

`func (o *CsvFileRequest) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *CsvFileRequest) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *CsvFileRequest) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *CsvFileRequest) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *CsvFileRequest) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *CsvFileRequest) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


