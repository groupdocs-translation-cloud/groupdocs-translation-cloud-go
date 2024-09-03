# HealthCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]HealthCheckEntity**](HealthCheckEntity.md) |  | [optional] 

## Methods

### NewHealthCheckStatus

`func NewHealthCheckStatus() *HealthCheckStatus`

NewHealthCheckStatus instantiates a new HealthCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckStatusWithDefaults

`func NewHealthCheckStatusWithDefaults() *HealthCheckStatus`

NewHealthCheckStatusWithDefaults instantiates a new HealthCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *HealthCheckStatus) GetItems() []HealthCheckEntity`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HealthCheckStatus) GetItemsOk() (*[]HealthCheckEntity, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HealthCheckStatus) SetItems(v []HealthCheckEntity)`

SetItems sets Items field to given value.

### HasItems

`func (o *HealthCheckStatus) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *HealthCheckStatus) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *HealthCheckStatus) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


