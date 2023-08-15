# CloudFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**HttpStatusCode**](HttpStatusCode.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Urls** | Pointer to [**map[string]UrlFileInfo**](UrlFileInfo.md) |  | [optional] 
**Scores** | Pointer to **map[string]float32** |  | [optional] 

## Methods

### NewCloudFileResponse

`func NewCloudFileResponse() *CloudFileResponse`

NewCloudFileResponse instantiates a new CloudFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFileResponseWithDefaults

`func NewCloudFileResponseWithDefaults() *CloudFileResponse`

NewCloudFileResponseWithDefaults instantiates a new CloudFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudFileResponse) GetStatus() HttpStatusCode`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudFileResponse) GetStatusOk() (*HttpStatusCode, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudFileResponse) SetStatus(v HttpStatusCode)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudFileResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CloudFileResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudFileResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudFileResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudFileResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CloudFileResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CloudFileResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetUrls

`func (o *CloudFileResponse) GetUrls() map[string]UrlFileInfo`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CloudFileResponse) GetUrlsOk() (*map[string]UrlFileInfo, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CloudFileResponse) SetUrls(v map[string]UrlFileInfo)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *CloudFileResponse) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### SetUrlsNil

`func (o *CloudFileResponse) SetUrlsNil(b bool)`

 SetUrlsNil sets the value for Urls to be an explicit nil

### UnsetUrls
`func (o *CloudFileResponse) UnsetUrls()`

UnsetUrls ensures that no value is present for Urls, not even an explicit nil
### GetScores

`func (o *CloudFileResponse) GetScores() map[string]float32`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *CloudFileResponse) GetScoresOk() (*map[string]float32, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *CloudFileResponse) SetScores(v map[string]float32)`

SetScores sets Scores field to given value.

### HasScores

`func (o *CloudFileResponse) HasScores() bool`

HasScores returns a boolean if a field has been set.

### SetScoresNil

`func (o *CloudFileResponse) SetScoresNil(b bool)`

 SetScoresNil sets the value for Scores to be an explicit nil

### UnsetScores
`func (o *CloudFileResponse) UnsetScores()`

UnsetScores ensures that no value is present for Scores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


