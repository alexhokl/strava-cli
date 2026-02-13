# UpdatableActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commute** | Pointer to **bool** | Whether this activity is a commute | [optional] 
**Trainer** | Pointer to **bool** | Whether this activity was recorded on a training machine | [optional] 
**HideFromHome** | Pointer to **bool** | Whether this activity is muted | [optional] 
**Description** | Pointer to **string** | The description of the activity | [optional] 
**Name** | Pointer to **string** | The name of the activity | [optional] 
**Type** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**SportType** | Pointer to [**SportType**](SportType.md) |  | [optional] 
**GearId** | Pointer to **string** | Identifier for the gear associated with the activity. ‘none’ clears gear from activity | [optional] 

## Methods

### NewUpdatableActivity

`func NewUpdatableActivity() *UpdatableActivity`

NewUpdatableActivity instantiates a new UpdatableActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatableActivityWithDefaults

`func NewUpdatableActivityWithDefaults() *UpdatableActivity`

NewUpdatableActivityWithDefaults instantiates a new UpdatableActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommute

`func (o *UpdatableActivity) GetCommute() bool`

GetCommute returns the Commute field if non-nil, zero value otherwise.

### GetCommuteOk

`func (o *UpdatableActivity) GetCommuteOk() (*bool, bool)`

GetCommuteOk returns a tuple with the Commute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommute

`func (o *UpdatableActivity) SetCommute(v bool)`

SetCommute sets Commute field to given value.

### HasCommute

`func (o *UpdatableActivity) HasCommute() bool`

HasCommute returns a boolean if a field has been set.

### GetTrainer

`func (o *UpdatableActivity) GetTrainer() bool`

GetTrainer returns the Trainer field if non-nil, zero value otherwise.

### GetTrainerOk

`func (o *UpdatableActivity) GetTrainerOk() (*bool, bool)`

GetTrainerOk returns a tuple with the Trainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainer

`func (o *UpdatableActivity) SetTrainer(v bool)`

SetTrainer sets Trainer field to given value.

### HasTrainer

`func (o *UpdatableActivity) HasTrainer() bool`

HasTrainer returns a boolean if a field has been set.

### GetHideFromHome

`func (o *UpdatableActivity) GetHideFromHome() bool`

GetHideFromHome returns the HideFromHome field if non-nil, zero value otherwise.

### GetHideFromHomeOk

`func (o *UpdatableActivity) GetHideFromHomeOk() (*bool, bool)`

GetHideFromHomeOk returns a tuple with the HideFromHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromHome

`func (o *UpdatableActivity) SetHideFromHome(v bool)`

SetHideFromHome sets HideFromHome field to given value.

### HasHideFromHome

`func (o *UpdatableActivity) HasHideFromHome() bool`

HasHideFromHome returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatableActivity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatableActivity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatableActivity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatableActivity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdatableActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatableActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatableActivity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatableActivity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UpdatableActivity) GetType() ActivityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdatableActivity) GetTypeOk() (*ActivityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdatableActivity) SetType(v ActivityType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdatableActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSportType

`func (o *UpdatableActivity) GetSportType() SportType`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *UpdatableActivity) GetSportTypeOk() (*SportType, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *UpdatableActivity) SetSportType(v SportType)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *UpdatableActivity) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetGearId

`func (o *UpdatableActivity) GetGearId() string`

GetGearId returns the GearId field if non-nil, zero value otherwise.

### GetGearIdOk

`func (o *UpdatableActivity) GetGearIdOk() (*string, bool)`

GetGearIdOk returns a tuple with the GearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearId

`func (o *UpdatableActivity) SetGearId(v string)`

SetGearId sets GearId field to given value.

### HasGearId

`func (o *UpdatableActivity) HasGearId() bool`

HasGearId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


