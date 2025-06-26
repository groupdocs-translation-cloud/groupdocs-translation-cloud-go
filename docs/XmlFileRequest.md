# XmlFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLanguage** | **string** | Language of original file | [default to "en"]
**TargetLanguages** | **[]string** | List of target languages | 
**OriginalFileName** | Pointer to **NullableString** | Type in the file name. If null will be as request ID. | [optional] 
**Url** | **string** | Link to file for translation. Ignore, if \&quot;file\&quot; property not null | 
**Origin** | Pointer to **NullableString** | Url or name of the application using this SDK. Not required. | [optional] 
**IsNeedAlignment** | Pointer to **bool** | Do result formating like the source. This option needs more expensive requests. | [optional] 
**TranslationDictionary** | Pointer to **map[string]string** | Set a specific translation between source and target words. | [optional] 
**SavingMode** | Pointer to **string** | Toggle file saving mode for storage.  Is Files by default. | [optional] 
**IgnoreList** | Pointer to **[]string** | List of elements for Xml, Json and Yaml formats. Determines which items should be blacklisted or whitelisted for processing depending on GroupDocs.Translation.ApiGateway.DTO.XmlFileRequest.IsWhiteList. | [optional] 
**IsWhiteList** | Pointer to **bool** | Determines to which list the items in GroupDocs.Translation.ApiGateway.DTO.XmlFileRequest.IgnoreList should be allocated. The default is the black list. | [optional] 

## Methods

### NewXmlFileRequest

`func NewXmlFileRequest(sourceLanguage string, targetLanguages []string, url string, ) *XmlFileRequest`

NewXmlFileRequest instantiates a new XmlFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlFileRequestWithDefaults

`func NewXmlFileRequestWithDefaults() *XmlFileRequest`

NewXmlFileRequestWithDefaults instantiates a new XmlFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLanguage

`func (o *XmlFileRequest) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *XmlFileRequest) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *XmlFileRequest) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.


### GetTargetLanguages

`func (o *XmlFileRequest) GetTargetLanguages() []string`

GetTargetLanguages returns the TargetLanguages field if non-nil, zero value otherwise.

### GetTargetLanguagesOk

`func (o *XmlFileRequest) GetTargetLanguagesOk() (*[]string, bool)`

GetTargetLanguagesOk returns a tuple with the TargetLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLanguages

`func (o *XmlFileRequest) SetTargetLanguages(v []string)`

SetTargetLanguages sets TargetLanguages field to given value.


### GetOriginalFileName

`func (o *XmlFileRequest) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *XmlFileRequest) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *XmlFileRequest) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *XmlFileRequest) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### SetOriginalFileNameNil

`func (o *XmlFileRequest) SetOriginalFileNameNil(b bool)`

 SetOriginalFileNameNil sets the value for OriginalFileName to be an explicit nil

### UnsetOriginalFileName
`func (o *XmlFileRequest) UnsetOriginalFileName()`

UnsetOriginalFileName ensures that no value is present for OriginalFileName, not even an explicit nil
### GetUrl

`func (o *XmlFileRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *XmlFileRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *XmlFileRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetOrigin

`func (o *XmlFileRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *XmlFileRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *XmlFileRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *XmlFileRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *XmlFileRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *XmlFileRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetIsNeedAlignment

`func (o *XmlFileRequest) GetIsNeedAlignment() bool`

GetIsNeedAlignment returns the IsNeedAlignment field if non-nil, zero value otherwise.

### GetIsNeedAlignmentOk

`func (o *XmlFileRequest) GetIsNeedAlignmentOk() (*bool, bool)`

GetIsNeedAlignmentOk returns a tuple with the IsNeedAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeedAlignment

`func (o *XmlFileRequest) SetIsNeedAlignment(v bool)`

SetIsNeedAlignment sets IsNeedAlignment field to given value.

### HasIsNeedAlignment

`func (o *XmlFileRequest) HasIsNeedAlignment() bool`

HasIsNeedAlignment returns a boolean if a field has been set.

### GetTranslationDictionary

`func (o *XmlFileRequest) GetTranslationDictionary() map[string]string`

GetTranslationDictionary returns the TranslationDictionary field if non-nil, zero value otherwise.

### GetTranslationDictionaryOk

`func (o *XmlFileRequest) GetTranslationDictionaryOk() (*map[string]string, bool)`

GetTranslationDictionaryOk returns a tuple with the TranslationDictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationDictionary

`func (o *XmlFileRequest) SetTranslationDictionary(v map[string]string)`

SetTranslationDictionary sets TranslationDictionary field to given value.

### HasTranslationDictionary

`func (o *XmlFileRequest) HasTranslationDictionary() bool`

HasTranslationDictionary returns a boolean if a field has been set.

### SetTranslationDictionaryNil

`func (o *XmlFileRequest) SetTranslationDictionaryNil(b bool)`

 SetTranslationDictionaryNil sets the value for TranslationDictionary to be an explicit nil

### UnsetTranslationDictionary
`func (o *XmlFileRequest) UnsetTranslationDictionary()`

UnsetTranslationDictionary ensures that no value is present for TranslationDictionary, not even an explicit nil
### GetSavingMode

`func (o *XmlFileRequest) GetSavingMode() string`

GetSavingMode returns the SavingMode field if non-nil, zero value otherwise.

### GetSavingModeOk

`func (o *XmlFileRequest) GetSavingModeOk() (*string, bool)`

GetSavingModeOk returns a tuple with the SavingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingMode

`func (o *XmlFileRequest) SetSavingMode(v string)`

SetSavingMode sets SavingMode field to given value.

### HasSavingMode

`func (o *XmlFileRequest) HasSavingMode() bool`

HasSavingMode returns a boolean if a field has been set.

### GetIgnoreList

`func (o *XmlFileRequest) GetIgnoreList() []string`

GetIgnoreList returns the IgnoreList field if non-nil, zero value otherwise.

### GetIgnoreListOk

`func (o *XmlFileRequest) GetIgnoreListOk() (*[]string, bool)`

GetIgnoreListOk returns a tuple with the IgnoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreList

`func (o *XmlFileRequest) SetIgnoreList(v []string)`

SetIgnoreList sets IgnoreList field to given value.

### HasIgnoreList

`func (o *XmlFileRequest) HasIgnoreList() bool`

HasIgnoreList returns a boolean if a field has been set.

### SetIgnoreListNil

`func (o *XmlFileRequest) SetIgnoreListNil(b bool)`

 SetIgnoreListNil sets the value for IgnoreList to be an explicit nil

### UnsetIgnoreList
`func (o *XmlFileRequest) UnsetIgnoreList()`

UnsetIgnoreList ensures that no value is present for IgnoreList, not even an explicit nil
### GetIsWhiteList

`func (o *XmlFileRequest) GetIsWhiteList() bool`

GetIsWhiteList returns the IsWhiteList field if non-nil, zero value otherwise.

### GetIsWhiteListOk

`func (o *XmlFileRequest) GetIsWhiteListOk() (*bool, bool)`

GetIsWhiteListOk returns a tuple with the IsWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteList

`func (o *XmlFileRequest) SetIsWhiteList(v bool)`

SetIsWhiteList sets IsWhiteList field to given value.

### HasIsWhiteList

`func (o *XmlFileRequest) HasIsWhiteList() bool`

HasIsWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


