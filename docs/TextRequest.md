# TextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguage** | Pointer to **NullableString** | Language of original text | [optional] 
**TargetLanguages** | Pointer to **[]string** | List of target languages | [optional] 
**Texts** | Pointer to **[]string** | Text array to translate | [optional] 
**Origin** | Pointer to **NullableString** | For analysis only | [optional] 
**ContainsMarkdown** | Pointer to **bool** | Set to true if you want to handle markdown in text | [optional] 
**TranslationDictionary** | Pointer to **map[string]string** | Set a specific translation between source and target words. | [optional] 

## Methods

### NewTextRequest

`func NewTextRequest() *TextRequest`

NewTextRequest instantiates a new TextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextRequestWithDefaults

`func NewTextRequestWithDefaults() *TextRequest`

NewTextRequestWithDefaults instantiates a new TextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *TextRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *TextRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *TextRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.

### HasSourceLanguage

`func (o *TextRequest) HasSourceLanguage() bool`

HasSourceLanguage returns a boolean if a field has been set.

### SetSourceLanguageNil

`func (o *TextRequest) SetSourceLanguageNil(b bool)`

 SetSourceLanguageNil sets the value for SourceLanguage to be an explicit nil

### UnsetSourceLanguage
`func (o *TextRequest) UnsetSourceLanguage()`

UnsetSourceLanguage ensures that no value is present for SourceLanguage, not even an explicit nil
### GetTargetLanguages

`func (o *TextRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *TextRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *TextRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.

### HasTargetLanguages

`func (o *TextRequest) HasTargetLanguages() bool`

HasTargetLanguages returns a boolean if a field has been set.

### SetTargetLanguagesNil

`func (o *TextRequest) SetTargetLanguagesNil(b bool)`

 SetTargetLanguagesNil sets the value for TargetLanguages to be an explicit nil

### UnsetTargetLanguages
`func (o *TextRequest) UnsetTargetLanguages()`

UnsetTargetLanguages ensures that no value is present for TargetLanguages, not even an explicit nil
### GetTexts

`func (o *TextRequest) GetTexts() []string`

GetTexts returns the Texts field if non-nil, zero value otherwise.

### GetTextsOk

`func (o *TextRequest) GetTextsOk() (*[]string, bool)`

GetTextsOk returns a tuple with the Texts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexts

`func (o *TextRequest) SetTexts(v []string)`

SetTexts sets Texts field to given value.

### HasTexts

`func (o *TextRequest) HasTexts() bool`

HasTexts returns a boolean if a field has been set.

### SetTextsNil

`func (o *TextRequest) SetTextsNil(b bool)`

 SetTextsNil sets the value for Texts to be an explicit nil

### UnsetTexts
`func (o *TextRequest) UnsetTexts()`

UnsetTexts ensures that no value is present for Texts, not even an explicit nil
### GetOrigin

`func (o *TextRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *TextRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *TextRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *TextRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *TextRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *TextRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetContainsMarkdown

`func (o *TextRequest) GetContainsMarkdown() bool`

GetContainsMarkdown returns the ContainsMarkdown field if non-nil, zero value otherwise.

### GetContainsMarkdownOk

`func (o *TextRequest) GetContainsMarkdownOk() (*bool, bool)`

GetContainsMarkdownOk returns a tuple with the ContainsMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsMarkdown

`func (o *TextRequest) SetContainsMarkdown(v bool)`

SetContainsMarkdown sets ContainsMarkdown field to given value.

### HasContainsMarkdown

`func (o *TextRequest) HasContainsMarkdown() bool`

HasContainsMarkdown returns a boolean if a field has been set.

### GetTranslationDictionary

`func (o *TextRequest) GetTranslationDictionary() map[string]string`

GetTranslationDictionary returns the TranslationDictionary field if non-nil, zero value otherwise.

### GetTranslationDictionaryOk

`func (o *TextRequest) GetTranslationDictionaryOk() (*map[string]string, bool)`

GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationDictionary

`func (o *TextRequest) SetTranslationDictionary(v map[string]string)`

SetTranslationDictionary sets TranslationDictionary field to given value.

### HasTranslationDictionary

`func (o *TextRequest) HasTranslationDictionary() bool`

HasTranslationDictionary returns a boolean if a field has been set.

### SetTranslationDictionaryNil

`func (o *TextRequest) SetTranslationDictionaryNil(b bool)`

 SetTranslationDictionaryNil sets the value for TranslationDictionary to be an explicit nil

### UnsetTranslationDictionary
`func (o *TextRequest) UnsetTranslationDictionary()`

UnsetTranslationDictionary ensures that no value is present for TranslationDictionary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


