# DistanceStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalSize** | Pointer to **int32** | The number of data points in this stream | [optional] 
**Resolution** | Pointer to **string** | The level of detail (sampling) in which this stream was returned | [optional] 
**SeriesType** | Pointer to **string** | The base series used in the case the stream was downsampled | [optional] 
**Data** | Pointer to **[]float32** | The sequence of distance values for this stream, in meters | [optional] 

## Methods

### NewDistanceStream

`func NewDistanceStream() *DistanceStream`

NewDistanceStream instantiates a new DistanceStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistanceStreamWithDefaults

`func NewDistanceStreamWithDefaults() *DistanceStream`

NewDistanceStreamWithDefaults instantiates a new DistanceStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalSize

`func (o *DistanceStream) GetOriginalSize() int32`

GetOriginalSize returns the OriginalSize field if non-nil, zero value otherwise.

### GetOriginalSizeOk

`func (o *DistanceStream) GetOriginalSizeOk() (*int32, bool)`

GetOriginalSizeOk returns a tuple with the OriginalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSize

`func (o *DistanceStream) SetOriginalSize(v int32)`

SetOriginalSize sets OriginalSize field to given value.

### HasOriginalSize

`func (o *DistanceStream) HasOriginalSize() bool`

HasOriginalSize returns a boolean if a field has been set.

### GetResolution

`func (o *DistanceStream) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *DistanceStream) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *DistanceStream) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *DistanceStream) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSeriesType

`func (o *DistanceStream) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *DistanceStream) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *DistanceStream) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *DistanceStream) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetData

`func (o *DistanceStream) GetData() []float32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DistanceStream) GetDataOk() (*[]float32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DistanceStream) SetData(v []float32)`

SetData sets Data field to given value.

### HasData

`func (o *DistanceStream) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


