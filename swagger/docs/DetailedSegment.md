# DetailedSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of this segment | [optional] 
**Name** | Pointer to **string** | The name of this segment | [optional] 
**ActivityType** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **float32** | The segment&#39;s distance, in meters | [optional] 
**AverageGrade** | Pointer to **float32** | The segment&#39;s average grade, in percents | [optional] 
**MaximumGrade** | Pointer to **float32** | The segments&#39;s maximum grade, in percents | [optional] 
**ElevationHigh** | Pointer to **float32** | The segments&#39;s highest elevation, in meters | [optional] 
**ElevationLow** | Pointer to **float32** | The segments&#39;s lowest elevation, in meters | [optional] 
**StartLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**EndLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**ClimbCategory** | Pointer to **int32** | The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category. | [optional] 
**City** | Pointer to **string** | The segments&#39;s city. | [optional] 
**State** | Pointer to **string** | The segments&#39;s state or geographical region. | [optional] 
**Country** | Pointer to **string** | The segment&#39;s country. | [optional] 
**Private** | Pointer to **bool** | Whether this segment is private. | [optional] 
**AthletePrEffort** | Pointer to [**SummaryPRSegmentEffort**](SummaryPRSegmentEffort.md) |  | [optional] 
**AthleteSegmentStats** | Pointer to [**SummarySegmentEffort**](SummarySegmentEffort.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which the segment was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which the segment was last updated. | [optional] 
**TotalElevationGain** | Pointer to **float32** | The segment&#39;s total elevation gain. | [optional] 
**Map** | Pointer to [**PolylineMap**](PolylineMap.md) |  | [optional] 
**EffortCount** | Pointer to **int32** | The total number of efforts for this segment | [optional] 
**AthleteCount** | Pointer to **int32** | The number of unique athletes who have an effort for this segment | [optional] 
**Hazardous** | Pointer to **bool** | Whether this segment is considered hazardous | [optional] 
**StarCount** | Pointer to **int32** | The number of stars for this segment | [optional] 

## Methods

### NewDetailedSegment

`func NewDetailedSegment() *DetailedSegment`

NewDetailedSegment instantiates a new DetailedSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedSegmentWithDefaults

`func NewDetailedSegmentWithDefaults() *DetailedSegment`

NewDetailedSegmentWithDefaults instantiates a new DetailedSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedSegment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedSegment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedSegment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DetailedSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedSegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedSegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActivityType

`func (o *DetailedSegment) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *DetailedSegment) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *DetailedSegment) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *DetailedSegment) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetDistance

`func (o *DetailedSegment) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *DetailedSegment) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *DetailedSegment) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *DetailedSegment) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetAverageGrade

`func (o *DetailedSegment) GetAverageGrade() float32`

GetAverageGrade returns the AverageGrade field if non-nil, zero value otherwise.

### GetAverageGradeOk

`func (o *DetailedSegment) GetAverageGradeOk() (*float32, bool)`

GetAverageGradeOk returns a tuple with the AverageGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageGrade

`func (o *DetailedSegment) SetAverageGrade(v float32)`

SetAverageGrade sets AverageGrade field to given value.

### HasAverageGrade

`func (o *DetailedSegment) HasAverageGrade() bool`

HasAverageGrade returns a boolean if a field has been set.

### GetMaximumGrade

`func (o *DetailedSegment) GetMaximumGrade() float32`

GetMaximumGrade returns the MaximumGrade field if non-nil, zero value otherwise.

### GetMaximumGradeOk

`func (o *DetailedSegment) GetMaximumGradeOk() (*float32, bool)`

GetMaximumGradeOk returns a tuple with the MaximumGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumGrade

`func (o *DetailedSegment) SetMaximumGrade(v float32)`

SetMaximumGrade sets MaximumGrade field to given value.

### HasMaximumGrade

`func (o *DetailedSegment) HasMaximumGrade() bool`

HasMaximumGrade returns a boolean if a field has been set.

### GetElevationHigh

`func (o *DetailedSegment) GetElevationHigh() float32`

