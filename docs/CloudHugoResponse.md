# CloudHugoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**HttpStatusCode**](HttpStatusCode.md) |  | [optional] 
**Message** | Pointer to **NullableString** | If file was parsed correctly or text of error | [optional] 
**Frontmatters** | Pointer to **[][]string** | Structure of front matter syntax | [optional] 
**Shortcodes** | Pointer to **[][]string** | Structure of short code syntax | [optional] 

## Methods

### NewCloudHugoResponse

`func NewCloudHugoResponse() *CloudHugoResponse`

NewCloudHugoResponse instantiates a new CloudHugoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHugoResponseWithDefaults

`func NewCloudHugoResponseWithDefaults() *CloudHugoResponse`

NewCloudHugoResponseWithDefaults instantiates a new CloudHugoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudHugoResponse) GetStatus() HttpStatusCode`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudHugoResponse) GetStatusOk() (*HttpStatusCode, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudHugoResponse) SetStatus(v HttpStatusCode)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudHugoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CloudHugoResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudHugoResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudHugoResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudHugoResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CloudHugoResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CloudHugoResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetFrontmatters

`func (o *CloudHugoResponse) GetFrontmatters() [][]string`

GetFrontmatters returns the Frontmatters field if non-nil, zero value otherwise.

### GetFrontmattersOk

`func (o *CloudHugoResponse) GetFrontmattersOk() (*[][]string, bool)`

GetFrontmattersOk returns a tuple with the Frontmatters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontmatters

`func (o *CloudHugoResponse) SetFrontmatters(v [][]string)`

SetFrontmatters sets Frontmatters field to given value.

### HasFrontmatters

`func (o *CloudHugoResponse) HasFrontmatters() bool`

HasFrontmatters returns a boolean if a field has been set.

### SetFrontmattersNil

`func (o *CloudHugoResponse) SetFrontmattersNil(b bool)`

 SetFrontmattersNil sets the value for Frontmatters to be an explicit nil

### UnsetFrontmatters
`func (o *CloudHugoResponse) UnsetFrontmatters()`

UnsetFrontmatters ensures that no value is present for Frontmatters, not even an explicit nil
### GetShortcodes

`func (o *CloudHugoResponse) GetShortcodes() [][]string`

GetShortcodes returns the Shortcodes field if non-nil, zero value otherwise.

### GetShortcodesOk

`func (o *CloudHugoResponse) GetShortcodesOk() (*[][]string, bool)`

GetShortcodesOk returns a tuple with the Shortcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcodes

`func (o *CloudHugoResponse) SetShortcodes(v [][]string)`

SetShortcodes sets Shortcodes field to given value.

### HasShortcodes

`func (o *CloudHugoResponse) HasShortcodes() bool`

HasShortcodes returns a boolean if a field has been set.

### SetShortcodesNil

`func (o *CloudHugoResponse) SetShortcodesNil(b bool)`

 SetShortcodesNil sets the value for Shortcodes to be an explicit nil

### UnsetShortcodes
`func (o *CloudHugoResponse) UnsetShortcodes()`

UnsetShortcodes ensures that no value is present for Shortcodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


