/*
GroupDocs.Translation SDK

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 23.8
*/

package groupdocs_translation_api

import (
	"encoding/json"
)

// checks if the AutoPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoPostRequest{}

// AutoPostRequest struct for AutoPostRequest
type AutoPostRequest struct {
	// Input file format
	Format string `json:"Format"`
	// output file format
	OutputFormat string `json:"OutputFormat"`
	// If translate master slides
	Masters *bool `json:"Masters,omitempty"`
	// If document's formatting should be preserved, default true
	Formatting *bool `json:"Formatting,omitempty"`
	// Endpoint route
	Route *string `json:"Route,omitempty"`
	// Separator in files
	Separator *string `json:"Separator,omitempty"`
	// List of slides to translate
	Elements []int32 `json:"Elements,omitempty"`
	// Dictionary of ranges in Excel workbooks
	Ranges *map[string]WorksheetData `json:"Ranges,omitempty"`
	// Dictionary of short code names and parameters names to translate
	ShortCodeList *map[string][][]string `json:"ShortCodeList,omitempty"`
	// Dictionary where key is zero-based front matter index and value is list of lists of front matter paths
	FrontMatterList [][]string `json:"FrontMatterList,omitempty"`
	// Language of original file
	SourceLanguage string `json:"SourceLanguage"`
	// List of target languages
	TargetLanguages []string `json:"TargetLanguages"`
	// File as byte array
	File *string `json:"File,omitempty"`
	// Type in the file name. If null will be as request ID.
	OriginalFileName *string `json:"OriginalFileName,omitempty"`
	// Link to file for translation. Ignore, if \"file\" property not null
	Url *string `json:"Url,omitempty"`
	// Url or name of application using this SDK. Not required.
	Origin *string `json:"Origin,omitempty"`
	// Toggle file saving mode for storage.  Is Files by default.
	SavingMode *string `json:"SavingMode,omitempty"`
}

// NewAutoPostRequest instantiates a new AutoPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoPostRequest(format string, outputFormat string, sourceLanguage string, targetLanguages []string) *AutoPostRequest {
	this := AutoPostRequest{}
	this.Format = format
	this.OutputFormat = outputFormat
	var masters bool = false
	this.Masters = &masters
	var formatting bool = true
	this.Formatting = &formatting
	this.SourceLanguage = sourceLanguage
	this.TargetLanguages = targetLanguages
	return &this
}

// NewAutoPostRequestWithDefaults instantiates a new AutoPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoPostRequestWithDefaults() *AutoPostRequest {
	this := AutoPostRequest{}
	var format string = "Unknown"
	this.Format = format
	var masters bool = false
	this.Masters = &masters
	var formatting bool = true
	this.Formatting = &formatting
	var sourceLanguage string = "en"
	this.SourceLanguage = sourceLanguage
	return &this
}

// GetFormat returns the Format field value
func (o *AutoPostRequest) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *AutoPostRequest) SetFormat(v string) {
	o.Format = v
}

// GetOutputFormat returns the OutputFormat field value
func (o *AutoPostRequest) GetOutputFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetOutputFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputFormat, true
}

// SetOutputFormat sets field value
func (o *AutoPostRequest) SetOutputFormat(v string) {
	o.OutputFormat = v
}

// GetMasters returns the Masters field value if set, zero value otherwise.
func (o *AutoPostRequest) GetMasters() bool {
	if o == nil || IsNil(o.Masters) {
		var ret bool
		return ret
	}
	return *o.Masters
}

// GetMastersOk returns a tuple with the Masters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetMastersOk() (*bool, bool) {
	if o == nil || IsNil(o.Masters) {
		return nil, false
	}
	return o.Masters, true
}

// HasMasters returns a boolean if a field has been set.
func (o *AutoPostRequest) HasMasters() bool {
	if o != nil && !IsNil(o.Masters) {
		return true
	}

	return false
}

// SetMasters gets a reference to the given bool and assigns it to the Masters field.
func (o *AutoPostRequest) SetMasters(v bool) {
	o.Masters = &v
}

// GetFormatting returns the Formatting field value if set, zero value otherwise.
func (o *AutoPostRequest) GetFormatting() bool {
	if o == nil || IsNil(o.Formatting) {
		var ret bool
		return ret
	}
	return *o.Formatting
}

// GetFormattingOk returns a tuple with the Formatting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetFormattingOk() (*bool, bool) {
	if o == nil || IsNil(o.Formatting) {
		return nil, false
	}
	return o.Formatting, true
}

