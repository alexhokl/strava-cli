# Zones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeartRate** | Pointer to [**HeartRateZoneRanges**](HeartRateZoneRanges.md) |  | [optional] 
**Power** | Pointer to [**PowerZoneRanges**](PowerZoneRanges.md) |  | [optional] 

## Methods

### NewZones

`func NewZones() *Zones`

NewZones instantiates a new Zones object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonesWithDefaults

`func NewZonesWithDefaults() *Zones`

NewZonesWithDefaults instantiates a new Zones object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeartRate

`func (o *Zones) GetHeartRate() HeartRateZoneRanges`

GetHeartRate returns the HeartRate field if non-nil, zero value otherwise.

### GetHeartRateOk

`func (o *Zones) GetHeartRateOk() (*HeartRateZoneRanges, bool)`

GetHeartRateOk returns a tuple with the HeartRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartRate

`func (o *Zones) SetHeartRate(v HeartRateZoneRanges)`

SetHeartRate sets HeartRate field to given value.

### HasHeartRate

`func (o *Zones) HasHeartRate() bool`

HasHeartRate returns a boolean if a field has been set.

### GetPower

`func (o *Zones) GetPower() PowerZoneRanges`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *Zones) GetPowerOk() (*PowerZoneRanges, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *Zones) SetPower(v PowerZoneRanges)`

SetPower sets Power field to given value.

### HasPower

`func (o *Zones) HasPower() bool`

HasPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


