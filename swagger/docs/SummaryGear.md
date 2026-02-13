# SummaryGear

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The gear&#39;s unique identifier. | [optional] 
**ResourceState** | Pointer to **int32** | Resource state, indicates level of detail. Possible values: 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] 
**Primary** | Pointer to **bool** | Whether this gear&#39;s is the owner&#39;s default one. | [optional] 
**Name** | Pointer to **string** | The gear&#39;s name. | [optional] 
**Distance** | Pointer to **float32** | The distance logged with this gear. | [optional] 

## Methods

### NewSummaryGear

`func NewSummaryGear() *SummaryGear`

NewSummaryGear instantiates a new SummaryGear object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryGearWithDefaults

`func NewSummaryGearWithDefaults() *SummaryGear`

NewSummaryGearWithDefaults instantiates a new SummaryGear object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SummaryGear) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummaryGear) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummaryGear) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SummaryGear) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceState

`func (o *SummaryGear) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *SummaryGear) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *SummaryGear) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *SummaryGear) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetPrimary

`func (o *SummaryGear) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *SummaryGear) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *SummaryGear) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *SummaryGear) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetName

`func (o *SummaryGear) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummaryGear) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummaryGear) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummaryGear) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistance

`func (o *SummaryGear) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SummaryGear) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SummaryGear) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SummaryGear) HasDistance() bool`

HasDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


