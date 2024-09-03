# Go API client for groupdocs_translation_cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview

- API version: 24.3.0
- Package version: 24.3.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://products.groupdocs.cloud/translation/](https://products.groupdocs.cloud/translation/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import groupdocs_translation_cloud "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), groupdocs_translation_cloud.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), groupdocs_translation_cloud.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), groupdocs_translation_cloud.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), groupdocs_translation_cloud.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.groupdocs.cloud/v2.0/translation*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FileAPI* | [**FileUploadPost**](docs/FileAPI.md#fileuploadpost) | **Post** /file/upload | 
*SwaggerAPI* | [**SwaggerSpecGet**](docs/SwaggerAPI.md#swaggerspecget) | **Get** /swagger/spec | 
*TranslationAPI* | [**AutoPost**](docs/TranslationAPI.md#autopost) | **Post** /auto | Translate any supported file
*TranslationAPI* | [**CsvPost**](docs/TranslationAPI.md#csvpost) | **Post** /csv | Translate CSV and TSV files
*TranslationAPI* | [**DocumentPost**](docs/TranslationAPI.md#documentpost) | **Post** /document | Translate Microsoft Word documents, rtf, txt, odt
*TranslationAPI* | [**DocumentRequestIdGet**](docs/TranslationAPI.md#documentrequestidget) | **Get** /document/{requestId} | Return document translation status.  Also return URLs for downloading of translated document if translation was successful
*TranslationAPI* | [**DocumentTrialGet**](docs/TranslationAPI.md#documenttrialget) | **Get** /document/trial | Return document translation status for trial request.  Also return URLs for downloading of translated document if translation was successful
*TranslationAPI* | [**DocumentTrialPost**](docs/TranslationAPI.md#documenttrialpost) | **Post** /document/trial | Trial translate Microsoft Word documents, rtf, txt, odt without conversation. Translate only first page or 1000 symbols.
*TranslationAPI* | [**HcGet**](docs/TranslationAPI.md#hcget) | **Get** /hc | Health check for all services.
*TranslationAPI* | [**HtmlPost**](docs/TranslationAPI.md#htmlpost) | **Post** /html | Translate HTML files
*TranslationAPI* | [**HugoGet**](docs/TranslationAPI.md#hugoget) | **Get** /hugo | Get hugo syntax structure from markdown file
*TranslationAPI* | [**HugoPost**](docs/TranslationAPI.md#hugopost) | **Post** /hugo | Run hugo syntax structure analyzing from markdown file
*TranslationAPI* | [**ImageToFilePost**](docs/TranslationAPI.md#imagetofilepost) | **Post** /image-to-file | Translate image or scanned pdf and return file
*TranslationAPI* | [**ImageToTextPost**](docs/TranslationAPI.md#imagetotextpost) | **Post** /image-to-text | Translate text on image or scanned pdf
*TranslationAPI* | [**LanguagesGet**](docs/TranslationAPI.md#languagesget) | **Get** /languages | Return list of available language pairs
*TranslationAPI* | [**MarkdownPost**](docs/TranslationAPI.md#markdownpost) | **Post** /markdown | Translate Markdown files
*TranslationAPI* | [**PdfPost**](docs/TranslationAPI.md#pdfpost) | **Post** /pdf | Translate pdf files
*TranslationAPI* | [**PdfTrialPost**](docs/TranslationAPI.md#pdftrialpost) | **Post** /pdf/trial | Trial pdf translation. Translate only first page without conversion to another format.
*TranslationAPI* | [**PresentationPost**](docs/TranslationAPI.md#presentationpost) | **Post** /presentation | Translate Microsoft PowerPoint presentations, odp
*TranslationAPI* | [**ResxPost**](docs/TranslationAPI.md#resxpost) | **Post** /resx | Translate Resx files
*TranslationAPI* | [**SpreadsheetPost**](docs/TranslationAPI.md#spreadsheetpost) | **Post** /spreadsheet | Translate Microsoft Excel workbooks, ods
*TranslationAPI* | [**SrtPost**](docs/TranslationAPI.md#srtpost) | **Post** /srt | Translate Srt files
*TranslationAPI* | [**TextPost**](docs/TranslationAPI.md#textpost) | **Post** /text | Translate text
*TranslationAPI* | [**TextRequestIdGet**](docs/TranslationAPI.md#textrequestidget) | **Get** /text/{requestId} | Return text translation status.  Also return translated text if translation was successful
*TranslationAPI* | [**TextTrialGet**](docs/TranslationAPI.md#texttrialget) | **Get** /text/trial | Return text translation status for trial requests.  Also return translated text if translation was successful
*TranslationAPI* | [**TextTrialPost**](docs/TranslationAPI.md#texttrialpost) | **Post** /text/trial | Trial translate text. Translate only 1000 symbols.


## Documentation For Models

 - [CloudFileRequest](docs/CloudFileRequest.md)
 - [CloudFileResponse](docs/CloudFileResponse.md)
 - [CloudFileResponseWithAdditionalInfo](docs/CloudFileResponseWithAdditionalInfo.md)
 - [CloudHugoResponse](docs/CloudHugoResponse.md)
 - [CloudTextResponse](docs/CloudTextResponse.md)
 - [CsvFileRequest](docs/CsvFileRequest.md)
 - [FileRequest](docs/FileRequest.md)
 - [HealthCheckEntity](docs/HealthCheckEntity.md)
 - [HealthCheckStatus](docs/HealthCheckStatus.md)
 - [HtmlFileRequest](docs/HtmlFileRequest.md)
 - [HttpStatusCode](docs/HttpStatusCode.md)
 - [HugoRequest](docs/HugoRequest.md)
 - [ImageToFileRequest](docs/ImageToFileRequest.md)
 - [ImageToTextRequest](docs/ImageToTextRequest.md)
 - [LanguagePairData](docs/LanguagePairData.md)
 - [MarkdownFileRequest](docs/MarkdownFileRequest.md)
 - [PdfFileRequest](docs/PdfFileRequest.md)
 - [PresentationFileRequest](docs/PresentationFileRequest.md)
 - [SpreadsheetFileRequest](docs/SpreadsheetFileRequest.md)
 - [SrtFileRequest](docs/SrtFileRequest.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [StringStringTuple](docs/StringStringTuple.md)
 - [TextDocumentFileRequest](docs/TextDocumentFileRequest.md)
 - [TextRequest](docs/TextRequest.md)
 - [UrlFileInfo](docs/UrlFileInfo.md)
 - [WorksheetData](docs/WorksheetData.md)


## Documentation For Authorization



### JWT


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: https://id.groupdocs.cloud/
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

anton.perhunov@aspose.com

