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

// checks if the FileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileRequest{}

// FileRequest struct for FileRequest
type FileRequest struct {
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
	// output file format
	OutputFormat *string `json:"outputFormat,omitempty"`
	// If translate master slides
	Masters *bool `json:"masters,omitempty"`
	// If document's formatting should be preserved, default true
	Formatting *bool `json:"formatting,omitempty"`
	// Endpoint route
	Route NullableString `json:"route,omitempty"`
	// Separator in files
	Separator NullableString `json:"separator,omitempty"`
	// List of slides to translate
	Elements []int32 `json:"elements,omitempty"`
	// Dictionary of ranges in Excel workbooks
	Ranges map[string]WorksheetData `json:"ranges,omitempty"`
	// Dictionary of short code names and parameters names to translate
	ShortCodeList map[string][]string `json:"shortCodeList,omitempty"`
	// Dictionary where key is zero-based front matter index and value is list of lists of front matter paths
	FrontMatterList [][]string `json:"frontMatterList,omitempty"`
}

// NewFileRequest instantiates a new FileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileRequest() *FileRequest {
	this := FileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = &sourceLanguage
	var format string = "Unknown"
	this.Format = &format
	var masters bool = false
	this.Masters = &masters
	var formatting bool = true
	this.Formatting = &formatting
	return &this
}

// NewFileRequestWithDefaults instantiates a new FileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileRequestWithDefaults() *FileRequest {
	this := FileRequest{}
	var sourceLanguage string = "en"
	this.SourceLanguage = &sourceLanguage
	var format string = "Unknown"
	this.Format = &format
	var masters bool = false
	this.Masters = &masters
	var formatting bool = true
	this.Formatting = &formatting
	return &this
}

// GetSourceLanguage returns the SourceLanguage field value if set, zero value otherwise.
func (o *FileRequest) GetSourceLanguage() string {
	if o == nil || IsNil(o.SourceLanguage) {
		var ret string
		return ret
	}
	return *o.SourceLanguage
}

// GetSourceLanguageOk returns a tuple with the SourceLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetSourceLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.SourceLanguage) {
		return nil, false
	}
	return o.SourceLanguage, true
}

// HasSourceLanguage returns a boolean if a field has been set.
func (o *FileRequest) HasSourceLanguage() bool {
	if o != nil && !IsNil(o.SourceLanguage) {
		return true
	}

	return false
}

// SetSourceLanguage gets a reference to the given string and assigns it to the SourceLanguage field.
func (o *FileRequest) SetSourceLanguage(v string) {
	o.SourceLanguage = &v
}

// GetTargetLanguages returns the TargetLanguages field value if set, zero value otherwise.
func (o *FileRequest) GetTargetLanguages() []string {
	if o == nil || IsNil(o.TargetLanguages) {
		var ret []string
		return ret
	}
	return o.TargetLanguages
}

// GetTargetLanguagesOk returns a tuple with the TargetLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetTargetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetLanguages) {
		return nil, false
	}
	return o.TargetLanguages, true
}

// HasTargetLanguages returns a boolean if a field has been set.
func (o *FileRequest) HasTargetLanguages() bool {
	if o != nil && !IsNil(o.TargetLanguages) {
		return true
	}

	return false
}

// SetTargetLanguages gets a reference to the given []string and assigns it to the TargetLanguages field.
func (o *FileRequest) SetTargetLanguages(v []string) {
	o.TargetLanguages = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetFile() string {
	if o == nil || IsNil(o.File.Get()) {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *FileRequest) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *FileRequest) SetFile(v string) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *FileRequest) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *FileRequest) UnsetFile() {
	o.File.Unset()
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetOriginalFileName() string {
	if o == nil || IsNil(o.OriginalFileName.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalFileName.Get()
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetOriginalFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFileName.Get(), o.OriginalFileName.IsSet()
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *FileRequest) HasOriginalFileName() bool {
	if o != nil && o.OriginalFileName.IsSet() {
		return true
	}

	return false
}

// SetOriginalFileName gets a reference to the given NullableString and assigns it to the OriginalFileName field.
func (o *FileRequest) SetOriginalFileName(v string) {
	o.OriginalFileName.Set(&v)
}
// SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil
func (o *FileRequest) SetOriginalFileNameNil() {
	o.OriginalFileName.Set(nil)
}

// UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
func (o *FileRequest) UnsetOriginalFileName() {
	o.OriginalFileName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *FileRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *FileRequest) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *FileRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *FileRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *FileRequest) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *FileRequest) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *FileRequest) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *FileRequest) UnsetOrigin() {
	o.Origin.Unset()
}

