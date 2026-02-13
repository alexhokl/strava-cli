# DetailedSegmentEffort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of this effort | [optional] 
**ActivityId** | Pointer to **int64** | The unique identifier of the activity related to this effort | [optional] 
**ElapsedTime** | Pointer to **int32** | The effort&#39;s elapsed time | [optional] 
**StartDate** | Pointer to **time.Time** | The time at which the effort was started. | [optional] 
**StartDateLocal** | Pointer to **time.Time** | The time at which the effort was started in the local timezone. | [optional] 
**Distance** | Pointer to **float32** | The effort&#39;s distance in meters | [optional] 
**IsKom** | Pointer to **bool** | Whether this effort is the current best on the leaderboard | [optional] 
**Name** | Pointer to **string** | The name of the segment on which this effort was performed | [optional] 
**Activity** | Pointer to [**MetaActivity**](MetaActivity.md) |  | [optional] 
**Athlete** | Pointer to [**MetaAthlete**](MetaAthlete.md) |  | [optional] 
**MovingTime** | Pointer to **int32** | The effort&#39;s moving time | [optional] 
**StartIndex** | Pointer to **int32** | The start index of this effort in its activity&#39;s stream | [optional] 
**EndIndex** | Pointer to **int32** | The end index of this effort in its activity&#39;s stream | [optional] 
**AverageCadence** | Pointer to **float32** | The effort&#39;s average cadence | [optional] 
**AverageWatts** | Pointer to **float32** | The average wattage of this effort | [optional] 
**DeviceWatts** | Pointer to **bool** | For riding efforts, whether the wattage was reported by a dedicated recording device | [optional] 
**AverageHeartrate** | Pointer to **float32** | The heart heart rate of the athlete during this effort | [optional] 
**MaxHeartrate** | Pointer to **float32** | The maximum heart rate of the athlete during this effort | [optional] 
**Segment** | Pointer to [**SummarySegment**](SummarySegment.md) |  | [optional] 
**KomRank** | Pointer to **int32** | The rank of the effort on the global leaderboard if it belongs in the top 10 at the time of upload | [optional] 
**PrRank** | Pointer to **int32** | The rank of the effort on the athlete&#39;s leaderboard if it belongs in the top 3 at the time of upload | [optional] 
**Hidden** | Pointer to **bool** | Whether this effort should be hidden when viewed within an activity | [optional] 

## Methods

### NewDetailedSegmentEffort

`func NewDetailedSegmentEffort() *DetailedSegmentEffort`

NewDetailedSegmentEffort instantiates a new DetailedSegmentEffort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedSegmentEffortWithDefaults

`func NewDetailedSegmentEffortWithDefaults() *DetailedSegmentEffort`

NewDetailedSegmentEffortWithDefaults instantiates a new DetailedSegmentEffort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedSegmentEffort) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedSegmentEffort) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedSegmentEffort) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedSegmentEffort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *DetailedSegmentEffort) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *DetailedSegmentEffort) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *DetailedSegmentEffort) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *DetailedSegmentEffort) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetElapsedTime

`func (o *DetailedSegmentEffort) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *DetailedSegmentEffort) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *DetailedSegmentEffort) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *DetailedSegmentEffort) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetStartDate

`func (o *DetailedSegmentEffort) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DetailedSegmentEffort) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DetailedSegmentEffort) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DetailedSegmentEffort) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateLocal

`func (o *DetailedSegmentEffort) GetStartDateLocal() time.Time`

GetStartDateLocal returns the StartDateLocal field if non-nil, zero value otherwise.

### GetStartDateLocalOk

`func (o *DetailedSegmentEffort) GetStartDateLocalOk() (*time.Time, bool)`

GetStartDateLocalOk returns a tuple with the StartDateLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLocal

`func (o *DetailedSegmentEffort) SetStartDateLocal(v time.Time)`

SetStartDateLocal sets StartDateLocal field to given value.

