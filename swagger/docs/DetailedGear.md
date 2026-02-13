# DetailedGear

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The gear&#39;s unique identifier. | [optional] 
**ResourceState** | Pointer to **int32** | Resource state, indicates level of detail. Possible values: 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] 
**Primary** | Pointer to **bool** | Whether this gear&#39;s is the owner&#39;s default one. | [optional] 
**Name** | Pointer to **string** | The gear&#39;s name. | [optional] 
**Distance** | Pointer to **float32** | The distance logged with this gear. | [optional] 
**BrandName** | Pointer to **string** | The gear&#39;s brand name. | [optional] 
**ModelName** | Pointer to **string** | The gear&#39;s model name. | [optional] 
**FrameType** | Pointer to **int32** | The gear&#39;s frame type (bike only). | [optional] 
**Description** | Pointer to **string** | The gear&#39;s description. | [optional] 

## Methods

### NewDetailedGear

`func NewDetailedGear() *DetailedGear`

NewDetailedGear instantiates a new DetailedGear object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedGearWithDefaults

`func NewDetailedGearWithDefaults() *DetailedGear`

NewDetailedGearWithDefaults instantiates a new DetailedGear object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedGear) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedGear) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedGear) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedGear) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceState

`func (o *DetailedGear) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *DetailedGear) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *DetailedGear) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *DetailedGear) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetPrimary

`func (o *DetailedGear) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *DetailedGear) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *DetailedGear) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *DetailedGear) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetName

`func (o *DetailedGear) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedGear) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedGear) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedGear) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistance

`func (o *DetailedGear) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *DetailedGear) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *DetailedGear) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *DetailedGear) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetBrandName

`func (o *DetailedGear) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *DetailedGear) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *DetailedGear) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *DetailedGear) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetModelName

`func (o *DetailedGear) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *DetailedGear) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *DetailedGear) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *DetailedGear) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetFrameType

`func (o *DetailedGear) GetFrameType() int32`

GetFrameType returns the FrameType field if non-nil, zero value otherwise.

### GetFrameTypeOk

`func (o *DetailedGear) GetFrameTypeOk() (*int32, bool)`

GetFrameTypeOk returns a tuple with the FrameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameType

`func (o *DetailedGear) SetFrameType(v int32)`

SetFrameType sets FrameType field to given value.

### HasFrameType

`func (o *DetailedGear) HasFrameType() bool`

HasFrameType returns a boolean if a field has been set.

### GetDescription

`func (o *DetailedGear) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DetailedGear) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DetailedGear) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DetailedGear) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


