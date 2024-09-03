# \SwaggerAPI

All URIs are relative to *https://api.groupdocs.cloud/v2.0/translation*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwaggerSpecGet**](SwaggerAPI.md#SwaggerSpecGet) | **Get** /swagger/spec | 



## SwaggerSpecGet

> SwaggerSpecGet(ctx).IsYaml(isYaml).SerialaizeAsV2(serialaizeAsV2).Execute()



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
    isYaml := true // bool |  (optional) (default to false)
    serialaizeAsV2 := true // bool |  (optional) (default to false)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    r, err := apiClient.SwaggerAPI.SwaggerSpecGet(context.Background()).IsYaml(isYaml).SerialaizeAsV2(serialaizeAsV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwaggerAPI.SwaggerSpecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwaggerSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isYaml** | **bool** |  | [default to false]
 **serialaizeAsV2** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

