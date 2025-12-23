# MessageHeaderTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the recipient. | [optional] 
**Email** | Pointer to **string** | Email address of the recipient. | [optional] 
**Cc** | Pointer to **[]string** | List of CC recipients | [optional] 
**Bcc** | Pointer to **[]string** | List of BCC recipients | [optional] 
**CustomFields** | Pointer to **map[string]string** | Key-Value pair of custom fields. | [optional] 

## Methods

### NewMessageHeaderTo

`func NewMessageHeaderTo() *MessageHeaderTo`

NewMessageHeaderTo instantiates a new MessageHeaderTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageHeaderToWithDefaults

`func NewMessageHeaderToWithDefaults() *MessageHeaderTo`

NewMessageHeaderToWithDefaults instantiates a new MessageHeaderTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MessageHeaderTo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageHeaderTo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageHeaderTo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MessageHeaderTo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *MessageHeaderTo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MessageHeaderTo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MessageHeaderTo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MessageHeaderTo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCc

`func (o *MessageHeaderTo) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *MessageHeaderTo) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *MessageHeaderTo) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *MessageHeaderTo) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *MessageHeaderTo) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *MessageHeaderTo) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *MessageHeaderTo) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *MessageHeaderTo) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCustomFields

`func (o *MessageHeaderTo) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MessageHeaderTo) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MessageHeaderTo) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MessageHeaderTo) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


