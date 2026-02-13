# ExplorerSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of this segment | [optional] 
**Name** | Pointer to **string** | The name of this segment | [optional] 
**ClimbCategory** | Pointer to **int32** | The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category. If climb_category &#x3D; 5, climb_category_desc &#x3D; HC. If climb_category &#x3D; 2, climb_category_desc &#x3D; 3. | [optional] 
**ClimbCategoryDesc** | Pointer to **string** | The description for the category of the climb | [optional] 
**AvgGrade** | Pointer to **float32** | The segment&#39;s average grade, in percents | [optional] 
**StartLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**EndLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**ElevDifference** | Pointer to **float32** | The segments&#39;s evelation difference, in meters | [optional] 
**Distance** | Pointer to **float32** | The segment&#39;s distance, in meters | [optional] 
**Points** | Pointer to **string** | The polyline of the segment | [optional] 

## Methods

### NewExplorerSegment

`func NewExplorerSegment() *ExplorerSegment`

NewExplorerSegment instantiates a new ExplorerSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplorerSegmentWithDefaults

`func NewExplorerSegmentWithDefaults() *ExplorerSegment`

NewExplorerSegmentWithDefaults instantiates a new ExplorerSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExplorerSegment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExplorerSegment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExplorerSegment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExplorerSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExplorerSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExplorerSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExplorerSegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExplorerSegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClimbCategory

`func (o *ExplorerSegment) GetClimbCategory() int32`

GetClimbCategory returns the ClimbCategory field if non-nil, zero value otherwise.

### GetClimbCategoryOk

`func (o *ExplorerSegment) GetClimbCategoryOk() (*int32, bool)`

GetClimbCategoryOk returns a tuple with the ClimbCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimbCategory

`func (o *ExplorerSegment) SetClimbCategory(v int32)`

SetClimbCategory sets ClimbCategory field to given value.

### HasClimbCategory

`func (o *ExplorerSegment) HasClimbCategory() bool`

HasClimbCategory returns a boolean if a field has been set.

### GetClimbCategoryDesc

`func (o *ExplorerSegment) GetClimbCategoryDesc() string`

GetClimbCategoryDesc returns the ClimbCategoryDesc field if non-nil, zero value otherwise.

### GetClimbCategoryDescOk

`func (o *ExplorerSegment) GetClimbCategoryDescOk() (*string, bool)`

GetClimbCategoryDescOk returns a tuple with the ClimbCategoryDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClimbCategoryDesc

`func (o *ExplorerSegment) SetClimbCategoryDesc(v string)`

SetClimbCategoryDesc sets ClimbCategoryDesc field to given value.

### HasClimbCategoryDesc

`func (o *ExplorerSegment) HasClimbCategoryDesc() bool`

HasClimbCategoryDesc returns a boolean if a field has been set.

### GetAvgGrade

`func (o *ExplorerSegment) GetAvgGrade() float32`

GetAvgGrade returns the AvgGrade field if non-nil, zero value otherwise.

### GetAvgGradeOk

`func (o *ExplorerSegment) GetAvgGradeOk() (*float32, bool)`

GetAvgGradeOk returns a tuple with the AvgGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgGrade

`func (o *ExplorerSegment) SetAvgGrade(v float32)`

SetAvgGrade sets AvgGrade field to given value.

### HasAvgGrade

`func (o *ExplorerSegment) HasAvgGrade() bool`

HasAvgGrade returns a boolean if a field has been set.

### GetStartLatlng

`func (o *ExplorerSegment) GetStartLatlng() []float32`

GetStartLatlng returns the StartLatlng field if non-nil, zero value otherwise.

### GetStartLatlngOk

`func (o *ExplorerSegment) GetStartLatlngOk() (*[]float32, bool)`

GetStartLatlngOk returns a tuple with the StartLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLatlng

`func (o *ExplorerSegment) SetStartLatlng(v []float32)`

SetStartLatlng sets StartLatlng field to given value.

### HasStartLatlng

`func (o *ExplorerSegment) HasStartLatlng() bool`

HasStartLatlng returns a boolean if a field has been set.

### GetEndLatlng

`func (o *ExplorerSegment) GetEndLatlng() []float32`

GetEndLatlng returns the EndLatlng field if non-nil, zero value otherwise.

### GetEndLatlngOk

`func (o *ExplorerSegment) GetEndLatlngOk() (*[]float32, bool)`

GetEndLatlngOk returns a tuple with the EndLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLatlng

`func (o *ExplorerSegment) SetEndLatlng(v []float32)`

SetEndLatlng sets EndLatlng field to given value.

### HasEndLatlng

`func (o *ExplorerSegment) HasEndLatlng() bool`

HasEndLatlng returns a boolean if a field has been set.

### GetElevDifference

`func (o *ExplorerSegment) GetElevDifference() float32`

GetElevDifference returns the ElevDifference field if non-nil, zero value otherwise.

### GetElevDifferenceOk

`func (o *ExplorerSegment) GetElevDifferenceOk() (*float32, bool)`

GetElevDifferenceOk returns a tuple with the ElevDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevDifference

`func (o *ExplorerSegment) SetElevDifference(v float32)`

SetElevDifference sets ElevDifference field to given value.

### HasElevDifference

`func (o *ExplorerSegment) HasElevDifference() bool`

HasElevDifference returns a boolean if a field has been set.

### GetDistance

`func (o *ExplorerSegment) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *ExplorerSegment) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *ExplorerSegment) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *ExplorerSegment) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetPoints

`func (o *ExplorerSegment) GetPoints() string`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *ExplorerSegment) GetPointsOk() (*string, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *ExplorerSegment) SetPoints(v string)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *ExplorerSegment) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


