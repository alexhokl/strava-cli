# Lap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of this lap | [optional] 
**Activity** | Pointer to [**MetaActivity**](MetaActivity.md) |  | [optional] 
**Athlete** | Pointer to [**MetaAthlete**](MetaAthlete.md) |  | [optional] 
**AverageCadence** | Pointer to **float32** | The lap&#39;s average cadence | [optional] 
**AverageSpeed** | Pointer to **float32** | The lap&#39;s average speed | [optional] 
**Distance** | Pointer to **float32** | The lap&#39;s distance, in meters | [optional] 
**ElapsedTime** | Pointer to **int32** | The lap&#39;s elapsed time, in seconds | [optional] 
**StartIndex** | Pointer to **int32** | The start index of this effort in its activity&#39;s stream | [optional] 
**EndIndex** | Pointer to **int32** | The end index of this effort in its activity&#39;s stream | [optional] 
**LapIndex** | Pointer to **int32** | The index of this lap in the activity it belongs to | [optional] 
**MaxSpeed** | Pointer to **float32** | The maximum speed of this lat, in meters per second | [optional] 
**MovingTime** | Pointer to **int32** | The lap&#39;s moving time, in seconds | [optional] 
**Name** | Pointer to **string** | The name of the lap | [optional] 
**PaceZone** | Pointer to **int32** | The athlete&#39;s pace zone during this lap | [optional] 
**Split** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **time.Time** | The time at which the lap was started. | [optional] 
**StartDateLocal** | Pointer to **time.Time** | The time at which the lap was started in the local timezone. | [optional] 
**TotalElevationGain** | Pointer to **float32** | The elevation gain of this lap, in meters | [optional] 

## Methods

### NewLap

`func NewLap() *Lap`

NewLap instantiates a new Lap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapWithDefaults

`func NewLapWithDefaults() *Lap`

NewLapWithDefaults instantiates a new Lap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Lap) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lap) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lap) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Lap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivity

`func (o *Lap) GetActivity() MetaActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Lap) GetActivityOk() (*MetaActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Lap) SetActivity(v MetaActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Lap) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetAthlete

`func (o *Lap) GetAthlete() MetaAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *Lap) GetAthleteOk() (*MetaAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *Lap) SetAthlete(v MetaAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *Lap) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetAverageCadence

`func (o *Lap) GetAverageCadence() float32`

GetAverageCadence returns the AverageCadence field if non-nil, zero value otherwise.

### GetAverageCadenceOk

`func (o *Lap) GetAverageCadenceOk() (*float32, bool)`

GetAverageCadenceOk returns a tuple with the AverageCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCadence

`func (o *Lap) SetAverageCadence(v float32)`

SetAverageCadence sets AverageCadence field to given value.

### HasAverageCadence

`func (o *Lap) HasAverageCadence() bool`

HasAverageCadence returns a boolean if a field has been set.

### GetAverageSpeed

`func (o *Lap) GetAverageSpeed() float32`

GetAverageSpeed returns the AverageSpeed field if non-nil, zero value otherwise.

### GetAverageSpeedOk

`func (o *Lap) GetAverageSpeedOk() (*float32, bool)`

GetAverageSpeedOk returns a tuple with the AverageSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSpeed

`func (o *Lap) SetAverageSpeed(v float32)`

SetAverageSpeed sets AverageSpeed field to given value.

### HasAverageSpeed

`func (o *Lap) HasAverageSpeed() bool`

HasAverageSpeed returns a boolean if a field has been set.

### GetDistance

`func (o *Lap) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Lap) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Lap) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Lap) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetElapsedTime

`func (o *Lap) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *Lap) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *Lap) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *Lap) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetStartIndex

`func (o *Lap) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *Lap) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *Lap) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *Lap) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetEndIndex

`func (o *Lap) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *Lap) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *Lap) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *Lap) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetLapIndex

`func (o *Lap) GetLapIndex() int32`

GetLapIndex returns the LapIndex field if non-nil, zero value otherwise.

### GetLapIndexOk

`func (o *Lap) GetLapIndexOk() (*int32, bool)`

GetLapIndexOk returns a tuple with the LapIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLapIndex

`func (o *Lap) SetLapIndex(v int32)`

SetLapIndex sets LapIndex field to given value.

### HasLapIndex

`func (o *Lap) HasLapIndex() bool`

HasLapIndex returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *Lap) GetMaxSpeed() float32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *Lap) GetMaxSpeedOk() (*float32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *Lap) SetMaxSpeed(v float32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *Lap) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetMovingTime

`func (o *Lap) GetMovingTime() int32`

GetMovingTime returns the MovingTime field if non-nil, zero value otherwise.

### GetMovingTimeOk

`func (o *Lap) GetMovingTimeOk() (*int32, bool)`

GetMovingTimeOk returns a tuple with the MovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingTime

`func (o *Lap) SetMovingTime(v int32)`

SetMovingTime sets MovingTime field to given value.

### HasMovingTime

`func (o *Lap) HasMovingTime() bool`

HasMovingTime returns a boolean if a field has been set.

### GetName

`func (o *Lap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Lap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Lap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Lap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaceZone

`func (o *Lap) GetPaceZone() int32`

GetPaceZone returns the PaceZone field if non-nil, zero value otherwise.

### GetPaceZoneOk

`func (o *Lap) GetPaceZoneOk() (*int32, bool)`

GetPaceZoneOk returns a tuple with the PaceZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaceZone

`func (o *Lap) SetPaceZone(v int32)`

SetPaceZone sets PaceZone field to given value.

### HasPaceZone

`func (o *Lap) HasPaceZone() bool`

HasPaceZone returns a boolean if a field has been set.

### GetSplit

`func (o *Lap) GetSplit() int32`

GetSplit returns the Split field if non-nil, zero value otherwise.

### GetSplitOk

`func (o *Lap) GetSplitOk() (*int32, bool)`

GetSplitOk returns a tuple with the Split field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplit

`func (o *Lap) SetSplit(v int32)`

SetSplit sets Split field to given value.

### HasSplit

`func (o *Lap) HasSplit() bool`

HasSplit returns a boolean if a field has been set.

### GetStartDate

`func (o *Lap) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Lap) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Lap) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Lap) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateLocal

`func (o *Lap) GetStartDateLocal() time.Time`

GetStartDateLocal returns the StartDateLocal field if non-nil, zero value otherwise.

### GetStartDateLocalOk

`func (o *Lap) GetStartDateLocalOk() (*time.Time, bool)`

GetStartDateLocalOk returns a tuple with the StartDateLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLocal

`func (o *Lap) SetStartDateLocal(v time.Time)`

SetStartDateLocal sets StartDateLocal field to given value.

### HasStartDateLocal

`func (o *Lap) HasStartDateLocal() bool`

HasStartDateLocal returns a boolean if a field has been set.

### GetTotalElevationGain

`func (o *Lap) GetTotalElevationGain() float32`

GetTotalElevationGain returns the TotalElevationGain field if non-nil, zero value otherwise.

### GetTotalElevationGainOk

`func (o *Lap) GetTotalElevationGainOk() (*float32, bool)`

GetTotalElevationGainOk returns a tuple with the TotalElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElevationGain

`func (o *Lap) SetTotalElevationGain(v float32)`

SetTotalElevationGain sets TotalElevationGain field to given value.

### HasTotalElevationGain

`func (o *Lap) HasTotalElevationGain() bool`

HasTotalElevationGain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


