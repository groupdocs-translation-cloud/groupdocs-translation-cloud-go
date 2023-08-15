/*
GroupDocs.Translation SDK

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 23.8
*/

package groupdocs_translation_api

import (
	"encoding/json"
)

// checks if the ResxFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResxFileRequest{}

// ResxFileRequest Request for resources files like resx
type ResxFileRequest struct {
	// Language of original file
	SourceLanguage string `json:"sourceLanguage"`
	// List of target languages
	TargetLanguages []string `json:"targetLanguages"`
	// File as byte array
	File NullableString `json:"file,omitempty"`
	// Type in the file name. If null will be as request ID.
	OriginalFileName NullableString `json:"originalFileName,omitempty"`
	// Link to file for translation. Ignore, if \"file\" property not null
	Url NullableString `json:"url,omitempty"`
	// Url or name of application using this SDK. Not required.
	Origin NullableString `json:"origin,omitempty"`
	// Toggle file saving mode for storage.  Is Files by default.
	SavingMode *string `json:"savingMode,omitempty"`
}

// NewResxFileRequest instantiates a new ResxFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResxFileRequest(sourceLanguage string, targetLanguages []string) *ResxFileRequest {
	this := ResxFileRequest{}
	this.SourceLanguage = sourceLanguage
	this.TargetLanguages = targetLanguages
	return &this
}

// NewResxFileRequestWithDefaults instantiates a new ResxFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResxFileRequestWithDefaults() *ResxFileRequest {
	this := ResxFileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = sourceLanguage
	return &this
}

// GetSourceLanguage returns the SourceLanguage field value
func (o *ResxFileRequest) GetSourceLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceLanguage
}

// GetSourceLanguageOk returns a tuple with the SourceLanguage field value
// and a boolean to check if the value has been set.
func (o *ResxFileRequest) GetSourceLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceLanguage, true
}

// SetSourceLanguage sets field value
func (o *ResxFileRequest) SetSourceLanguage(v string) {
	o.SourceLanguage = v
}

// GetTargetLanguages returns the TargetLanguages field value
func (o *ResxFileRequest) GetTargetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetLanguages
}

// GetTargetLanguagesOk returns a tuple with the TargetLanguages field value
// and a boolean to check if the value has been set.
func (o *ResxFileRequest) GetTargetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetLanguages, true
}

// SetTargetLanguages sets field value
func (o *ResxFileRequest) SetTargetLanguages(v []string) {
	o.TargetLanguages = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResxFileRequest) GetFile() string {
	if o == nil || IsNil(o.File.Get()) {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResxFileRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *ResxFileRequest) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *ResxFileRequest) SetFile(v string) {
	o.File.Set(&v)
}

// SetFileNil sets the value for File to be an explicit nil
func (o *ResxFileRequest) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *ResxFileRequest) UnsetFile() {
	o.File.Unset()
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResxFileRequest) GetOriginalFileName() string {
	if o == nil || IsNil(o.OriginalFileName.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalFileName.Get()
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResxFileRequest) GetOriginalFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFileName.Get(), o.OriginalFileName.IsSet()
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *ResxFileRequest) HasOriginalFileName() bool {
	if o != nil && o.OriginalFileName.IsSet() {
		return true
	}

	return false
}

// SetOriginalFileName gets a reference to the given NullableString and assigns it to the OriginalFileName field.
func (o *ResxFileRequest) SetOriginalFileName(v string) {
	o.OriginalFileName.Set(&v)
}

// SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil
func (o *ResxFileRequest) SetOriginalFileNameNil() {
	o.OriginalFileName.Set(nil)
}

// UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
func (o *ResxFileRequest) UnsetOriginalFileName() {
	o.OriginalFileName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResxFileRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResxFileRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ResxFileRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ResxFileRequest) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *ResxFileRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ResxFileRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResxFileRequest) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResxFileRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *ResxFileRequest) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *ResxFileRequest) SetOrigin(v string) {
	o.Origin.Set(&v)
}

// SetOriginNil sets the value for Origin to be an explicit nil
func (o *ResxFileRequest) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *ResxFileRequest) UnsetOrigin() {
	o.Origin.Unset()
}

// GetSavingMode returns the SavingMode field value if set, zero value otherwise.
func (o *ResxFileRequest) GetSavingMode() string {
	if o == nil || IsNil(o.SavingMode) {
		var ret string
		return ret
	}
	return *o.SavingMode
}

// GetSavingModeOk returns a tuple with the SavingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResxFileRequest) GetSavingModeOk() (*string, bool) {
	if o == nil || IsNil(o.SavingMode) {
		return nil, false
	}
	return o.SavingMode, true
}

// HasSavingMode returns a boolean if a field has been set.
func (o *ResxFileRequest) HasSavingMode() bool {
	if o != nil && !IsNil(o.SavingMode) {
		return true
	}

	return false
}

// SetSavingMode gets a reference to the given string and assigns it to the SavingMode field.
func (o *ResxFileRequest) SetSavingMode(v string) {
	o.SavingMode = &v
}

func (o ResxFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResxFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceLanguage"] = o.SourceLanguage
	toSerialize["targetLanguages"] = o.TargetLanguages
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	if o.OriginalFileName.IsSet() {
		toSerialize["originalFileName"] = o.OriginalFileName.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	if !IsNil(o.SavingMode) {
		toSerialize["savingMode"] = o.SavingMode
	}
	return toSerialize, nil
}

type NullableResxFileRequest struct {
	value *ResxFileRequest
	isSet bool
}

func (v NullableResxFileRequest) Get() *ResxFileRequest {
	return v.value
}

func (v *NullableResxFileRequest) Set(val *ResxFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResxFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResxFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResxFileRequest(val *ResxFileRequest) *NullableResxFileRequest {
	return &NullableResxFileRequest{value: val, isSet: true}
}

func (v NullableResxFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResxFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
