# PresentationFileRequest

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
**Format** | Pointer to **string** | Input file format | [optional] [default to "Pptx"]
**OutputFormat** | Pointer to **string** | Output file format | [optional] 
**Masters** | Pointer to **bool** | If translate master slides | [optional] [default to false]
**Slides** | Pointer to **[]int32** | List of slides to translate | [optional] 

## Methods

### NewPresentationFileRequest

`func NewPresentationFileRequest() *PresentationFileRequest`

NewPresentationFileRequest instantiates a new PresentationFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresentationFileRequestWithDefaults

`func NewPresentationFileRequestWithDefaults() *PresentationFileRequest`

NewPresentationFileRequestWithDefaults instantiates a new PresentationFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *PresentationFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *PresentationFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *PresentationFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.

### HasSourceLanguage

`func (o *PresentationFileRequest) HasSourceLanguage() bool`

HasSourceLanguage returns a boolean if a field has been set.

### GetTargetLanguages

`func (o *PresentationFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *PresentationFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *PresentationFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.

### HasTargetLanguages

`func (o *PresentationFileRequest) HasTargetLanguages() bool`

HasTargetLanguages returns a boolean if a field has been set.

### GetFile

`func (o *PresentationFileRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PresentationFileRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PresentationFileRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *PresentationFileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *PresentationFileRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *PresentationFileRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetOriginalFileName

`func (o *PresentationFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *PresentationFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *PresentationFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *PresentationFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *PresentationFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *PresentationFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *PresentationFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PresentationFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PresentationFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PresentationFileRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PresentationFileRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PresentationFileRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetOrigin

`func (o *PresentationFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PresentationFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PresentationFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *PresentationFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *PresentationFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *PresentationFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetSavingMode

`func (o *PresentationFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *PresentationFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *PresentationFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *PresentationFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetFormat

`func (o *PresentationFileRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PresentationFileRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PresentationFileRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PresentationFileRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOutputFormat

`func (o *PresentationFileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *PresentationFileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *PresentationFileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *PresentationFileRequest) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetMasters

`func (o *PresentationFileRequest) GetMasters() bool`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *PresentationFileRequest) GetMastersOk() (*bool, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *PresentationFileRequest) SetMasters(v bool)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *PresentationFileRequest) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetSlides

`func (o *PresentationFileRequest) GetSlides() []int32`

GetSlides returns the Slides field if non-nil, zero value otherwise.

### GetSlidesOk

`func (o *PresentationFileRequest) GetSlidesOk() (*[]int32, bool)`

GetSlidesOk returns a tuple with the Slides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlides

`func (o *PresentationFileRequest) SetSlides(v []int32)`

SetSlides sets Slides field to given value.

### HasSlides

`func (o *PresentationFileRequest) HasSlides() bool`

HasSlides returns a boolean if a field has been set.

### SetSlidesNil

`func (o *PresentationFileRequest) SetSlidesNil(b bool)`

 SetSlidesNil sets the value for Slides to be an explicit nil

### UnsetSlides
`func (o *PresentationFileRequest) UnsetSlides()`

UnsetSlides ensures that no value is present for Slides, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


