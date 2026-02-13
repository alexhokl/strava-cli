# DetailedAthlete

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
**FollowerCount** | Pointer to **int32** | The athlete&#39;s follower count. | [optional] 
**FriendCount** | Pointer to **int32** | The athlete&#39;s friend count. | [optional] 
**MeasurementPreference** | Pointer to **string** | The athlete&#39;s preferred unit system. | [optional] 
**Ftp** | Pointer to **int32** | The athlete&#39;s FTP (Functional Threshold Power). | [optional] 
**Weight** | Pointer to **float32** | The athlete&#39;s weight. | [optional] 
**Clubs** | Pointer to [**[]SummaryClub**](SummaryClub.md) | The athlete&#39;s clubs. | [optional] 
**Bikes** | Pointer to [**[]SummaryGear**](SummaryGear.md) | The athlete&#39;s bikes. | [optional] 
**Shoes** | Pointer to [**[]SummaryGear**](SummaryGear.md) | The athlete&#39;s shoes. | [optional] 

## Methods

### NewDetailedAthlete

`func NewDetailedAthlete() *DetailedAthlete`

NewDetailedAthlete instantiates a new DetailedAthlete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedAthleteWithDefaults

`func NewDetailedAthleteWithDefaults() *DetailedAthlete`

NewDetailedAthleteWithDefaults instantiates a new DetailedAthlete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedAthlete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedAthlete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedAthlete) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedAthlete) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceState

`func (o *DetailedAthlete) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *DetailedAthlete) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *DetailedAthlete) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *DetailedAthlete) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetFirstname

`func (o *DetailedAthlete) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *DetailedAthlete) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *DetailedAthlete) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *DetailedAthlete) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *DetailedAthlete) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *DetailedAthlete) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *DetailedAthlete) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *DetailedAthlete) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetProfileMedium

`func (o *DetailedAthlete) GetProfileMedium() string`

GetProfileMedium returns the ProfileMedium field if non-nil, zero value otherwise.

### GetProfileMediumOk

`func (o *DetailedAthlete) GetProfileMediumOk() (*string, bool)`

GetProfileMediumOk returns a tuple with the ProfileMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMedium

`func (o *DetailedAthlete) SetProfileMedium(v string)`

SetProfileMedium sets ProfileMedium field to given value.

### HasProfileMedium

`func (o *DetailedAthlete) HasProfileMedium() bool`

HasProfileMedium returns a boolean if a field has been set.

### GetProfile

`func (o *DetailedAthlete) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DetailedAthlete) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DetailedAthlete) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *DetailedAthlete) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCity

`func (o *DetailedAthlete) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DetailedAthlete) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DetailedAthlete) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DetailedAthlete) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *DetailedAthlete) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DetailedAthlete) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DetailedAthlete) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DetailedAthlete) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *DetailedAthlete) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DetailedAthlete) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DetailedAthlete) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DetailedAthlete) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSex

`func (o *DetailedAthlete) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *DetailedAthlete) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *DetailedAthlete) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *DetailedAthlete) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetPremium

`func (o *DetailedAthlete) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *DetailedAthlete) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *DetailedAthlete) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *DetailedAthlete) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetSummit

`func (o *DetailedAthlete) GetSummit() bool`

GetSummit returns the Summit field if non-nil, zero value otherwise.

### GetSummitOk

`func (o *DetailedAthlete) GetSummitOk() (*bool, bool)`

GetSummitOk returns a tuple with the Summit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummit

`func (o *DetailedAthlete) SetSummit(v bool)`

SetSummit sets Summit field to given value.

### HasSummit

`func (o *DetailedAthlete) HasSummit() bool`

HasSummit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DetailedAthlete) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DetailedAthlete) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DetailedAthlete) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DetailedAthlete) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DetailedAthlete) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DetailedAthlete) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DetailedAthlete) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DetailedAthlete) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFollowerCount

`func (o *DetailedAthlete) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *DetailedAthlete) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *DetailedAthlete) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *DetailedAthlete) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetFriendCount

`func (o *DetailedAthlete) GetFriendCount() int32`

GetFriendCount returns the FriendCount field if non-nil, zero value otherwise.

### GetFriendCountOk

`func (o *DetailedAthlete) GetFriendCountOk() (*int32, bool)`

GetFriendCountOk returns a tuple with the FriendCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendCount

`func (o *DetailedAthlete) SetFriendCount(v int32)`

SetFriendCount sets FriendCount field to given value.

### HasFriendCount

`func (o *DetailedAthlete) HasFriendCount() bool`

HasFriendCount returns a boolean if a field has been set.

### GetMeasurementPreference

`func (o *DetailedAthlete) GetMeasurementPreference() string`

GetMeasurementPreference returns the MeasurementPreference field if non-nil, zero value otherwise.

### GetMeasurementPreferenceOk

`func (o *DetailedAthlete) GetMeasurementPreferenceOk() (*string, bool)`

GetMeasurementPreferenceOk returns a tuple with the MeasurementPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPreference

`func (o *DetailedAthlete) SetMeasurementPreference(v string)`

SetMeasurementPreference sets MeasurementPreference field to given value.

### HasMeasurementPreference

`func (o *DetailedAthlete) HasMeasurementPreference() bool`

HasMeasurementPreference returns a boolean if a field has been set.

### GetFtp

`func (o *DetailedAthlete) GetFtp() int32`

GetFtp returns the Ftp field if non-nil, zero value otherwise.

### GetFtpOk

`func (o *DetailedAthlete) GetFtpOk() (*int32, bool)`

GetFtpOk returns a tuple with the Ftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtp

`func (o *DetailedAthlete) SetFtp(v int32)`

SetFtp sets Ftp field to given value.

### HasFtp

`func (o *DetailedAthlete) HasFtp() bool`

HasFtp returns a boolean if a field has been set.

### GetWeight

`func (o *DetailedAthlete) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DetailedAthlete) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DetailedAthlete) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DetailedAthlete) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetClubs

`func (o *DetailedAthlete) GetClubs() []SummaryClub`

GetClubs returns the Clubs field if non-nil, zero value otherwise.

### GetClubsOk

`func (o *DetailedAthlete) GetClubsOk() (*[]SummaryClub, bool)`

GetClubsOk returns a tuple with the Clubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClubs

`func (o *DetailedAthlete) SetClubs(v []SummaryClub)`

SetClubs sets Clubs field to given value.

### HasClubs

`func (o *DetailedAthlete) HasClubs() bool`

HasClubs returns a boolean if a field has been set.

### GetBikes

`func (o *DetailedAthlete) GetBikes() []SummaryGear`

GetBikes returns the Bikes field if non-nil, zero value otherwise.

### GetBikesOk

`func (o *DetailedAthlete) GetBikesOk() (*[]SummaryGear, bool)`

GetBikesOk returns a tuple with the Bikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBikes

`func (o *DetailedAthlete) SetBikes(v []SummaryGear)`

SetBikes sets Bikes field to given value.

### HasBikes

`func (o *DetailedAthlete) HasBikes() bool`

HasBikes returns a boolean if a field has been set.

### GetShoes

`func (o *DetailedAthlete) GetShoes() []SummaryGear`

GetShoes returns the Shoes field if non-nil, zero value otherwise.

### GetShoesOk

`func (o *DetailedAthlete) GetShoesOk() (*[]SummaryGear, bool)`

GetShoesOk returns a tuple with the Shoes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoes

`func (o *DetailedAthlete) SetShoes(v []SummaryGear)`

SetShoes sets Shoes field to given value.

### HasShoes

`func (o *DetailedAthlete) HasShoes() bool`

HasShoes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


