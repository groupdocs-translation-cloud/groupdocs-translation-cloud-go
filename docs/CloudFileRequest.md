# CloudFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **NullableString** | \&quot;doc(x|m)\&quot; if Word document, \&quot;xls(x|m)\&quot; if Excel workbook | [optional] 
**OutFormat** | Pointer to **NullableString** | output file format | [optional] 
**RequestId** | Pointer to **NullableString** | id of request | [optional] 
**Ids** | Pointer to **[]int32** | Language pairs to translate | [optional] 
**Url** | Pointer to **NullableString** | Link to file for translation | [optional] 
**Size** | Pointer to **int64** | File size | [optional] 
**Masters** | Pointer to **bool** | If translate master slides | [optional] 
**Formatting** | Pointer to **bool** | If document&#39;s formatting should be preserved, default true | [optional] 
**Origin** | Pointer to **NullableString** | for analysis only | [optional] 
**Elements** | Pointer to **[]int32** | List of slides to translate | [optional] 
**Ranges** | Pointer to [**map[string]WorksheetData**](WorksheetData.md) | Dictionary of ranges in Excel workbooks | [optional] 
**ShortCodeDict** | Pointer to **map[string][]string** | Dictiory of short code names and parameters names to translate | [optional] 
**FrontMatterList** | Pointer to **[][]string** | Dictionary where key is zero-based front matter index and value is list of lists of front matter paths | [optional] 
**OriginalFileName** | Pointer to **NullableString** | Original name of file | [optional] 
**Separator** | Pointer to **NullableString** | Separator in files | [optional] 
**IsPaid** | Pointer to **bool** | Set true if paid user | [optional] 
**SavingMode** | Pointer to **string** | Toggle files saving mode | [optional] 
**Details** | Pointer to **map[string]string** | Details of the requests. Using for e2e tracking | [optional] 

## Methods

### NewCloudFileRequest

`func NewCloudFileRequest() *CloudFileRequest`

NewCloudFileRequest instantiates a new CloudFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFileRequestWithDefaults

`func NewCloudFileRequestWithDefaults() *CloudFileRequest`

NewCloudFileRequestWithDefaults instantiates a new CloudFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *CloudFileRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CloudFileRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CloudFileRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CloudFileRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *CloudFileRequest) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *CloudFileRequest) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetOutFormat

`func (o *CloudFileRequest) GetOutFormat() string`

GetOutFormat returns the OutFormat field if non-nil, zero value otherwise.

### GetOutFormatOk

`func (o *CloudFileRequest) GetOutFormatOk() (*string, bool)`

GetOutFormatOk returns a tuple with the OutFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutFormat

`func (o *CloudFileRequest) SetOutFormat(v string)`

SetOutFormat sets OutFormat field to given value.

### HasOutFormat

`func (o *CloudFileRequest) HasOutFormat() bool`

HasOutFormat returns a boolean if a field has been set.

### SetOutFormatNil

`func (o *CloudFileRequest) SetOutFormatNil(b bool)`

 SetOutFormatNil sets the value for OutFormat to be an explicit nil

### UnsetOutFormat
`func (o *CloudFileRequest) UnsetOutFormat()`

UnsetOutFormat ensures that no value is present for OutFormat, not even an explicit nil
### GetRequestId

`func (o *CloudFileRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CloudFileRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CloudFileRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CloudFileRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *CloudFileRequest) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *CloudFileRequest) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetIds

`func (o *CloudFileRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *CloudFileRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *CloudFileRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *CloudFileRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *CloudFileRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *CloudFileRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetUrl

`func (o *CloudFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudFileRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CloudFileRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CloudFileRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSize

`func (o *CloudFileRequest) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudFileRequest) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudFileRequest) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudFileRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMasters

`func (o *CloudFileRequest) GetMasters() bool`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *CloudFileRequest) GetMastersOk() (*bool, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *CloudFileRequest) SetMasters(v bool)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *CloudFileRequest) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetFormatting

`func (o *CloudFileRequest) GetFormatting() bool`

GetFormatting returns the Formatting field if non-nil, zero value otherwise.

### GetFormattingOk

`func (o *CloudFileRequest) GetFormattingOk() (*bool, bool)`

GetFormattingOk returns a tuple with the Formatting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatting

`func (o *CloudFileRequest) SetFormatting(v bool)`

SetFormatting sets Formatting field to given value.

### HasFormatting

