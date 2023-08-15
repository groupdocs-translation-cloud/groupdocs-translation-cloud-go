# AutoPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | Input file format | [default to "Unknown"]
**OutputFormat** | **string** | output file format | 
**Masters** | Pointer to **bool** | If translate master slides | [optional] [default to false]
**Formatting** | Pointer to **bool** | If document&#39;s formatting should be preserved, default true | [optional] [default to true]
**Route** | Pointer to **string** | Endpoint route | [optional] 
**Separator** | Pointer to **string** | Separator in files | [optional] 
**Elements** | Pointer to **[]int32** | List of slides to translate | [optional] 
**Ranges** | Pointer to [**map[string]WorksheetData**](WorksheetData.md) | Dictionary of ranges in Excel workbooks | [optional] 
**ShortCodeList** | Pointer to [**map[string][][]string**](array.md) | Dictionary of short code names and parameters names to translate | [optional] 
**FrontMatterList** | Pointer to **[][]string** | Dictionary where key is zero-based front matter index and value is list of lists of front matter paths | [optional] 
**SourceLanguage** | **string** | Language of original file | [default to "en"]
**TargetLanguages** | **[]string** | List of target languages | 
**File** | Pointer to **string** | File as byte array | [optional] 
**OriginalFileName** | Pointer to **string** | Type in the file name. If null will be as request ID. | [optional] 
**Url** | Pointer to **string** | Link to file for translation. Ignore, if \&quot;file\&quot; property not null | [optional] 
**Origin** | Pointer to **string** | Url or name of application using this SDK. Not required. | [optional] 
**SavingMode** | Pointer to **string** | Toggle file saving mode for storage.  Is Files by default. | [optional] 

## Methods

### NewAutoPostRequest

`func NewAutoPostRequest(format string, outputFormat string, sourceLanguage string, targetLanguages []string, ) *AutoPostRequest`

NewAutoPostRequest instantiates a new AutoPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPostRequestWithDefaults

`func NewAutoPostRequestWithDefaults() *AutoPostRequest`

NewAutoPostRequestWithDefaults instantiates a new AutoPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *AutoPostRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AutoPostRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AutoPostRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetOutputFormat

`func (o *AutoPostRequest) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *AutoPostRequest) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *AutoPostRequest) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.


### GetMasters

`func (o *AutoPostRequest) GetMasters() bool`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *AutoPostRequest) GetMastersOk() (*bool, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *AutoPostRequest) SetMasters(v bool)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *AutoPostRequest) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetFormatting

`func (o *AutoPostRequest) GetFormatting() bool`

GetFormatting returns the Formatting field if non-nil, zero value otherwise.

### GetFormattingOk

`func (o *AutoPostRequest) GetFormattingOk() (*bool, bool)`

GetFormattingOk returns a tuple with the Formatting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatting

`func (o *AutoPostRequest) SetFormatting(v bool)`

SetFormatting sets Formatting field to given value.

### HasFormatting

`func (o *AutoPostRequest) HasFormatting() bool`

HasFormatting returns a boolean if a field has been set.

### GetRoute

`func (o *AutoPostRequest) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *AutoPostRequest) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *AutoPostRequest) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *AutoPostRequest) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetSeparator

`func (o *AutoPostRequest) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *AutoPostRequest) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *AutoPostRequest) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *AutoPostRequest) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetElements

`func (o *AutoPostRequest) GetElements() []int32`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *AutoPostRequest) GetElementsOk() (*[]int32, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *AutoPostRequest) SetElements(v []int32)`

SetElements sets Elements field to given value.

### HasElements

`func (o *AutoPostRequest) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetRanges

`func (o *AutoPostRequest) GetRanges() map[string]WorksheetData`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AutoPostRequest) GetRangesOk() (*map[string]WorksheetData, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AutoPostRequest) SetRanges(v map[string]WorksheetData)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AutoPostRequest) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetShortCodeList

`func (o *AutoPostRequest) GetShortCodeList() map[string][][]string`

GetShortCodeList returns the ShortCodeList field if non-nil, zero value otherwise.

### GetShortCodeListOk

`func (o *AutoPostRequest) GetShortCodeListOk() (*map[string][][]string, bool)`

GetShortCodeListOk returns a tuple with the ShortCodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCodeList

`func (o *AutoPostRequest) SetShortCodeList(v map[string][][]string)`

SetShortCodeList sets ShortCodeList field to given value.

### HasShortCodeList

`func (o *AutoPostRequest) HasShortCodeList() bool`

HasShortCodeList returns a boolean if a field has been set.

### GetFrontMatterList

`func (o *AutoPostRequest) GetFrontMatterList() [][]string`

GetFrontMatterList returns the FrontMatterList field if non-nil, zero value otherwise.

### GetFrontMatterListOk

`func (o *AutoPostRequest) GetFrontMatterListOk() (*[][]string, bool)`

GetFrontMatterListOk returns a tuple with the FrontMatterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontMatterList

`func (o *AutoPostRequest) SetFrontMatterList(v [][]string)`

SetFrontMatterList sets FrontMatterList field to given value.

### HasFrontMatterList

`func (o *AutoPostRequest) HasFrontMatterList() bool`

HasFrontMatterList returns a boolean if a field has been set.

### GetSourceLanguage

`func (o *AutoPostRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *AutoPostRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *AutoPostRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.


### GetTargetLanguages

`func (o *AutoPostRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *AutoPostRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *AutoPostRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.


### GetFile

`func (o *AutoPostRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AutoPostRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AutoPostRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *AutoPostRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *AutoPostRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *AutoPostRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *AutoPostRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *AutoPostRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### GetUrl

`func (o *AutoPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutoPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutoPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutoPostRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOrigin

`func (o *AutoPostRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AutoPostRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AutoPostRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AutoPostRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetSavingMode

`func (o *AutoPostRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *AutoPostRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *AutoPostRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *AutoPostRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


