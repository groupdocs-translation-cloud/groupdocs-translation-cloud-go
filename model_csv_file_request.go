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

// checks if the CsvFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsvFileRequest{}

// CsvFileRequest Request for comma / tab separated files like csv, tsv
type CsvFileRequest struct {
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
	// Input file format
	Format *string `json:"format,omitempty"`
	// output file format
	OutputFormat string `json:"outputFormat"`
	// Separator in files
	Separator NullableString `json:"separator,omitempty"`
}

type _CsvFileRequest CsvFileRequest

// NewCsvFileRequest instantiates a new CsvFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsvFileRequest(sourceLanguage string, targetLanguages []string, url string, outputFormat string) *CsvFileRequest {
	this := CsvFileRequest{}
	this.SourceLanguage = sourceLanguage
	this.TargetLanguages = targetLanguages
	this.Url = url
	var format string = "Csv"
	this.Format = &format
	this.OutputFormat = outputFormat
	return &this
}

// NewCsvFileRequestWithDefaults instantiates a new CsvFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsvFileRequestWithDefaults() *CsvFileRequest {
	this := CsvFileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = sourceLanguage
	var format string = "Csv"
	this.Format = &format
	return &this
}

// GetSourceLanguage returns the SourceLanguage field value
func (o *CsvFileRequest) GetSourceLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceLanguage
}

// GetSourceLanguageOk returns a tuple with the SourceLanguage field value
// and a boolean to check if the value has been set.
func (o *CsvFileRequest) GetSourceLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceLanguage, true
}

// SetSourceLanguage sets field value
func (o *CsvFileRequest) SetSourceLanguage(v string) {
	o.SourceLanguage = v
}

// GetTargetLanguages returns the TargetLanguages field value
func (o *CsvFileRequest) GetTargetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetLanguages
}

// GetTargetLanguagesOk returns a tuple with the TargetLanguages field value
// and a boolean to check if the value has been set.
func (o *CsvFileRequest) GetTargetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetLanguages, true
}

// SetTargetLanguages sets field value
func (o *CsvFileRequest) SetTargetLanguages(v []string) {
	o.TargetLanguages = v
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsvFileRequest) GetOriginalFileName() string {
	if o == nil || IsNil(o.OriginalFileName.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalFileName.Get()
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CsvFileRequest) GetOriginalFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFileName.Get(), o.OriginalFileName.IsSet()
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *CsvFileRequest) HasOriginalFileName() bool {
	if o != nil && o.OriginalFileName.IsSet() {
		return true
	}

	return false
}

// SetOriginalFileName gets a reference to the given NullableString and assigns it to the OriginalFileName field.
func (o *CsvFileRequest) SetOriginalFileName(v string) {
	o.OriginalFileName.Set(&v)
}
// SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil
func (o *CsvFileRequest) SetOriginalFileNameNil() {
	o.OriginalFileName.Set(nil)
}

// UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
func (o *CsvFileRequest) UnsetOriginalFileName() {
	o.OriginalFileName.Unset()
}

// GetUrl returns the Url field value
func (o *CsvFileRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CsvFileRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CsvFileRequest) SetUrl(v string) {
	o.Url = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsvFileRequest) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CsvFileRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *CsvFileRequest) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *CsvFileRequest) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *CsvFileRequest) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *CsvFileRequest) UnsetOrigin() {
	o.Origin.Unset()
}

// GetIsNeedAlignment returns the IsNeedAlignment field value if set, zero value otherwise.
func (o *CsvFileRequest) GetIsNeedAlignment() bool {
	if o == nil || IsNil(o.IsNeedAlignment) {
		var ret bool
		return ret
	}
	return *o.IsNeedAlignment
}

// GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvFileRequest) GetIsNeedAlignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNeedAlignment) {
		return nil, false
	}
	return o.IsNeedAlignment, true
}

// HasIsNeedAlignment returns a boolean if a field has been set.
func (o *CsvFileRequest) HasIsNeedAlignment() bool {
	if o != nil && !IsNil(o.IsNeedAlignment) {
		return true
	}

	return false
}

