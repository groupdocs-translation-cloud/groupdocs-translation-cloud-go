# FileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguage** | Pointer to **string** | Language of original file | [optional] [default to "en"]
**TargetLanguages** | Pointer to **[]string** | List of target languages | [optional] 
**OriginalFileName** | Pointer to **NullableString** | Type in the file name. If null will be as request ID. | [optional] 
**Url** | Pointer to **string** | Link to file for translation. Ignore, if \&quot;file\&quot; property not null | [optional] 
**Origin** | Pointer to **NullableString** | Url or name of the application using this SDK. Not required. | [optional] 
**IsNeedAlignment** | Pointer to **bool** | Do result formating like the source. This option needs more expensive requests. | [optional] 
**TranslationDictionary** | Pointer to **map[string]string** | Set a specific translation between source and target words. | [optional] 
**SavingMode** | Pointer to **string** | Toggle file saving mode for storage.  Is Files by default. | [optional] 
**Format** | Pointer to **string** | Input file format | [optional] [default to "Unknown"]
**OutputFormat** | Pointer to **string** | output file format | [optional] 
**Masters** | Pointer to **bool** | If translate master slides | [optional] [default to false]
**Formatting** | Pointer to **bool** | If document&#39;s formatting should be preserved, default true | [optional] [default to true]
**Route** | Pointer to **NullableString** | Endpoint route | [optional] 
**Separator** | Pointer to **NullableString** | Separator in files | [optional] 
**Elements** | Pointer to **[]int32** | List of slides to translate (1-based index). If not present, translate all elements (page, slide, worksheet) | [optional] 
**Ranges** | Pointer to [**map[string]WorksheetData**](WorksheetData.md) | Dictionary of ranges in Excel workbooks | [optional] 
**Shortcodedict** | Pointer to **map[string][]string** | Dictionary of short code names and parameters names to translate | [optional] 
**FrontMatterList** | Pointer to **[][]string** | Dictionary where key is zero-based front matter index and value is list of lists of front matter paths | [optional] 
**IgnoreList** | Pointer to **[]string** | List of elements for Xml, Json and Yaml formats. Determines which items should be blacklisted or whitelisted for processing depending on GroupDocs.Translation.ApiGateway.DTO.FileRequest.IsWhiteList. | [optional] 
**IsWhiteList** | Pointer to **bool** | Determines to which list the items in GroupDocs.Translation.ApiGateway.DTO.FileRequest.IgnoreList should be allocated. The default is the black list. | [optional] 

## Methods

### NewFileRequest

`func NewFileRequest() *FileRequest`

NewFileRequest instantiates a new FileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRequestWithDefaults

`func NewFileRequestWithDefaults() *FileRequest`

NewFileRequestWithDefaults instantiates a new FileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *FileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *FileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *FileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.

### HasSourceLanguage

`func (o *FileRequest) HasSourceLanguage() bool`

HasSourceLanguage returns a boolean if a field has been set.

### GetTargetLanguages

`func (o *FileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *FileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *FileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.

### HasTargetLanguages

`func (o *FileRequest) HasTargetLanguages() bool`

HasTargetLanguages returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *FileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *FileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *FileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *FileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *FileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *FileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *FileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOrigin

`func (o *FileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *FileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *FileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *FileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *FileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *FileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetIsNeedAlignment

`func (o *FileRequest) GetIsNeedAlignment() bool`

GetIsNeedAlignment returns the IsNeedAlignment field if non-nil, zero value otherwise.

### GetIsNeedAlignmentOk

`func (o *FileRequest) GetIsNeedAlignmentOk() (*bool, bool)`

GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeedAlignment

`func (o *FileRequest) SetIsNeedAlignment(v bool)`

SetIsNeedAlignment sets IsNeedAlignment field to given value.

### HasIsNeedAlignment

`func (o *FileRequest) HasIsNeedAlignment() bool`

HasIsNeedAlignment returns a boolean if a field has been set.

### GetTranslationDictionary

`func (o *FileRequest) GetTranslationDictionary() map[string]string`

GetTranslationDictionary returns the TranslationDictionary field if non-nil, zero value otherwise.

### GetTranslationDictionaryOk

`func (o *FileRequest) GetTranslationDictionaryOk() (*map[string]string, bool)`

GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationDictionary

`func (o *FileRequest) SetTranslationDictionary(v map[string]string)`

SetTranslationDictionary sets TranslationDictionary field to given value.

### HasTranslationDictionary

`func (o *FileRequest) HasTranslationDictionary() bool`

HasTranslationDictionary returns a boolean if a field has been set.

### SetTranslationDictionaryNil

`func (o *FileRequest) SetTranslationDictionaryNil(b bool)`

 SetTranslationDictionaryNil sets the value for TranslationDictionary to be an explicit nil

### UnsetTranslationDictionary
`func (o *FileRequest) UnsetTranslationDictionary()`

UnsetTranslationDictionary ensures that no value is present for TranslationDictionary, not even an explicit nil
### GetSavingMode

`func (o *FileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *FileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *FileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *FileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetFormat

`func (o *FileRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *FileRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *FileRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *FileRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOutputFormat

`func (o *FileRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *FileRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *FileRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *FileRequest) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetMasters

`func (o *FileRequest) GetMasters() bool`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *FileRequest) GetMastersOk() (*bool, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *FileRequest) SetMasters(v bool)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *FileRequest) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetFormatting

`func (o *FileRequest) GetFormatting() bool`

GetFormatting returns the Formatting field if non-nil, zero value otherwise.

### GetFormattingOk

`func (o *FileRequest) GetFormattingOk() (*bool, bool)`

GetFormattingOk returns a tuple with the Formatting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatting

`func (o *FileRequest) SetFormatting(v bool)`

SetFormatting sets Formatting field to given value.

### HasFormatting

`func (o *FileRequest) HasFormatting() bool`

HasFormatting returns a boolean if a field has been set.

### GetRoute

`func (o *FileRequest) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *FileRequest) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *FileRequest) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *FileRequest) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### SetRouteNil

