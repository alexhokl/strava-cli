# \RoutesAPI

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteAsGPX**](RoutesAPI.md#GetRouteAsGPX) | **Get** /routes/{id}/export_gpx | Export Route GPX
[**GetRouteAsTCX**](RoutesAPI.md#GetRouteAsTCX) | **Get** /routes/{id}/export_tcx | Export Route TCX
[**GetRouteById**](RoutesAPI.md#GetRouteById) | **Get** /routes/{id} | Get Route
[**GetRoutesByAthleteId**](RoutesAPI.md#GetRoutesByAthleteId) | **Get** /athletes/{id}/routes | List Athlete Routes



## GetRouteAsGPX

> *os.File GetRouteAsGPX(ctx, id).Execute()

Export Route GPX



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
	id := int64(789) // int64 | The identifier of the route.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRouteAsGPX(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRouteAsGPX``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAsGPX`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRouteAsGPX`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAsGPXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/gpx+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAsTCX

> *os.File GetRouteAsTCX(ctx, id).Execute()

Export Route TCX



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
	id := int64(789) // int64 | The identifier of the route.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRouteAsTCX(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRouteAsTCX``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteAsTCX`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRouteAsTCX`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAsTCXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/tcx+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteById

> Route GetRouteById(ctx, id).Execute()

Get Route



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
	id := int64(789) // int64 | The identifier of the route.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRouteById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRouteById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteById`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRouteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Route**](Route.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutesByAthleteId

> []Route GetRoutesByAthleteId(ctx).Page(page).PerPage(perPage).Execute()

List Athlete Routes



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
	resp, r, err := apiClient.RoutesAPI.GetRoutesByAthleteId(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRoutesByAthleteId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutesByAthleteId`: []Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRoutesByAthleteId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesByAthleteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]Route**](Route.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

