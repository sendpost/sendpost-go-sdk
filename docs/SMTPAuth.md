# SMTPAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the SMTP credentials | [optional] 
**Username** | Pointer to **string** | SMTP username for authentication. Format: {identifier}@{subaccount_id}.sendpost.io  | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when credentials were created | [optional] 

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

`func (o *SMTPAuth) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SMTPAuth) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SMTPAuth) SetId(v int64)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