GetElevationHigh returns the ElevationHigh field if non-nil, zero value otherwise.

### GetElevationHighOk

`func (o *DetailedSegment) GetElevationHighOk() (*float32, bool)`

GetElevationHighOk returns a tuple with the ElevationHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevationHigh

`func (o *DetailedSegment) SetElevationHigh(v float32)`

SetElevationHigh sets ElevationHigh field to given value.

### HasElevationHigh

`func (o *DetailedSegment) HasElevationHigh() bool`

HasElevationHigh returns a boolean if a field has been set.

### GetElevationLow

`func (o *DetailedSegment) GetElevationLow() float32`

GetElevationLow returns the ElevationLow field if non-nil, zero value otherwise.

### GetElevationLowOk

`func (o *DetailedSegment) GetElevationLowOk() (*float32, bool)`

GetElevationLowOk returns a tuple with the ElevationLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevationLow

`func (o *DetailedSegment) SetElevationLow(v float32)`

SetElevationLow sets ElevationLow field to given value.

### HasElevationLow

`func (o *DetailedSegment) HasElevationLow() bool`

HasElevationLow returns a boolean if a field has been set.

### GetStartLatlng

`func (o *DetailedSegment) GetStartLatlng() []float32`

GetStartLatlng returns the StartLatlng field if non-nil, zero value otherwise.

### GetStartLatlngOk

`func (o *DetailedSegment) GetStartLatlngOk() (*[]float32, bool)`

GetStartLatlngOk returns a tuple with the StartLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLatlng

`func (o *DetailedSegment) SetStartLatlng(v []float32)`

SetStartLatlng sets StartLatlng field to given value.

### HasStartLatlng

`func (o *DetailedSegment) HasStartLatlng() bool`

HasStartLatlng returns a boolean if a field has been set.

### GetEndLatlng

`func (o *DetailedSegment) GetEndLatlng() []float32`

GetEndLatlng returns the EndLatlng field if non-nil, zero value otherwise.

### GetEndLatlngOk

`func (o *DetailedSegment) GetEndLatlngOk() (*[]float32, bool)`

GetEndLatlngOk returns a tuple with the EndLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLatlng

`func (o *DetailedSegment) SetEndLatlng(v []float32)`

SetEndLatlng sets EndLatlng field to given value.

### HasEndLatlng

`func (o *DetailedSegment) HasEndLatlng() bool`

HasEndLatlng returns a boolean if a field has been set.

### GetClimbCategory

`func (o *DetailedSegment) GetClimbCategory() int32`

GetClimbCategory returns the ClimbCategory field if non-nil, zero value otherwise.

### GetClimbCategoryOk

`func (o *DetailedSegment) GetClimbCategoryOk() (*int32, bool)`

GetClimbCategoryOk returns a tuple with the ClimbCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimbCategory

`func (o *DetailedSegment) SetClimbCategory(v int32)`

SetClimbCategory sets ClimbCategory field to given value.

### HasClimbCategory

`func (o *DetailedSegment) HasClimbCategory() bool`

HasClimbCategory returns a boolean if a field has been set.

### GetCity

`func (o *DetailedSegment) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DetailedSegment) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DetailedSegment) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DetailedSegment) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *DetailedSegment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DetailedSegment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DetailedSegment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DetailedSegment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *DetailedSegment) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DetailedSegment) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DetailedSegment) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DetailedSegment) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPrivate

`func (o *DetailedSegment) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *DetailedSegment) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *DetailedSegment) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *DetailedSegment) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetAthletePrEffort

`func (o *DetailedSegment) GetAthletePrEffort() SummaryPRSegmentEffort`

GetAthletePrEffort returns the AthletePrEffort field if non-nil, zero value otherwise.

### GetAthletePrEffortOk

`func (o *DetailedSegment) GetAthletePrEffortOk() (*SummaryPRSegmentEffort, bool)`

GetAthletePrEffortOk returns a tuple with the AthletePrEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthletePrEffort

`func (o *DetailedSegment) SetAthletePrEffort(v SummaryPRSegmentEffort)`

SetAthletePrEffort sets AthletePrEffort field to given value.

