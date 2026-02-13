# SummarySegmentEffort

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

## Methods

### NewSummarySegmentEffort

`func NewSummarySegmentEffort() *SummarySegmentEffort`

NewSummarySegmentEffort instantiates a new SummarySegmentEffort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarySegmentEffortWithDefaults

`func NewSummarySegmentEffortWithDefaults() *SummarySegmentEffort`

NewSummarySegmentEffortWithDefaults instantiates a new SummarySegmentEffort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SummarySegmentEffort) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummarySegmentEffort) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummarySegmentEffort) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SummarySegmentEffort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *SummarySegmentEffort) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *SummarySegmentEffort) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *SummarySegmentEffort) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *SummarySegmentEffort) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetElapsedTime

`func (o *SummarySegmentEffort) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SummarySegmentEffort) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SummarySegmentEffort) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *SummarySegmentEffort) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetStartDate

`func (o *SummarySegmentEffort) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SummarySegmentEffort) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SummarySegmentEffort) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SummarySegmentEffort) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateLocal

`func (o *SummarySegmentEffort) GetStartDateLocal() time.Time`

GetStartDateLocal returns the StartDateLocal field if non-nil, zero value otherwise.

### GetStartDateLocalOk

`func (o *SummarySegmentEffort) GetStartDateLocalOk() (*time.Time, bool)`

GetStartDateLocalOk returns a tuple with the StartDateLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLocal

`func (o *SummarySegmentEffort) SetStartDateLocal(v time.Time)`

SetStartDateLocal sets StartDateLocal field to given value.

### HasStartDateLocal

`func (o *SummarySegmentEffort) HasStartDateLocal() bool`

HasStartDateLocal returns a boolean if a field has been set.

### GetDistance

`func (o *SummarySegmentEffort) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SummarySegmentEffort) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SummarySegmentEffort) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SummarySegmentEffort) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetIsKom

`func (o *SummarySegmentEffort) GetIsKom() bool`

GetIsKom returns the IsKom field if non-nil, zero value otherwise.

### GetIsKomOk

`func (o *SummarySegmentEffort) GetIsKomOk() (*bool, bool)`

GetIsKomOk returns a tuple with the IsKom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKom

`func (o *SummarySegmentEffort) SetIsKom(v bool)`

SetIsKom sets IsKom field to given value.

### HasIsKom

`func (o *SummarySegmentEffort) HasIsKom() bool`

HasIsKom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


