# EmailMessageToInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]string** |  | [optional] 
**Cc** | Pointer to [**[]EmailMessageToInnerCcInner**](EmailMessageToInnerCcInner.md) |  | [optional] 
**Bcc** | Pointer to [**[]EmailMessageToInnerBccInner**](EmailMessageToInnerBccInner.md) |  | [optional] 

## Methods

### NewEmailMessageToInner

`func NewEmailMessageToInner() *EmailMessageToInner`

NewEmailMessageToInner instantiates a new EmailMessageToInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageToInnerWithDefaults

`func NewEmailMessageToInnerWithDefaults() *EmailMessageToInner`

NewEmailMessageToInnerWithDefaults instantiates a new EmailMessageToInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailMessageToInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailMessageToInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailMessageToInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailMessageToInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *EmailMessageToInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailMessageToInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailMessageToInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailMessageToInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCustomFields

`func (o *EmailMessageToInner) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EmailMessageToInner) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EmailMessageToInner) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EmailMessageToInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCc

`func (o *EmailMessageToInner) GetCc() []EmailMessageToInnerCcInner`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *EmailMessageToInner) GetCcOk() (*[]EmailMessageToInnerCcInner, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *EmailMessageToInner) SetCc(v []EmailMessageToInnerCcInner)`

SetCc sets Cc field to given value.

### HasCc

`func (o *EmailMessageToInner) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *EmailMessageToInner) GetBcc() []EmailMessageToInnerBccInner`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *EmailMessageToInner) GetBccOk() (*[]EmailMessageToInnerBccInner, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *EmailMessageToInner) SetBcc(v []EmailMessageToInnerBccInner)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *EmailMessageToInner) HasBcc() bool`

HasBcc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


