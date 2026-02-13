# Waypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**TargetLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**Categories** | Pointer to **[]string** | Categories that the waypoint belongs to | [optional] 
**Title** | Pointer to **string** | A title for the waypoint | [optional] 
**Description** | Pointer to **string** | A description of the waypoint (optional) | [optional] 
**DistanceIntoRoute** | Pointer to **float32** | The number meters along the route that the waypoint is located | [optional] 

## Methods

### NewWaypoint

`func NewWaypoint() *Waypoint`

NewWaypoint instantiates a new Waypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaypointWithDefaults

`func NewWaypointWithDefaults() *Waypoint`

NewWaypointWithDefaults instantiates a new Waypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatlng

`func (o *Waypoint) GetLatlng() []float32`

GetLatlng returns the Latlng field if non-nil, zero value otherwise.

### GetLatlngOk

`func (o *Waypoint) GetLatlngOk() (*[]float32, bool)`

GetLatlngOk returns a tuple with the Latlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlng

`func (o *Waypoint) SetLatlng(v []float32)`

SetLatlng sets Latlng field to given value.

### HasLatlng

`func (o *Waypoint) HasLatlng() bool`

HasLatlng returns a boolean if a field has been set.

### GetTargetLatlng

`func (o *Waypoint) GetTargetLatlng() []float32`

GetTargetLatlng returns the TargetLatlng field if non-nil, zero value otherwise.

### GetTargetLatlngOk

`func (o *Waypoint) GetTargetLatlngOk() (*[]float32, bool)`

GetTargetLatlngOk returns a tuple with the TargetLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLatlng

`func (o *Waypoint) SetTargetLatlng(v []float32)`

SetTargetLatlng sets TargetLatlng field to given value.

### HasTargetLatlng

`func (o *Waypoint) HasTargetLatlng() bool`

HasTargetLatlng returns a boolean if a field has been set.

### GetCategories

`func (o *Waypoint) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Waypoint) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Waypoint) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Waypoint) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTitle

`func (o *Waypoint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Waypoint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Waypoint) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Waypoint) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Waypoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Waypoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Waypoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Waypoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistanceIntoRoute

`func (o *Waypoint) GetDistanceIntoRoute() float32`

GetDistanceIntoRoute returns the DistanceIntoRoute field if non-nil, zero value otherwise.

### GetDistanceIntoRouteOk

`func (o *Waypoint) GetDistanceIntoRouteOk() (*float32, bool)`

GetDistanceIntoRouteOk returns a tuple with the DistanceIntoRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceIntoRoute

`func (o *Waypoint) SetDistanceIntoRoute(v float32)`

SetDistanceIntoRoute sets DistanceIntoRoute field to given value.

### HasDistanceIntoRoute

`func (o *Waypoint) HasDistanceIntoRoute() bool`

HasDistanceIntoRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


