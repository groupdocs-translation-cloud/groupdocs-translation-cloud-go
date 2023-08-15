# HealthCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaDeliveryStatus** | Pointer to **NullableString** |  | [optional] 
**CloudStatus** | Pointer to **NullableString** |  | [optional] 

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

### GetKafkaDeliveryStatus

`func (o *HealthCheckStatus) GetKafkaDeliveryStatus() string`

GetKafkaDeliveryStatus returns the KafkaDeliveryStatus field if non-nil, zero value otherwise.

### GetKafkaDeliveryStatusOk

`func (o *HealthCheckStatus) GetKafkaDeliveryStatusOk() (*string, bool)`

GetKafkaDeliveryStatusOk returns a tuple with the KafkaDeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaDeliveryStatus

`func (o *HealthCheckStatus) SetKafkaDeliveryStatus(v string)`

SetKafkaDeliveryStatus sets KafkaDeliveryStatus field to given value.

### HasKafkaDeliveryStatus

`func (o *HealthCheckStatus) HasKafkaDeliveryStatus() bool`

HasKafkaDeliveryStatus returns a boolean if a field has been set.

### SetKafkaDeliveryStatusNil

`func (o *HealthCheckStatus) SetKafkaDeliveryStatusNil(b bool)`

 SetKafkaDeliveryStatusNil sets the value for KafkaDeliveryStatus to be an explicit nil

### UnsetKafkaDeliveryStatus
`func (o *HealthCheckStatus) UnsetKafkaDeliveryStatus()`

UnsetKafkaDeliveryStatus ensures that no value is present for KafkaDeliveryStatus, not even an explicit nil
### GetCloudStatus

`func (o *HealthCheckStatus) GetCloudStatus() string`

GetCloudStatus returns the CloudStatus field if non-nil, zero value otherwise.

### GetCloudStatusOk

`func (o *HealthCheckStatus) GetCloudStatusOk() (*string, bool)`

GetCloudStatusOk returns a tuple with the CloudStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStatus

`func (o *HealthCheckStatus) SetCloudStatus(v string)`

SetCloudStatus sets CloudStatus field to given value.

### HasCloudStatus

`func (o *HealthCheckStatus) HasCloudStatus() bool`

HasCloudStatus returns a boolean if a field has been set.

### SetCloudStatusNil

`func (o *HealthCheckStatus) SetCloudStatusNil(b bool)`

 SetCloudStatusNil sets the value for CloudStatus to be an explicit nil

### UnsetCloudStatus
`func (o *HealthCheckStatus) UnsetCloudStatus()`

UnsetCloudStatus ensures that no value is present for CloudStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


