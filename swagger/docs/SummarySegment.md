# SummarySegment

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

## Methods

### NewSummarySegment

`func NewSummarySegment() *SummarySegment`

NewSummarySegment instantiates a new SummarySegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarySegmentWithDefaults

`func NewSummarySegmentWithDefaults() *SummarySegment`

NewSummarySegmentWithDefaults instantiates a new SummarySegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SummarySegment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummarySegment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummarySegment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SummarySegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SummarySegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummarySegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummarySegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummarySegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActivityType

`func (o *SummarySegment) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *SummarySegment) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *SummarySegment) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *SummarySegment) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetDistance

`func (o *SummarySegment) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SummarySegment) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SummarySegment) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SummarySegment) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetAverageGrade

`func (o *SummarySegment) GetAverageGrade() float32`

GetAverageGrade returns the AverageGrade field if non-nil, zero value otherwise.

### GetAverageGradeOk

`func (o *SummarySegment) GetAverageGradeOk() (*float32, bool)`

GetAverageGradeOk returns a tuple with the AverageGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageGrade

`func (o *SummarySegment) SetAverageGrade(v float32)`

SetAverageGrade sets AverageGrade field to given value.

### HasAverageGrade

`func (o *SummarySegment) HasAverageGrade() bool`

HasAverageGrade returns a boolean if a field has been set.

### GetMaximumGrade

`func (o *SummarySegment) GetMaximumGrade() float32`

GetMaximumGrade returns the MaximumGrade field if non-nil, zero value otherwise.

### GetMaximumGradeOk

`func (o *SummarySegment) GetMaximumGradeOk() (*float32, bool)`

GetMaximumGradeOk returns a tuple with the MaximumGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumGrade

`func (o *SummarySegment) SetMaximumGrade(v float32)`

SetMaximumGrade sets MaximumGrade field to given value.

### HasMaximumGrade

`func (o *SummarySegment) HasMaximumGrade() bool`

HasMaximumGrade returns a boolean if a field has been set.

### GetElevationHigh

`func (o *SummarySegment) GetElevationHigh() float32`

GetElevationHigh returns the ElevationHigh field if non-nil, zero value otherwise.

### GetElevationHighOk

`func (o *SummarySegment) GetElevationHighOk() (*float32, bool)`

GetElevationHighOk returns a tuple with the ElevationHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevationHigh

`func (o *SummarySegment) SetElevationHigh(v float32)`

SetElevationHigh sets ElevationHigh field to given value.

### HasElevationHigh

`func (o *SummarySegment) HasElevationHigh() bool`

HasElevationHigh returns a boolean if a field has been set.

### GetElevationLow

`func (o *SummarySegment) GetElevationLow() float32`

GetElevationLow returns the ElevationLow field if non-nil, zero value otherwise.

### GetElevationLowOk

`func (o *SummarySegment) GetElevationLowOk() (*float32, bool)`

GetElevationLowOk returns a tuple with the ElevationLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevationLow

`func (o *SummarySegment) SetElevationLow(v float32)`

SetElevationLow sets ElevationLow field to given value.

### HasElevationLow

`func (o *SummarySegment) HasElevationLow() bool`

HasElevationLow returns a boolean if a field has been set.

### GetStartLatlng

`func (o *SummarySegment) GetStartLatlng() []float32`

GetStartLatlng returns the StartLatlng field if non-nil, zero value otherwise.

### GetStartLatlngOk

`func (o *SummarySegment) GetStartLatlngOk() (*[]float32, bool)`

GetStartLatlngOk returns a tuple with the StartLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLatlng

`func (o *SummarySegment) SetStartLatlng(v []float32)`

SetStartLatlng sets StartLatlng field to given value.

### HasStartLatlng

`func (o *SummarySegment) HasStartLatlng() bool`

HasStartLatlng returns a boolean if a field has been set.

### GetEndLatlng

`func (o *SummarySegment) GetEndLatlng() []float32`

GetEndLatlng returns the EndLatlng field if non-nil, zero value otherwise.

### GetEndLatlngOk

`func (o *SummarySegment) GetEndLatlngOk() (*[]float32, bool)`

GetEndLatlngOk returns a tuple with the EndLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLatlng

`func (o *SummarySegment) SetEndLatlng(v []float32)`

SetEndLatlng sets EndLatlng field to given value.

### HasEndLatlng

`func (o *SummarySegment) HasEndLatlng() bool`

HasEndLatlng returns a boolean if a field has been set.

### GetClimbCategory

`func (o *SummarySegment) GetClimbCategory() int32`

GetClimbCategory returns the ClimbCategory field if non-nil, zero value otherwise.

### GetClimbCategoryOk

`func (o *SummarySegment) GetClimbCategoryOk() (*int32, bool)`

GetClimbCategoryOk returns a tuple with the ClimbCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimbCategory

`func (o *SummarySegment) SetClimbCategory(v int32)`

SetClimbCategory sets ClimbCategory field to given value.

### HasClimbCategory

`func (o *SummarySegment) HasClimbCategory() bool`

HasClimbCategory returns a boolean if a field has been set.

### GetCity

`func (o *SummarySegment) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SummarySegment) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SummarySegment) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SummarySegment) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *SummarySegment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SummarySegment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SummarySegment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SummarySegment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *SummarySegment) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SummarySegment) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SummarySegment) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SummarySegment) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPrivate

`func (o *SummarySegment) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SummarySegment) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SummarySegment) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SummarySegment) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetAthletePrEffort

`func (o *SummarySegment) GetAthletePrEffort() SummaryPRSegmentEffort`

GetAthletePrEffort returns the AthletePrEffort field if non-nil, zero value otherwise.

### GetAthletePrEffortOk

`func (o *SummarySegment) GetAthletePrEffortOk() (*SummaryPRSegmentEffort, bool)`

GetAthletePrEffortOk returns a tuple with the AthletePrEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthletePrEffort

`func (o *SummarySegment) SetAthletePrEffort(v SummaryPRSegmentEffort)`

SetAthletePrEffort sets AthletePrEffort field to given value.

### HasAthletePrEffort

`func (o *SummarySegment) HasAthletePrEffort() bool`

HasAthletePrEffort returns a boolean if a field has been set.

### GetAthleteSegmentStats

`func (o *SummarySegment) GetAthleteSegmentStats() SummarySegmentEffort`

GetAthleteSegmentStats returns the AthleteSegmentStats field if non-nil, zero value otherwise.

### GetAthleteSegmentStatsOk

`func (o *SummarySegment) GetAthleteSegmentStatsOk() (*SummarySegmentEffort, bool)`

GetAthleteSegmentStatsOk returns a tuple with the AthleteSegmentStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthleteSegmentStats

`func (o *SummarySegment) SetAthleteSegmentStats(v SummarySegmentEffort)`

SetAthleteSegmentStats sets AthleteSegmentStats field to given value.

### HasAthleteSegmentStats

`func (o *SummarySegment) HasAthleteSegmentStats() bool`

HasAthleteSegmentStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


