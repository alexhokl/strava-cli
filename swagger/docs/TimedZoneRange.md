# TimedZoneRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Min** | Pointer to **int32** | The minimum value in the range. | [optional] 
**Max** | Pointer to **int32** | The maximum value in the range. | [optional] 
**Time** | Pointer to **int32** | The number of seconds spent in this zone | [optional] 

## Methods

### NewTimedZoneRange

`func NewTimedZoneRange() *TimedZoneRange`

NewTimedZoneRange instantiates a new TimedZoneRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimedZoneRangeWithDefaults

`func NewTimedZoneRangeWithDefaults() *TimedZoneRange`

NewTimedZoneRangeWithDefaults instantiates a new TimedZoneRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMin

`func (o *TimedZoneRange) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *TimedZoneRange) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *TimedZoneRange) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *TimedZoneRange) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *TimedZoneRange) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *TimedZoneRange) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *TimedZoneRange) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *TimedZoneRange) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetTime

`func (o *TimedZoneRange) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TimedZoneRange) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TimedZoneRange) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *TimedZoneRange) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


