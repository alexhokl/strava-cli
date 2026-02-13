# StreamSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**TimeStream**](TimeStream.md) |  | [optional] 
**Distance** | Pointer to [**DistanceStream**](DistanceStream.md) |  | [optional] 
**Latlng** | Pointer to [**LatLngStream**](LatLngStream.md) |  | [optional] 
**Altitude** | Pointer to [**AltitudeStream**](AltitudeStream.md) |  | [optional] 
**VelocitySmooth** | Pointer to [**SmoothVelocityStream**](SmoothVelocityStream.md) |  | [optional] 
**Heartrate** | Pointer to [**HeartrateStream**](HeartrateStream.md) |  | [optional] 
**Cadence** | Pointer to [**CadenceStream**](CadenceStream.md) |  | [optional] 
**Watts** | Pointer to [**PowerStream**](PowerStream.md) |  | [optional] 
**Temp** | Pointer to [**TemperatureStream**](TemperatureStream.md) |  | [optional] 
**Moving** | Pointer to [**MovingStream**](MovingStream.md) |  | [optional] 
**GradeSmooth** | Pointer to [**SmoothGradeStream**](SmoothGradeStream.md) |  | [optional] 

## Methods

### NewStreamSet

`func NewStreamSet() *StreamSet`

NewStreamSet instantiates a new StreamSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSetWithDefaults

`func NewStreamSetWithDefaults() *StreamSet`

NewStreamSetWithDefaults instantiates a new StreamSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *StreamSet) GetTime() TimeStream`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *StreamSet) GetTimeOk() (*TimeStream, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *StreamSet) SetTime(v TimeStream)`

SetTime sets Time field to given value.

### HasTime

`func (o *StreamSet) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDistance

`func (o *StreamSet) GetDistance() DistanceStream`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *StreamSet) GetDistanceOk() (*DistanceStream, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *StreamSet) SetDistance(v DistanceStream)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *StreamSet) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetLatlng

`func (o *StreamSet) GetLatlng() LatLngStream`

GetLatlng returns the Latlng field if non-nil, zero value otherwise.

### GetLatlngOk

`func (o *StreamSet) GetLatlngOk() (*LatLngStream, bool)`

GetLatlngOk returns a tuple with the Latlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlng

`func (o *StreamSet) SetLatlng(v LatLngStream)`

SetLatlng sets Latlng field to given value.

### HasLatlng

`func (o *StreamSet) HasLatlng() bool`

HasLatlng returns a boolean if a field has been set.

### GetAltitude

`func (o *StreamSet) GetAltitude() AltitudeStream`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *StreamSet) GetAltitudeOk() (*AltitudeStream, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *StreamSet) SetAltitude(v AltitudeStream)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *StreamSet) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetVelocitySmooth

`func (o *StreamSet) GetVelocitySmooth() SmoothVelocityStream`

GetVelocitySmooth returns the VelocitySmooth field if non-nil, zero value otherwise.

### GetVelocitySmoothOk

`func (o *StreamSet) GetVelocitySmoothOk() (*SmoothVelocityStream, bool)`

GetVelocitySmoothOk returns a tuple with the VelocitySmooth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocitySmooth

`func (o *StreamSet) SetVelocitySmooth(v SmoothVelocityStream)`

SetVelocitySmooth sets VelocitySmooth field to given value.

### HasVelocitySmooth

`func (o *StreamSet) HasVelocitySmooth() bool`

HasVelocitySmooth returns a boolean if a field has been set.

### GetHeartrate

`func (o *StreamSet) GetHeartrate() HeartrateStream`

GetHeartrate returns the Heartrate field if non-nil, zero value otherwise.

### GetHeartrateOk

`func (o *StreamSet) GetHeartrateOk() (*HeartrateStream, bool)`

GetHeartrateOk returns a tuple with the Heartrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartrate

`func (o *StreamSet) SetHeartrate(v HeartrateStream)`

SetHeartrate sets Heartrate field to given value.

### HasHeartrate

`func (o *StreamSet) HasHeartrate() bool`

HasHeartrate returns a boolean if a field has been set.

### GetCadence

`func (o *StreamSet) GetCadence() CadenceStream`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *StreamSet) GetCadenceOk() (*CadenceStream, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *StreamSet) SetCadence(v CadenceStream)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *StreamSet) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetWatts

`func (o *StreamSet) GetWatts() PowerStream`

GetWatts returns the Watts field if non-nil, zero value otherwise.

### GetWattsOk

`func (o *StreamSet) GetWattsOk() (*PowerStream, bool)`

GetWattsOk returns a tuple with the Watts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatts

`func (o *StreamSet) SetWatts(v PowerStream)`

SetWatts sets Watts field to given value.

### HasWatts

`func (o *StreamSet) HasWatts() bool`

HasWatts returns a boolean if a field has been set.

### GetTemp

`func (o *StreamSet) GetTemp() TemperatureStream`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *StreamSet) GetTempOk() (*TemperatureStream, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *StreamSet) SetTemp(v TemperatureStream)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *StreamSet) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetMoving

`func (o *StreamSet) GetMoving() MovingStream`

GetMoving returns the Moving field if non-nil, zero value otherwise.

### GetMovingOk

`func (o *StreamSet) GetMovingOk() (*MovingStream, bool)`

GetMovingOk returns a tuple with the Moving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoving

`func (o *StreamSet) SetMoving(v MovingStream)`

SetMoving sets Moving field to given value.

### HasMoving

`func (o *StreamSet) HasMoving() bool`

HasMoving returns a boolean if a field has been set.

### GetGradeSmooth

`func (o *StreamSet) GetGradeSmooth() SmoothGradeStream`

GetGradeSmooth returns the GradeSmooth field if non-nil, zero value otherwise.

### GetGradeSmoothOk

`func (o *StreamSet) GetGradeSmoothOk() (*SmoothGradeStream, bool)`

GetGradeSmoothOk returns a tuple with the GradeSmooth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeSmooth

`func (o *StreamSet) SetGradeSmooth(v SmoothGradeStream)`

SetGradeSmooth sets GradeSmooth field to given value.

### HasGradeSmooth

`func (o *StreamSet) HasGradeSmooth() bool`

HasGradeSmooth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