// GetSavingMode returns the SavingMode field value if set, zero value otherwise.
func (o *FileRequest) GetSavingMode() string {
	if o == nil || IsNil(o.SavingMode) {
		var ret string
		return ret
	}
	return *o.SavingMode
}

// GetSavingModeOk returns a tuple with the SavingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetSavingModeOk() (*string, bool) {
	if o == nil || IsNil(o.SavingMode) {
		return nil, false
	}
	return o.SavingMode, true
}

// HasSavingMode returns a boolean if a field has been set.
func (o *FileRequest) HasSavingMode() bool {
	if o != nil && !IsNil(o.SavingMode) {
		return true
	}

	return false
}

// SetSavingMode gets a reference to the given string and assigns it to the SavingMode field.
func (o *FileRequest) SetSavingMode(v string) {
	o.SavingMode = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *FileRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *FileRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *FileRequest) SetFormat(v string) {
	o.Format = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *FileRequest) GetOutputFormat() string {
	if o == nil || IsNil(o.OutputFormat) {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetOutputFormatOk() (*string, bool) {
	if o == nil || IsNil(o.OutputFormat) {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *FileRequest) HasOutputFormat() bool {
	if o != nil && !IsNil(o.OutputFormat) {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *FileRequest) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

// GetMasters returns the Masters field value if set, zero value otherwise.
func (o *FileRequest) GetMasters() bool {
	if o == nil || IsNil(o.Masters) {
		var ret bool
		return ret
	}
	return *o.Masters
}

// GetMastersOk returns a tuple with the Masters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetMastersOk() (*bool, bool) {
	if o == nil || IsNil(o.Masters) {
		return nil, false
	}
	return o.Masters, true
}

// HasMasters returns a boolean if a field has been set.
func (o *FileRequest) HasMasters() bool {
	if o != nil && !IsNil(o.Masters) {
		return true
	}

	return false
}

// SetMasters gets a reference to the given bool and assigns it to the Masters field.
func (o *FileRequest) SetMasters(v bool) {
	o.Masters = &v
}

// GetFormatting returns the Formatting field value if set, zero value otherwise.
func (o *FileRequest) GetFormatting() bool {
	if o == nil || IsNil(o.Formatting) {
		var ret bool
		return ret
	}
	return *o.Formatting
}

// GetFormattingOk returns a tuple with the Formatting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRequest) GetFormattingOk() (*bool, bool) {
	if o == nil || IsNil(o.Formatting) {
		return nil, false
	}
	return o.Formatting, true
}

// HasFormatting returns a boolean if a field has been set.
func (o *FileRequest) HasFormatting() bool {
	if o != nil && !IsNil(o.Formatting) {
		return true
	}

	return false
}

// SetFormatting gets a reference to the given bool and assigns it to the Formatting field.
func (o *FileRequest) SetFormatting(v bool) {
	o.Formatting = &v
}

// GetRoute returns the Route field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetRoute() string {
	if o == nil || IsNil(o.Route.Get()) {
		var ret string
		return ret
	}
	return *o.Route.Get()
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetRouteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Route.Get(), o.Route.IsSet()
}

// HasRoute returns a boolean if a field has been set.
func (o *FileRequest) HasRoute() bool {
	if o != nil && o.Route.IsSet() {
		return true
	}

	return false
}

// SetRoute gets a reference to the given NullableString and assigns it to the Route field.
func (o *FileRequest) SetRoute(v string) {
	o.Route.Set(&v)
}
// SetRouteNil sets the value for Route to be an explicit nil
func (o *FileRequest) SetRouteNil() {
	o.Route.Set(nil)
}

// UnsetRoute ensures that no value is present for Route, not even an explicit nil
func (o *FileRequest) UnsetRoute() {
	o.Route.Unset()
}

// GetSeparator returns the Separator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetSeparator() string {
	if o == nil || IsNil(o.Separator.Get()) {
		var ret string
		return ret
	}
	return *o.Separator.Get()
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetSeparatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Separator.Get(), o.Separator.IsSet()
}

