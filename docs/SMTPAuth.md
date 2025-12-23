# SMTPAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the SMTP Auth | [optional] 
**Username** | Pointer to **string** | Username for the SMTP Auth | [optional] 
**Password** | Pointer to **string** | Password for the SMTP Auth | [optional] 
**Created** | Pointer to **int64** | UNIX epoch nano timestamp when the SMTP Auth was created | [optional] 
**Updated** | Pointer to **int64** | UNIX epoch nano timestamp when the SMTP Auth was updated | [optional] 

## Methods

### NewSMTPAuth

`func NewSMTPAuth() *SMTPAuth`

NewSMTPAuth instantiates a new SMTPAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPAuthWithDefaults

`func NewSMTPAuthWithDefaults() *SMTPAuth`

NewSMTPAuthWithDefaults instantiates a new SMTPAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SMTPAuth) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SMTPAuth) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SMTPAuth) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SMTPAuth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *SMTPAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SMTPAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SMTPAuth) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SMTPAuth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *SMTPAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SMTPAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SMTPAuth) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SMTPAuth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCreated

`func (o *SMTPAuth) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SMTPAuth) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SMTPAuth) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SMTPAuth) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *SMTPAuth) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SMTPAuth) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SMTPAuth) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SMTPAuth) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


