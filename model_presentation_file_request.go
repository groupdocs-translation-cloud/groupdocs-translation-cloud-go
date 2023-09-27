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

// checks if the PresentationFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresentationFileRequest{}

// PresentationFileRequest Request for presentation files like ppt, pptx, pptm, odp
type PresentationFileRequest struct {
	// Language of original file
	SourceLanguage *string `json:"sourceLanguage,omitempty"`
	// List of target languages
	TargetLanguages []string `json:"targetLanguages,omitempty"`
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
	// Input file format
	Format *string `json:"format,omitempty"`
	// Output file format
	OutputFormat *string `json:"outputFormat,omitempty"`
	// If translate master slides
	Masters *bool `json:"masters,omitempty"`
	// List of slides to translate
	Slides []int32 `json:"slides,omitempty"`
}

// NewPresentationFileRequest instantiates a new PresentationFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresentationFileRequest() *PresentationFileRequest {
	this := PresentationFileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = &sourceLanguage
	var format string = "Pptx"
	this.Format = &format
	var masters bool = false
	this.Masters = &masters
	return &this
}

// NewPresentationFileRequestWithDefaults instantiates a new PresentationFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresentationFileRequestWithDefaults() *PresentationFileRequest {
	this := PresentationFileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = &sourceLanguage
	var format string = "Pptx"
	this.Format = &format
	var masters bool = false
	this.Masters = &masters
	return &this
}

// GetSourceLanguage returns the SourceLanguage field value if set, zero value otherwise.
func (o *PresentationFileRequest) GetSourceLanguage() string {
	if o == nil || IsNil(o.SourceLanguage) {
		var ret string
		return ret
	}
	return *o.SourceLanguage
}

// GetSourceLanguageOk returns a tuple with the SourceLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresentationFileRequest) GetSourceLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.SourceLanguage) {
		return nil, false
	}
	return o.SourceLanguage, true
}

// HasSourceLanguage returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasSourceLanguage() bool {
	if o != nil && !IsNil(o.SourceLanguage) {
		return true
	}

	return false
}

// SetSourceLanguage gets a reference to the given string and assigns it to the SourceLanguage field.
func (o *PresentationFileRequest) SetSourceLanguage(v string) {
	o.SourceLanguage = &v
}

// GetTargetLanguages returns the TargetLanguages field value if set, zero value otherwise.
func (o *PresentationFileRequest) GetTargetLanguages() []string {
	if o == nil || IsNil(o.TargetLanguages) {
		var ret []string
		return ret
	}
	return o.TargetLanguages
}

// GetTargetLanguagesOk returns a tuple with the TargetLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresentationFileRequest) GetTargetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetLanguages) {
		return nil, false
	}
	return o.TargetLanguages, true
}

// HasTargetLanguages returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasTargetLanguages() bool {
	if o != nil && !IsNil(o.TargetLanguages) {
		return true
	}

	return false
}

// SetTargetLanguages gets a reference to the given []string and assigns it to the TargetLanguages field.
func (o *PresentationFileRequest) SetTargetLanguages(v []string) {
	o.TargetLanguages = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresentationFileRequest) GetFile() string {
	if o == nil || IsNil(o.File.Get()) {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresentationFileRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *PresentationFileRequest) SetFile(v string) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *PresentationFileRequest) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *PresentationFileRequest) UnsetFile() {
	o.File.Unset()
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresentationFileRequest) GetOriginalFileName() string {
	if o == nil || IsNil(o.OriginalFileName.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalFileName.Get()
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresentationFileRequest) GetOriginalFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFileName.Get(), o.OriginalFileName.IsSet()
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasOriginalFileName() bool {
	if o != nil && o.OriginalFileName.IsSet() {
		return true
	}

	return false
}

// SetOriginalFileName gets a reference to the given NullableString and assigns it to the OriginalFileName field.
func (o *PresentationFileRequest) SetOriginalFileName(v string) {
	o.OriginalFileName.Set(&v)
}
// SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil
func (o *PresentationFileRequest) SetOriginalFileNameNil() {
	o.OriginalFileName.Set(nil)
}

// UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
func (o *PresentationFileRequest) UnsetOriginalFileName() {
	o.OriginalFileName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresentationFileRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresentationFileRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PresentationFileRequest) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PresentationFileRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PresentationFileRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresentationFileRequest) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresentationFileRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *PresentationFileRequest) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *PresentationFileRequest) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *PresentationFileRequest) UnsetOrigin() {
	o.Origin.Unset()
}

