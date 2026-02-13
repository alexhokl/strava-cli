# SmoothGradeStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalSize** | Pointer to **int32** | The number of data points in this stream | [optional] 
**Resolution** | Pointer to **string** | The level of detail (sampling) in which this stream was returned | [optional] 
**SeriesType** | Pointer to **string** | The base series used in the case the stream was downsampled | [optional] 
**Data** | Pointer to **[]float32** | The sequence of grade values for this stream, as percents of a grade | [optional] 

## Methods

### NewSmoothGradeStream

`func NewSmoothGradeStream() *SmoothGradeStream`

NewSmoothGradeStream instantiates a new SmoothGradeStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmoothGradeStreamWithDefaults

`func NewSmoothGradeStreamWithDefaults() *SmoothGradeStream`

NewSmoothGradeStreamWithDefaults instantiates a new SmoothGradeStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalSize

`func (o *SmoothGradeStream) GetOriginalSize() int32`

GetOriginalSize returns the OriginalSize field if non-nil, zero value otherwise.

### GetOriginalSizeOk

`func (o *SmoothGradeStream) GetOriginalSizeOk() (*int32, bool)`

GetOriginalSizeOk returns a tuple with the OriginalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSize

`func (o *SmoothGradeStream) SetOriginalSize(v int32)`

SetOriginalSize sets OriginalSize field to given value.

### HasOriginalSize

`func (o *SmoothGradeStream) HasOriginalSize() bool`

HasOriginalSize returns a boolean if a field has been set.

### GetResolution

`func (o *SmoothGradeStream) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *SmoothGradeStream) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *SmoothGradeStream) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *SmoothGradeStream) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSeriesType

`func (o *SmoothGradeStream) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *SmoothGradeStream) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *SmoothGradeStream) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *SmoothGradeStream) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetData

`func (o *SmoothGradeStream) GetData() []float32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SmoothGradeStream) GetDataOk() (*[]float32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SmoothGradeStream) SetData(v []float32)`

SetData sets Data field to given value.

### HasData

`func (o *SmoothGradeStream) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


