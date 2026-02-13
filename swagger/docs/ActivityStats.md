# ActivityStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiggestRideDistance** | Pointer to **float64** | The longest distance ridden by the athlete. | [optional] 
**BiggestClimbElevationGain** | Pointer to **float64** | The highest climb ridden by the athlete. | [optional] 
**RecentRideTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**RecentRunTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**RecentSwimTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**YtdRideTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**YtdRunTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**YtdSwimTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**AllRideTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**AllRunTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 
**AllSwimTotals** | Pointer to [**ActivityTotal**](ActivityTotal.md) |  | [optional] 

## Methods

### NewActivityStats

`func NewActivityStats() *ActivityStats`

NewActivityStats instantiates a new ActivityStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityStatsWithDefaults

`func NewActivityStatsWithDefaults() *ActivityStats`

NewActivityStatsWithDefaults instantiates a new ActivityStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiggestRideDistance

`func (o *ActivityStats) GetBiggestRideDistance() float64`

GetBiggestRideDistance returns the BiggestRideDistance field if non-nil, zero value otherwise.

### GetBiggestRideDistanceOk

`func (o *ActivityStats) GetBiggestRideDistanceOk() (*float64, bool)`

GetBiggestRideDistanceOk returns a tuple with the BiggestRideDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiggestRideDistance

`func (o *ActivityStats) SetBiggestRideDistance(v float64)`

SetBiggestRideDistance sets BiggestRideDistance field to given value.

### HasBiggestRideDistance

`func (o *ActivityStats) HasBiggestRideDistance() bool`

HasBiggestRideDistance returns a boolean if a field has been set.

### GetBiggestClimbElevationGain

`func (o *ActivityStats) GetBiggestClimbElevationGain() float64`

GetBiggestClimbElevationGain returns the BiggestClimbElevationGain field if non-nil, zero value otherwise.

### GetBiggestClimbElevationGainOk

`func (o *ActivityStats) GetBiggestClimbElevationGainOk() (*float64, bool)`

GetBiggestClimbElevationGainOk returns a tuple with the BiggestClimbElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiggestClimbElevationGain

`func (o *ActivityStats) SetBiggestClimbElevationGain(v float64)`

SetBiggestClimbElevationGain sets BiggestClimbElevationGain field to given value.

### HasBiggestClimbElevationGain

`func (o *ActivityStats) HasBiggestClimbElevationGain() bool`

HasBiggestClimbElevationGain returns a boolean if a field has been set.

### GetRecentRideTotals

`func (o *ActivityStats) GetRecentRideTotals() ActivityTotal`

GetRecentRideTotals returns the RecentRideTotals field if non-nil, zero value otherwise.

### GetRecentRideTotalsOk

`func (o *ActivityStats) GetRecentRideTotalsOk() (*ActivityTotal, bool)`

GetRecentRideTotalsOk returns a tuple with the RecentRideTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentRideTotals

`func (o *ActivityStats) SetRecentRideTotals(v ActivityTotal)`

SetRecentRideTotals sets RecentRideTotals field to given value.

### HasRecentRideTotals

`func (o *ActivityStats) HasRecentRideTotals() bool`

HasRecentRideTotals returns a boolean if a field has been set.

### GetRecentRunTotals

`func (o *ActivityStats) GetRecentRunTotals() ActivityTotal`

GetRecentRunTotals returns the RecentRunTotals field if non-nil, zero value otherwise.

### GetRecentRunTotalsOk

`func (o *ActivityStats) GetRecentRunTotalsOk() (*ActivityTotal, bool)`

GetRecentRunTotalsOk returns a tuple with the RecentRunTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentRunTotals

`func (o *ActivityStats) SetRecentRunTotals(v ActivityTotal)`

SetRecentRunTotals sets RecentRunTotals field to given value.

### HasRecentRunTotals

`func (o *ActivityStats) HasRecentRunTotals() bool`

HasRecentRunTotals returns a boolean if a field has been set.

### GetRecentSwimTotals

`func (o *ActivityStats) GetRecentSwimTotals() ActivityTotal`

GetRecentSwimTotals returns the RecentSwimTotals field if non-nil, zero value otherwise.

### GetRecentSwimTotalsOk

`func (o *ActivityStats) GetRecentSwimTotalsOk() (*ActivityTotal, bool)`

GetRecentSwimTotalsOk returns a tuple with the RecentSwimTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentSwimTotals

`func (o *ActivityStats) SetRecentSwimTotals(v ActivityTotal)`

SetRecentSwimTotals sets RecentSwimTotals field to given value.

