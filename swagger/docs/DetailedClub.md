# DetailedClub

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
**Membership** | Pointer to **string** | The membership status of the logged-in athlete. | [optional] 
**Admin** | Pointer to **bool** | Whether the currently logged-in athlete is an administrator of this club. | [optional] 
**Owner** | Pointer to **bool** | Whether the currently logged-in athlete is the owner of this club. | [optional] 
**FollowingCount** | Pointer to **int32** | The number of athletes in the club that the logged-in athlete follows. | [optional] 

## Methods

### NewDetailedClub

`func NewDetailedClub() *DetailedClub`

NewDetailedClub instantiates a new DetailedClub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedClubWithDefaults

`func NewDetailedClubWithDefaults() *DetailedClub`

NewDetailedClubWithDefaults instantiates a new DetailedClub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedClub) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedClub) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedClub) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedClub) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceState

`func (o *DetailedClub) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *DetailedClub) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *DetailedClub) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *DetailedClub) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetName

`func (o *DetailedClub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedClub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedClub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedClub) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileMedium

`func (o *DetailedClub) GetProfileMedium() string`

GetProfileMedium returns the ProfileMedium field if non-nil, zero value otherwise.

### GetProfileMediumOk

`func (o *DetailedClub) GetProfileMediumOk() (*string, bool)`

GetProfileMediumOk returns a tuple with the ProfileMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMedium

`func (o *DetailedClub) SetProfileMedium(v string)`

SetProfileMedium sets ProfileMedium field to given value.

### HasProfileMedium

`func (o *DetailedClub) HasProfileMedium() bool`

HasProfileMedium returns a boolean if a field has been set.

### GetCoverPhoto

`func (o *DetailedClub) GetCoverPhoto() string`

GetCoverPhoto returns the CoverPhoto field if non-nil, zero value otherwise.

### GetCoverPhotoOk

`func (o *DetailedClub) GetCoverPhotoOk() (*string, bool)`

GetCoverPhotoOk returns a tuple with the CoverPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverPhoto

`func (o *DetailedClub) SetCoverPhoto(v string)`

SetCoverPhoto sets CoverPhoto field to given value.

### HasCoverPhoto

`func (o *DetailedClub) HasCoverPhoto() bool`

HasCoverPhoto returns a boolean if a field has been set.

### GetCoverPhotoSmall

`func (o *DetailedClub) GetCoverPhotoSmall() string`

GetCoverPhotoSmall returns the CoverPhotoSmall field if non-nil, zero value otherwise.

### GetCoverPhotoSmallOk

`func (o *DetailedClub) GetCoverPhotoSmallOk() (*string, bool)`

GetCoverPhotoSmallOk returns a tuple with the CoverPhotoSmall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverPhotoSmall

`func (o *DetailedClub) SetCoverPhotoSmall(v string)`

SetCoverPhotoSmall sets CoverPhotoSmall field to given value.

### HasCoverPhotoSmall

`func (o *DetailedClub) HasCoverPhotoSmall() bool`

HasCoverPhotoSmall returns a boolean if a field has been set.

### GetSportType

`func (o *DetailedClub) GetSportType() string`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *DetailedClub) GetSportTypeOk() (*string, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *DetailedClub) SetSportType(v string)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *DetailedClub) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetActivityTypes

`func (o *DetailedClub) GetActivityTypes() []ActivityType`

GetActivityTypes returns the ActivityTypes field if non-nil, zero value otherwise.

### GetActivityTypesOk

`func (o *DetailedClub) GetActivityTypesOk() (*[]ActivityType, bool)`

GetActivityTypesOk returns a tuple with the ActivityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypes

`func (o *DetailedClub) SetActivityTypes(v []ActivityType)`

SetActivityTypes sets ActivityTypes field to given value.

### HasActivityTypes

`func (o *DetailedClub) HasActivityTypes() bool`

HasActivityTypes returns a boolean if a field has been set.

### GetCity

`func (o *DetailedClub) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DetailedClub) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DetailedClub) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DetailedClub) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *DetailedClub) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DetailedClub) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DetailedClub) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DetailedClub) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *DetailedClub) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DetailedClub) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DetailedClub) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DetailedClub) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPrivate

`func (o *DetailedClub) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *DetailedClub) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *DetailedClub) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *DetailedClub) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetMemberCount

`func (o *DetailedClub) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *DetailedClub) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *DetailedClub) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *DetailedClub) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetFeatured

`func (o *DetailedClub) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *DetailedClub) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *DetailedClub) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *DetailedClub) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVerified

`func (o *DetailedClub) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *DetailedClub) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *DetailedClub) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *DetailedClub) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetUrl

`func (o *DetailedClub) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DetailedClub) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DetailedClub) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DetailedClub) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMembership

`func (o *DetailedClub) GetMembership() string`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *DetailedClub) GetMembershipOk() (*string, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *DetailedClub) SetMembership(v string)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *DetailedClub) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetAdmin

`func (o *DetailedClub) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *DetailedClub) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *DetailedClub) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *DetailedClub) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetOwner

`func (o *DetailedClub) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DetailedClub) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DetailedClub) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DetailedClub) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFollowingCount

`func (o *DetailedClub) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *DetailedClub) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *DetailedClub) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *DetailedClub) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