`func (o *CloudFileRequest) HasFormatting() bool`

HasFormatting returns a boolean if a field has been set.

### GetOrigin

`func (o *CloudFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CloudFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CloudFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CloudFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *CloudFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *CloudFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetElements

`func (o *CloudFileRequest) GetElements() []int32`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CloudFileRequest) GetElementsOk() (*[]int32, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CloudFileRequest) SetElements(v []int32)`

SetElements sets Elements field to given value.

### HasElements

`func (o *CloudFileRequest) HasElements() bool`

HasElements returns a boolean if a field has been set.

### SetElementsNil

`func (o *CloudFileRequest) SetElementsNil(b bool)`

 SetElementsNil sets the value for Elements to be an explicit nil

### UnsetElements
`func (o *CloudFileRequest) UnsetElements()`

UnsetElements ensures that no value is present for Elements, not even an explicit nil
### GetRanges

`func (o *CloudFileRequest) GetRanges() map[string]WorksheetData`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *CloudFileRequest) GetRangesOk() (*map[string]WorksheetData, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *CloudFileRequest) SetRanges(v map[string]WorksheetData)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *CloudFileRequest) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### SetRangesNil

`func (o *CloudFileRequest) SetRangesNil(b bool)`

 SetRangesNil sets the value for Ranges to be an explicit nil

### UnsetRanges
`func (o *CloudFileRequest) UnsetRanges()`

UnsetRanges ensures that no value is present for Ranges, not even an explicit nil
### GetShortCodeDict

`func (o *CloudFileRequest) GetShortCodeDict() map[string][]string`

GetShortCodeDict returns the ShortCodeDict field if non-nil, zero value otherwise.

### GetShortCodeDictOk

`func (o *CloudFileRequest) GetShortCodeDictOk() (*map[string][]string, bool)`

GetShortCodeDictOk returns a tuple with the ShortCodeDict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCodeDict

`func (o *CloudFileRequest) SetShortCodeDict(v map[string][]string)`

SetShortCodeDict sets ShortCodeDict field to given value.

### HasShortCodeDict

`func (o *CloudFileRequest) HasShortCodeDict() bool`

HasShortCodeDict returns a boolean if a field has been set.

### SetShortCodeDictNil

`func (o *CloudFileRequest) SetShortCodeDictNil(b bool)`

 SetShortCodeDictNil sets the value for ShortCodeDict to be an explicit nil

### UnsetShortCodeDict
`func (o *CloudFileRequest) UnsetShortCodeDict()`

UnsetShortCodeDict ensures that no value is present for ShortCodeDict, not even an explicit nil
### GetFrontMatterList

`func (o *CloudFileRequest) GetFrontMatterList() [][]string`

GetFrontMatterList returns the FrontMatterList field if non-nil, zero value otherwise.

### GetFrontMatterListOk

`func (o *CloudFileRequest) GetFrontMatterListOk() (*[][]string, bool)`

GetFrontMatterListOk returns a tuple with the FrontMatterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontMatterList

`func (o *CloudFileRequest) SetFrontMatterList(v [][]string)`

SetFrontMatterList sets FrontMatterList field to given value.

### HasFrontMatterList

`func (o *CloudFileRequest) HasFrontMatterList() bool`

HasFrontMatterList returns a boolean if a field has been set.

### SetFrontMatterListNil

`func (o *CloudFileRequest) SetFrontMatterListNil(b bool)`

 SetFrontMatterListNil sets the value for FrontMatterList to be an explicit nil

### UnsetFrontMatterList
`func (o *CloudFileRequest) UnsetFrontMatterList()`

UnsetFrontMatterList ensures that no value is present for FrontMatterList, not even an explicit nil
### GetOriginalFileName

`func (o *CloudFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *CloudFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *CloudFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *CloudFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *CloudFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *CloudFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetSeparator

`func (o *CloudFileRequest) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *CloudFileRequest) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *CloudFileRequest) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *CloudFileRequest) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *CloudFileRequest) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *CloudFileRequest) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
### GetIsPaid

`func (o *CloudFileRequest) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *CloudFileRequest) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *CloudFileRequest) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *CloudFileRequest) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetSavingMode

`func (o *CloudFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *CloudFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *CloudFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *CloudFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetDetails

`func (o *CloudFileRequest) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CloudFileRequest) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CloudFileRequest) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CloudFileRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CloudFileRequest) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CloudFileRequest) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


