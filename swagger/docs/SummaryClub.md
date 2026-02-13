# SummaryClub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The club&#39;s unique identifier. | [optional] 
**ResourceState** | Pointer to **int32** | Resource state, indicates level of detail. Possible values: 1 -&gt; \&quot;meta\&quot;, 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] 
**Name** | Pointer to **string** | The club&#39;s name. | [optional] 
**ProfileMedium** | Pointer to **string** | URL to a 60x60 pixel profile picture. | [optional] 
**CoverPhoto** | Pointer to **string** | URL to a ~1185x580 pixel cover photo. | [optional] 
**CoverPhotoSmall** | Pointer to **string** | URL to a ~360x176  pixel cover photo. | [optional] 
**SportType** | Pointer to **string** | Deprecated. Prefer to use activity_types. | [optional] 
**ActivityTypes** | Pointer to [**[]ActivityType**](ActivityType.md) | The activity types that count for a club. This takes precedence over sport_type. | [optional] 
**City** | Pointer to **string** | The club&#39;s city. | [optional] 
**State** | Pointer to **string** | The club&#39;s state or geographical region. | [optional] 
**Country** | Pointer to **string** | The club&#39;s country. | [optional] 
**Private** | Pointer to **bool** | Whether the club is private. | [optional] 
**MemberCount** | Pointer to **int32** | The club&#39;s member count. | [optional] 
**Featured** | Pointer to **bool** | Whether the club is featured or not. | [optional] 
**Verified** | Pointer to **bool** | Whether the club is verified or not. | [optional] 
**Url** | Pointer to **string** | The club&#39;s vanity URL. | [optional] 

## Methods

### NewSummaryClub

`func NewSummaryClub() *SummaryClub`

NewSummaryClub instantiates a new SummaryClub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryClubWithDefaults

`func NewSummaryClubWithDefaults() *SummaryClub`

NewSummaryClubWithDefaults instantiates a new SummaryClub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SummaryClub) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummaryClub) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummaryClub) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SummaryClub) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceState

`func (o *SummaryClub) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *SummaryClub) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *SummaryClub) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *SummaryClub) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetName

`func (o *SummaryClub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummaryClub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummaryClub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummaryClub) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileMedium

`func (o *SummaryClub) GetProfileMedium() string`

GetProfileMedium returns the ProfileMedium field if non-nil, zero value otherwise.

### GetProfileMediumOk

`func (o *SummaryClub) GetProfileMediumOk() (*string, bool)`

GetProfileMediumOk returns a tuple with the ProfileMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMedium

`func (o *SummaryClub) SetProfileMedium(v string)`

SetProfileMedium sets ProfileMedium field to given value.

### HasProfileMedium

`func (o *SummaryClub) HasProfileMedium() bool`

HasProfileMedium returns a boolean if a field has been set.

### GetCoverPhoto

`func (o *SummaryClub) GetCoverPhoto() string`

GetCoverPhoto returns the CoverPhoto field if non-nil, zero value otherwise.

### GetCoverPhotoOk

`func (o *SummaryClub) GetCoverPhotoOk() (*string, bool)`

GetCoverPhotoOk returns a tuple with the CoverPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverPhoto

`func (o *SummaryClub) SetCoverPhoto(v string)`

SetCoverPhoto sets CoverPhoto field to given value.

### HasCoverPhoto

`func (o *SummaryClub) HasCoverPhoto() bool`

HasCoverPhoto returns a boolean if a field has been set.

### GetCoverPhotoSmall

`func (o *SummaryClub) GetCoverPhotoSmall() string`

GetCoverPhotoSmall returns the CoverPhotoSmall field if non-nil, zero value otherwise.

### GetCoverPhotoSmallOk

`func (o *SummaryClub) GetCoverPhotoSmallOk() (*string, bool)`

GetCoverPhotoSmallOk returns a tuple with the CoverPhotoSmall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverPhotoSmall

`func (o *SummaryClub) SetCoverPhotoSmall(v string)`

SetCoverPhotoSmall sets CoverPhotoSmall field to given value.

### HasCoverPhotoSmall

`func (o *SummaryClub) HasCoverPhotoSmall() bool`

HasCoverPhotoSmall returns a boolean if a field has been set.

### GetSportType

`func (o *SummaryClub) GetSportType() string`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *SummaryClub) GetSportTypeOk() (*string, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *SummaryClub) SetSportType(v string)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *SummaryClub) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetActivityTypes

`func (o *SummaryClub) GetActivityTypes() []ActivityType`

GetActivityTypes returns the ActivityTypes field if non-nil, zero value otherwise.

### GetActivityTypesOk

`func (o *SummaryClub) GetActivityTypesOk() (*[]ActivityType, bool)`

GetActivityTypesOk returns a tuple with the ActivityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypes

`func (o *SummaryClub) SetActivityTypes(v []ActivityType)`

SetActivityTypes sets ActivityTypes field to given value.

### HasActivityTypes

`func (o *SummaryClub) HasActivityTypes() bool`

HasActivityTypes returns a boolean if a field has been set.

### GetCity

`func (o *SummaryClub) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SummaryClub) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SummaryClub) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SummaryClub) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *SummaryClub) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SummaryClub) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SummaryClub) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SummaryClub) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *SummaryClub) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SummaryClub) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SummaryClub) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SummaryClub) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPrivate

`func (o *SummaryClub) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SummaryClub) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SummaryClub) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SummaryClub) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetMemberCount

`func (o *SummaryClub) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *SummaryClub) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *SummaryClub) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *SummaryClub) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetFeatured

`func (o *SummaryClub) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *SummaryClub) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *SummaryClub) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *SummaryClub) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVerified

`func (o *SummaryClub) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SummaryClub) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SummaryClub) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SummaryClub) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetUrl

`func (o *SummaryClub) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SummaryClub) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SummaryClub) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SummaryClub) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


