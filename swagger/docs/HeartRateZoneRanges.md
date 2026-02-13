# HeartRateZoneRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomZones** | Pointer to **bool** | Whether the athlete has set their own custom heart rate zones | [optional] 
**Zones** | Pointer to [**[]ZoneRange**](ZoneRange.md) |  | [optional] 

## Methods

### NewHeartRateZoneRanges

`func NewHeartRateZoneRanges() *HeartRateZoneRanges`

NewHeartRateZoneRanges instantiates a new HeartRateZoneRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeartRateZoneRangesWithDefaults

`func NewHeartRateZoneRangesWithDefaults() *HeartRateZoneRanges`

NewHeartRateZoneRangesWithDefaults instantiates a new HeartRateZoneRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomZones

`func (o *HeartRateZoneRanges) GetCustomZones() bool`

GetCustomZones returns the CustomZones field if non-nil, zero value otherwise.

### GetCustomZonesOk

`func (o *HeartRateZoneRanges) GetCustomZonesOk() (*bool, bool)`

GetCustomZonesOk returns a tuple with the CustomZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomZones

`func (o *HeartRateZoneRanges) SetCustomZones(v bool)`

SetCustomZones sets CustomZones field to given value.

### HasCustomZones

`func (o *HeartRateZoneRanges) HasCustomZones() bool`

HasCustomZones returns a boolean if a field has been set.

### GetZones

`func (o *HeartRateZoneRanges) GetZones() []ZoneRange`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *HeartRateZoneRanges) GetZonesOk() (*[]ZoneRange, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *HeartRateZoneRanges) SetZones(v []ZoneRange)`

SetZones sets Zones field to given value.

### HasZones

`func (o *HeartRateZoneRanges) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


