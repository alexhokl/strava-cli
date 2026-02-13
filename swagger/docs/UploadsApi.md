# \UploadsAPI

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUpload**](UploadsAPI.md#CreateUpload) | **Post** /uploads | Upload Activity
[**GetUploadById**](UploadsAPI.md#GetUploadById) | **Get** /uploads/{uploadId} | Get Upload



## CreateUpload

> Upload CreateUpload(ctx).File(file).Name(name).Description(description).Trainer(trainer).Commute(commute).DataType(dataType).ExternalId(externalId).Execute()

Upload Activity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/strava-cli/swagger"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | The uploaded file. (optional)
	name := "name_example" // string | The desired name of the resulting activity. (optional)
	description := "description_example" // string | The desired description of the resulting activity. (optional)
	trainer := "trainer_example" // string | Whether the resulting activity should be marked as having been performed on a trainer. (optional)
	commute := "commute_example" // string | Whether the resulting activity should be tagged as a commute. (optional)
	dataType := "dataType_example" // string | The format of the uploaded file. (optional)
	externalId := "externalId_example" // string | The desired external identifier of the resulting activity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadsAPI.CreateUpload(context.Background()).File(file).Name(name).Description(description).Trainer(trainer).Commute(commute).DataType(dataType).ExternalId(externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadsAPI.CreateUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpload`: Upload
	fmt.Fprintf(os.Stdout, "Response from `UploadsAPI.CreateUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The uploaded file. | 
 **name** | **string** | The desired name of the resulting activity. | 
 **description** | **string** | The desired description of the resulting activity. | 
 **trainer** | **string** | Whether the resulting activity should be marked as having been performed on a trainer. | 
 **commute** | **string** | Whether the resulting activity should be tagged as a commute. | 
 **dataType** | **string** | The format of the uploaded file. | 
 **externalId** | **string** | The desired external identifier of the resulting activity. | 

### Return type

[**Upload**](Upload.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadById

> Upload GetUploadById(ctx, uploadId).Execute()

Get Upload



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/strava-cli/swagger"
)

func main() {
	uploadId := int64(789) // int64 | The identifier of the upload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadsAPI.GetUploadById(context.Background(), uploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadsAPI.GetUploadById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUploadById`: Upload
	fmt.Fprintf(os.Stdout, "Response from `UploadsAPI.GetUploadById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadId** | **int64** | The identifier of the upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Upload**](Upload.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

