# HealthCheckEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**IsHealthy** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHealthCheckEntity

`func NewHealthCheckEntity() *HealthCheckEntity`

NewHealthCheckEntity instantiates a new HealthCheckEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckEntityWithDefaults

`func NewHealthCheckEntityWithDefaults() *HealthCheckEntity`

NewHealthCheckEntityWithDefaults instantiates a new HealthCheckEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HealthCheckEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheckEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheckEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthCheckEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HealthCheckEntity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HealthCheckEntity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsHealthy

`func (o *HealthCheckEntity) GetIsHealthy() bool`

GetIsHealthy returns the IsHealthy field if non-nil, zero value otherwise.

### GetIsHealthyOk

`func (o *HealthCheckEntity) GetIsHealthyOk() (*bool, bool)`

GetIsHealthyOk returns a tuple with the IsHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHealthy

`func (o *HealthCheckEntity) SetIsHealthy(v bool)`

SetIsHealthy sets IsHealthy field to given value.

### HasIsHealthy

`func (o *HealthCheckEntity) HasIsHealthy() bool`

HasIsHealthy returns a boolean if a field has been set.

### GetMessage

`func (o *HealthCheckEntity) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HealthCheckEntity) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HealthCheckEntity) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HealthCheckEntity) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *HealthCheckEntity) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *HealthCheckEntity) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