// GetSavingMode returns the SavingMode field value if set, zero value otherwise.
func (o *PresentationFileRequest) GetSavingMode() string {
	if o == nil || IsNil(o.SavingMode) {
		var ret string
		return ret
	}
	return *o.SavingMode
}

// GetSavingModeOk returns a tuple with the SavingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresentationFileRequest) GetSavingModeOk() (*string, bool) {
	if o == nil || IsNil(o.SavingMode) {
		return nil, false
	}
	return o.SavingMode, true
}

// HasSavingMode returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasSavingMode() bool {
	if o != nil && !IsNil(o.SavingMode) {
		return true
	}

	return false
}

// SetSavingMode gets a reference to the given string and assigns it to the SavingMode field.
func (o *PresentationFileRequest) SetSavingMode(v string) {
	o.SavingMode = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PresentationFileRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresentationFileRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PresentationFileRequest) SetFormat(v string) {
	o.Format = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *PresentationFileRequest) GetOutputFormat() string {
	if o == nil || IsNil(o.OutputFormat) {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresentationFileRequest) GetOutputFormatOk() (*string, bool) {
	if o == nil || IsNil(o.OutputFormat) {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasOutputFormat() bool {
	if o != nil && !IsNil(o.OutputFormat) {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *PresentationFileRequest) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

// GetMasters returns the Masters field value if set, zero value otherwise.
func (o *PresentationFileRequest) GetMasters() bool {
	if o == nil || IsNil(o.Masters) {
		var ret bool
		return ret
	}
	return *o.Masters
}

// GetMastersOk returns a tuple with the Masters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresentationFileRequest) GetMastersOk() (*bool, bool) {
	if o == nil || IsNil(o.Masters) {
		return nil, false
	}
	return o.Masters, true
}

// HasMasters returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasMasters() bool {
	if o != nil && !IsNil(o.Masters) {
		return true
	}

	return false
}

// SetMasters gets a reference to the given bool and assigns it to the Masters field.
func (o *PresentationFileRequest) SetMasters(v bool) {
	o.Masters = &v
}

// GetSlides returns the Slides field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresentationFileRequest) GetSlides() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Slides
}

// GetSlidesOk returns a tuple with the Slides field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresentationFileRequest) GetSlidesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Slides) {
		return nil, false
	}
	return o.Slides, true
}

// HasSlides returns a boolean if a field has been set.
func (o *PresentationFileRequest) HasSlides() bool {
	if o != nil && IsNil(o.Slides) {
		return true
	}

	return false
}

// SetSlides gets a reference to the given []int32 and assigns it to the Slides field.
func (o *PresentationFileRequest) SetSlides(v []int32) {
	o.Slides = v
}

func (o PresentationFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresentationFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceLanguage) {
		toSerialize["sourceLanguage"] = o.SourceLanguage
	}
	if !IsNil(o.TargetLanguages) {
		toSerialize["targetLanguages"] = o.TargetLanguages
	}
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
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.OutputFormat) {
		toSerialize["outputFormat"] = o.OutputFormat
	}
	if !IsNil(o.Masters) {
		toSerialize["masters"] = o.Masters
	}
	if o.Slides != nil {
		toSerialize["slides"] = o.Slides
	}
	return toSerialize, nil
}

type NullablePresentationFileRequest struct {
	value *PresentationFileRequest
	isSet bool
}

func (v NullablePresentationFileRequest) Get() *PresentationFileRequest {
	return v.value
}

func (v *NullablePresentationFileRequest) Set(val *PresentationFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePresentationFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePresentationFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresentationFileRequest(val *PresentationFileRequest) *NullablePresentationFileRequest {
	return &NullablePresentationFileRequest{value: val, isSet: true}
}

func (v NullablePresentationFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresentationFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


