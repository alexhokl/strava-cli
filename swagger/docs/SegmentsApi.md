# \SegmentsAPI

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExploreSegments**](SegmentsAPI.md#ExploreSegments) | **Get** /segments/explore | Explore segments
[**GetLoggedInAthleteStarredSegments**](SegmentsAPI.md#GetLoggedInAthleteStarredSegments) | **Get** /segments/starred | List Starred Segments
[**GetSegmentById**](SegmentsAPI.md#GetSegmentById) | **Get** /segments/{id} | Get Segment
[**StarSegment**](SegmentsAPI.md#StarSegment) | **Put** /segments/{id}/starred | Star Segment



## ExploreSegments

> ExplorerResponse ExploreSegments(ctx).Bounds(bounds).ActivityType(activityType).MinCat(minCat).MaxCat(maxCat).Execute()

Explore segments



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
	bounds := []float32{float32(123)} // []float32 | The latitude and longitude for two points describing a rectangular boundary for the search: [southwest corner latitutde, southwest corner longitude, northeast corner latitude, northeast corner longitude]
	activityType := "activityType_example" // string | Desired activity type. (optional)
	minCat := int32(56) // int32 | The minimum climbing category. (optional)
	maxCat := int32(56) // int32 | The maximum climbing category. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.ExploreSegments(context.Background()).Bounds(bounds).ActivityType(activityType).MinCat(minCat).MaxCat(maxCat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.ExploreSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExploreSegments`: ExplorerResponse
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.ExploreSegments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExploreSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bounds** | **[]float32** | The latitude and longitude for two points describing a rectangular boundary for the search: [southwest corner latitutde, southwest corner longitude, northeast corner latitude, northeast corner longitude] | 
 **activityType** | **string** | Desired activity type. | 
 **minCat** | **int32** | The minimum climbing category. | 
 **maxCat** | **int32** | The maximum climbing category. | 

### Return type

[**ExplorerResponse**](ExplorerResponse.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggedInAthleteStarredSegments

> []SummarySegment GetLoggedInAthleteStarredSegments(ctx).Page(page).PerPage(perPage).Execute()

List Starred Segments



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
	page := int32(56) // int32 | Page number. Defaults to 1. (optional)
	perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.GetLoggedInAthleteStarredSegments(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.GetLoggedInAthleteStarredSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoggedInAthleteStarredSegments`: []SummarySegment
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.GetLoggedInAthleteStarredSegments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggedInAthleteStarredSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]SummarySegment**](SummarySegment.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentById

> DetailedSegment GetSegmentById(ctx, id).Execute()

Get Segment



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
	id := int64(789) // int64 | The identifier of the segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.GetSegmentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.GetSegmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSegmentById`: DetailedSegment
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.GetSegmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetailedSegment**](DetailedSegment.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StarSegment

> DetailedSegment StarSegment(ctx, id).Starred(starred).Execute()

Star Segment



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
	id := int64(789) // int64 | The identifier of the segment to star.
	starred := true // bool | If true, star the segment; if false, unstar the segment. (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsAPI.StarSegment(context.Background(), id).Starred(starred).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsAPI.StarSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StarSegment`: DetailedSegment
	fmt.Fprintf(os.Stdout, "Response from `SegmentsAPI.StarSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the segment to star. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStarSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **starred** | **bool** | If true, star the segment; if false, unstar the segment. | [default to false]

### Return type

[**DetailedSegment**](DetailedSegment.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

