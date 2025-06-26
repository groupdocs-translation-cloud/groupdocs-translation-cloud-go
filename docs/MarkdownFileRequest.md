# MarkdownFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguage** | **string** | Language of original file | [default to "en"]
**TargetLanguages** | **[]string** | List of target languages | 
**OriginalFileName** | Pointer to **NullableString** | Type in the file name. If null will be as request ID. | [optional] 
**Url** | **string** | Link to file for translation. Ignore, if \&quot;file\&quot; property not null | 
**Origin** | Pointer to **NullableString** | Url or name of the application using this SDK. Not required. | [optional] 
**IsNeedAlignment** | Pointer to **bool** | Do result formating like the source. This option needs more expensive requests. | [optional] 
**TranslationDictionary** | Pointer to **map[string]string** | Set a specific translation between source and target words. | [optional] 
**SavingMode** | Pointer to **string** | Toggle file saving mode for storage.  Is Files by default. | [optional] 
**OutputFormat** | **string** | output file format | 
**ShortCodeList** | Pointer to **map[string][]string** | Dictionary of short code names and parameters names to translate | [optional] 
**FrontMatterList** | Pointer to **[][]string** | List of lists of frontmatter paths | [optional] 

## Methods

### NewMarkdownFileRequest

`func NewMarkdownFileRequest(sourceLanguage string, targetLanguages []string, url string, outputFormat string, ) *MarkdownFileRequest`

NewMarkdownFileRequest instantiates a new MarkdownFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkdownFileRequestWithDefaults

`func NewMarkdownFileRequestWithDefaults() *MarkdownFileRequest`

NewMarkdownFileRequestWithDefaults instantiates a new MarkdownFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *MarkdownFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *MarkdownFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *MarkdownFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.


### GetTargetLanguages

`func (o *MarkdownFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *MarkdownFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *MarkdownFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.


### GetOriginalFileName

`func (o *MarkdownFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *MarkdownFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *MarkdownFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *MarkdownFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *MarkdownFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *MarkdownFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *MarkdownFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarkdownFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarkdownFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetOrigin

`func (o *MarkdownFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MarkdownFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MarkdownFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *MarkdownFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *MarkdownFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *MarkdownFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetIsNeedAlignment

`func (o *MarkdownFileRequest) GetIsNeedAlignment() bool`

GetIsNeedAlignment returns the IsNeedAlignment field if non-nil, zero value otherwise.

### GetIsNeedAlignmentOk

`func (o *MarkdownFileRequest) GetIsNeedAlignmentOk() (*bool, bool)`

GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeedAlignment

`func (o *MarkdownFileRequest) SetIsNeedAlignment(v bool)`

SetIsNeedAlignment sets IsNeedAlignment field to given value.

### HasIsNeedAlignment

`func (o *MarkdownFileRequest) HasIsNeedAlignment() bool`

HasIsNeedAlignment returns a boolean if a field has been set.

### GetTranslationDictionary

`func (o *MarkdownFileRequest) GetTranslationDictionary() map[string]string`

GetTranslationDictionary returns the TranslationDictionary field if non-nil, zero value otherwise.

### GetTranslationDictionaryOk

`func (o *MarkdownFileRequest) GetTranslationDictionaryOk() (*map[string]string, bool)`

GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationDictionary

`func (o *MarkdownFileRequest) SetTranslationDictionary(v map[string]string)`

SetTranslationDictionary sets TranslationDictionary field to given value.

### HasTranslationDictionary

`func (o *MarkdownFileRequest) HasTranslationDictionary() bool`

HasTranslationDictionary returns a boolean if a field has been set.

### SetTranslationDictionaryNil

`func (o *MarkdownFileRequest) SetTranslationDictionaryNil(b bool)`

 SetTranslationDictionaryNil sets the value for TranslationDictionary to be an explicit nil

### UnsetTranslationDictionary
`func (o *MarkdownFileRequest) UnsetTranslationDictionary()`

UnsetTranslationDictionary ensures that no value is present for TranslationDictionary, not even an explicit nil
### GetSavingMode

`func (o *MarkdownFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *MarkdownFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *MarkdownFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *MarkdownFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetOutputFormat

`func (o *MarkdownFileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *MarkdownFileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *MarkdownFileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.


### GetShortCodeList

`func (o *MarkdownFileRequest) GetShortCodeList() map[string][]string`

GetShortCodeList returns the ShortCodeList field if non-nil, zero value otherwise.

### GetShortCodeListOk

`func (o *MarkdownFileRequest) GetShortCodeListOk() (*map[string][]string, bool)`

GetShortCodeListOk returns a tuple with the ShortCodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCodeList

`func (o *MarkdownFileRequest) SetShortCodeList(v map[string][]string)`

SetShortCodeList sets ShortCodeList field to given value.

### HasShortCodeList

`func (o *MarkdownFileRequest) HasShortCodeList() bool`

HasShortCodeList returns a boolean if a field has been set.

### SetShortCodeListNil

`func (o *MarkdownFileRequest) SetShortCodeListNil(b bool)`

 SetShortCodeListNil sets the value for ShortCodeList to be an explicit nil

### UnsetShortCodeList
`func (o *MarkdownFileRequest) UnsetShortCodeList()`

UnsetShortCodeList ensures that no value is present for ShortCodeList, not even an explicit nil
### GetFrontMatterList

`func (o *MarkdownFileRequest) GetFrontMatterList() [][]string`

GetFrontMatterList returns the FrontMatterList field if non-nil, zero value otherwise.

### GetFrontMatterListOk

`func (o *MarkdownFileRequest) GetFrontMatterListOk() (*[][]string, bool)`

GetFrontMatterListOk returns a tuple with the FrontMatterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontMatterList

`func (o *MarkdownFileRequest) SetFrontMatterList(v [][]string)`

SetFrontMatterList sets FrontMatterList field to given value.

### HasFrontMatterList

`func (o *MarkdownFileRequest) HasFrontMatterList() bool`

HasFrontMatterList returns a boolean if a field has been set.

### SetFrontMatterListNil

`func (o *MarkdownFileRequest) SetFrontMatterListNil(b bool)`

 SetFrontMatterListNil sets the value for FrontMatterList to be an explicit nil

### UnsetFrontMatterList
`func (o *MarkdownFileRequest) UnsetFrontMatterList()`

UnsetFrontMatterList ensures that no value is present for FrontMatterList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