### HasStartDateLocal

`func (o *DetailedSegmentEffort) HasStartDateLocal() bool`

HasStartDateLocal returns a boolean if a field has been set.

### GetDistance

`func (o *DetailedSegmentEffort) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *DetailedSegmentEffort) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *DetailedSegmentEffort) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *DetailedSegmentEffort) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetIsKom

`func (o *DetailedSegmentEffort) GetIsKom() bool`

GetIsKom returns the IsKom field if non-nil, zero value otherwise.

### GetIsKomOk

`func (o *DetailedSegmentEffort) GetIsKomOk() (*bool, bool)`

GetIsKomOk returns a tuple with the IsKom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKom

`func (o *DetailedSegmentEffort) SetIsKom(v bool)`

SetIsKom sets IsKom field to given value.

### HasIsKom

`func (o *DetailedSegmentEffort) HasIsKom() bool`

HasIsKom returns a boolean if a field has been set.

### GetName

`func (o *DetailedSegmentEffort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedSegmentEffort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedSegmentEffort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedSegmentEffort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActivity

`func (o *DetailedSegmentEffort) GetActivity() MetaActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *DetailedSegmentEffort) GetActivityOk() (*MetaActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *DetailedSegmentEffort) SetActivity(v MetaActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *DetailedSegmentEffort) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetAthlete

`func (o *DetailedSegmentEffort) GetAthlete() MetaAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *DetailedSegmentEffort) GetAthleteOk() (*MetaAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *DetailedSegmentEffort) SetAthlete(v MetaAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *DetailedSegmentEffort) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetMovingTime

`func (o *DetailedSegmentEffort) GetMovingTime() int32`

GetMovingTime returns the MovingTime field if non-nil, zero value otherwise.

### GetMovingTimeOk

`func (o *DetailedSegmentEffort) GetMovingTimeOk() (*int32, bool)`

GetMovingTimeOk returns a tuple with the MovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingTime

`func (o *DetailedSegmentEffort) SetMovingTime(v int32)`

SetMovingTime sets MovingTime field to given value.

### HasMovingTime

`func (o *DetailedSegmentEffort) HasMovingTime() bool`

HasMovingTime returns a boolean if a field has been set.

### GetStartIndex

`func (o *DetailedSegmentEffort) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *DetailedSegmentEffort) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *DetailedSegmentEffort) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *DetailedSegmentEffort) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetEndIndex

`func (o *DetailedSegmentEffort) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *DetailedSegmentEffort) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *DetailedSegmentEffort) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *DetailedSegmentEffort) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetAverageCadence

`func (o *DetailedSegmentEffort) GetAverageCadence() float32`

GetAverageCadence returns the AverageCadence field if non-nil, zero value otherwise.

### GetAverageCadenceOk

`func (o *DetailedSegmentEffort) GetAverageCadenceOk() (*float32, bool)`

GetAverageCadenceOk returns a tuple with the AverageCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCadence

`func (o *DetailedSegmentEffort) SetAverageCadence(v float32)`

SetAverageCadence sets AverageCadence field to given value.

### HasAverageCadence

`func (o *DetailedSegmentEffort) HasAverageCadence() bool`

HasAverageCadence returns a boolean if a field has been set.

### GetAverageWatts

`func (o *DetailedSegmentEffort) GetAverageWatts() float32`

GetAverageWatts returns the AverageWatts field if non-nil, zero value otherwise.

### GetAverageWattsOk

`func (o *DetailedSegmentEffort) GetAverageWattsOk() (*float32, bool)`

GetAverageWattsOk returns a tuple with the AverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatts

`func (o *DetailedSegmentEffort) SetAverageWatts(v float32)`

SetAverageWatts sets AverageWatts field to given value.

### HasAverageWatts

`func (o *DetailedSegmentEffort) HasAverageWatts() bool`

