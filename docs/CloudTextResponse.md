# CloudTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**HttpStatusCode**](HttpStatusCode.md) |  | [optional] 
**Message** | Pointer to **NullableString** | If file was translated correctly or text of error | [optional] 
**Translations** | Pointer to **map[string][]string** | Translated texts | [optional] 

## Methods

### NewCloudTextResponse

`func NewCloudTextResponse() *CloudTextResponse`

NewCloudTextResponse instantiates a new CloudTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTextResponseWithDefaults

`func NewCloudTextResponseWithDefaults() *CloudTextResponse`

NewCloudTextResponseWithDefaults instantiates a new CloudTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudTextResponse) GetStatus() HttpStatusCode`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudTextResponse) GetStatusOk() (*HttpStatusCode, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudTextResponse) SetStatus(v HttpStatusCode)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudTextResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CloudTextResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudTextResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudTextResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudTextResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CloudTextResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CloudTextResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTranslations

`func (o *CloudTextResponse) GetTranslations() map[string][]string`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *CloudTextResponse) GetTranslationsOk() (*map[string][]string, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *CloudTextResponse) SetTranslations(v map[string][]string)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *CloudTextResponse) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### SetTranslationsNil

`func (o *CloudTextResponse) SetTranslationsNil(b bool)`

 SetTranslationsNil sets the value for Translations to be an explicit nil

### UnsetTranslations
`func (o *CloudTextResponse) UnsetTranslations()`

UnsetTranslations ensures that no value is present for Translations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


