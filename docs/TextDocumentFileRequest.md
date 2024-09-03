# TextDocumentFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguage** | **string** | Language of original file | [default to "en"]
**TargetLanguages** | **[]string** | List of target languages | 
**File** | Pointer to **NullableString** | File as byte array | [optional] 
**OriginalFileName** | Pointer to **NullableString** | Type in the file name. If null will be as request ID. | [optional] 
**Url** | Pointer to **NullableString** | Link to file for translation. Ignore, if \&quot;file\&quot; property not null | [optional] 
**SavingMode** | Pointer to **string** | Toggle file saving mode for storage.  Is Files by default. | [optional] 
**Format** | **string** | Input file format | [default to "Docx"]
**OutputFormat** | **string** | output file format | 
**PreserveFormatting** | Pointer to **bool** | If document&#39;s formatting should be preserved, default true | [optional] [default to true]
**Origin** | Pointer to **NullableString** | for analysis only | [optional] 
**Pages** | Pointer to **[]int32** | Choose pages for translation (1-based index). If not present, translate all pages | [optional] 

## Methods

### NewTextDocumentFileRequest

`func NewTextDocumentFileRequest(sourceLanguage string, targetLanguages []string, format string, outputFormat string, ) *TextDocumentFileRequest`

NewTextDocumentFileRequest instantiates a new TextDocumentFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextDocumentFileRequestWithDefaults

`func NewTextDocumentFileRequestWithDefaults() *TextDocumentFileRequest`

NewTextDocumentFileRequestWithDefaults instantiates a new TextDocumentFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *TextDocumentFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *TextDocumentFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *TextDocumentFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.


### GetTargetLanguages

`func (o *TextDocumentFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *TextDocumentFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *TextDocumentFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.


### GetFile

`func (o *TextDocumentFileRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TextDocumentFileRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TextDocumentFileRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *TextDocumentFileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *TextDocumentFileRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *TextDocumentFileRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetOriginalFileName

`func (o *TextDocumentFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *TextDocumentFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *TextDocumentFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *TextDocumentFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *TextDocumentFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *TextDocumentFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *TextDocumentFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TextDocumentFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TextDocumentFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TextDocumentFileRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TextDocumentFileRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TextDocumentFileRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSavingMode

`func (o *TextDocumentFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *TextDocumentFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *TextDocumentFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *TextDocumentFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetFormat

`func (o *TextDocumentFileRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TextDocumentFileRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TextDocumentFileRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetOutputFormat

`func (o *TextDocumentFileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *TextDocumentFileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *TextDocumentFileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.


### GetPreserveFormatting

`func (o *TextDocumentFileRequest) GetPreserveFormatting() bool`

GetPreserveFormatting returns the PreserveFormatting field if non-nil, zero value otherwise.

### GetPreserveFormattingOk

`func (o *TextDocumentFileRequest) GetPreserveFormattingOk() (*bool, bool)`

GetPreserveFormattingOk returns a tuple with the PreserveFormatting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFormatting

`func (o *TextDocumentFileRequest) SetPreserveFormatting(v bool)`

SetPreserveFormatting sets PreserveFormatting field to given value.

### HasPreserveFormatting

`func (o *TextDocumentFileRequest) HasPreserveFormatting() bool`

HasPreserveFormatting returns a boolean if a field has been set.

### GetOrigin

`func (o *TextDocumentFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *TextDocumentFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *TextDocumentFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *TextDocumentFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *TextDocumentFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *TextDocumentFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetPages

`func (o *TextDocumentFileRequest) GetPages() []int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *TextDocumentFileRequest) GetPagesOk() (*[]int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *TextDocumentFileRequest) SetPages(v []int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *TextDocumentFileRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPagesNil

`func (o *TextDocumentFileRequest) SetPagesNil(b bool)`

 SetPagesNil sets the value for Pages to be an explicit nil

### UnsetPages
`func (o *TextDocumentFileRequest) UnsetPages()`

UnsetPages ensures that no value is present for Pages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


