# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of this comment | [optional] 
**ActivityId** | Pointer to **int64** | The identifier of the activity this comment is related to | [optional] 
**Text** | Pointer to **string** | The content of the comment | [optional] 
**Athlete** | Pointer to [**SummaryAthlete**](SummaryAthlete.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which this comment was created. | [optional] 

## Methods

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *Comment) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *Comment) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *Comment) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *Comment) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetText

`func (o *Comment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Comment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Comment) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Comment) HasText() bool`

HasText returns a boolean if a field has been set.

### GetAthlete

`func (o *Comment) GetAthlete() SummaryAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *Comment) GetAthleteOk() (*SummaryAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *Comment) SetAthlete(v SummaryAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *Comment) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Comment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Comment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Comment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Comment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


