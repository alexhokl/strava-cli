# AltitudeStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalSize** | Pointer to **int32** | The number of data points in this stream | [optional] 
**Resolution** | Pointer to **string** | The level of detail (sampling) in which this stream was returned | [optional] 
**SeriesType** | Pointer to **string** | The base series used in the case the stream was downsampled | [optional] 
**Data** | Pointer to **[]float32** | The sequence of altitude values for this stream, in meters | [optional] 

## Methods

### NewAltitudeStream

`func NewAltitudeStream() *AltitudeStream`

NewAltitudeStream instantiates a new AltitudeStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAltitudeStreamWithDefaults

`func NewAltitudeStreamWithDefaults() *AltitudeStream`

NewAltitudeStreamWithDefaults instantiates a new AltitudeStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalSize

`func (o *AltitudeStream) GetOriginalSize() int32`

GetOriginalSize returns the OriginalSize field if non-nil, zero value otherwise.

### GetOriginalSizeOk

`func (o *AltitudeStream) GetOriginalSizeOk() (*int32, bool)`

GetOriginalSizeOk returns a tuple with the OriginalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSize

`func (o *AltitudeStream) SetOriginalSize(v int32)`

SetOriginalSize sets OriginalSize field to given value.

### HasOriginalSize

`func (o *AltitudeStream) HasOriginalSize() bool`

HasOriginalSize returns a boolean if a field has been set.

### GetResolution

`func (o *AltitudeStream) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *AltitudeStream) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *AltitudeStream) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *AltitudeStream) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSeriesType

`func (o *AltitudeStream) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *AltitudeStream) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *AltitudeStream) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *AltitudeStream) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetData

`func (o *AltitudeStream) GetData() []float32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AltitudeStream) GetDataOk() (*[]float32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AltitudeStream) SetData(v []float32)`

SetData sets Data field to given value.

### HasData

`func (o *AltitudeStream) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


