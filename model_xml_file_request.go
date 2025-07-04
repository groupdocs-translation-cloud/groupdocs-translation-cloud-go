/*
GroupDocs.Translation SDK

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 25.3.0
Contact: anton.perhunov@aspose.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package groupdocs_translation_cloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the XmlFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XmlFileRequest{}

// XmlFileRequest struct for XmlFileRequest
type XmlFileRequest struct {
	// Language of original file
	SourceLanguage string `json:"sourceLanguage"`
	// List of target languages
	TargetLanguages []string `json:"targetLanguages"`
	// Type in the file name. If null will be as request ID.
	OriginalFileName NullableString `json:"originalFileName,omitempty"`
	// Link to file for translation. Ignore, if \"file\" property not null
	Url string `json:"url"`
	// Url or name of the application using this SDK. Not required.
	Origin NullableString `json:"origin,omitempty"`
	// Do result formating like the source. This option needs more expensive requests.
	IsNeedAlignment *bool `json:"isNeedAlignment,omitempty"`
	// Set a specific translation between source and target words.
	TranslationDictionary map[string]string `json:"translationDictionary,omitempty"`
	// Toggle file saving mode for storage.  Is Files by default.
	SavingMode *string `json:"savingMode,omitempty"`
	// List of elements for Xml, Json and Yaml formats. Determines which items should be blacklisted or whitelisted for processing depending on GroupDocs.Translation.ApiGateway.DTO.XmlFileRequest.IsWhiteList.
	IgnoreList []string `json:"ignoreList,omitempty"`
	// Determines to which list the items in GroupDocs.Translation.ApiGateway.DTO.XmlFileRequest.IgnoreList should be allocated. The default is the black list.
	IsWhiteList *bool `json:"isWhiteList,omitempty"`
}

type _XmlFileRequest XmlFileRequest

// NewXmlFileRequest instantiates a new XmlFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlFileRequest(sourceLanguage string, targetLanguages []string, url string) *XmlFileRequest {
	this := XmlFileRequest{}
	this.SourceLanguage = sourceLanguage
	this.TargetLanguages = targetLanguages
	this.Url = url
	return &this
}

// NewXmlFileRequestWithDefaults instantiates a new XmlFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlFileRequestWithDefaults() *XmlFileRequest {
	this := XmlFileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = sourceLanguage
	return &this
}

// GetSourceLanguage returns the SourceLanguage field value
func (o *XmlFileRequest) GetSourceLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceLanguage
}

// GetSourceLanguageOk returns a tuple with the SourceLanguage field value
// and a boolean to check if the value has been set.
func (o *XmlFileRequest) GetSourceLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceLanguage, true
}

// SetSourceLanguage sets field value
func (o *XmlFileRequest) SetSourceLanguage(v string) {
	o.SourceLanguage = v
}

// GetTargetLanguages returns the TargetLanguages field value
func (o *XmlFileRequest) GetTargetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetLanguages
}

// GetTargetLanguagesOk returns a tuple with the TargetLanguages field value
// and a boolean to check if the value has been set.
func (o *XmlFileRequest) GetTargetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetLanguages, true
}

// SetTargetLanguages sets field value
func (o *XmlFileRequest) SetTargetLanguages(v []string) {
	o.TargetLanguages = v
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *XmlFileRequest) GetOriginalFileName() string {
	if o == nil || IsNil(o.OriginalFileName.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalFileName.Get()
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *XmlFileRequest) GetOriginalFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFileName.Get(), o.OriginalFileName.IsSet()
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *XmlFileRequest) HasOriginalFileName() bool {
	if o != nil && o.OriginalFileName.IsSet() {
		return true
	}

	return false
}

// SetOriginalFileName gets a reference to the given NullableString and assigns it to the OriginalFileName field.
func (o *XmlFileRequest) SetOriginalFileName(v string) {
	o.OriginalFileName.Set(&v)
}
// SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil
func (o *XmlFileRequest) SetOriginalFileNameNil() {
	o.OriginalFileName.Set(nil)
}

// UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
func (o *XmlFileRequest) UnsetOriginalFileName() {
	o.OriginalFileName.Unset()
}

// GetUrl returns the Url field value
func (o *XmlFileRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *XmlFileRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *XmlFileRequest) SetUrl(v string) {
	o.Url = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *XmlFileRequest) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *XmlFileRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *XmlFileRequest) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *XmlFileRequest) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *XmlFileRequest) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *XmlFileRequest) UnsetOrigin() {
	o.Origin.Unset()
}

// GetIsNeedAlignment returns the IsNeedAlignment field value if set, zero value otherwise.
func (o *XmlFileRequest) GetIsNeedAlignment() bool {
	if o == nil || IsNil(o.IsNeedAlignment) {
		var ret bool
		return ret
	}
	return *o.IsNeedAlignment
}

// GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlFileRequest) GetIsNeedAlignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNeedAlignment) {
		return nil, false
	}
	return o.IsNeedAlignment, true
}

// HasIsNeedAlignment returns a boolean if a field has been set.
func (o *XmlFileRequest) HasIsNeedAlignment() bool {
	if o != nil && !IsNil(o.IsNeedAlignment) {
		return true
	}

	return false
}

// SetIsNeedAlignment gets a reference to the given bool and assigns it to the IsNeedAlignment field.
func (o *XmlFileRequest) SetIsNeedAlignment(v bool) {
	o.IsNeedAlignment = &v
}

// GetTranslationDictionary returns the TranslationDictionary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *XmlFileRequest) GetTranslationDictionary() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.TranslationDictionary
}

// GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *XmlFileRequest) GetTranslationDictionaryOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TranslationDictionary) {
		return nil, false
	}
	return &o.TranslationDictionary, true
}

// HasTranslationDictionary returns a boolean if a field has been set.
func (o *XmlFileRequest) HasTranslationDictionary() bool {
	if o != nil && !IsNil(o.TranslationDictionary) {
		return true
	}

	return false
}

// SetTranslationDictionary gets a reference to the given map[string]string and assigns it to the TranslationDictionary field.
func (o *XmlFileRequest) SetTranslationDictionary(v map[string]string) {
	o.TranslationDictionary = v
}

// GetSavingMode returns the SavingMode field value if set, zero value otherwise.
func (o *XmlFileRequest) GetSavingMode() string {
	if o == nil || IsNil(o.SavingMode) {
		var ret string
		return ret
	}
	return *o.SavingMode
}

// GetSavingModeOk returns a tuple with the SavingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlFileRequest) GetSavingModeOk() (*string, bool) {
	if o == nil || IsNil(o.SavingMode) {
		return nil, false
	}
	return o.SavingMode, true
}

// HasSavingMode returns a boolean if a field has been set.
func (o *XmlFileRequest) HasSavingMode() bool {
	if o != nil && !IsNil(o.SavingMode) {
		return true
	}

	return false
}

// SetSavingMode gets a reference to the given string and assigns it to the SavingMode field.
func (o *XmlFileRequest) SetSavingMode(v string) {
	o.SavingMode = &v
}

// GetIgnoreList returns the IgnoreList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *XmlFileRequest) GetIgnoreList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IgnoreList
}

// GetIgnoreListOk returns a tuple with the IgnoreList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *XmlFileRequest) GetIgnoreListOk() ([]string, bool) {
	if o == nil || IsNil(o.IgnoreList) {
		return nil, false
	}
	return o.IgnoreList, true
}

// HasIgnoreList returns a boolean if a field has been set.
func (o *XmlFileRequest) HasIgnoreList() bool {
	if o != nil && !IsNil(o.IgnoreList) {
		return true
	}

	return false
}

// SetIgnoreList gets a reference to the given []string and assigns it to the IgnoreList field.
func (o *XmlFileRequest) SetIgnoreList(v []string) {
	o.IgnoreList = v
}

// GetIsWhiteList returns the IsWhiteList field value if set, zero value otherwise.
func (o *XmlFileRequest) GetIsWhiteList() bool {
	if o == nil || IsNil(o.IsWhiteList) {
		var ret bool
		return ret
	}
	return *o.IsWhiteList
}

// GetIsWhiteListOk returns a tuple with the IsWhiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlFileRequest) GetIsWhiteListOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWhiteList) {
		return nil, false
	}
	return o.IsWhiteList, true
}

// HasIsWhiteList returns a boolean if a field has been set.
func (o *XmlFileRequest) HasIsWhiteList() bool {
	if o != nil && !IsNil(o.IsWhiteList) {
		return true
	}

	return false
}

// SetIsWhiteList gets a reference to the given bool and assigns it to the IsWhiteList field.
func (o *XmlFileRequest) SetIsWhiteList(v bool) {
	o.IsWhiteList = &v
}

func (o XmlFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XmlFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceLanguage"] = o.SourceLanguage
	toSerialize["targetLanguages"] = o.TargetLanguages
	if o.OriginalFileName.IsSet() {
		toSerialize["originalFileName"] = o.OriginalFileName.Get()
	}
	toSerialize["url"] = o.Url
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	if !IsNil(o.IsNeedAlignment) {
		toSerialize["isNeedAlignment"] = o.IsNeedAlignment
	}
	if o.TranslationDictionary != nil {
		toSerialize["translationDictionary"] = o.TranslationDictionary
	}
	if !IsNil(o.SavingMode) {
		toSerialize["savingMode"] = o.SavingMode
	}
	if o.IgnoreList != nil {
		toSerialize["ignoreList"] = o.IgnoreList
	}
	if !IsNil(o.IsWhiteList) {
		toSerialize["isWhiteList"] = o.IsWhiteList
	}
	return toSerialize, nil
}

func (o *XmlFileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sourceLanguage",
		"targetLanguages",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varXmlFileRequest := _XmlFileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varXmlFileRequest)

	if err != nil {
		return err
	}

	*o = XmlFileRequest(varXmlFileRequest)

	return err
}

type NullableXmlFileRequest struct {
	value *XmlFileRequest
	isSet bool
}

func (v NullableXmlFileRequest) Get() *XmlFileRequest {
	return v.value
}

func (v *NullableXmlFileRequest) Set(val *XmlFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlFileRequest(val *XmlFileRequest) *NullableXmlFileRequest {
	return &NullableXmlFileRequest{value: val, isSet: true}
}

func (v NullableXmlFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