### HasAthletePrEffort

`func (o *DetailedSegment) HasAthletePrEffort() bool`

HasAthletePrEffort returns a boolean if a field has been set.

### GetAthleteSegmentStats

`func (o *DetailedSegment) GetAthleteSegmentStats() SummarySegmentEffort`

GetAthleteSegmentStats returns the AthleteSegmentStats field if non-nil, zero value otherwise.

### GetAthleteSegmentStatsOk

`func (o *DetailedSegment) GetAthleteSegmentStatsOk() (*SummarySegmentEffort, bool)`

GetAthleteSegmentStatsOk returns a tuple with the AthleteSegmentStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthleteSegmentStats

`func (o *DetailedSegment) SetAthleteSegmentStats(v SummarySegmentEffort)`

SetAthleteSegmentStats sets AthleteSegmentStats field to given value.

### HasAthleteSegmentStats

`func (o *DetailedSegment) HasAthleteSegmentStats() bool`

HasAthleteSegmentStats returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DetailedSegment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DetailedSegment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DetailedSegment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DetailedSegment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DetailedSegment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DetailedSegment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DetailedSegment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DetailedSegment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTotalElevationGain

`func (o *DetailedSegment) GetTotalElevationGain() float32`

GetTotalElevationGain returns the TotalElevationGain field if non-nil, zero value otherwise.

### GetTotalElevationGainOk

`func (o *DetailedSegment) GetTotalElevationGainOk() (*float32, bool)`

GetTotalElevationGainOk returns a tuple with the TotalElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElevationGain

`func (o *DetailedSegment) SetTotalElevationGain(v float32)`

SetTotalElevationGain sets TotalElevationGain field to given value.

### HasTotalElevationGain

`func (o *DetailedSegment) HasTotalElevationGain() bool`

HasTotalElevationGain returns a boolean if a field has been set.

### GetMap

`func (o *DetailedSegment) GetMap() PolylineMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *DetailedSegment) GetMapOk() (*PolylineMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *DetailedSegment) SetMap(v PolylineMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *DetailedSegment) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetEffortCount

`func (o *DetailedSegment) GetEffortCount() int32`

GetEffortCount returns the EffortCount field if non-nil, zero value otherwise.

### GetEffortCountOk

`func (o *DetailedSegment) GetEffortCountOk() (*int32, bool)`

GetEffortCountOk returns a tuple with the EffortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortCount

`func (o *DetailedSegment) SetEffortCount(v int32)`

SetEffortCount sets EffortCount field to given value.

### HasEffortCount

`func (o *DetailedSegment) HasEffortCount() bool`

HasEffortCount returns a boolean if a field has been set.

### GetAthleteCount

`func (o *DetailedSegment) GetAthleteCount() int32`

GetAthleteCount returns the AthleteCount field if non-nil, zero value otherwise.

### GetAthleteCountOk

`func (o *DetailedSegment) GetAthleteCountOk() (*int32, bool)`

GetAthleteCountOk returns a tuple with the AthleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthleteCount

`func (o *DetailedSegment) SetAthleteCount(v int32)`

SetAthleteCount sets AthleteCount field to given value.

### HasAthleteCount

`func (o *DetailedSegment) HasAthleteCount() bool`

HasAthleteCount returns a boolean if a field has been set.

### GetHazardous

`func (o *DetailedSegment) GetHazardous() bool`

GetHazardous returns the Hazardous field if non-nil, zero value otherwise.

### GetHazardousOk

`func (o *DetailedSegment) GetHazardousOk() (*bool, bool)`

GetHazardousOk returns a tuple with the Hazardous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHazardous

`func (o *DetailedSegment) SetHazardous(v bool)`

SetHazardous sets Hazardous field to given value.

### HasHazardous

`func (o *DetailedSegment) HasHazardous() bool`

HasHazardous returns a boolean if a field has been set.

### GetStarCount

`func (o *DetailedSegment) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *DetailedSegment) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *DetailedSegment) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.

### HasStarCount

`func (o *DetailedSegment) HasStarCount() bool`

HasStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


