# WorksheetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to **[]int32** |  | [optional] 
**Columns** | Pointer to **[]int32** |  | [optional] 
**Ranges** | Pointer to [**[]StringStringTuple**](StringStringTuple.md) |  | [optional] 

## Methods

### NewWorksheetData

`func NewWorksheetData() *WorksheetData`

NewWorksheetData instantiates a new WorksheetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorksheetDataWithDefaults

`func NewWorksheetDataWithDefaults() *WorksheetData`

NewWorksheetDataWithDefaults instantiates a new WorksheetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *WorksheetData) GetRows() []int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WorksheetData) GetRowsOk() (*[]int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WorksheetData) SetRows(v []int32)`

SetRows sets Rows field to given value.

### HasRows

`func (o *WorksheetData) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRowsNil

`func (o *WorksheetData) SetRowsNil(b bool)`

 SetRowsNil sets the value for Rows to be an explicit nil

### UnsetRows
`func (o *WorksheetData) UnsetRows()`

UnsetRows ensures that no value is present for Rows, not even an explicit nil
### GetColumns

`func (o *WorksheetData) GetColumns() []int32`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WorksheetData) GetColumnsOk() (*[]int32, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *WorksheetData) SetColumns(v []int32)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *WorksheetData) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *WorksheetData) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *WorksheetData) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetRanges

`func (o *WorksheetData) GetRanges() []StringStringTuple`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *WorksheetData) GetRangesOk() (*[]StringStringTuple, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *WorksheetData) SetRanges(v []StringStringTuple)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *WorksheetData) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### SetRangesNil

`func (o *WorksheetData) SetRangesNil(b bool)`

 SetRangesNil sets the value for Ranges to be an explicit nil

### UnsetRanges
`func (o *WorksheetData) UnsetRanges()`

UnsetRanges ensures that no value is present for Ranges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


