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

// checks if the MarkdownFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkdownFileRequest{}

// MarkdownFileRequest Request for markdown files or markdown files with Hugo syntax
type MarkdownFileRequest struct {
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
	// output file format
	OutputFormat string `json:"outputFormat"`
	// Dictionary of short code names and parameters names to translate
	ShortCodeList map[string][]string `json:"shortCodeList,omitempty"`
	// List of lists of frontmatter paths
	FrontMatterList [][]string `json:"frontMatterList,omitempty"`
}

// NewMarkdownFileRequest instantiates a new MarkdownFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkdownFileRequest(sourceLanguage string, targetLanguages []string, outputFormat string) *MarkdownFileRequest {
	this := MarkdownFileRequest{}
	this.SourceLanguage = sourceLanguage
	this.TargetLanguages = targetLanguages
	this.OutputFormat = outputFormat
	return &this
}

// NewMarkdownFileRequestWithDefaults instantiates a new MarkdownFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkdownFileRequestWithDefaults() *MarkdownFileRequest {
	this := MarkdownFileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = sourceLanguage
	return &this
}

// GetSourceLanguage returns the SourceLanguage field value
func (o *MarkdownFileRequest) GetSourceLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceLanguage
}

// GetSourceLanguageOk returns a tuple with the SourceLanguage field value
// and a boolean to check if the value has been set.
func (o *MarkdownFileRequest) GetSourceLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceLanguage, true
}

// SetSourceLanguage sets field value
func (o *MarkdownFileRequest) SetSourceLanguage(v string) {
	o.SourceLanguage = v
}

// GetTargetLanguages returns the TargetLanguages field value
func (o *MarkdownFileRequest) GetTargetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetLanguages
}

// GetTargetLanguagesOk returns a tuple with the TargetLanguages field value
// and a boolean to check if the value has been set.
func (o *MarkdownFileRequest) GetTargetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetLanguages, true
}

// SetTargetLanguages sets field value
func (o *MarkdownFileRequest) SetTargetLanguages(v []string) {
	o.TargetLanguages = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkdownFileRequest) GetFile() string {
	if o == nil || IsNil(o.File.Get()) {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkdownFileRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *MarkdownFileRequest) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *MarkdownFileRequest) SetFile(v string) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *MarkdownFileRequest) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *MarkdownFileRequest) UnsetFile() {
	o.File.Unset()
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkdownFileRequest) GetOriginalFileName() string {
	if o == nil || IsNil(o.OriginalFileName.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalFileName.Get()
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkdownFileRequest) GetOriginalFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFileName.Get(), o.OriginalFileName.IsSet()
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *MarkdownFileRequest) HasOriginalFileName() bool {
	if o != nil && o.OriginalFileName.IsSet() {
		return true
	}

	return false
}

// SetOriginalFileName gets a reference to the given NullableString and assigns it to the OriginalFileName field.
func (o *MarkdownFileRequest) SetOriginalFileName(v string) {
	o.OriginalFileName.Set(&v)
}
// SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil
func (o *MarkdownFileRequest) SetOriginalFileNameNil() {
	o.OriginalFileName.Set(nil)
}

// UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
func (o *MarkdownFileRequest) UnsetOriginalFileName() {
	o.OriginalFileName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkdownFileRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkdownFileRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MarkdownFileRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MarkdownFileRequest) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *MarkdownFileRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MarkdownFileRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkdownFileRequest) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkdownFileRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *MarkdownFileRequest) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *MarkdownFileRequest) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *MarkdownFileRequest) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *MarkdownFileRequest) UnsetOrigin() {
	o.Origin.Unset()
}

// GetSavingMode returns the SavingMode field value if set, zero value otherwise.
func (o *MarkdownFileRequest) GetSavingMode() string {
	if o == nil || IsNil(o.SavingMode) {
		var ret string
		return ret
	}
	return *o.SavingMode
}

// GetSavingModeOk returns a tuple with the SavingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkdownFileRequest) GetSavingModeOk() (*string, bool) {
	if o == nil || IsNil(o.SavingMode) {
		return nil, false
	}
	return o.SavingMode, true
}

// HasSavingMode returns a boolean if a field has been set.
func (o *MarkdownFileRequest) HasSavingMode() bool {
	if o != nil && !IsNil(o.SavingMode) {
		return true
	}

	return false
}

// SetSavingMode gets a reference to the given string and assigns it to the SavingMode field.
func (o *MarkdownFileRequest) SetSavingMode(v string) {
	o.SavingMode = &v
}

// GetOutputFormat returns the OutputFormat field value
func (o *MarkdownFileRequest) GetOutputFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value
// and a boolean to check if the value has been set.
func (o *MarkdownFileRequest) GetOutputFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputFormat, true
}

// SetOutputFormat sets field value
func (o *MarkdownFileRequest) SetOutputFormat(v string) {
	o.OutputFormat = v
}

// GetShortCodeList returns the ShortCodeList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkdownFileRequest) GetShortCodeList() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.ShortCodeList
}

// GetShortCodeListOk returns a tuple with the ShortCodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkdownFileRequest) GetShortCodeListOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ShortCodeList) {
		return nil, false
	}
	return &o.ShortCodeList, true
}

// HasShortCodeList returns a boolean if a field has been set.
func (o *MarkdownFileRequest) HasShortCodeList() bool {
	if o != nil && IsNil(o.ShortCodeList) {
		return true
	}

	return false
}

// SetShortCodeList gets a reference to the given map[string][]string and assigns it to the ShortCodeList field.
func (o *MarkdownFileRequest) SetShortCodeList(v map[string][]string) {
	o.ShortCodeList = v
}

// GetFrontMatterList returns the FrontMatterList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkdownFileRequest) GetFrontMatterList() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.FrontMatterList
}

// GetFrontMatterListOk returns a tuple with the FrontMatterList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkdownFileRequest) GetFrontMatterListOk() ([][]string, bool) {
	if o == nil || IsNil(o.FrontMatterList) {
		return nil, false
	}
	return o.FrontMatterList, true
}

// HasFrontMatterList returns a boolean if a field has been set.
func (o *MarkdownFileRequest) HasFrontMatterList() bool {
	if o != nil && IsNil(o.FrontMatterList) {
		return true
	}

	return false
}

// SetFrontMatterList gets a reference to the given [][]string and assigns it to the FrontMatterList field.
func (o *MarkdownFileRequest) SetFrontMatterList(v [][]string) {
	o.FrontMatterList = v
}

func (o MarkdownFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkdownFileRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["outputFormat"] = o.OutputFormat
	if o.ShortCodeList != nil {
		toSerialize["shortCodeList"] = o.ShortCodeList
	}
	if o.FrontMatterList != nil {
		toSerialize["frontMatterList"] = o.FrontMatterList
	}
	return toSerialize, nil
}

type NullableMarkdownFileRequest struct {
	value *MarkdownFileRequest
	isSet bool
}

func (v NullableMarkdownFileRequest) Get() *MarkdownFileRequest {
	return v.value
}

func (v *NullableMarkdownFileRequest) Set(val *MarkdownFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkdownFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkdownFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkdownFileRequest(val *MarkdownFileRequest) *NullableMarkdownFileRequest {
	return &NullableMarkdownFileRequest{value: val, isSet: true}
}

func (v NullableMarkdownFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkdownFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


