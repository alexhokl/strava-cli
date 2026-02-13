# PhotosSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of photos | [optional] 
**Primary** | Pointer to [**PhotosSummaryPrimary**](PhotosSummaryPrimary.md) |  | [optional] 

## Methods

### NewPhotosSummary

`func NewPhotosSummary() *PhotosSummary`

NewPhotosSummary instantiates a new PhotosSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhotosSummaryWithDefaults

`func NewPhotosSummaryWithDefaults() *PhotosSummary`

NewPhotosSummaryWithDefaults instantiates a new PhotosSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PhotosSummary) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PhotosSummary) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PhotosSummary) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PhotosSummary) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPrimary

`func (o *PhotosSummary) GetPrimary() PhotosSummaryPrimary`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PhotosSummary) GetPrimaryOk() (*PhotosSummaryPrimary, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PhotosSummary) SetPrimary(v PhotosSummaryPrimary)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PhotosSummary) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