// HasSeparator returns a boolean if a field has been set.
func (o *FileRequest) HasSeparator() bool {
	if o != nil && o.Separator.IsSet() {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given NullableString and assigns it to the Separator field.
func (o *FileRequest) SetSeparator(v string) {
	o.Separator.Set(&v)
}
// SetSeparatorNil sets the value for Separator to be an explicit nil
func (o *FileRequest) SetSeparatorNil() {
	o.Separator.Set(nil)
}

// UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
func (o *FileRequest) UnsetSeparator() {
	o.Separator.Unset()
}

// GetElements returns the Elements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetElements() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetElementsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *FileRequest) HasElements() bool {
	if o != nil && IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []int32 and assigns it to the Elements field.
func (o *FileRequest) SetElements(v []int32) {
	o.Elements = v
}

// GetRanges returns the Ranges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetRanges() map[string]WorksheetData {
	if o == nil {
		var ret map[string]WorksheetData
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetRangesOk() (*map[string]WorksheetData, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return &o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *FileRequest) HasRanges() bool {
	if o != nil && IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given map[string]WorksheetData and assigns it to the Ranges field.
func (o *FileRequest) SetRanges(v map[string]WorksheetData) {
	o.Ranges = v
}

// GetShortCodeList returns the ShortCodeList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetShortCodeList() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.ShortCodeList
}

// GetShortCodeListOk returns a tuple with the ShortCodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetShortCodeListOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ShortCodeList) {
		return nil, false
	}
	return &o.ShortCodeList, true
}

// HasShortCodeList returns a boolean if a field has been set.
func (o *FileRequest) HasShortCodeList() bool {
	if o != nil && IsNil(o.ShortCodeList) {
		return true
	}

	return false
}

// SetShortCodeList gets a reference to the given map[string][]string and assigns it to the ShortCodeList field.
func (o *FileRequest) SetShortCodeList(v map[string][]string) {
	o.ShortCodeList = v
}

// GetFrontMatterList returns the FrontMatterList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileRequest) GetFrontMatterList() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.FrontMatterList
}

// GetFrontMatterListOk returns a tuple with the FrontMatterList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileRequest) GetFrontMatterListOk() ([][]string, bool) {
	if o == nil || IsNil(o.FrontMatterList) {
		return nil, false
	}
	return o.FrontMatterList, true
}

// HasFrontMatterList returns a boolean if a field has been set.
func (o *FileRequest) HasFrontMatterList() bool {
	if o != nil && IsNil(o.FrontMatterList) {
		return true
	}

	return false
}

// SetFrontMatterList gets a reference to the given [][]string and assigns it to the FrontMatterList field.
func (o *FileRequest) SetFrontMatterList(v [][]string) {
	o.FrontMatterList = v
}

func (o FileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Formatting) {
		toSerialize["formatting"] = o.Formatting
	}
	if o.Route.IsSet() {
		toSerialize["route"] = o.Route.Get()
	}
	if o.Separator.IsSet() {
		toSerialize["separator"] = o.Separator.Get()
	}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	if o.ShortCodeList != nil {
		toSerialize["shortCodeList"] = o.ShortCodeList
	}
	if o.FrontMatterList != nil {
		toSerialize["frontMatterList"] = o.FrontMatterList
	}
	return toSerialize, nil
}

type NullableFileRequest struct {
	value *FileRequest
	isSet bool
}

func (v NullableFileRequest) Get() *FileRequest {
	return v.value
}

func (v *NullableFileRequest) Set(val *FileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileRequest(val *FileRequest) *NullableFileRequest {
	return &NullableFileRequest{value: val, isSet: true}
}

func (v NullableFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