HasAverageWatts returns a boolean if a field has been set.

### GetDeviceWatts

`func (o *DetailedSegmentEffort) GetDeviceWatts() bool`

GetDeviceWatts returns the DeviceWatts field if non-nil, zero value otherwise.

### GetDeviceWattsOk

`func (o *DetailedSegmentEffort) GetDeviceWattsOk() (*bool, bool)`

GetDeviceWattsOk returns a tuple with the DeviceWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWatts

`func (o *DetailedSegmentEffort) SetDeviceWatts(v bool)`

SetDeviceWatts sets DeviceWatts field to given value.

### HasDeviceWatts

`func (o *DetailedSegmentEffort) HasDeviceWatts() bool`

HasDeviceWatts returns a boolean if a field has been set.

### GetAverageHeartrate

`func (o *DetailedSegmentEffort) GetAverageHeartrate() float32`

GetAverageHeartrate returns the AverageHeartrate field if non-nil, zero value otherwise.

### GetAverageHeartrateOk

`func (o *DetailedSegmentEffort) GetAverageHeartrateOk() (*float32, bool)`

GetAverageHeartrateOk returns a tuple with the AverageHeartrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageHeartrate

`func (o *DetailedSegmentEffort) SetAverageHeartrate(v float32)`

SetAverageHeartrate sets AverageHeartrate field to given value.

### HasAverageHeartrate

`func (o *DetailedSegmentEffort) HasAverageHeartrate() bool`

HasAverageHeartrate returns a boolean if a field has been set.

### GetMaxHeartrate

`func (o *DetailedSegmentEffort) GetMaxHeartrate() float32`

GetMaxHeartrate returns the MaxHeartrate field if non-nil, zero value otherwise.

### GetMaxHeartrateOk

`func (o *DetailedSegmentEffort) GetMaxHeartrateOk() (*float32, bool)`

GetMaxHeartrateOk returns a tuple with the MaxHeartrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeartrate

`func (o *DetailedSegmentEffort) SetMaxHeartrate(v float32)`

SetMaxHeartrate sets MaxHeartrate field to given value.

### HasMaxHeartrate

`func (o *DetailedSegmentEffort) HasMaxHeartrate() bool`

HasMaxHeartrate returns a boolean if a field has been set.

### GetSegment

`func (o *DetailedSegmentEffort) GetSegment() SummarySegment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *DetailedSegmentEffort) GetSegmentOk() (*SummarySegment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *DetailedSegmentEffort) SetSegment(v SummarySegment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *DetailedSegmentEffort) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetKomRank

`func (o *DetailedSegmentEffort) GetKomRank() int32`

GetKomRank returns the KomRank field if non-nil, zero value otherwise.

### GetKomRankOk

`func (o *DetailedSegmentEffort) GetKomRankOk() (*int32, bool)`

GetKomRankOk returns a tuple with the KomRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKomRank

`func (o *DetailedSegmentEffort) SetKomRank(v int32)`

SetKomRank sets KomRank field to given value.

### HasKomRank

`func (o *DetailedSegmentEffort) HasKomRank() bool`

HasKomRank returns a boolean if a field has been set.

### GetPrRank

`func (o *DetailedSegmentEffort) GetPrRank() int32`

GetPrRank returns the PrRank field if non-nil, zero value otherwise.

### GetPrRankOk

`func (o *DetailedSegmentEffort) GetPrRankOk() (*int32, bool)`

GetPrRankOk returns a tuple with the PrRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrRank

`func (o *DetailedSegmentEffort) SetPrRank(v int32)`

SetPrRank sets PrRank field to given value.

### HasPrRank

`func (o *DetailedSegmentEffort) HasPrRank() bool`

HasPrRank returns a boolean if a field has been set.

### GetHidden

`func (o *DetailedSegmentEffort) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *DetailedSegmentEffort) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *DetailedSegmentEffort) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *DetailedSegmentEffort) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


