# \FileAPI

All URIs are relative to *https://api.groupdocs.cloud/v2.0/translation*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileUploadPost**](FileAPI.md#FileUploadPost) | **Post** /file/upload | 



## FileUploadPost

> string FileUploadPost(ctx).Format(format).File(file).Execute()



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
	format := "format_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPI.FileUploadPost(context.Background()).Format(format).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.FileUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileUploadPost`: string
	fmt.Fprintf(os.Stdout, "Response from `FileAPI.FileUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFileUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

