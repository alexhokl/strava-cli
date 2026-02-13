# Upload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the upload | [optional] 
**IdStr** | Pointer to **string** | The unique identifier of the upload in string format | [optional] 
**ExternalId** | Pointer to **string** | The external identifier of the upload | [optional] 
**Error** | Pointer to **string** | The error associated with this upload | [optional] 
**Status** | Pointer to **string** | The status of this upload | [optional] 
**ActivityId** | Pointer to **int64** | The identifier of the activity this upload resulted into | [optional] 

## Methods

### NewUpload

`func NewUpload() *Upload`

NewUpload instantiates a new Upload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWithDefaults

`func NewUploadWithDefaults() *Upload`

NewUploadWithDefaults instantiates a new Upload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Upload) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Upload) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Upload) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Upload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdStr

`func (o *Upload) GetIdStr() string`

GetIdStr returns the IdStr field if non-nil, zero value otherwise.

### GetIdStrOk

`func (o *Upload) GetIdStrOk() (*string, bool)`

GetIdStrOk returns a tuple with the IdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStr

`func (o *Upload) SetIdStr(v string)`

SetIdStr sets IdStr field to given value.

### HasIdStr

`func (o *Upload) HasIdStr() bool`

HasIdStr returns a boolean if a field has been set.

### GetExternalId

`func (o *Upload) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Upload) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Upload) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Upload) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetError

`func (o *Upload) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Upload) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Upload) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Upload) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *Upload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Upload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Upload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Upload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActivityId

`func (o *Upload) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *Upload) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *Upload) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *Upload) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


