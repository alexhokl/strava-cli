# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Athlete** | Pointer to [**SummaryAthlete**](SummaryAthlete.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the route | [optional] 
**Distance** | Pointer to **float32** | The route&#39;s distance, in meters | [optional] 
**ElevationGain** | Pointer to **float32** | The route&#39;s elevation gain. | [optional] 
**Id** | Pointer to **int64** | The unique identifier of this route | [optional] 
**IdStr** | Pointer to **string** | The unique identifier of the route in string format | [optional] 
**Map** | Pointer to [**PolylineMap**](PolylineMap.md) |  | [optional] 
**Name** | Pointer to **string** | The name of this route | [optional] 
**Private** | Pointer to **bool** | Whether this route is private | [optional] 
**Starred** | Pointer to **bool** | Whether this route is starred by the logged-in athlete | [optional] 
**Timestamp** | Pointer to **int32** | An epoch timestamp of when the route was created | [optional] 
**Type** | Pointer to **int32** | This route&#39;s type (1 for ride, 2 for runs) | [optional] 
**SubType** | Pointer to **int32** | This route&#39;s sub-type (1 for road, 2 for mountain bike, 3 for cross, 4 for trail, 5 for mixed) | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which the route was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which the route was last updated | [optional] 
**EstimatedMovingTime** | Pointer to **int32** | Estimated time in seconds for the authenticated athlete to complete route | [optional] 
**Segments** | Pointer to [**[]SummarySegment**](SummarySegment.md) | The segments traversed by this route | [optional] 
**Waypoints** | Pointer to [**[]Waypoint**](Waypoint.md) | The custom waypoints along this route | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAthlete

`func (o *Route) GetAthlete() SummaryAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *Route) GetAthleteOk() (*SummaryAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *Route) SetAthlete(v SummaryAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *Route) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetDescription

`func (o *Route) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistance

`func (o *Route) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Route) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Route) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Route) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetElevationGain

`func (o *Route) GetElevationGain() float32`

GetElevationGain returns the ElevationGain field if non-nil, zero value otherwise.

### GetElevationGainOk

`func (o *Route) GetElevationGainOk() (*float32, bool)`

GetElevationGainOk returns a tuple with the ElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevationGain

`func (o *Route) SetElevationGain(v float32)`

SetElevationGain sets ElevationGain field to given value.

### HasElevationGain

`func (o *Route) HasElevationGain() bool`

HasElevationGain returns a boolean if a field has been set.

### GetId

`func (o *Route) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Route) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdStr

`func (o *Route) GetIdStr() string`

GetIdStr returns the IdStr field if non-nil, zero value otherwise.

### GetIdStrOk

`func (o *Route) GetIdStrOk() (*string, bool)`

GetIdStrOk returns a tuple with the IdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStr

`func (o *Route) SetIdStr(v string)`

SetIdStr sets IdStr field to given value.

### HasIdStr

`func (o *Route) HasIdStr() bool`

HasIdStr returns a boolean if a field has been set.

### GetMap

`func (o *Route) GetMap() PolylineMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *Route) GetMapOk() (*PolylineMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *Route) SetMap(v PolylineMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *Route) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetName

`func (o *Route) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Route) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *Route) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Route) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Route) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Route) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetStarred

`func (o *Route) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *Route) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *Route) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *Route) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetTimestamp

`func (o *Route) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Route) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Route) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Route) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *Route) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Route) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Route) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Route) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubType

`func (o *Route) GetSubType() int32`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Route) GetSubTypeOk() (*int32, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Route) SetSubType(v int32)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Route) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Route) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Route) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Route) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Route) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Route) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Route) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Route) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Route) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEstimatedMovingTime

`func (o *Route) GetEstimatedMovingTime() int32`

GetEstimatedMovingTime returns the EstimatedMovingTime field if non-nil, zero value otherwise.

### GetEstimatedMovingTimeOk

`func (o *Route) GetEstimatedMovingTimeOk() (*int32, bool)`

GetEstimatedMovingTimeOk returns a tuple with the EstimatedMovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMovingTime

`func (o *Route) SetEstimatedMovingTime(v int32)`

SetEstimatedMovingTime sets EstimatedMovingTime field to given value.

### HasEstimatedMovingTime

`func (o *Route) HasEstimatedMovingTime() bool`

HasEstimatedMovingTime returns a boolean if a field has been set.

### GetSegments

`func (o *Route) GetSegments() []SummarySegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Route) GetSegmentsOk() (*[]SummarySegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Route) SetSegments(v []SummarySegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Route) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetWaypoints

`func (o *Route) GetWaypoints() []Waypoint`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *Route) GetWaypointsOk() (*[]Waypoint, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *Route) SetWaypoints(v []Waypoint)`

SetWaypoints sets Waypoints field to given value.

### HasWaypoints

`func (o *Route) HasWaypoints() bool`

HasWaypoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


