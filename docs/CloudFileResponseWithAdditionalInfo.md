# CloudFileResponseWithAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**HttpStatusCode**](HttpStatusCode.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Urls** | Pointer to [**map[string]UrlFileInfo**](UrlFileInfo.md) |  | [optional] 
**Scores** | Pointer to **map[string]float32** |  | [optional] 
**Request** | Pointer to [**CloudFileRequest**](CloudFileRequest.md) |  | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCloudFileResponseWithAdditionalInfo

`func NewCloudFileResponseWithAdditionalInfo() *CloudFileResponseWithAdditionalInfo`

NewCloudFileResponseWithAdditionalInfo instantiates a new CloudFileResponseWithAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFileResponseWithAdditionalInfoWithDefaults

`func NewCloudFileResponseWithAdditionalInfoWithDefaults() *CloudFileResponseWithAdditionalInfo`

NewCloudFileResponseWithAdditionalInfoWithDefaults instantiates a new CloudFileResponseWithAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudFileResponseWithAdditionalInfo) GetStatus() HttpStatusCode`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudFileResponseWithAdditionalInfo) GetStatusOk() (*HttpStatusCode, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudFileResponseWithAdditionalInfo) SetStatus(v HttpStatusCode)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudFileResponseWithAdditionalInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CloudFileResponseWithAdditionalInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudFileResponseWithAdditionalInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudFileResponseWithAdditionalInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudFileResponseWithAdditionalInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CloudFileResponseWithAdditionalInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CloudFileResponseWithAdditionalInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetUrls

`func (o *CloudFileResponseWithAdditionalInfo) GetUrls() map[string]UrlFileInfo`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CloudFileResponseWithAdditionalInfo) GetUrlsOk() (*map[string]UrlFileInfo, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CloudFileResponseWithAdditionalInfo) SetUrls(v map[string]UrlFileInfo)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *CloudFileResponseWithAdditionalInfo) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### SetUrlsNil

`func (o *CloudFileResponseWithAdditionalInfo) SetUrlsNil(b bool)`

 SetUrlsNil sets the value for Urls to be an explicit nil

### UnsetUrls
`func (o *CloudFileResponseWithAdditionalInfo) UnsetUrls()`

UnsetUrls ensures that no value is present for Urls, not even an explicit nil
### GetScores

`func (o *CloudFileResponseWithAdditionalInfo) GetScores() map[string]float32`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *CloudFileResponseWithAdditionalInfo) GetScoresOk() (*map[string]float32, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *CloudFileResponseWithAdditionalInfo) SetScores(v map[string]float32)`

SetScores sets Scores field to given value.

### HasScores

`func (o *CloudFileResponseWithAdditionalInfo) HasScores() bool`

HasScores returns a boolean if a field has been set.

### SetScoresNil

`func (o *CloudFileResponseWithAdditionalInfo) SetScoresNil(b bool)`

 SetScoresNil sets the value for Scores to be an explicit nil

### UnsetScores
`func (o *CloudFileResponseWithAdditionalInfo) UnsetScores()`

UnsetScores ensures that no value is present for Scores, not even an explicit nil
### GetRequest

`func (o *CloudFileResponseWithAdditionalInfo) GetRequest() CloudFileRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CloudFileResponseWithAdditionalInfo) GetRequestOk() (*CloudFileRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CloudFileResponseWithAdditionalInfo) SetRequest(v CloudFileRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *CloudFileResponseWithAdditionalInfo) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetDetails

`func (o *CloudFileResponseWithAdditionalInfo) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CloudFileResponseWithAdditionalInfo) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CloudFileResponseWithAdditionalInfo) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CloudFileResponseWithAdditionalInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CloudFileResponseWithAdditionalInfo) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CloudFileResponseWithAdditionalInfo) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