// SetIsNeedAlignment gets a reference to the given bool and assigns it to the IsNeedAlignment field.
func (o *CsvFileRequest) SetIsNeedAlignment(v bool) {
	o.IsNeedAlignment = &v
}

// GetTranslationDictionary returns the TranslationDictionary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsvFileRequest) GetTranslationDictionary() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.TranslationDictionary
}

// GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CsvFileRequest) GetTranslationDictionaryOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TranslationDictionary) {
		return nil, false
	}
	return &o.TranslationDictionary, true
}

// HasTranslationDictionary returns a boolean if a field has been set.
func (o *CsvFileRequest) HasTranslationDictionary() bool {
	if o != nil && !IsNil(o.TranslationDictionary) {
		return true
	}

	return false
}

// SetTranslationDictionary gets a reference to the given map[string]string and assigns it to the TranslationDictionary field.
func (o *CsvFileRequest) SetTranslationDictionary(v map[string]string) {
	o.TranslationDictionary = v
}

// GetSavingMode returns the SavingMode field value if set, zero value otherwise.
func (o *CsvFileRequest) GetSavingMode() string {
	if o == nil || IsNil(o.SavingMode) {
		var ret string
		return ret
	}
	return *o.SavingMode
}

// GetSavingModeOk returns a tuple with the SavingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvFileRequest) GetSavingModeOk() (*string, bool) {
	if o == nil || IsNil(o.SavingMode) {
		return nil, false
	}
	return o.SavingMode, true
}

// HasSavingMode returns a boolean if a field has been set.
func (o *CsvFileRequest) HasSavingMode() bool {
	if o != nil && !IsNil(o.SavingMode) {
		return true
	}

	return false
}

// SetSavingMode gets a reference to the given string and assigns it to the SavingMode field.
func (o *CsvFileRequest) SetSavingMode(v string) {
	o.SavingMode = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CsvFileRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsvFileRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CsvFileRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CsvFileRequest) SetFormat(v string) {
	o.Format = &v
}

// GetOutputFormat returns the OutputFormat field value
func (o *CsvFileRequest) GetOutputFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value
// and a boolean to check if the value has been set.
func (o *CsvFileRequest) GetOutputFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputFormat, true
}

// SetOutputFormat sets field value
func (o *CsvFileRequest) SetOutputFormat(v string) {
	o.OutputFormat = v
}

// GetSeparator returns the Separator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsvFileRequest) GetSeparator() string {
	if o == nil || IsNil(o.Separator.Get()) {
		var ret string
		return ret
	}
	return *o.Separator.Get()
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CsvFileRequest) GetSeparatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Separator.Get(), o.Separator.IsSet()
}

// HasSeparator returns a boolean if a field has been set.
func (o *CsvFileRequest) HasSeparator() bool {
	if o != nil && o.Separator.IsSet() {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given NullableString and assigns it to the Separator field.
func (o *CsvFileRequest) SetSeparator(v string) {
	o.Separator.Set(&v)
}
// SetSeparatorNil sets the value for Separator to be an explicit nil
func (o *CsvFileRequest) SetSeparatorNil() {
	o.Separator.Set(nil)
}

// UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
func (o *CsvFileRequest) UnsetSeparator() {
	o.Separator.Unset()
}

func (o CsvFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsvFileRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	toSerialize["outputFormat"] = o.OutputFormat
	if o.Separator.IsSet() {
		toSerialize["separator"] = o.Separator.Get()
	}
	return toSerialize, nil
}

func (o *CsvFileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sourceLanguage",
		"targetLanguages",
		"url",
		"outputFormat",
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

	varCsvFileRequest := _CsvFileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCsvFileRequest)

	if err != nil {
		return err
	}

	*o = CsvFileRequest(varCsvFileRequest)

	return err
}

type NullableCsvFileRequest struct {
	value *CsvFileRequest
	isSet bool
}

func (v NullableCsvFileRequest) Get() *CsvFileRequest {
	return v.value
}

func (v *NullableCsvFileRequest) Set(val *CsvFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCsvFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCsvFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsvFileRequest(val *CsvFileRequest) *NullableCsvFileRequest {
	return &NullableCsvFileRequest{value: val, isSet: true}
}

func (v NullableCsvFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsvFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


