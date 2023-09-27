/*
GroupDocs.Translation SDK

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 23.9.4
Contact: anton.perhunov@aspose.com
*/


package groupdocs_translation_api

import (
	"encoding/json"
)

// checks if the UrlFileInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrlFileInfo{}

// UrlFileInfo struct for UrlFileInfo
type UrlFileInfo struct {
	Url NullableString `json:"url,omitempty"`
	Size *int64 `json:"size,omitempty"`
}

// NewUrlFileInfo instantiates a new UrlFileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlFileInfo() *UrlFileInfo {
	this := UrlFileInfo{}
	return &this
}

// NewUrlFileInfoWithDefaults instantiates a new UrlFileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlFileInfoWithDefaults() *UrlFileInfo {
	this := UrlFileInfo{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UrlFileInfo) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UrlFileInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *UrlFileInfo) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *UrlFileInfo) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *UrlFileInfo) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *UrlFileInfo) UnsetUrl() {
	o.Url.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UrlFileInfo) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlFileInfo) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UrlFileInfo) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *UrlFileInfo) SetSize(v int64) {
	o.Size = &v
}

func (o UrlFileInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlFileInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableUrlFileInfo struct {
	value *UrlFileInfo
	isSet bool
}

func (v NullableUrlFileInfo) Get() *UrlFileInfo {
	return v.value
}

func (v *NullableUrlFileInfo) Set(val *UrlFileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlFileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlFileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlFileInfo(val *UrlFileInfo) *NullableUrlFileInfo {
	return &NullableUrlFileInfo{value: val, isSet: true}
}

func (v NullableUrlFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlFileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


