# \AthletesAPI

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoggedInAthlete**](AthletesAPI.md#GetLoggedInAthlete) | **Get** /athlete | Get Authenticated Athlete
[**GetLoggedInAthleteZones**](AthletesAPI.md#GetLoggedInAthleteZones) | **Get** /athlete/zones | Get Zones
[**GetStats**](AthletesAPI.md#GetStats) | **Get** /athletes/{id}/stats | Get Athlete Stats
[**UpdateLoggedInAthlete**](AthletesAPI.md#UpdateLoggedInAthlete) | **Put** /athlete | Update Athlete



## GetLoggedInAthlete

> DetailedAthlete GetLoggedInAthlete(ctx).Execute()

Get Authenticated Athlete



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AthletesAPI.GetLoggedInAthlete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AthletesAPI.GetLoggedInAthlete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoggedInAthlete`: DetailedAthlete
	fmt.Fprintf(os.Stdout, "Response from `AthletesAPI.GetLoggedInAthlete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggedInAthleteRequest struct via the builder pattern


### Return type

[**DetailedAthlete**](DetailedAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggedInAthleteZones

> Zones GetLoggedInAthleteZones(ctx).Execute()

Get Zones



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AthletesAPI.GetLoggedInAthleteZones(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AthletesAPI.GetLoggedInAthleteZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoggedInAthleteZones`: Zones
	fmt.Fprintf(os.Stdout, "Response from `AthletesAPI.GetLoggedInAthleteZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggedInAthleteZonesRequest struct via the builder pattern


### Return type

[**Zones**](Zones.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStats

> ActivityStats GetStats(ctx, id).Execute()

Get Athlete Stats



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
	id := int64(789) // int64 | The identifier of the athlete. Must match the authenticated athlete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AthletesAPI.GetStats(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AthletesAPI.GetStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStats`: ActivityStats
	fmt.Fprintf(os.Stdout, "Response from `AthletesAPI.GetStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the athlete. Must match the authenticated athlete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityStats**](ActivityStats.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoggedInAthlete

> DetailedAthlete UpdateLoggedInAthlete(ctx, weight).Execute()

Update Athlete



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
	weight := float32(3.4) // float32 | The weight of the athlete in kilograms.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AthletesAPI.UpdateLoggedInAthlete(context.Background(), weight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AthletesAPI.UpdateLoggedInAthlete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoggedInAthlete`: DetailedAthlete
	fmt.Fprintf(os.Stdout, "Response from `AthletesAPI.UpdateLoggedInAthlete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**weight** | **float32** | The weight of the athlete in kilograms. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoggedInAthleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetailedAthlete**](DetailedAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

