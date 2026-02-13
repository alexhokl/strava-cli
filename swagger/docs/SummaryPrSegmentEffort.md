# SummaryPRSegmentEffort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrActivityId** | Pointer to **int64** | The unique identifier of the activity related to the PR effort. | [optional] 
**PrElapsedTime** | Pointer to **int32** | The elapsed time ot the PR effort. | [optional] 
**PrDate** | Pointer to **time.Time** | The time at which the PR effort was started. | [optional] 
**EffortCount** | Pointer to **int32** | Number of efforts by the authenticated athlete on this segment. | [optional] 

## Methods

### NewSummaryPRSegmentEffort

`func NewSummaryPRSegmentEffort() *SummaryPRSegmentEffort`

NewSummaryPRSegmentEffort instantiates a new SummaryPRSegmentEffort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryPRSegmentEffortWithDefaults

`func NewSummaryPRSegmentEffortWithDefaults() *SummaryPRSegmentEffort`

NewSummaryPRSegmentEffortWithDefaults instantiates a new SummaryPRSegmentEffort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrActivityId

`func (o *SummaryPRSegmentEffort) GetPrActivityId() int64`

GetPrActivityId returns the PrActivityId field if non-nil, zero value otherwise.

### GetPrActivityIdOk

`func (o *SummaryPRSegmentEffort) GetPrActivityIdOk() (*int64, bool)`

GetPrActivityIdOk returns a tuple with the PrActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrActivityId

`func (o *SummaryPRSegmentEffort) SetPrActivityId(v int64)`

SetPrActivityId sets PrActivityId field to given value.

### HasPrActivityId

`func (o *SummaryPRSegmentEffort) HasPrActivityId() bool`

HasPrActivityId returns a boolean if a field has been set.

### GetPrElapsedTime

`func (o *SummaryPRSegmentEffort) GetPrElapsedTime() int32`

GetPrElapsedTime returns the PrElapsedTime field if non-nil, zero value otherwise.

### GetPrElapsedTimeOk

`func (o *SummaryPRSegmentEffort) GetPrElapsedTimeOk() (*int32, bool)`

GetPrElapsedTimeOk returns a tuple with the PrElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrElapsedTime

`func (o *SummaryPRSegmentEffort) SetPrElapsedTime(v int32)`

SetPrElapsedTime sets PrElapsedTime field to given value.

### HasPrElapsedTime

`func (o *SummaryPRSegmentEffort) HasPrElapsedTime() bool`

HasPrElapsedTime returns a boolean if a field has been set.

### GetPrDate

`func (o *SummaryPRSegmentEffort) GetPrDate() time.Time`

GetPrDate returns the PrDate field if non-nil, zero value otherwise.

### GetPrDateOk

`func (o *SummaryPRSegmentEffort) GetPrDateOk() (*time.Time, bool)`

GetPrDateOk returns a tuple with the PrDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrDate

`func (o *SummaryPRSegmentEffort) SetPrDate(v time.Time)`

SetPrDate sets PrDate field to given value.

### HasPrDate

`func (o *SummaryPRSegmentEffort) HasPrDate() bool`

HasPrDate returns a boolean if a field has been set.

### GetEffortCount

`func (o *SummaryPRSegmentEffort) GetEffortCount() int32`

GetEffortCount returns the EffortCount field if non-nil, zero value otherwise.

### GetEffortCountOk

`func (o *SummaryPRSegmentEffort) GetEffortCountOk() (*int32, bool)`

GetEffortCountOk returns a tuple with the EffortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortCount

`func (o *SummaryPRSegmentEffort) SetEffortCount(v int32)`

SetEffortCount sets EffortCount field to given value.

### HasEffortCount

`func (o *SummaryPRSegmentEffort) HasEffortCount() bool`

HasEffortCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


