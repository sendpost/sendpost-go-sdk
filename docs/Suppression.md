# Suppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the suppression record | [optional] 
**Email** | Pointer to **string** | The suppressed email address | [optional] 
**Reason** | Pointer to **int32** | Reason code for the suppression: - &#x60;0&#x60; &#x3D; Manual (added via API or dashboard) - &#x60;1&#x60; &#x3D; Unsubscribe (recipient clicked unsubscribe link) - &#x60;2&#x60; &#x3D; Hard Bounce (permanent delivery failure) - &#x60;3&#x60; &#x3D; Spam Complaint (recipient marked as spam) - &#x60;4&#x60; &#x3D; Hard Bounce (detected by post-send validation)  | [optional] 
**ReasonText** | Pointer to **string** | Human-readable suppression reason | [optional] 
**SmtpError** | Pointer to **string** | SMTP error message from the receiving server (only for hard bounce suppressions). Useful for diagnosing delivery issues.  | [optional] 
**MessageUUID** | Pointer to **string** | UUID of the message whose bounce/complaint caused this suppression. Empty for manually added suppressions. Useful for tracing the origin.  | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the suppression was added | [optional] 

## Methods

### NewSuppression

`func NewSuppression() *Suppression`

NewSuppression instantiates a new Suppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionWithDefaults

`func NewSuppressionWithDefaults() *Suppression`

NewSuppressionWithDefaults instantiates a new Suppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Suppression) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Suppression) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Suppression) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Suppression) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *Suppression) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Suppression) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Suppression) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Suppression) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetReason

`func (o *Suppression) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Suppression) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Suppression) SetReason(v int32)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Suppression) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonText

`func (o *Suppression) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *Suppression) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *Suppression) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *Suppression) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetSmtpError

`func (o *Suppression) GetSmtpError() string`

GetSmtpError returns the SmtpError field if non-nil, zero value otherwise.

### GetSmtpErrorOk

`func (o *Suppression) GetSmtpErrorOk() (*string, bool)`

GetSmtpErrorOk returns a tuple with the SmtpError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpError

`func (o *Suppression) SetSmtpError(v string)`

SetSmtpError sets SmtpError field to given value.

### HasSmtpError

`func (o *Suppression) HasSmtpError() bool`

HasSmtpError returns a boolean if a field has been set.

### GetMessageUUID

`func (o *Suppression) GetMessageUUID() string`

GetMessageUUID returns the MessageUUID field if non-nil, zero value otherwise.

### GetMessageUUIDOk

`func (o *Suppression) GetMessageUUIDOk() (*string, bool)`

GetMessageUUIDOk returns a tuple with the MessageUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageUUID

`func (o *Suppression) SetMessageUUID(v string)`

SetMessageUUID sets MessageUUID field to given value.

### HasMessageUUID

`func (o *Suppression) HasMessageUUID() bool`

HasMessageUUID returns a boolean if a field has been set.

### GetCreated

`func (o *Suppression) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Suppression) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Suppression) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Suppression) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


