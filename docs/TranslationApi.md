# \TranslationApi

All URIs are relative to *https://api.groupdocs.cloud/v2.0/translation*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoPost**](TranslationApi.md#AutoPost) | **Post** /auto | Translate any supported file
[**CsvPost**](TranslationApi.md#CsvPost) | **Post** /csv | Translate CSV and TSV files
[**DocumentPost**](TranslationApi.md#DocumentPost) | **Post** /document | Translate Microsoft Word documents, rtf, txt, odt
[**DocumentRequestIdGet**](TranslationApi.md#DocumentRequestIdGet) | **Get** /document/{requestId} | Return document translation status.  Also return URLs for downloading of translated document if translation was successful
[**DocumentTrialPost**](TranslationApi.md#DocumentTrialPost) | **Post** /document/trial | Trial translate Microsoft Word documents, rtf, txt, odt without conversation. Translate only first page or 1000 symbols.
[**HcGet**](TranslationApi.md#HcGet) | **Get** /hc | Health check for all services.
[**HtmlPost**](TranslationApi.md#HtmlPost) | **Post** /html | Translate HTML files
[**HugoGet**](TranslationApi.md#HugoGet) | **Get** /hugo | Get hugo syntax structure from markdown file
[**HugoPost**](TranslationApi.md#HugoPost) | **Post** /hugo | Run hugo syntax structure analyzing from markdown file
[**ImageToFilePost**](TranslationApi.md#ImageToFilePost) | **Post** /image-to-file | Translate image or scanned pdf and return file
[**ImageToTextPost**](TranslationApi.md#ImageToTextPost) | **Post** /image-to-text | Translate text on image or scanned pdf
[**LanguagesGet**](TranslationApi.md#LanguagesGet) | **Get** /languages | Return list of available language pairs
[**MarkdownPost**](TranslationApi.md#MarkdownPost) | **Post** /markdown | Translate Markdown files
[**PdfPost**](TranslationApi.md#PdfPost) | **Post** /pdf | Translate pdf files
[**PdfTrialPost**](TranslationApi.md#PdfTrialPost) | **Post** /pdf/trial | Trial pdf translation. Translate only first page without conversion to another format.
[**PresentationPost**](TranslationApi.md#PresentationPost) | **Post** /presentation | Translate Microsoft PowerPoint presentations, odp
[**ResxPost**](TranslationApi.md#ResxPost) | **Post** /resx | Translate RESX files
[**SpreadsheetPost**](TranslationApi.md#SpreadsheetPost) | **Post** /spreadsheet | Translate Microsoft Excel workbooks, ods
[**TextPost**](TranslationApi.md#TextPost) | **Post** /text | Translate text
[**TextRequestIdGet**](TranslationApi.md#TextRequestIdGet) | **Get** /text/{requestId} | Return text translation status.  Also return translated text if translation was successful
[**TextTrialPost**](TranslationApi.md#TextTrialPost) | **Post** /text/trial | Trial translate text. Translate only 1000 symbols.



## AutoPost

> StatusResponse AutoPost(ctx).FileRequest(fileRequest).Execute()

Translate any supported file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    fileRequest := *openapiclient.NewFileRequest() // FileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.AutoPost(context.Background()).FileRequest(fileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.AutoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.AutoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileRequest** | [**FileRequest**](FileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CsvPost

> StatusResponse CsvPost(ctx).CsvFileRequest(csvFileRequest).Execute()

Translate CSV and TSV files

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    csvFileRequest := *openapiclient.NewCsvFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "OutFormat_example") // CsvFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.CsvPost(context.Background()).CsvFileRequest(csvFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.CsvPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CsvPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.CsvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCsvPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **csvFileRequest** | [**CsvFileRequest**](CsvFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentPost

> StatusResponse DocumentPost(ctx).TextDocumentFileRequest(textDocumentFileRequest).Execute()

Translate Microsoft Word documents, rtf, txt, odt

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    textDocumentFileRequest := *openapiclient.NewTextDocumentFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Unknown", "OutputFormat_example") // TextDocumentFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.DocumentPost(context.Background()).TextDocumentFileRequest(textDocumentFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.DocumentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocumentPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.DocumentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textDocumentFileRequest** | [**TextDocumentFileRequest**](TextDocumentFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentRequestIdGet

> CloudFileResponse DocumentRequestIdGet(ctx, requestId).Execute()

Return document translation status.  Also return URLs for downloading of translated document if translation was successful

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    requestId := "requestId_example" // string | GUID which got from /v3/translation/document response

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.DocumentRequestIdGet(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.DocumentRequestIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocumentRequestIdGet`: CloudFileResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.DocumentRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | GUID which got from /v3/translation/document response | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumentRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudFileResponse**](CloudFileResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentTrialPost

> StatusResponse DocumentTrialPost(ctx).TextDocumentFileRequest(textDocumentFileRequest).Execute()

Trial translate Microsoft Word documents, rtf, txt, odt without conversation. Translate only first page or 1000 symbols.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    textDocumentFileRequest := *openapiclient.NewTextDocumentFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Unknown", "OutputFormat_example") // TextDocumentFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.DocumentTrialPost(context.Background()).TextDocumentFileRequest(textDocumentFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.DocumentTrialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocumentTrialPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.DocumentTrialPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentTrialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textDocumentFileRequest** | [**TextDocumentFileRequest**](TextDocumentFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HcGet

> HealthCheckStatus HcGet(ctx).Execute()

Health check for all services.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.HcGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.HcGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HcGet`: HealthCheckStatus
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.HcGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHcGetRequest struct via the builder pattern


### Return type

[**HealthCheckStatus**](HealthCheckStatus.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HtmlPost

> StatusResponse HtmlPost(ctx).HtmlFileRequest(htmlFileRequest).Execute()

Translate HTML files

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    htmlFileRequest := *openapiclient.NewHtmlFileRequest() // HtmlFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.HtmlPost(context.Background()).HtmlFileRequest(htmlFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.HtmlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HtmlPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.HtmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHtmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **htmlFileRequest** | [**HtmlFileRequest**](HtmlFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HugoGet

> CloudHugoResponse HugoGet(ctx).Id(id).Execute()

Get hugo syntax structure from markdown file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    id := "id_example" // string | id from PostHugo> (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.HugoGet(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.HugoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HugoGet`: CloudHugoResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.HugoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHugoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | id from PostHugo&gt; | 

### Return type

[**CloudHugoResponse**](CloudHugoResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HugoPost

> StatusResponse HugoPost(ctx).HugoRequest(hugoRequest).Execute()

Run hugo syntax structure analyzing from markdown file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    hugoRequest := *openapiclient.NewHugoRequest() // HugoRequest |  (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.HugoPost(context.Background()).HugoRequest(hugoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.HugoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HugoPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.HugoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHugoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hugoRequest** | [**HugoRequest**](HugoRequest.md) |  | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageToFilePost

> StatusResponse ImageToFilePost(ctx).ImageToFileRequest(imageToFileRequest).Execute()

Translate image or scanned pdf and return file

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    imageToFileRequest := *openapiclient.NewImageToFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Ocrformat_example", "OutputFormat_example") // ImageToFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.ImageToFilePost(context.Background()).ImageToFileRequest(imageToFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.ImageToFilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageToFilePost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.ImageToFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageToFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageToFileRequest** | [**ImageToFileRequest**](ImageToFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageToTextPost

> StatusResponse ImageToTextPost(ctx).ImageToTextRequest(imageToTextRequest).Execute()

Translate text on image or scanned pdf

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    imageToTextRequest := *openapiclient.NewImageToTextRequest() // ImageToTextRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.ImageToTextPost(context.Background()).ImageToTextRequest(imageToTextRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.ImageToTextPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageToTextPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.ImageToTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageToTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageToTextRequest** | [**ImageToTextRequest**](ImageToTextRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LanguagesGet

> []LanguagePairData LanguagesGet(ctx).Execute()

Return list of available language pairs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.LanguagesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.LanguagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LanguagesGet`: []LanguagePairData
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.LanguagesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLanguagesGetRequest struct via the builder pattern


### Return type

[**[]LanguagePairData**](LanguagePairData.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkdownPost

> StatusResponse MarkdownPost(ctx).MarkdownFileRequest(markdownFileRequest).Execute()

Translate Markdown files

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    markdownFileRequest := *openapiclient.NewMarkdownFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "OutputFormat_example") // MarkdownFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.MarkdownPost(context.Background()).MarkdownFileRequest(markdownFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.MarkdownPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkdownPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.MarkdownPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markdownFileRequest** | [**MarkdownFileRequest**](MarkdownFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PdfPost

> StatusResponse PdfPost(ctx).PdfFileRequest(pdfFileRequest).Execute()

Translate pdf files

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    pdfFileRequest := *openapiclient.NewPdfFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "OutputFormat_example") // PdfFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.PdfPost(context.Background()).PdfFileRequest(pdfFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.PdfPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PdfPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.PdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFileRequest** | [**PdfFileRequest**](PdfFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PdfTrialPost

> StatusResponse PdfTrialPost(ctx).PdfFileRequest(pdfFileRequest).Execute()

Trial pdf translation. Translate only first page without conversion to another format.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    pdfFileRequest := *openapiclient.NewPdfFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "OutputFormat_example") // PdfFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.PdfTrialPost(context.Background()).PdfFileRequest(pdfFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.PdfTrialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PdfTrialPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.PdfTrialPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPdfTrialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfFileRequest** | [**PdfFileRequest**](PdfFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PresentationPost

> StatusResponse PresentationPost(ctx).PresentationFileRequest(presentationFileRequest).Execute()

Translate Microsoft PowerPoint presentations, odp

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    presentationFileRequest := *openapiclient.NewPresentationFileRequest() // PresentationFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.PresentationPost(context.Background()).PresentationFileRequest(presentationFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.PresentationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PresentationPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.PresentationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPresentationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presentationFileRequest** | [**PresentationFileRequest**](PresentationFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResxPost

> StatusResponse ResxPost(ctx).ResxFileRequest(resxFileRequest).Execute()

Translate RESX files

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    resxFileRequest := *openapiclient.NewResxFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}) // ResxFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.ResxPost(context.Background()).ResxFileRequest(resxFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.ResxPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResxPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.ResxPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resxFileRequest** | [**ResxFileRequest**](ResxFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpreadsheetPost

> StatusResponse SpreadsheetPost(ctx).SpreadsheetFileRequest(spreadsheetFileRequest).Execute()

Translate Microsoft Excel workbooks, ods

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    spreadsheetFileRequest := *openapiclient.NewSpreadsheetFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Unknown", "OutputFormat_example") // SpreadsheetFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.SpreadsheetPost(context.Background()).SpreadsheetFileRequest(spreadsheetFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.SpreadsheetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpreadsheetPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.SpreadsheetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpreadsheetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spreadsheetFileRequest** | [**SpreadsheetFileRequest**](SpreadsheetFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TextPost

> StatusResponse TextPost(ctx).TextRequest(textRequest).Execute()

Translate text

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.TextPost(context.Background()).TextRequest(textRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.TextPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TextPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.TextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textRequest** | [**TextRequest**](TextRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TextRequestIdGet

> CloudTextResponse TextRequestIdGet(ctx, requestId).Execute()

Return text translation status.  Also return translated text if translation was successful

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    requestId := "requestId_example" // string | GUID which got from /v3/translation/text response

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.TextRequestIdGet(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.TextRequestIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TextRequestIdGet`: CloudTextResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.TextRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | GUID which got from /v3/translation/text response | 

### Other Parameters

Other parameters are passed through a pointer to a apiTextRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudTextResponse**](CloudTextResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TextTrialPost

> StatusResponse TextTrialPost(ctx).TextRequest(textRequest).Execute()

Trial translate text. Translate only 1000 symbols.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.TextTrialPost(context.Background()).TextRequest(textRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.TextTrialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TextTrialPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.TextTrialPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTextTrialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textRequest** | [**TextRequest**](TextRequest.md) | String in body of request, containing JSON with parameters for translation. | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

