# SummaryAthlete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the athlete | [optional] 
**ResourceState** | Pointer to **int32** | Resource state, indicates level of detail. Possible values: 1 -&gt; \&quot;meta\&quot;, 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] 
**Firstname** | Pointer to **string** | The athlete&#39;s first name. | [optional] 
**Lastname** | Pointer to **string** | The athlete&#39;s last name. | [optional] 
**ProfileMedium** | Pointer to **string** | URL to a 62x62 pixel profile picture. | [optional] 
**Profile** | Pointer to **string** | URL to a 124x124 pixel profile picture. | [optional] 
**City** | Pointer to **string** | The athlete&#39;s city. | [optional] 
**State** | Pointer to **string** | The athlete&#39;s state or geographical region. | [optional] 
**Country** | Pointer to **string** | The athlete&#39;s country. | [optional] 
**Sex** | Pointer to **string** | The athlete&#39;s sex. | [optional] 
**Premium** | Pointer to **bool** | Deprecated.  Use summit field instead. Whether the athlete has any Summit subscription. | [optional] 
**Summit** | Pointer to **bool** | Whether the athlete has any Summit subscription. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which the athlete was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which the athlete was last updated. | [optional] 

## Methods

### NewSummaryAthlete

`func NewSummaryAthlete() *SummaryAthlete`

NewSummaryAthlete instantiates a new SummaryAthlete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryAthleteWithDefaults

`func NewSummaryAthleteWithDefaults() *SummaryAthlete`

NewSummaryAthleteWithDefaults instantiates a new SummaryAthlete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SummaryAthlete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummaryAthlete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummaryAthlete) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SummaryAthlete) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceState

`func (o *SummaryAthlete) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *SummaryAthlete) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *SummaryAthlete) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *SummaryAthlete) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetFirstname

`func (o *SummaryAthlete) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *SummaryAthlete) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *SummaryAthlete) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *SummaryAthlete) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *SummaryAthlete) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *SummaryAthlete) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *SummaryAthlete) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *SummaryAthlete) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetProfileMedium

`func (o *SummaryAthlete) GetProfileMedium() string`

GetProfileMedium returns the ProfileMedium field if non-nil, zero value otherwise.

### GetProfileMediumOk

`func (o *SummaryAthlete) GetProfileMediumOk() (*string, bool)`

GetProfileMediumOk returns a tuple with the ProfileMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMedium

`func (o *SummaryAthlete) SetProfileMedium(v string)`

SetProfileMedium sets ProfileMedium field to given value.

### HasProfileMedium

`func (o *SummaryAthlete) HasProfileMedium() bool`

HasProfileMedium returns a boolean if a field has been set.

### GetProfile

`func (o *SummaryAthlete) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SummaryAthlete) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SummaryAthlete) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SummaryAthlete) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCity

`func (o *SummaryAthlete) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SummaryAthlete) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SummaryAthlete) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SummaryAthlete) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *SummaryAthlete) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SummaryAthlete) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SummaryAthlete) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SummaryAthlete) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *SummaryAthlete) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SummaryAthlete) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SummaryAthlete) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SummaryAthlete) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSex

`func (o *SummaryAthlete) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SummaryAthlete) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SummaryAthlete) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SummaryAthlete) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetPremium

`func (o *SummaryAthlete) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *SummaryAthlete) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *SummaryAthlete) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *SummaryAthlete) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetSummit

`func (o *SummaryAthlete) GetSummit() bool`

GetSummit returns the Summit field if non-nil, zero value otherwise.

### GetSummitOk

`func (o *SummaryAthlete) GetSummitOk() (*bool, bool)`

GetSummitOk returns a tuple with the Summit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummit

`func (o *SummaryAthlete) SetSummit(v bool)`

SetSummit sets Summit field to given value.

### HasSummit

`func (o *SummaryAthlete) HasSummit() bool`

HasSummit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SummaryAthlete) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SummaryAthlete) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SummaryAthlete) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SummaryAthlete) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SummaryAthlete) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SummaryAthlete) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SummaryAthlete) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SummaryAthlete) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


