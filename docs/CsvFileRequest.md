# CsvFileRequest

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
**Format** | Pointer to **string** | Input file format | [optional] [default to "Csv"]
**OutputFormat** | **string** | output file format | 
**Separator** | Pointer to **NullableString** | Separator in files | [optional] 

## Methods

### NewCsvFileRequest

`func NewCsvFileRequest(sourceLanguage string, targetLanguages []string, url string, outputFormat string, ) *CsvFileRequest`

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
### GetIsNeedAlignment

`func (o *CsvFileRequest) GetIsNeedAlignment() bool`

GetIsNeedAlignment returns the IsNeedAlignment field if non-nil, zero value otherwise.

### GetIsNeedAlignmentOk

`func (o *CsvFileRequest) GetIsNeedAlignmentOk() (*bool, bool)`

GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeedAlignment

`func (o *CsvFileRequest) SetIsNeedAlignment(v bool)`

SetIsNeedAlignment sets IsNeedAlignment field to given value.

### HasIsNeedAlignment

`func (o *CsvFileRequest) HasIsNeedAlignment() bool`

HasIsNeedAlignment returns a boolean if a field has been set.

### GetTranslationDictionary

`func (o *CsvFileRequest) GetTranslationDictionary() map[string]string`

GetTranslationDictionary returns the TranslationDictionary field if non-nil, zero value otherwise.

### GetTranslationDictionaryOk

`func (o *CsvFileRequest) GetTranslationDictionaryOk() (*map[string]string, bool)`

GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationDictionary

`func (o *CsvFileRequest) SetTranslationDictionary(v map[string]string)`

SetTranslationDictionary sets TranslationDictionary field to given value.

### HasTranslationDictionary

`func (o *CsvFileRequest) HasTranslationDictionary() bool`

HasTranslationDictionary returns a boolean if a field has been set.

### SetTranslationDictionaryNil

`func (o *CsvFileRequest) SetTranslationDictionaryNil(b bool)`

 SetTranslationDictionaryNil sets the value for TranslationDictionary to be an explicit nil

### UnsetTranslationDictionary
`func (o *CsvFileRequest) UnsetTranslationDictionary()`

UnsetTranslationDictionary ensures that no value is present for TranslationDictionary, not even an explicit nil
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

### GetOutputFormat

`func (o *CsvFileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *CsvFileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *CsvFileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.


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


