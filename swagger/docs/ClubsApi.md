# \ClubsAPI

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClubActivitiesById**](ClubsAPI.md#GetClubActivitiesById) | **Get** /clubs/{id}/activities | List Club Activities
[**GetClubAdminsById**](ClubsAPI.md#GetClubAdminsById) | **Get** /clubs/{id}/admins | List Club Administrators
[**GetClubById**](ClubsAPI.md#GetClubById) | **Get** /clubs/{id} | Get Club
[**GetClubMembersById**](ClubsAPI.md#GetClubMembersById) | **Get** /clubs/{id}/members | List Club Members
[**GetLoggedInAthleteClubs**](ClubsAPI.md#GetLoggedInAthleteClubs) | **Get** /athlete/clubs | List Athlete Clubs



## GetClubActivitiesById

> []ClubActivity GetClubActivitiesById(ctx, id).Page(page).PerPage(perPage).Execute()

List Club Activities



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
	id := int64(789) // int64 | The identifier of the club.
	page := int32(56) // int32 | Page number. Defaults to 1. (optional)
	perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClubsAPI.GetClubActivitiesById(context.Background(), id).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClubsAPI.GetClubActivitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClubActivitiesById`: []ClubActivity
	fmt.Fprintf(os.Stdout, "Response from `ClubsAPI.GetClubActivitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the club. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClubActivitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]ClubActivity**](ClubActivity.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClubAdminsById

> []SummaryAthlete GetClubAdminsById(ctx, id).Page(page).PerPage(perPage).Execute()

List Club Administrators



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
	id := int64(789) // int64 | The identifier of the club.
	page := int32(56) // int32 | Page number. Defaults to 1. (optional)
	perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClubsAPI.GetClubAdminsById(context.Background(), id).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClubsAPI.GetClubAdminsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClubAdminsById`: []SummaryAthlete
	fmt.Fprintf(os.Stdout, "Response from `ClubsAPI.GetClubAdminsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the club. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClubAdminsByIdRequest struct via the builder pattern


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


## GetClubById

> DetailedClub GetClubById(ctx, id).Execute()

Get Club



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
	id := int64(789) // int64 | The identifier of the club.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClubsAPI.GetClubById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClubsAPI.GetClubById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClubById`: DetailedClub
	fmt.Fprintf(os.Stdout, "Response from `ClubsAPI.GetClubById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the club. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClubByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DetailedClub**](DetailedClub.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClubMembersById

> []ClubAthlete GetClubMembersById(ctx, id).Page(page).PerPage(perPage).Execute()

List Club Members



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
	id := int64(789) // int64 | The identifier of the club.
	page := int32(56) // int32 | Page number. Defaults to 1. (optional)
	perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClubsAPI.GetClubMembersById(context.Background(), id).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClubsAPI.GetClubMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClubMembersById`: []ClubAthlete
	fmt.Fprintf(os.Stdout, "Response from `ClubsAPI.GetClubMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the club. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClubMembersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]ClubAthlete**](ClubAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoggedInAthleteClubs

> []SummaryClub GetLoggedInAthleteClubs(ctx).Page(page).PerPage(perPage).Execute()

List Athlete Clubs



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
	resp, r, err := apiClient.ClubsAPI.GetLoggedInAthleteClubs(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClubsAPI.GetLoggedInAthleteClubs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoggedInAthleteClubs`: []SummaryClub
	fmt.Fprintf(os.Stdout, "Response from `ClubsAPI.GetLoggedInAthleteClubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoggedInAthleteClubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]SummaryClub**](SummaryClub.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

