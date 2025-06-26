# \TranslationAPI

All URIs are relative to *https://api.groupdocs.cloud/v2.0/translation*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoClassificationPost**](TranslationAPI.md#AutoClassificationPost) | **Post** /auto/classification | Domain Ttanslation of any supported file based on classification
[**AutoPost**](TranslationAPI.md#AutoPost) | **Post** /auto | Translate any supported file
[**AutoSlangPost**](TranslationAPI.md#AutoSlangPost) | **Post** /auto/slang | Slang translation of any supported file
[**CsvPost**](TranslationAPI.md#CsvPost) | **Post** /csv | Translate CSV and TSV files
[**DocumentPost**](TranslationAPI.md#DocumentPost) | **Post** /document | Translate Microsoft Word documents, rtf, txt, odt
[**DocumentRequestIdGet**](TranslationAPI.md#DocumentRequestIdGet) | **Get** /document/{requestId} | Return document translation status.  Also return URLs for downloading of translated document if translation was successful
[**DocumentTrialGet**](TranslationAPI.md#DocumentTrialGet) | **Get** /document/trial | Return document translation status for trial request.  Also return URLs for downloading of translated document if translation was successful
[**DocumentTrialPost**](TranslationAPI.md#DocumentTrialPost) | **Post** /document/trial | Trial translate Microsoft Word documents, rtf, txt, odt without conversation. Translate only first page or 1000 symbols.
[**HcGet**](TranslationAPI.md#HcGet) | **Get** /hc | Health check for all services.
[**HtmlPost**](TranslationAPI.md#HtmlPost) | **Post** /html | Translate HTML files
[**HugoGet**](TranslationAPI.md#HugoGet) | **Get** /hugo | Get hugo syntax structure from markdown file
[**HugoPost**](TranslationAPI.md#HugoPost) | **Post** /hugo | Run hugo syntax structure analyzing from markdown file
[**ImageToFilePost**](TranslationAPI.md#ImageToFilePost) | **Post** /image-to-file | Translate image or scanned pdf and return file
[**ImageToTextPost**](TranslationAPI.md#ImageToTextPost) | **Post** /image-to-text | Translate text on image or scanned pdf
[**LanguagesGet**](TranslationAPI.md#LanguagesGet) | **Get** /languages | Return list of available language pairs
[**MarkdownPost**](TranslationAPI.md#MarkdownPost) | **Post** /markdown | Translate Markdown files
[**MediaToFilePost**](TranslationAPI.md#MediaToFilePost) | **Post** /media-to-file | Translate audio or video and return file as a result
[**MediaToTextPost**](TranslationAPI.md#MediaToTextPost) | **Post** /media-to-text | Translate audio or video and return text as a result
[**PdfPost**](TranslationAPI.md#PdfPost) | **Post** /pdf | Translate pdf files
[**PdfTrialPost**](TranslationAPI.md#PdfTrialPost) | **Post** /pdf/trial | Trial pdf translation. Translate only first page without conversion to another format.
[**PresentationPost**](TranslationAPI.md#PresentationPost) | **Post** /presentation | Translate Microsoft PowerPoint presentations, odp
[**ResxPost**](TranslationAPI.md#ResxPost) | **Post** /resx | Translate Resx files
[**SpreadsheetPost**](TranslationAPI.md#SpreadsheetPost) | **Post** /spreadsheet | Translate Microsoft Excel workbooks, ods
[**SrtPost**](TranslationAPI.md#SrtPost) | **Post** /srt | Translate Srt files
[**SummarizeDocumentPost**](TranslationAPI.md#SummarizeDocumentPost) | **Post** /summarize-document | Summarize and translate document file
[**SummarizeImagePost**](TranslationAPI.md#SummarizeImagePost) | **Post** /summarize-image | Summarize and translate image or scanned pdf and return file
[**SummarizeMediaPost**](TranslationAPI.md#SummarizeMediaPost) | **Post** /summarize-media | Summarize audio or video, translate it  and return file as a result
[**SummarizeTextPost**](TranslationAPI.md#SummarizeTextPost) | **Post** /summarize-text | Summarize and translate text
[**TextClassificationPost**](TranslationAPI.md#TextClassificationPost) | **Post** /text/classification | Domain text translation based on classification
[**TextMedicalPost**](TranslationAPI.md#TextMedicalPost) | **Post** /text/medical | Translate medical
[**TextPost**](TranslationAPI.md#TextPost) | **Post** /text | Translate text
[**TextRequestIdGet**](TranslationAPI.md#TextRequestIdGet) | **Get** /text/{requestId} | Return text translation status.  Also return translated text if translation was successful
[**TextSlangPost**](TranslationAPI.md#TextSlangPost) | **Post** /text/slang | Translate slang
[**TextTrialGet**](TranslationAPI.md#TextTrialGet) | **Get** /text/trial | Return text translation status for trial requests.  Also return translated text if translation was successful
[**TextTrialPost**](TranslationAPI.md#TextTrialPost) | **Post** /text/trial | Trial translate text. Translate only 1000 symbols.
[**XmlPost**](TranslationAPI.md#XmlPost) | **Post** /xml | Translate XML files



## AutoClassificationPost

> StatusResponse AutoClassificationPost(ctx).FileRequest(fileRequest).Execute()

Domain Ttanslation of any supported file based on classification

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	fileRequest := *openapiclient.NewFileRequest() // FileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.AutoClassificationPost(context.Background()).FileRequest(fileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.AutoClassificationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoClassificationPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.AutoClassificationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoClassificationPostRequest struct via the builder pattern


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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	fileRequest := *openapiclient.NewFileRequest() // FileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.AutoPost(context.Background()).FileRequest(fileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.AutoPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.AutoPost`: %v\n", resp)
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


## AutoSlangPost

> StatusResponse AutoSlangPost(ctx).FileRequest(fileRequest).Execute()

Slang translation of any supported file

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	fileRequest := *openapiclient.NewFileRequest() // FileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.AutoSlangPost(context.Background()).FileRequest(fileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.AutoSlangPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoSlangPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.AutoSlangPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoSlangPostRequest struct via the builder pattern


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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	csvFileRequest := *openapiclient.NewCsvFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "OutputFormat_example") // CsvFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.CsvPost(context.Background()).CsvFileRequest(csvFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.CsvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CsvPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.CsvPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textDocumentFileRequest := *openapiclient.NewTextDocumentFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Unknown", "OutputFormat_example") // TextDocumentFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.DocumentPost(context.Background()).TextDocumentFileRequest(textDocumentFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.DocumentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.DocumentPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	requestId := "requestId_example" // string | GUID which got from /translation/document response

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.DocumentRequestIdGet(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.DocumentRequestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentRequestIdGet`: CloudFileResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.DocumentRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | GUID which got from /translation/document response | 

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


## DocumentTrialGet

> CloudFileResponse DocumentTrialGet(ctx).RequestId(requestId).Execute()

Return document translation status for trial request.  Also return URLs for downloading of translated document if translation was successful

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	requestId := "requestId_example" // string | GUID which got from /v3/translation/document response (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.DocumentTrialGet(context.Background()).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.DocumentTrialGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentTrialGet`: CloudFileResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.DocumentTrialGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentTrialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | GUID which got from /v3/translation/document response | 

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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textDocumentFileRequest := *openapiclient.NewTextDocumentFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Unknown", "OutputFormat_example") // TextDocumentFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.DocumentTrialPost(context.Background()).TextDocumentFileRequest(textDocumentFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.DocumentTrialPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentTrialPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.DocumentTrialPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.HcGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.HcGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HcGet`: HealthCheckStatus
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.HcGet`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	htmlFileRequest := *openapiclient.NewHtmlFileRequest() // HtmlFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.HtmlPost(context.Background()).HtmlFileRequest(htmlFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.HtmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HtmlPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.HtmlPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	id := "id_example" // string | id from PostHugo> (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.HugoGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.HugoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HugoGet`: CloudHugoResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.HugoGet`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	hugoRequest := *openapiclient.NewHugoRequest() // HugoRequest |  (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.HugoPost(context.Background()).HugoRequest(hugoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.HugoPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HugoPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.HugoPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	imageToFileRequest := *openapiclient.NewImageToFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Ocrformat_example", "OutputFormat_example") // ImageToFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.ImageToFilePost(context.Background()).ImageToFileRequest(imageToFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.ImageToFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageToFilePost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.ImageToFilePost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	imageToTextRequest := *openapiclient.NewImageToTextRequest() // ImageToTextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.ImageToTextPost(context.Background()).ImageToTextRequest(imageToTextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.ImageToTextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageToTextPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.ImageToTextPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.LanguagesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.LanguagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguagesGet`: []LanguagePairData
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.LanguagesGet`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	markdownFileRequest := *openapiclient.NewMarkdownFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "OutputFormat_example") // MarkdownFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.MarkdownPost(context.Background()).MarkdownFileRequest(markdownFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.MarkdownPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkdownPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.MarkdownPost`: %v\n", resp)
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


## MediaToFilePost

> StatusResponse MediaToFilePost(ctx).MediaToFileRequest(mediaToFileRequest).Execute()

Translate audio or video and return file as a result

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	mediaToFileRequest := *openapiclient.NewMediaToFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Unknown", "OutputFormat_example") // MediaToFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.MediaToFilePost(context.Background()).MediaToFileRequest(mediaToFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.MediaToFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaToFilePost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.MediaToFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaToFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaToFileRequest** | [**MediaToFileRequest**](MediaToFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

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


## MediaToTextPost

> StatusResponse MediaToTextPost(ctx).MediaToTextRequest(mediaToTextRequest).Execute()

Translate audio or video and return text as a result

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	mediaToTextRequest := *openapiclient.NewMediaToTextRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Unknown") // MediaToTextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.MediaToTextPost(context.Background()).MediaToTextRequest(mediaToTextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.MediaToTextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaToTextPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.MediaToTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaToTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaToTextRequest** | [**MediaToTextRequest**](MediaToTextRequest.md) | String in body of request, containing JSON with parameters for translation. | 

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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	pdfFileRequest := *openapiclient.NewPdfFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "OutputFormat_example") // PdfFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.PdfPost(context.Background()).PdfFileRequest(pdfFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.PdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PdfPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.PdfPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	pdfFileRequest := *openapiclient.NewPdfFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "OutputFormat_example") // PdfFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.PdfTrialPost(context.Background()).PdfFileRequest(pdfFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.PdfTrialPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PdfTrialPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.PdfTrialPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	presentationFileRequest := *openapiclient.NewPresentationFileRequest() // PresentationFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.PresentationPost(context.Background()).PresentationFileRequest(presentationFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.PresentationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PresentationPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.PresentationPost`: %v\n", resp)
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

> StatusResponse ResxPost(ctx).SrtFileRequest(srtFileRequest).Execute()

Translate Resx files

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	srtFileRequest := *openapiclient.NewSrtFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example") // SrtFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.ResxPost(context.Background()).SrtFileRequest(srtFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.ResxPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResxPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.ResxPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srtFileRequest** | [**SrtFileRequest**](SrtFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	spreadsheetFileRequest := *openapiclient.NewSpreadsheetFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Unknown", "OutputFormat_example") // SpreadsheetFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.SpreadsheetPost(context.Background()).SpreadsheetFileRequest(spreadsheetFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.SpreadsheetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpreadsheetPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.SpreadsheetPost`: %v\n", resp)
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


## SrtPost

> StatusResponse SrtPost(ctx).SrtFileRequest(srtFileRequest).Execute()

Translate Srt files

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	srtFileRequest := *openapiclient.NewSrtFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example") // SrtFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.SrtPost(context.Background()).SrtFileRequest(srtFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.SrtPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SrtPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.SrtPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSrtPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srtFileRequest** | [**SrtFileRequest**](SrtFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

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


## SummarizeDocumentPost

> StatusResponse SummarizeDocumentPost(ctx).FileRequest(fileRequest).Execute()

Summarize and translate document file

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	fileRequest := *openapiclient.NewFileRequest() // FileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.SummarizeDocumentPost(context.Background()).FileRequest(fileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.SummarizeDocumentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummarizeDocumentPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.SummarizeDocumentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeDocumentPostRequest struct via the builder pattern


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


## SummarizeImagePost

> StatusResponse SummarizeImagePost(ctx).ImageToFileRequest(imageToFileRequest).Execute()

Summarize and translate image or scanned pdf and return file

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	imageToFileRequest := *openapiclient.NewImageToFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Ocrformat_example", "OutputFormat_example") // ImageToFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.SummarizeImagePost(context.Background()).ImageToFileRequest(imageToFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.SummarizeImagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummarizeImagePost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.SummarizeImagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeImagePostRequest struct via the builder pattern


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


## SummarizeMediaPost

> StatusResponse SummarizeMediaPost(ctx).MediaToFileRequest(mediaToFileRequest).Execute()

Summarize audio or video, translate it  and return file as a result

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	mediaToFileRequest := *openapiclient.NewMediaToFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example", "Unknown", "OutputFormat_example") // MediaToFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.SummarizeMediaPost(context.Background()).MediaToFileRequest(mediaToFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.SummarizeMediaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummarizeMediaPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.SummarizeMediaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaToFileRequest** | [**MediaToFileRequest**](MediaToFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

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


## SummarizeTextPost

> StatusResponse SummarizeTextPost(ctx).TextRequest(textRequest).Execute()

Summarize and translate text

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.SummarizeTextPost(context.Background()).TextRequest(textRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.SummarizeTextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummarizeTextPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.SummarizeTextPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeTextPostRequest struct via the builder pattern


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


## TextClassificationPost

> StatusResponse TextClassificationPost(ctx).TextRequest(textRequest).Execute()

Domain text translation based on classification

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.TextClassificationPost(context.Background()).TextRequest(textRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.TextClassificationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TextClassificationPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.TextClassificationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTextClassificationPostRequest struct via the builder pattern


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


## TextMedicalPost

> StatusResponse TextMedicalPost(ctx).TextRequest(textRequest).Execute()

Translate medical

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.TextMedicalPost(context.Background()).TextRequest(textRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.TextMedicalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TextMedicalPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.TextMedicalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTextMedicalPostRequest struct via the builder pattern


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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.TextPost(context.Background()).TextRequest(textRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.TextPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TextPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.TextPost`: %v\n", resp)
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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	requestId := "requestId_example" // string | GUID which got from /v3/translation/text response

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.TextRequestIdGet(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.TextRequestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TextRequestIdGet`: CloudTextResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.TextRequestIdGet`: %v\n", resp)
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


## TextSlangPost

> StatusResponse TextSlangPost(ctx).TextRequest(textRequest).Execute()

Translate slang

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.TextSlangPost(context.Background()).TextRequest(textRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.TextSlangPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TextSlangPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.TextSlangPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTextSlangPostRequest struct via the builder pattern


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


## TextTrialGet

> CloudTextResponse TextTrialGet(ctx).RequestId(requestId).Execute()

Return text translation status for trial requests.  Also return translated text if translation was successful

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	requestId := "requestId_example" // string | GUID which got from /v3/translation/text response (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.TextTrialGet(context.Background()).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.TextTrialGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TextTrialGet`: CloudTextResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.TextTrialGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTextTrialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | GUID which got from /v3/translation/text response | 

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
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	textRequest := *openapiclient.NewTextRequest() // TextRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.TextTrialPost(context.Background()).TextRequest(textRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.TextTrialPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TextTrialPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.TextTrialPost`: %v\n", resp)
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


## XmlPost

> StatusResponse XmlPost(ctx).XmlFileRequest(xmlFileRequest).Execute()

Translate XML files

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/groupdocs-translation-cloud/groupdocs-translation-cloud-go"
)

func main() {
	xmlFileRequest := *openapiclient.NewXmlFileRequest("SourceLanguage_example", []string{"TargetLanguages_example"}, "Url_example") // XmlFileRequest | String in body of request, containing JSON with parameters for translation. (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationAPI.XmlPost(context.Background()).XmlFileRequest(xmlFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationAPI.XmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `XmlPost`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TranslationAPI.XmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xmlFileRequest** | [**XmlFileRequest**](XmlFileRequest.md) | String in body of request, containing JSON with parameters for translation. | 

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