### HasRecentSwimTotals

`func (o *ActivityStats) HasRecentSwimTotals() bool`

HasRecentSwimTotals returns a boolean if a field has been set.

### GetYtdRideTotals

`func (o *ActivityStats) GetYtdRideTotals() ActivityTotal`

GetYtdRideTotals returns the YtdRideTotals field if non-nil, zero value otherwise.

### GetYtdRideTotalsOk

`func (o *ActivityStats) GetYtdRideTotalsOk() (*ActivityTotal, bool)`

GetYtdRideTotalsOk returns a tuple with the YtdRideTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdRideTotals

`func (o *ActivityStats) SetYtdRideTotals(v ActivityTotal)`

SetYtdRideTotals sets YtdRideTotals field to given value.

### HasYtdRideTotals

`func (o *ActivityStats) HasYtdRideTotals() bool`

HasYtdRideTotals returns a boolean if a field has been set.

### GetYtdRunTotals

`func (o *ActivityStats) GetYtdRunTotals() ActivityTotal`

GetYtdRunTotals returns the YtdRunTotals field if non-nil, zero value otherwise.

### GetYtdRunTotalsOk

`func (o *ActivityStats) GetYtdRunTotalsOk() (*ActivityTotal, bool)`

GetYtdRunTotalsOk returns a tuple with the YtdRunTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdRunTotals

`func (o *ActivityStats) SetYtdRunTotals(v ActivityTotal)`

SetYtdRunTotals sets YtdRunTotals field to given value.

### HasYtdRunTotals

`func (o *ActivityStats) HasYtdRunTotals() bool`

HasYtdRunTotals returns a boolean if a field has been set.

### GetYtdSwimTotals

`func (o *ActivityStats) GetYtdSwimTotals() ActivityTotal`

GetYtdSwimTotals returns the YtdSwimTotals field if non-nil, zero value otherwise.

### GetYtdSwimTotalsOk

`func (o *ActivityStats) GetYtdSwimTotalsOk() (*ActivityTotal, bool)`

GetYtdSwimTotalsOk returns a tuple with the YtdSwimTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdSwimTotals

`func (o *ActivityStats) SetYtdSwimTotals(v ActivityTotal)`

SetYtdSwimTotals sets YtdSwimTotals field to given value.

### HasYtdSwimTotals

`func (o *ActivityStats) HasYtdSwimTotals() bool`

HasYtdSwimTotals returns a boolean if a field has been set.

### GetAllRideTotals

`func (o *ActivityStats) GetAllRideTotals() ActivityTotal`

GetAllRideTotals returns the AllRideTotals field if non-nil, zero value otherwise.

### GetAllRideTotalsOk

`func (o *ActivityStats) GetAllRideTotalsOk() (*ActivityTotal, bool)`

GetAllRideTotalsOk returns a tuple with the AllRideTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRideTotals

`func (o *ActivityStats) SetAllRideTotals(v ActivityTotal)`

SetAllRideTotals sets AllRideTotals field to given value.

### HasAllRideTotals

`func (o *ActivityStats) HasAllRideTotals() bool`

HasAllRideTotals returns a boolean if a field has been set.

### GetAllRunTotals

`func (o *ActivityStats) GetAllRunTotals() ActivityTotal`

GetAllRunTotals returns the AllRunTotals field if non-nil, zero value otherwise.

### GetAllRunTotalsOk

`func (o *ActivityStats) GetAllRunTotalsOk() (*ActivityTotal, bool)`

GetAllRunTotalsOk returns a tuple with the AllRunTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRunTotals

`func (o *ActivityStats) SetAllRunTotals(v ActivityTotal)`

SetAllRunTotals sets AllRunTotals field to given value.

### HasAllRunTotals

`func (o *ActivityStats) HasAllRunTotals() bool`

HasAllRunTotals returns a boolean if a field has been set.

### GetAllSwimTotals

`func (o *ActivityStats) GetAllSwimTotals() ActivityTotal`

GetAllSwimTotals returns the AllSwimTotals field if non-nil, zero value otherwise.

### GetAllSwimTotalsOk

`func (o *ActivityStats) GetAllSwimTotalsOk() (*ActivityTotal, bool)`

GetAllSwimTotalsOk returns a tuple with the AllSwimTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSwimTotals

`func (o *ActivityStats) SetAllSwimTotals(v ActivityTotal)`

SetAllSwimTotals sets AllSwimTotals field to given value.

### HasAllSwimTotals

`func (o *ActivityStats) HasAllSwimTotals() bool`

HasAllSwimTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


