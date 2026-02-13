# \SegmentEffortsAPI

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEffortsBySegmentId**](SegmentEffortsAPI.md#GetEffortsBySegmentId) | **Get** /segment_efforts | List Segment Efforts
[**GetSegmentEffortById**](SegmentEffortsAPI.md#GetSegmentEffortById) | **Get** /segment_efforts/{id} | Get Segment Effort



## GetEffortsBySegmentId

> []DetailedSegmentEffort GetEffortsBySegmentId(ctx).SegmentId(segmentId).StartDateLocal(startDateLocal).EndDateLocal(endDateLocal).PerPage(perPage).Execute()

List Segment Efforts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/alexhokl/strava-cli/swagger"
)

func main() {
	segmentId := int32(56) // int32 | The identifier of the segment.
	startDateLocal := time.Now() // time.Time | ISO 8601 formatted date time. (optional)
	endDateLocal := time.Now() // time.Time | ISO 8601 formatted date time. (optional)
	perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentEffortsAPI.GetEffortsBySegmentId(context.Background()).SegmentId(segmentId).StartDateLocal(startDateLocal).EndDateLocal(endDateLocal).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentEffortsAPI.GetEffortsBySegmentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEffortsBySegmentId`: []DetailedSegmentEffort
	fmt.Fprintf(os.Stdout, "Response from `SegmentEffortsAPI.GetEffortsBySegmentId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEffortsBySegmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **segmentId** | **int32** | The identifier of the segment. | 
 **startDateLocal** | **time.Time** | ISO 8601 formatted date time. | 
 **endDateLocal** | **time.Time** | ISO 8601 formatted date time. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]DetailedSegmentEffort**](DetailedSegmentEffort.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentEffortById

> DetailedSegmentEffort GetSegmentEffortById(ctx, id).Execute()

Get Segment Effort



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
	id := int64(789) // int64 | The identifier of the segment effort.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentEffortsAPI.GetSegmentEffortById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentEffortsAPI.GetSegmentEffortById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSegmentEffortById`: DetailedSegmentEffort
	fmt.Fprintf(os.Stdout, "Response from `SegmentEffortsAPI.GetSegmentEffortById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the segment effort. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentEffortByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetailedSegmentEffort**](DetailedSegmentEffort.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

