# Fault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error**](Error.md) | The set of specific errors associated with this fault, if any. | [optional] 
**Message** | Pointer to **string** | The message of the fault. | [optional] 

## Methods

### NewFault

`func NewFault() *Fault`

NewFault instantiates a new Fault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaultWithDefaults

`func NewFaultWithDefaults() *Fault`

NewFaultWithDefaults instantiates a new Fault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Fault) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Fault) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Fault) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Fault) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessage

`func (o *Fault) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Fault) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Fault) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Fault) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


