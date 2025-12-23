# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Cc** | Pointer to [**[]CopyTo**](CopyTo.md) |  | [optional] 
**Bcc** | Pointer to [**[]CopyTo**](CopyTo.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom fields for personalization | [optional] 

## Methods

### NewRecipient

`func NewRecipient() *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Recipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Recipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Recipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Recipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Recipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Recipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Recipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Recipient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCc

`func (o *Recipient) GetCc() []CopyTo`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *Recipient) GetCcOk() (*[]CopyTo, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *Recipient) SetCc(v []CopyTo)`

SetCc sets Cc field to given value.

### HasCc

`func (o *Recipient) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *Recipient) GetBcc() []CopyTo`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *Recipient) GetBccOk() (*[]CopyTo, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *Recipient) SetBcc(v []CopyTo)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *Recipient) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCustomFields

`func (o *Recipient) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Recipient) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Recipient) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Recipient) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


