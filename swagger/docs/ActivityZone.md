# ActivityZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int32** |  | [optional] 
**DistributionBuckets** | Pointer to [**[]TimedZoneRange**](TimedZoneRange.md) | Stores the exclusive ranges representing zones and the time spent in each. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SensorBased** | Pointer to **bool** |  | [optional] 
**Points** | Pointer to **int32** |  | [optional] 
**CustomZones** | Pointer to **bool** |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 

## Methods

### NewActivityZone

`func NewActivityZone() *ActivityZone`

NewActivityZone instantiates a new ActivityZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityZoneWithDefaults

`func NewActivityZoneWithDefaults() *ActivityZone`

NewActivityZoneWithDefaults instantiates a new ActivityZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *ActivityZone) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ActivityZone) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ActivityZone) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ActivityZone) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDistributionBuckets

`func (o *ActivityZone) GetDistributionBuckets() []TimedZoneRange`

GetDistributionBuckets returns the DistributionBuckets field if non-nil, zero value otherwise.

### GetDistributionBucketsOk

`func (o *ActivityZone) GetDistributionBucketsOk() (*[]TimedZoneRange, bool)`

GetDistributionBucketsOk returns a tuple with the DistributionBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionBuckets

`func (o *ActivityZone) SetDistributionBuckets(v []TimedZoneRange)`

SetDistributionBuckets sets DistributionBuckets field to given value.

### HasDistributionBuckets

`func (o *ActivityZone) HasDistributionBuckets() bool`

HasDistributionBuckets returns a boolean if a field has been set.

### GetType

`func (o *ActivityZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSensorBased

`func (o *ActivityZone) GetSensorBased() bool`

GetSensorBased returns the SensorBased field if non-nil, zero value otherwise.

### GetSensorBasedOk

`func (o *ActivityZone) GetSensorBasedOk() (*bool, bool)`

GetSensorBasedOk returns a tuple with the SensorBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorBased

`func (o *ActivityZone) SetSensorBased(v bool)`

SetSensorBased sets SensorBased field to given value.

### HasSensorBased

`func (o *ActivityZone) HasSensorBased() bool`

HasSensorBased returns a boolean if a field has been set.

### GetPoints

`func (o *ActivityZone) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *ActivityZone) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *ActivityZone) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *ActivityZone) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetCustomZones

`func (o *ActivityZone) GetCustomZones() bool`

GetCustomZones returns the CustomZones field if non-nil, zero value otherwise.

### GetCustomZonesOk

`func (o *ActivityZone) GetCustomZonesOk() (*bool, bool)`

GetCustomZonesOk returns a tuple with the CustomZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomZones

`func (o *ActivityZone) SetCustomZones(v bool)`

SetCustomZones sets CustomZones field to given value.

### HasCustomZones

`func (o *ActivityZone) HasCustomZones() bool`

HasCustomZones returns a boolean if a field has been set.

### GetMax

`func (o *ActivityZone) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ActivityZone) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ActivityZone) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *ActivityZone) HasMax() bool`

HasMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


