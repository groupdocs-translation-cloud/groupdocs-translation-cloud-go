# FileRequest

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
**Format** | Pointer to **string** | Input file format | [optional] [default to "Unknown"]
**OutputFormat** | Pointer to **string** | output file format | [optional] 
**Masters** | Pointer to **bool** | If translate master slides | [optional] [default to false]
**Formatting** | Pointer to **bool** | If document&#39;s formatting should be preserved, default true | [optional] [default to true]
**Route** | Pointer to **NullableString** | Endpoint route | [optional] 
**Separator** | Pointer to **NullableString** | Separator in files | [optional] 
**Elements** | Pointer to **[]int32** | List of slides to translate | [optional] 
**Ranges** | Pointer to [**map[string]WorksheetData**](WorksheetData.md) | Dictionary of ranges in Excel workbooks | [optional] 
**ShortCodeList** | Pointer to **map[string][]string** | Dictionary of short code names and parameters names to translate | [optional] 
**FrontMatterList** | Pointer to **[][]string** | Dictionary where key is zero-based front matter index and value is list of lists of front matter paths | [optional] 

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

### GetFile

`func (o *FileRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *FileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *FileRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *FileRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
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

### SetUrlNil

`func (o *FileRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FileRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
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
### GetShortCodeList

`func (o *FileRequest) GetShortCodeList() map[string][]string`

GetShortCodeList returns the ShortCodeList field if non-nil, zero value otherwise.

### GetShortCodeListOk

`func (o *FileRequest) GetShortCodeListOk() (*map[string][]string, bool)`

GetShortCodeListOk returns a tuple with the ShortCodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCodeList

`func (o *FileRequest) SetShortCodeList(v map[string][]string)`

SetShortCodeList sets ShortCodeList field to given value.

### HasShortCodeList

`func (o *FileRequest) HasShortCodeList() bool`

HasShortCodeList returns a boolean if a field has been set.

### SetShortCodeListNil

`func (o *FileRequest) SetShortCodeListNil(b bool)`

 SetShortCodeListNil sets the value for ShortCodeList to be an explicit nil

### UnsetShortCodeList
`func (o *FileRequest) UnsetShortCodeList()`

UnsetShortCodeList ensures that no value is present for ShortCodeList, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


