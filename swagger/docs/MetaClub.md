# MetaClub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The club&#39;s unique identifier. | [optional] 
**ResourceState** | Pointer to **int32** | Resource state, indicates level of detail. Possible values: 1 -&gt; \&quot;meta\&quot;, 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] 
**Name** | Pointer to **string** | The club&#39;s name. | [optional] 

## Methods

### NewMetaClub

`func NewMetaClub() *MetaClub`

NewMetaClub instantiates a new MetaClub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaClubWithDefaults

`func NewMetaClubWithDefaults() *MetaClub`

NewMetaClubWithDefaults instantiates a new MetaClub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetaClub) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetaClub) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetaClub) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetaClub) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceState

`func (o *MetaClub) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *MetaClub) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *MetaClub) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *MetaClub) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetName

`func (o *MetaClub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaClub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaClub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaClub) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


