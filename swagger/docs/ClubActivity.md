# ClubActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Athlete** | Pointer to [**MetaAthlete**](MetaAthlete.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the activity | [optional] 
**Distance** | Pointer to **float32** | The activity&#39;s distance, in meters | [optional] 
**MovingTime** | Pointer to **int32** | The activity&#39;s moving time, in seconds | [optional] 
**ElapsedTime** | Pointer to **int32** | The activity&#39;s elapsed time, in seconds | [optional] 
**TotalElevationGain** | Pointer to **float32** | The activity&#39;s total elevation gain. | [optional] 
**Type** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**SportType** | Pointer to [**SportType**](SportType.md) |  | [optional] 
**WorkoutType** | Pointer to **int32** | The activity&#39;s workout type | [optional] 

## Methods

### NewClubActivity

`func NewClubActivity() *ClubActivity`

NewClubActivity instantiates a new ClubActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClubActivityWithDefaults

`func NewClubActivityWithDefaults() *ClubActivity`

NewClubActivityWithDefaults instantiates a new ClubActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAthlete

`func (o *ClubActivity) GetAthlete() MetaAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *ClubActivity) GetAthleteOk() (*MetaAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *ClubActivity) SetAthlete(v MetaAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *ClubActivity) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetName

`func (o *ClubActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClubActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClubActivity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClubActivity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistance

`func (o *ClubActivity) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *ClubActivity) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *ClubActivity) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *ClubActivity) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetMovingTime

`func (o *ClubActivity) GetMovingTime() int32`

GetMovingTime returns the MovingTime field if non-nil, zero value otherwise.

### GetMovingTimeOk

`func (o *ClubActivity) GetMovingTimeOk() (*int32, bool)`

GetMovingTimeOk returns a tuple with the MovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingTime

`func (o *ClubActivity) SetMovingTime(v int32)`

SetMovingTime sets MovingTime field to given value.

### HasMovingTime

`func (o *ClubActivity) HasMovingTime() bool`

HasMovingTime returns a boolean if a field has been set.

### GetElapsedTime

`func (o *ClubActivity) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ClubActivity) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ClubActivity) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ClubActivity) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetTotalElevationGain

`func (o *ClubActivity) GetTotalElevationGain() float32`

GetTotalElevationGain returns the TotalElevationGain field if non-nil, zero value otherwise.

### GetTotalElevationGainOk

`func (o *ClubActivity) GetTotalElevationGainOk() (*float32, bool)`

GetTotalElevationGainOk returns a tuple with the TotalElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElevationGain

`func (o *ClubActivity) SetTotalElevationGain(v float32)`

SetTotalElevationGain sets TotalElevationGain field to given value.

### HasTotalElevationGain

`func (o *ClubActivity) HasTotalElevationGain() bool`

HasTotalElevationGain returns a boolean if a field has been set.

### GetType

`func (o *ClubActivity) GetType() ActivityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClubActivity) GetTypeOk() (*ActivityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClubActivity) SetType(v ActivityType)`

SetType sets Type field to given value.

### HasType

`func (o *ClubActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSportType

`func (o *ClubActivity) GetSportType() SportType`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *ClubActivity) GetSportTypeOk() (*SportType, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *ClubActivity) SetSportType(v SportType)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *ClubActivity) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetWorkoutType

`func (o *ClubActivity) GetWorkoutType() int32`

GetWorkoutType returns the WorkoutType field if non-nil, zero value otherwise.

### GetWorkoutTypeOk

`func (o *ClubActivity) GetWorkoutTypeOk() (*int32, bool)`

GetWorkoutTypeOk returns a tuple with the WorkoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutType

`func (o *ClubActivity) SetWorkoutType(v int32)`

SetWorkoutType sets WorkoutType field to given value.

### HasWorkoutType

`func (o *ClubActivity) HasWorkoutType() bool`

HasWorkoutType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


