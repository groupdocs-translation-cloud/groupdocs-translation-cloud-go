# SpreadsheetFileRequest

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
**Format** | **string** | Input file format | [default to "Xlsx"]
**OutputFormat** | **string** | output file format | 
**Worksheets** | Pointer to **[]int32** | List of Worksheets to translate by sequence number (1-based index). If not present, translate all worksheets | [optional] 
**Ranges** | Pointer to [**map[string]WorksheetData**](WorksheetData.md) | Dictionary of ranges in Excel workbooks | [optional] 

## Methods

### NewSpreadsheetFileRequest

`func NewSpreadsheetFileRequest(sourceLanguage string, targetLanguages []string, url string, format string, outputFormat string, ) *SpreadsheetFileRequest`

NewSpreadsheetFileRequest instantiates a new SpreadsheetFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpreadsheetFileRequestWithDefaults

`func NewSpreadsheetFileRequestWithDefaults() *SpreadsheetFileRequest`

NewSpreadsheetFileRequestWithDefaults instantiates a new SpreadsheetFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *SpreadsheetFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *SpreadsheetFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *SpreadsheetFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.


### GetTargetLanguages

`func (o *SpreadsheetFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *SpreadsheetFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *SpreadsheetFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.


### GetOriginalFileName

`func (o *SpreadsheetFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *SpreadsheetFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *SpreadsheetFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *SpreadsheetFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *SpreadsheetFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *SpreadsheetFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *SpreadsheetFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SpreadsheetFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SpreadsheetFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetOrigin

`func (o *SpreadsheetFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SpreadsheetFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SpreadsheetFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *SpreadsheetFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *SpreadsheetFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *SpreadsheetFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetIsNeedAlignment

`func (o *SpreadsheetFileRequest) GetIsNeedAlignment() bool`

GetIsNeedAlignment returns the IsNeedAlignment field if non-nil, zero value otherwise.

### GetIsNeedAlignmentOk

`func (o *SpreadsheetFileRequest) GetIsNeedAlignmentOk() (*bool, bool)`

GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeedAlignment

`func (o *SpreadsheetFileRequest) SetIsNeedAlignment(v bool)`

SetIsNeedAlignment sets IsNeedAlignment field to given value.

### HasIsNeedAlignment

`func (o *SpreadsheetFileRequest) HasIsNeedAlignment() bool`

HasIsNeedAlignment returns a boolean if a field has been set.

### GetTranslationDictionary

`func (o *SpreadsheetFileRequest) GetTranslationDictionary() map[string]string`

GetTranslationDictionary returns the TranslationDictionary field if non-nil, zero value otherwise.

### GetTranslationDictionaryOk

`func (o *SpreadsheetFileRequest) GetTranslationDictionaryOk() (*map[string]string, bool)`

GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationDictionary

`func (o *SpreadsheetFileRequest) SetTranslationDictionary(v map[string]string)`

SetTranslationDictionary sets TranslationDictionary field to given value.

### HasTranslationDictionary

`func (o *SpreadsheetFileRequest) HasTranslationDictionary() bool`

HasTranslationDictionary returns a boolean if a field has been set.

### SetTranslationDictionaryNil

`func (o *SpreadsheetFileRequest) SetTranslationDictionaryNil(b bool)`

 SetTranslationDictionaryNil sets the value for TranslationDictionary to be an explicit nil

### UnsetTranslationDictionary
`func (o *SpreadsheetFileRequest) UnsetTranslationDictionary()`

UnsetTranslationDictionary ensures that no value is present for TranslationDictionary, not even an explicit nil
### GetSavingMode

`func (o *SpreadsheetFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *SpreadsheetFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *SpreadsheetFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *SpreadsheetFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetFormat

`func (o *SpreadsheetFileRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SpreadsheetFileRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SpreadsheetFileRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetOutputFormat

`func (o *SpreadsheetFileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *SpreadsheetFileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *SpreadsheetFileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.


### GetWorksheets

`func (o *SpreadsheetFileRequest) GetWorksheets() []int32`

GetWorksheets returns the Worksheets field if non-nil, zero value otherwise.

### GetWorksheetsOk

`func (o *SpreadsheetFileRequest) GetWorksheetsOk() (*[]int32, bool)`

GetWorksheetsOk returns a tuple with the Worksheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheets

`func (o *SpreadsheetFileRequest) SetWorksheets(v []int32)`

SetWorksheets sets Worksheets field to given value.

### HasWorksheets

`func (o *SpreadsheetFileRequest) HasWorksheets() bool`

HasWorksheets returns a boolean if a field has been set.

### SetWorksheetsNil

`func (o *SpreadsheetFileRequest) SetWorksheetsNil(b bool)`

 SetWorksheetsNil sets the value for Worksheets to be an explicit nil

### UnsetWorksheets
`func (o *SpreadsheetFileRequest) UnsetWorksheets()`

UnsetWorksheets ensures that no value is present for Worksheets, not even an explicit nil
### GetRanges

`func (o *SpreadsheetFileRequest) GetRanges() map[string]WorksheetData`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *SpreadsheetFileRequest) GetRangesOk() (*map[string]WorksheetData, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *SpreadsheetFileRequest) SetRanges(v map[string]WorksheetData)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *SpreadsheetFileRequest) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### SetRangesNil

`func (o *SpreadsheetFileRequest) SetRangesNil(b bool)`

 SetRangesNil sets the value for Ranges to be an explicit nil

### UnsetRanges
`func (o *SpreadsheetFileRequest) UnsetRanges()`

UnsetRanges ensures that no value is present for Ranges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


