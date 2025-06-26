# ImageToFileRequest

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
**Format** | Pointer to **string** | Original file format | [optional] [default to "Unknown"]
**Ocrformat** | **string** | File format after recognition | [default to "Pdf"]
**OutputFormat** | **string** | output file format | 
**RotationAngle** | Pointer to **int32** | Left to write angle to rotate scanned image / pdf | [optional] 
**Formatting** | Pointer to **bool** | If document&#39;s formatting should be preserved, default true | [optional] [default to true]
**Route** | Pointer to **NullableString** | endpoints route | [optional] 
**Pages** | Pointer to **[]int32** | List of pages to translate for scanned pdf | [optional] 

## Methods

### NewImageToFileRequest

`func NewImageToFileRequest(sourceLanguage string, targetLanguages []string, url string, ocrformat string, outputFormat string, ) *ImageToFileRequest`

NewImageToFileRequest instantiates a new ImageToFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageToFileRequestWithDefaults

`func NewImageToFileRequestWithDefaults() *ImageToFileRequest`

NewImageToFileRequestWithDefaults instantiates a new ImageToFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *ImageToFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *ImageToFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *ImageToFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.


### GetTargetLanguages

`func (o *ImageToFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *ImageToFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *ImageToFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.


### GetOriginalFileName

`func (o *ImageToFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *ImageToFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *ImageToFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *ImageToFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *ImageToFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *ImageToFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *ImageToFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageToFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageToFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetOrigin

`func (o *ImageToFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ImageToFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ImageToFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ImageToFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ImageToFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ImageToFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetIsNeedAlignment

`func (o *ImageToFileRequest) GetIsNeedAlignment() bool`

GetIsNeedAlignment returns the IsNeedAlignment field if non-nil, zero value otherwise.

### GetIsNeedAlignmentOk

`func (o *ImageToFileRequest) GetIsNeedAlignmentOk() (*bool, bool)`

GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeedAlignment

`func (o *ImageToFileRequest) SetIsNeedAlignment(v bool)`

SetIsNeedAlignment sets IsNeedAlignment field to given value.

### HasIsNeedAlignment

`func (o *ImageToFileRequest) HasIsNeedAlignment() bool`

HasIsNeedAlignment returns a boolean if a field has been set.

### GetTranslationDictionary

`func (o *ImageToFileRequest) GetTranslationDictionary() map[string]string`

GetTranslationDictionary returns the TranslationDictionary field if non-nil, zero value otherwise.

### GetTranslationDictionaryOk

`func (o *ImageToFileRequest) GetTranslationDictionaryOk() (*map[string]string, bool)`

GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationDictionary

`func (o *ImageToFileRequest) SetTranslationDictionary(v map[string]string)`

SetTranslationDictionary sets TranslationDictionary field to given value.

### HasTranslationDictionary

`func (o *ImageToFileRequest) HasTranslationDictionary() bool`

HasTranslationDictionary returns a boolean if a field has been set.

### SetTranslationDictionaryNil

`func (o *ImageToFileRequest) SetTranslationDictionaryNil(b bool)`

 SetTranslationDictionaryNil sets the value for TranslationDictionary to be an explicit nil

### UnsetTranslationDictionary
`func (o *ImageToFileRequest) UnsetTranslationDictionary()`

UnsetTranslationDictionary ensures that no value is present for TranslationDictionary, not even an explicit nil
### GetSavingMode

`func (o *ImageToFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *ImageToFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *ImageToFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *ImageToFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetFormat

`func (o *ImageToFileRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ImageToFileRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ImageToFileRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ImageToFileRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOcrformat

`func (o *ImageToFileRequest) GetOcrformat() string`

GetOcrformat returns the Ocrformat field if non-nil, zero value otherwise.

### GetOcrformatOk

`func (o *ImageToFileRequest) GetOcrformatOk() (*string, bool)`

GetOcrformatOk returns a tuple with the Ocrformat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrformat

`func (o *ImageToFileRequest) SetOcrformat(v string)`

SetOcrformat sets Ocrformat field to given value.


### GetOutputFormat

`func (o *ImageToFileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *ImageToFileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *ImageToFileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.


### GetRotationAngle

`func (o *ImageToFileRequest) GetRotationAngle() int32`

GetRotationAngle returns the RotationAngle field if non-nil, zero value otherwise.

### GetRotationAngleOk

`func (o *ImageToFileRequest) GetRotationAngleOk() (*int32, bool)`

GetRotationAngleOk returns a tuple with the RotationAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationAngle

`func (o *ImageToFileRequest) SetRotationAngle(v int32)`

SetRotationAngle sets RotationAngle field to given value.

### HasRotationAngle

`func (o *ImageToFileRequest) HasRotationAngle() bool`

HasRotationAngle returns a boolean if a field has been set.

### GetFormatting

`func (o *ImageToFileRequest) GetFormatting() bool`

GetFormatting returns the Formatting field if non-nil, zero value otherwise.

### GetFormattingOk

`func (o *ImageToFileRequest) GetFormattingOk() (*bool, bool)`

GetFormattingOk returns a tuple with the Formatting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatting

`func (o *ImageToFileRequest) SetFormatting(v bool)`

SetFormatting sets Formatting field to given value.

### HasFormatting

`func (o *ImageToFileRequest) HasFormatting() bool`

HasFormatting returns a boolean if a field has been set.

### GetRoute

`func (o *ImageToFileRequest) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *ImageToFileRequest) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *ImageToFileRequest) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *ImageToFileRequest) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### SetRouteNil

`func (o *ImageToFileRequest) SetRouteNil(b bool)`

 SetRouteNil sets the value for Route to be an explicit nil

### UnsetRoute
`func (o *ImageToFileRequest) UnsetRoute()`

UnsetRoute ensures that no value is present for Route, not even an explicit nil
### GetPages

`func (o *ImageToFileRequest) GetPages() []int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ImageToFileRequest) GetPagesOk() (*[]int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ImageToFileRequest) SetPages(v []int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ImageToFileRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPagesNil

`func (o *ImageToFileRequest) SetPagesNil(b bool)`

 SetPagesNil sets the value for Pages to be an explicit nil

### UnsetPages
`func (o *ImageToFileRequest) UnsetPages()`

UnsetPages ensures that no value is present for Pages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