`func (o *FileRequest) SetRouteNil(b bool)`

 SetRouteNil sets the value for Route to be an explicit nil

### UnsetRoute
`func (o *FileRequest) UnsetRoute()`

UnsetRoute ensures that no value is present for Route, not even an explicit nil
### GetSeparator

`func (o *FileRequest) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *FileRequest) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *FileRequest) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *FileRequest) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *FileRequest) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *FileRequest) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
### GetElements

`func (o *FileRequest) GetElements() []int32`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *FileRequest) GetElementsOk() (*[]int32, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *FileRequest) SetElements(v []int32)`

SetElements sets Elements field to given value.

### HasElements

`func (o *FileRequest) HasElements() bool`

HasElements returns a boolean if a field has been set.

### SetElementsNil

`func (o *FileRequest) SetElementsNil(b bool)`

 SetElementsNil sets the value for Elements to be an explicit nil

### UnsetElements
`func (o *FileRequest) UnsetElements()`

UnsetElements ensures that no value is present for Elements, not even an explicit nil
### GetRanges

`func (o *FileRequest) GetRanges() map[string]WorksheetData`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *FileRequest) GetRangesOk() (*map[string]WorksheetData, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *FileRequest) SetRanges(v map[string]WorksheetData)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *FileRequest) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### SetRangesNil

`func (o *FileRequest) SetRangesNil(b bool)`

 SetRangesNil sets the value for Ranges to be an explicit nil

### UnsetRanges
`func (o *FileRequest) UnsetRanges()`

UnsetRanges ensures that no value is present for Ranges, not even an explicit nil
### GetShortcodedict

`func (o *FileRequest) GetShortcodedict() map[string][]string`

GetShortcodedict returns the Shortcodedict field if non-nil, zero value otherwise.

### GetShortcodedictOk

`func (o *FileRequest) GetShortcodedictOk() (*map[string][]string, bool)`

GetShortcodedictOk returns a tuple with the Shortcodedict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcodedict

`func (o *FileRequest) SetShortcodedict(v map[string][]string)`

SetShortcodedict sets Shortcodedict field to given value.

### HasShortcodedict

`func (o *FileRequest) HasShortcodedict() bool`

HasShortcodedict returns a boolean if a field has been set.

### SetShortcodedictNil

`func (o *FileRequest) SetShortcodedictNil(b bool)`

 SetShortcodedictNil sets the value for Shortcodedict to be an explicit nil

### UnsetShortcodedict
`func (o *FileRequest) UnsetShortcodedict()`

UnsetShortcodedict ensures that no value is present for Shortcodedict, not even an explicit nil
### GetFrontMatterList

`func (o *FileRequest) GetFrontMatterList() [][]string`

GetFrontMatterList returns the FrontMatterList field if non-nil, zero value otherwise.

### GetFrontMatterListOk

`func (o *FileRequest) GetFrontMatterListOk() (*[][]string, bool)`

GetFrontMatterListOk returns a tuple with the FrontMatterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontMatterList

`func (o *FileRequest) SetFrontMatterList(v [][]string)`

SetFrontMatterList sets FrontMatterList field to given value.

### HasFrontMatterList

`func (o *FileRequest) HasFrontMatterList() bool`

HasFrontMatterList returns a boolean if a field has been set.

### SetFrontMatterListNil

`func (o *FileRequest) SetFrontMatterListNil(b bool)`

 SetFrontMatterListNil sets the value for FrontMatterList to be an explicit nil

### UnsetFrontMatterList
`func (o *FileRequest) UnsetFrontMatterList()`

UnsetFrontMatterList ensures that no value is present for FrontMatterList, not even an explicit nil
### GetIgnoreList

`func (o *FileRequest) GetIgnoreList() []string`

GetIgnoreList returns the IgnoreList field if non-nil, zero value otherwise.

### GetIgnoreListOk

`func (o *FileRequest) GetIgnoreListOk() (*[]string, bool)`

GetIgnoreListOk returns a tuple with the IgnoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreList

`func (o *FileRequest) SetIgnoreList(v []string)`

SetIgnoreList sets IgnoreList field to given value.

### HasIgnoreList

`func (o *FileRequest) HasIgnoreList() bool`

HasIgnoreList returns a boolean if a field has been set.

### SetIgnoreListNil

`func (o *FileRequest) SetIgnoreListNil(b bool)`

 SetIgnoreListNil sets the value for IgnoreList to be an explicit nil

### UnsetIgnoreList
`func (o *FileRequest) UnsetIgnoreList()`

UnsetIgnoreList ensures that no value is present for IgnoreList, not even an explicit nil
### GetIsWhiteList

`func (o *FileRequest) GetIsWhiteList() bool`

GetIsWhiteList returns the IsWhiteList field if non-nil, zero value otherwise.

### GetIsWhiteListOk

`func (o *FileRequest) GetIsWhiteListOk() (*bool, bool)`

GetIsWhiteListOk returns a tuple with the IsWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteList

`func (o *FileRequest) SetIsWhiteList(v bool)`

SetIsWhiteList sets IsWhiteList field to given value.

### HasIsWhiteList

`func (o *FileRequest) HasIsWhiteList() bool`

HasIsWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


