# \ActivitiesAPI

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActivity**](ActivitiesAPI.md#CreateActivity) | **Post** /activities | Create an Activity
[**GetActivityById**](ActivitiesAPI.md#GetActivityById) | **Get** /activities/{id} | Get Activity
[**GetCommentsByActivityId**](ActivitiesAPI.md#GetCommentsByActivityId) | **Get** /activities/{id}/comments | List Activity Comments
[**GetKudoersByActivityId**](ActivitiesAPI.md#GetKudoersByActivityId) | **Get** /activities/{id}/kudos | List Activity Kudoers
[**GetLapsByActivityId**](ActivitiesAPI.md#GetLapsByActivityId) | **Get** /activities/{id}/laps | List Activity Laps
[**GetLoggedInAthleteActivities**](ActivitiesAPI.md#GetLoggedInAthleteActivities) | **Get** /athlete/activities | List Athlete Activities
[**GetZonesByActivityId**](ActivitiesAPI.md#GetZonesByActivityId) | **Get** /activities/{id}/zones | Get Activity Zones
[**UpdateActivityById**](ActivitiesAPI.md#UpdateActivityById) | **Put** /activities/{id} | Update Activity



## CreateActivity

> DetailedActivity CreateActivity(ctx).Name(name).SportType(sportType).StartDateLocal(startDateLocal).ElapsedTime(elapsedTime).Type_(type_).Description(description).Distance(distance).Trainer(trainer).Commute(commute).Execute()

Create an Activity



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
	name := "name_example" // string | The name of the activity.
	sportType := "sportType_example" // string | Sport type of activity. For example - Run, MountainBikeRide, Ride, etc.
	startDateLocal := time.Now() // time.Time | ISO 8601 formatted date time.
	elapsedTime := int32(56) // int32 | In seconds.
	type_ := "type__example" // string | Type of activity. For example - Run, Ride etc. (optional)
	description := "description_example" // string | Description of the activity. (optional)
	distance := float32(3.4) // float32 | In meters. (optional)
	trainer := int32(56) // int32 | Set to 1 to mark as a trainer activity. (optional)
	commute := int32(56) // int32 | Set to 1 to mark as commute. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.CreateActivity(context.Background()).Name(name).SportType(sportType).StartDateLocal(startDateLocal).ElapsedTime(elapsedTime).Type_(type_).Description(description).Distance(distance).Trainer(trainer).Commute(commute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.CreateActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActivity`: DetailedActivity
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.CreateActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the activity. | 
 **sportType** | **string** | Sport type of activity. For example - Run, MountainBikeRide, Ride, etc. | 
 **startDateLocal** | **time.Time** | ISO 8601 formatted date time. | 
 **elapsedTime** | **int32** | In seconds. | 
 **type_** | **string** | Type of activity. For example - Run, Ride etc. | 
 **description** | **string** | Description of the activity. | 
 **distance** | **float32** | In meters. | 
 **trainer** | **int32** | Set to 1 to mark as a trainer activity. | 
 **commute** | **int32** | Set to 1 to mark as commute. | 

### Return type

[**DetailedActivity**](DetailedActivity.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityById

> DetailedActivity GetActivityById(ctx, id).IncludeAllEfforts(includeAllEfforts).Execute()

Get Activity



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
	id := int64(789) // int64 | The identifier of the activity.
	includeAllEfforts := true // bool | To include all segments efforts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.GetActivityById(context.Background(), id).IncludeAllEfforts(includeAllEfforts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.GetActivityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityById`: DetailedActivity
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.GetActivityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAllEfforts** | **bool** | To include all segments efforts. | 

### Return type

[**DetailedActivity**](DetailedActivity.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentsByActivityId

> []Comment GetCommentsByActivityId(ctx, id).Page(page).PerPage(perPage).PageSize(pageSize).AfterCursor(afterCursor).Execute()

List Activity Comments



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
	id := int64(789) // int64 | The identifier of the activity.
	page := int32(56) // int32 | Deprecated. Prefer to use after_cursor. (optional)
	perPage := int32(56) // int32 | Deprecated. Prefer to use page_size. (optional) (default to 30)
	pageSize := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)
	afterCursor := "afterCursor_example" // string | Cursor of the last item in the previous page of results, used to request the subsequent page of results.  When omitted, the first page of results is fetched. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.GetCommentsByActivityId(context.Background(), id).Page(page).PerPage(perPage).PageSize(pageSize).AfterCursor(afterCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.GetCommentsByActivityId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentsByActivityId`: []Comment
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.GetCommentsByActivityId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsByActivityIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Deprecated. Prefer to use after_cursor. | 
 **perPage** | **int32** | Deprecated. Prefer to use page_size. | [default to 30]
 **pageSize** | **int32** | Number of items per page. Defaults to 30. | [default to 30]
 **afterCursor** | **string** | Cursor of the last item in the previous page of results, used to request the subsequent page of results.  When omitted, the first page of results is fetched. | 

### Return type

[**[]Comment**](Comment.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKudoersByActivityId

> []SummaryAthlete GetKudoersByActivityId(ctx, id).Page(page).PerPage(perPage).Execute()

List Activity Kudoers



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
	id := int64(789) // int64 | The identifier of the activity.
	page := int32(56) // int32 | Page number. Defaults to 1. (optional)
	perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.GetKudoersByActivityId(context.Background(), id).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.GetKudoersByActivityId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKudoersByActivityId`: []SummaryAthlete
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.GetKudoersByActivityId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKudoersByActivityIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]SummaryAthlete**](SummaryAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLapsByActivityId

> []Lap GetLapsByActivityId(ctx, id).Execute()

List Activity Laps



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
	id := int64(789) // int64 | The identifier of the activity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.GetLapsByActivityId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.GetLapsByActivityId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLapsByActivityId`: []Lap
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.GetLapsByActivityId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLapsByActivityIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Lap**](Lap.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggedInAthleteActivities

> []SummaryActivity GetLoggedInAthleteActivities(ctx).Before(before).After(after).Page(page).PerPage(perPage).Execute()

List Athlete Activities



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
	before := int32(56) // int32 | An epoch timestamp to use for filtering activities that have taken place before a certain time. (optional)
	after := int32(56) // int32 | An epoch timestamp to use for filtering activities that have taken place after a certain time. (optional)
	page := int32(56) // int32 | Page number. Defaults to 1. (optional)
	perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.GetLoggedInAthleteActivities(context.Background()).Before(before).After(after).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.GetLoggedInAthleteActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoggedInAthleteActivities`: []SummaryActivity
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.GetLoggedInAthleteActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggedInAthleteActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **int32** | An epoch timestamp to use for filtering activities that have taken place before a certain time. | 
 **after** | **int32** | An epoch timestamp to use for filtering activities that have taken place after a certain time. | 
 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]SummaryActivity**](SummaryActivity.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZonesByActivityId

> []ActivityZone GetZonesByActivityId(ctx, id).Execute()

Get Activity Zones



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
	id := int64(789) // int64 | The identifier of the activity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.GetZonesByActivityId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.GetZonesByActivityId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZonesByActivityId`: []ActivityZone
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.GetZonesByActivityId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesByActivityIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ActivityZone**](ActivityZone.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivityById

> DetailedActivity UpdateActivityById(ctx, id).Body(body).Execute()

Update Activity



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
	id := int64(789) // int64 | The identifier of the activity.
	body := *openapiclient.NewUpdatableActivity() // UpdatableActivity |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesAPI.UpdateActivityById(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesAPI.UpdateActivityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActivityById`: DetailedActivity
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesAPI.UpdateActivityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActivityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdatableActivity**](UpdatableActivity.md) |  | 

### Return type

[**DetailedActivity**](DetailedActivity.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