// HasFormatting returns a boolean if a field has been set.
func (o *AutoPostRequest) HasFormatting() bool {
	if o != nil && !IsNil(o.Formatting) {
		return true
	}

	return false
}

// SetFormatting gets a reference to the given bool and assigns it to the Formatting field.
func (o *AutoPostRequest) SetFormatting(v bool) {
	o.Formatting = &v
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *AutoPostRequest) GetRoute() string {
	if o == nil || IsNil(o.Route) {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetRouteOk() (*string, bool) {
	if o == nil || IsNil(o.Route) {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *AutoPostRequest) HasRoute() bool {
	if o != nil && !IsNil(o.Route) {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *AutoPostRequest) SetRoute(v string) {
	o.Route = &v
}

// GetSeparator returns the Separator field value if set, zero value otherwise.
func (o *AutoPostRequest) GetSeparator() string {
	if o == nil || IsNil(o.Separator) {
		var ret string
		return ret
	}
	return *o.Separator
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetSeparatorOk() (*string, bool) {
	if o == nil || IsNil(o.Separator) {
		return nil, false
	}
	return o.Separator, true
}

// HasSeparator returns a boolean if a field has been set.
func (o *AutoPostRequest) HasSeparator() bool {
	if o != nil && !IsNil(o.Separator) {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given string and assigns it to the Separator field.
func (o *AutoPostRequest) SetSeparator(v string) {
	o.Separator = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *AutoPostRequest) GetElements() []int32 {
	if o == nil || IsNil(o.Elements) {
		var ret []int32
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetElementsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *AutoPostRequest) HasElements() bool {
	if o != nil && !IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []int32 and assigns it to the Elements field.
func (o *AutoPostRequest) SetElements(v []int32) {
	o.Elements = v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *AutoPostRequest) GetRanges() map[string]WorksheetData {
	if o == nil || IsNil(o.Ranges) {
		var ret map[string]WorksheetData
		return ret
	}
	return *o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetRangesOk() (*map[string]WorksheetData, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *AutoPostRequest) HasRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given map[string]WorksheetData and assigns it to the Ranges field.
func (o *AutoPostRequest) SetRanges(v map[string]WorksheetData) {
	o.Ranges = &v
}

// GetShortCodeList returns the ShortCodeList field value if set, zero value otherwise.
func (o *AutoPostRequest) GetShortCodeList() map[string][][]string {
	if o == nil || IsNil(o.ShortCodeList) {
		var ret map[string][][]string
		return ret
	}
	return *o.ShortCodeList
}

// GetShortCodeListOk returns a tuple with the ShortCodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetShortCodeListOk() (*map[string][][]string, bool) {
	if o == nil || IsNil(o.ShortCodeList) {
		return nil, false
	}
	return o.ShortCodeList, true
}

// HasShortCodeList returns a boolean if a field has been set.
func (o *AutoPostRequest) HasShortCodeList() bool {
	if o != nil && !IsNil(o.ShortCodeList) {
		return true
	}

	return false
}

// SetShortCodeList gets a reference to the given map[string][][]string and assigns it to the ShortCodeList field.
func (o *AutoPostRequest) SetShortCodeList(v map[string][][]string) {
	o.ShortCodeList = &v
}

// GetFrontMatterList returns the FrontMatterList field value if set, zero value otherwise.
func (o *AutoPostRequest) GetFrontMatterList() [][]string {
	if o == nil || IsNil(o.FrontMatterList) {
		var ret [][]string
		return ret
	}
	return o.FrontMatterList
}

// GetFrontMatterListOk returns a tuple with the FrontMatterList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetFrontMatterListOk() ([][]string, bool) {
	if o == nil || IsNil(o.FrontMatterList) {
		return nil, false
	}
	return o.FrontMatterList, true
}

// HasFrontMatterList returns a boolean if a field has been set.
func (o *AutoPostRequest) HasFrontMatterList() bool {
	if o != nil && !IsNil(o.FrontMatterList) {
		return true
	}

	return false
}

// SetFrontMatterList gets a reference to the given [][]string and assigns it to the FrontMatterList field.
func (o *AutoPostRequest) SetFrontMatterList(v [][]string) {
	o.FrontMatterList = v
}

// GetSourceLanguage returns the SourceLanguage field value
func (o *AutoPostRequest) GetSourceLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceLanguage
}

// GetSourceLanguageOk returns a tuple with the SourceLanguage field value
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetSourceLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceLanguage, true
}

// SetSourceLanguage sets field value
func (o *AutoPostRequest) SetSourceLanguage(v string) {
	o.SourceLanguage = v
}

// GetTargetLanguages returns the TargetLanguages field value
func (o *AutoPostRequest) GetTargetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetLanguages
}

// GetTargetLanguagesOk returns a tuple with the TargetLanguages field value
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetTargetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetLanguages, true
}

// SetTargetLanguages sets field value
func (o *AutoPostRequest) SetTargetLanguages(v []string) {
	o.TargetLanguages = v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *AutoPostRequest) GetFile() string {
	if o == nil || IsNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetFileOk() (*string, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *AutoPostRequest) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *AutoPostRequest) SetFile(v string) {
	o.File = &v
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise.
func (o *AutoPostRequest) GetOriginalFileName() string {
	if o == nil || IsNil(o.OriginalFileName) {
		var ret string
		return ret
	}
	return *o.OriginalFileName
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetOriginalFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalFileName) {
		return nil, false
	}
	return o.OriginalFileName, true
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *AutoPostRequest) HasOriginalFileName() bool {
	if o != nil && !IsNil(o.OriginalFileName) {
		return true
	}

	return false
}

// SetOriginalFileName gets a reference to the given string and assigns it to the OriginalFileName field.
func (o *AutoPostRequest) SetOriginalFileName(v string) {
	o.OriginalFileName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AutoPostRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AutoPostRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AutoPostRequest) SetUrl(v string) {
	o.Url = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *AutoPostRequest) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *AutoPostRequest) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *AutoPostRequest) SetOrigin(v string) {
	o.Origin = &v
}

// GetSavingMode returns the SavingMode field value if set, zero value otherwise.
func (o *AutoPostRequest) GetSavingMode() string {
	if o == nil || IsNil(o.SavingMode) {
		var ret string
		return ret
	}
	return *o.SavingMode
}

// GetSavingModeOk returns a tuple with the SavingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPostRequest) GetSavingModeOk() (*string, bool) {
	if o == nil || IsNil(o.SavingMode) {
		return nil, false
	}
	return o.SavingMode, true
}

// HasSavingMode returns a boolean if a field has been set.
func (o *AutoPostRequest) HasSavingMode() bool {
	if o != nil && !IsNil(o.SavingMode) {
		return true
	}

	return false
}

// SetSavingMode gets a reference to the given string and assigns it to the SavingMode field.
func (o *AutoPostRequest) SetSavingMode(v string) {
	o.SavingMode = &v
}

func (o AutoPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Format"] = o.Format
	toSerialize["OutputFormat"] = o.OutputFormat
	if !IsNil(o.Masters) {
		toSerialize["Masters"] = o.Masters
	}
	if !IsNil(o.Formatting) {
		toSerialize["Formatting"] = o.Formatting
	}
	if !IsNil(o.Route) {
		toSerialize["Route"] = o.Route
	}
	if !IsNil(o.Separator) {
		toSerialize["Separator"] = o.Separator
	}
	if !IsNil(o.Elements) {
		toSerialize["Elements"] = o.Elements
	}
	if !IsNil(o.Ranges) {
		toSerialize["Ranges"] = o.Ranges
	}
	if !IsNil(o.ShortCodeList) {
		toSerialize["ShortCodeList"] = o.ShortCodeList
	}
	if !IsNil(o.FrontMatterList) {
		toSerialize["FrontMatterList"] = o.FrontMatterList
	}
	toSerialize["SourceLanguage"] = o.SourceLanguage
	toSerialize["TargetLanguages"] = o.TargetLanguages
	if !IsNil(o.File) {
		toSerialize["File"] = o.File
	}
	if !IsNil(o.OriginalFileName) {
		toSerialize["OriginalFileName"] = o.OriginalFileName
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.Origin) {
		toSerialize["Origin"] = o.Origin
	}
	if !IsNil(o.SavingMode) {
		toSerialize["SavingMode"] = o.SavingMode
	}
	return toSerialize, nil
}

type NullableAutoPostRequest struct {
	value *AutoPostRequest
	isSet bool
}

func (v NullableAutoPostRequest) Get() *AutoPostRequest {
	return v.value
}

func (v *NullableAutoPostRequest) Set(val *AutoPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoPostRequest(val *AutoPostRequest) *NullableAutoPostRequest {
	return &NullableAutoPostRequest{value: val, isSet: true}
}

func (v NullableAutoPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
