# \TransportApi

All URIs are relative to *https://api.groupdocs.cloud/v2.0/translation*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoPost**](TransportApi.md#AutoPost) | **Post** /auto | Translate any supported file
[**CsvPost**](TransportApi.md#CsvPost) | **Post** /csv | Translate CSV and TSV files
[**DocumentPost**](TransportApi.md#DocumentPost) | **Post** /document | Translate Microsoft Word documents, rtf, txt, odt
[**DocumentRequestIdGet**](TransportApi.md#DocumentRequestIdGet) | **Get** /document/{requestId} | Return document translation status.  Also return URLs for downloading of translated document if translation was successful
[**HcGet**](TransportApi.md#HcGet) | **Get** /hc | Health check for all services.
[**HtmlPost**](TransportApi.md#HtmlPost) | **Post** /html | Translate HTML files
[**HugoGet**](TransportApi.md#HugoGet) | **Get** /hugo | Get hugo syntax structure from markdown file
[**HugoPost**](TransportApi.md#HugoPost) | **Post** /hugo | Run hugo syntax structure analyzing from markdown file
[**ImageToFilePost**](TransportApi.md#ImageToFilePost) | **Post** /image-to-file | Translate image or scanned pdf and return file
[**ImageToTextPost**](TransportApi.md#ImageToTextPost) | **Post** /image-to-text | Translate text on image or scanned pdf
[**LanguagesGet**](TransportApi.md#LanguagesGet) | **Get** /languages | Return list of available language pairs
[**MarkdownPost**](TransportApi.md#MarkdownPost) | **Post** /markdown | Translate Markdown files
[**PdfPost**](TransportApi.md#PdfPost) | **Post** /pdf | Translate pdf files
[**PresentationPost**](TransportApi.md#PresentationPost) | **Post** /presentation | Translate Microsoft PowerPoint presentations, odp
[**ResxPost**](TransportApi.md#ResxPost) | **Post** /resx | Translate RESX files
[**SpreadsheetsPost**](TransportApi.md#SpreadsheetsPost) | **Post** /spreadsheets | Translate Microsoft Excel workbooks, ods
[**TextPost**](TransportApi.md#TextPost) | **Post** /text | Translate text
[**TextRequestIdGet**](TransportApi.md#TextRequestIdGet) | **Get** /text/{requestId} | Return text translation status.  Also return translated text if translation was successful



## AutoPost

> StatusResponse AutoPost(ctx).AutoPostRequest(autoPostRequest).Execute()

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
    autoPostRequest := *openapiclient.NewAutoPostRequest("Unknown", "OutputFormat_example", "SourceLanguage_example", []string{"TargetLanguages_example"}) // AutoPostRequest |  (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportApi.AutoPost(context.Background()).AutoPostRequest(autoPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.AutoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.AutoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoPostRequest** | [**AutoPostRequest**](AutoPostRequest.md) |  | 

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
    resp, r, err := apiClient.TransportApi.CsvPost(context.Background()).CsvFileRequest(csvFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.CsvPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CsvPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.CsvPost`: %v\n", resp)
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
    resp, r, err := apiClient.TransportApi.DocumentPost(context.Background()).TextDocumentFileRequest(textDocumentFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.DocumentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocumentPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.DocumentPost`: %v\n", resp)
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
    resp, r, err := apiClient.TransportApi.DocumentRequestIdGet(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.DocumentRequestIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocumentRequestIdGet`: CloudFileResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.DocumentRequestIdGet`: %v\n", resp)
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
    resp, r, err := apiClient.TransportApi.HcGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.HcGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HcGet`: HealthCheckStatus
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.HcGet`: %v\n", resp)
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

> StatusResponse HtmlPost(ctx).SourceLanguage(sourceLanguage).TargetLanguages(targetLanguages).File(file).OriginalFileName(originalFileName).Url(url).Origin(origin).SavingMode(savingMode).OutputFormat(outputFormat).Execute()

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
    sourceLanguage := "sourceLanguage_example" // string | Language of original file (optional) (default to "en")
    targetLanguages := []string{"Inner_example"} // []string | List of target languages (optional)
    file := string(BYTE_ARRAY_DATA_HERE) // string | File as byte array (optional)
    originalFileName := "originalFileName_example" // string | Type in the file name. If null will be as request ID. (optional)
    url := "url_example" // string | Link to file for translation. Ignore, if \\\"file\\\" property not null (optional)
    origin := "origin_example" // string | Url or name of application using this SDK. Not required. (optional)
    savingMode := "savingMode_example" // string | Toggle file saving mode for storage.  Is Files by default. (optional)
    outputFormat := "outputFormat_example" // string | output file format (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportApi.HtmlPost(context.Background()).SourceLanguage(sourceLanguage).TargetLanguages(targetLanguages).File(file).OriginalFileName(originalFileName).Url(url).Origin(origin).SavingMode(savingMode).OutputFormat(outputFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.HtmlPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HtmlPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.HtmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHtmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLanguage** | **string** | Language of original file | [default to &quot;en&quot;]
 **targetLanguages** | **[]string** | List of target languages | 
 **file** | **string** | File as byte array | 
 **originalFileName** | **string** | Type in the file name. If null will be as request ID. | 
 **url** | **string** | Link to file for translation. Ignore, if \\\&quot;file\\\&quot; property not null | 
 **origin** | **string** | Url or name of application using this SDK. Not required. | 
 **savingMode** | **string** | Toggle file saving mode for storage.  Is Files by default. | 
 **outputFormat** | **string** | output file format | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
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
    resp, r, err := apiClient.TransportApi.HugoGet(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.HugoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HugoGet`: CloudHugoResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.HugoGet`: %v\n", resp)
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

> StatusResponse HugoPost(ctx).File(file).Url(url).Execute()

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
    file := os.NewFile(1234, "some_file") // *os.File | File as byte array
    url := "url_example" // string | Link to file for translation (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportApi.HugoPost(context.Background()).File(file).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.HugoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HugoPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.HugoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHugoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | File as byte array | 
 **url** | **string** | Link to file for translation | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageToFilePost

> StatusResponse ImageToFilePost(ctx).OcrFileRequest(ocrFileRequest).Execute()

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
    ocrFileRequest := *openapiclient.NewOcrFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Ocrformat_example", "OutputFormat_example") // OcrFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportApi.ImageToFilePost(context.Background()).OcrFileRequest(ocrFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.ImageToFilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageToFilePost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.ImageToFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageToFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocrFileRequest** | [**OcrFileRequest**](OcrFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

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

> StatusResponse ImageToTextPost(ctx).Source(source).Targets(targets).Format(format).Url(url).Rotate(rotate).IsHandwritten(isHandwritten).Origin(origin).Route(route).File(file).Execute()

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
    source := "source_example" // string | Language of original file (default to "en")
    targets := []string{"Inner_example"} // []string | List of target languages
    format := "format_example" // string | Originnal file format (optional) (default to "Unknown")
    url := "url_example" // string | Link to file for translation (optional)
    rotate := int32(56) // int32 | Left to write angle to rotate scanned image / pdf (optional)
    isHandwritten := true // bool | is handwritten text (optional)
    origin := "origin_example" // string | for analysis only (optional)
    route := "route_example" // string | endpoints route (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportApi.ImageToTextPost(context.Background()).Source(source).Targets(targets).Format(format).Url(url).Rotate(rotate).IsHandwritten(isHandwritten).Origin(origin).Route(route).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.ImageToTextPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageToTextPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.ImageToTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageToTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | Language of original file | [default to &quot;en&quot;]
 **targets** | **[]string** | List of target languages | 
 **format** | **string** | Originnal file format | [default to &quot;Unknown&quot;]
 **url** | **string** | Link to file for translation | 
 **rotate** | **int32** | Left to write angle to rotate scanned image / pdf | 
 **isHandwritten** | **bool** | is handwritten text | 
 **origin** | **string** | for analysis only | 
 **route** | **string** | endpoints route | 
 **file** | ***os.File** |  | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
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
    resp, r, err := apiClient.TransportApi.LanguagesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.LanguagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LanguagesGet`: []LanguagePairData
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.LanguagesGet`: %v\n", resp)
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

> StatusResponse MarkdownPost(ctx).SourceLanguage(sourceLanguage).TargetLanguages(targetLanguages).OutputFormat(outputFormat).File(file).OriginalFileName(originalFileName).Url(url).Origin(origin).SavingMode(savingMode).FrontMatterList(frontMatterList).Execute()

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
    sourceLanguage := "sourceLanguage_example" // string | Language of original file (default to "en")
    targetLanguages := []string{"Inner_example"} // []string | List of target languages
    outputFormat := "outputFormat_example" // string | output file format
    file := string(BYTE_ARRAY_DATA_HERE) // string | File as byte array (optional)
    originalFileName := "originalFileName_example" // string | Type in the file name. If null will be as request ID. (optional)
    url := "url_example" // string | Link to file for translation. Ignore, if \\\"file\\\" property not null (optional)
    origin := "origin_example" // string | Url or name of application using this SDK. Not required. (optional)
    savingMode := "savingMode_example" // string | Toggle file saving mode for storage.  Is Files by default. (optional)
    frontMatterList := [][]string{[]string{"Inner_example"}} // [][]string | List of lists of frontmatter paths (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportApi.MarkdownPost(context.Background()).SourceLanguage(sourceLanguage).TargetLanguages(targetLanguages).OutputFormat(outputFormat).File(file).OriginalFileName(originalFileName).Url(url).Origin(origin).SavingMode(savingMode).FrontMatterList(frontMatterList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.MarkdownPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkdownPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.MarkdownPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLanguage** | **string** | Language of original file | [default to &quot;en&quot;]
 **targetLanguages** | **[]string** | List of target languages | 
 **outputFormat** | **string** | output file format | 
 **file** | **string** | File as byte array | 
 **originalFileName** | **string** | Type in the file name. If null will be as request ID. | 
 **url** | **string** | Link to file for translation. Ignore, if \\\&quot;file\\\&quot; property not null | 
 **origin** | **string** | Url or name of application using this SDK. Not required. | 
 **savingMode** | **string** | Toggle file saving mode for storage.  Is Files by default. | 
 **frontMatterList** | **[][]string** | List of lists of frontmatter paths | 

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
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
    resp, r, err := apiClient.TransportApi.PdfPost(context.Background()).PdfFileRequest(pdfFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.PdfPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PdfPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.PdfPost`: %v\n", resp)
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
    resp, r, err := apiClient.TransportApi.PresentationPost(context.Background()).PresentationFileRequest(presentationFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.PresentationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PresentationPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.PresentationPost`: %v\n", resp)
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
    resp, r, err := apiClient.TransportApi.ResxPost(context.Background()).ResxFileRequest(resxFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.ResxPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResxPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.ResxPost`: %v\n", resp)
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


## SpreadsheetsPost

> StatusResponse SpreadsheetsPost(ctx).SpreadsheetFileRequest(spreadsheetFileRequest).Execute()

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
    resp, r, err := apiClient.TransportApi.SpreadsheetsPost(context.Background()).SpreadsheetFileRequest(spreadsheetFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.SpreadsheetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpreadsheetsPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.SpreadsheetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpreadsheetsPostRequest struct via the builder pattern


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
    resp, r, err := apiClient.TransportApi.TextPost(context.Background()).TextRequest(textRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.TextPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TextPost`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.TextPost`: %v\n", resp)
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
    resp, r, err := apiClient.TransportApi.TextRequestIdGet(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportApi.TextRequestIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TextRequestIdGet`: CloudTextResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportApi.TextRequestIdGet`: %v\n", resp)
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

