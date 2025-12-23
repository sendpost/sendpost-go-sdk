# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is the webhook active or in a paused state? | [optional] 
**Url** | Pointer to **string** | URL endpoint to which webhook calls are sent. | [optional] 
**Processed** | Pointer to **bool** | Trigger webhook on email message being processed. | [optional] 
**Delivered** | Pointer to **bool** | Trigger webhook on email message being delivered. | [optional] 
**Dropped** | Pointer to **bool** | Trigger webhook on email message being dropped. | [optional] 
**SoftBounced** | Pointer to **bool** | Trigger webhook on email message being soft bounced. | [optional] 
**HardBounced** | Pointer to **bool** | Trigger webhook on email message being hard bounced. | [optional] 
**Opened** | Pointer to **bool** | Trigger webhook on email message being opened. | [optional] 
**Clicked** | Pointer to **bool** | Trigger webhook on email message link being clicked. | [optional] 
**Unsubscribed** | Pointer to **bool** | Trigger webhook on email message being unsubscribed. | [optional] 
**Spam** | Pointer to **bool** | Trigger webhook on email message being marked as spam. | [optional] 
**Sent** | Pointer to **bool** | Trigger webhook on email message being sent. | [optional] 
**SmtpDropped** | Pointer to **bool** | Trigger webhook on email message being dropped by SMTP. | [optional] 
**UniqueOpen** | Pointer to **bool** | Trigger webhook on unique email opens. | [optional] 
**UniqueClick** | Pointer to **bool** | Trigger webhook on unique email clicks. | [optional] 

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest() *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *CreateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateWebhookRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetProcessed

`func (o *CreateWebhookRequest) GetProcessed() bool`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *CreateWebhookRequest) GetProcessedOk() (*bool, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *CreateWebhookRequest) SetProcessed(v bool)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *CreateWebhookRequest) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetDelivered

`func (o *CreateWebhookRequest) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *CreateWebhookRequest) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *CreateWebhookRequest) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *CreateWebhookRequest) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDropped

`func (o *CreateWebhookRequest) GetDropped() bool`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *CreateWebhookRequest) GetDroppedOk() (*bool, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *CreateWebhookRequest) SetDropped(v bool)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *CreateWebhookRequest) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSoftBounced

`func (o *CreateWebhookRequest) GetSoftBounced() bool`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *CreateWebhookRequest) GetSoftBouncedOk() (*bool, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *CreateWebhookRequest) SetSoftBounced(v bool)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *CreateWebhookRequest) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetHardBounced

`func (o *CreateWebhookRequest) GetHardBounced() bool`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *CreateWebhookRequest) GetHardBouncedOk() (*bool, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *CreateWebhookRequest) SetHardBounced(v bool)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *CreateWebhookRequest) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetOpened

`func (o *CreateWebhookRequest) GetOpened() bool`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *CreateWebhookRequest) GetOpenedOk() (*bool, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *CreateWebhookRequest) SetOpened(v bool)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *CreateWebhookRequest) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *CreateWebhookRequest) GetClicked() bool`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *CreateWebhookRequest) GetClickedOk() (*bool, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *CreateWebhookRequest) SetClicked(v bool)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *CreateWebhookRequest) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *CreateWebhookRequest) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *CreateWebhookRequest) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *CreateWebhookRequest) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *CreateWebhookRequest) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *CreateWebhookRequest) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *CreateWebhookRequest) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *CreateWebhookRequest) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *CreateWebhookRequest) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetSent

`func (o *CreateWebhookRequest) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *CreateWebhookRequest) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *CreateWebhookRequest) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *CreateWebhookRequest) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *CreateWebhookRequest) GetSmtpDropped() bool`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *CreateWebhookRequest) GetSmtpDroppedOk() (*bool, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *CreateWebhookRequest) SetSmtpDropped(v bool)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *CreateWebhookRequest) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetUniqueOpen

`func (o *CreateWebhookRequest) GetUniqueOpen() bool`

GetUniqueOpen returns the UniqueOpen field if non-nil, zero value otherwise.

### GetUniqueOpenOk

`func (o *CreateWebhookRequest) GetUniqueOpenOk() (*bool, bool)`

GetUniqueOpenOk returns a tuple with the UniqueOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueOpen

`func (o *CreateWebhookRequest) SetUniqueOpen(v bool)`

SetUniqueOpen sets UniqueOpen field to given value.

### HasUniqueOpen

`func (o *CreateWebhookRequest) HasUniqueOpen() bool`

HasUniqueOpen returns a boolean if a field has been set.

### GetUniqueClick

`func (o *CreateWebhookRequest) GetUniqueClick() bool`

GetUniqueClick returns the UniqueClick field if non-nil, zero value otherwise.

### GetUniqueClickOk

`func (o *CreateWebhookRequest) GetUniqueClickOk() (*bool, bool)`

GetUniqueClickOk returns a tuple with the UniqueClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClick

`func (o *CreateWebhookRequest) SetUniqueClick(v bool)`

SetUniqueClick sets UniqueClick field to given value.

### HasUniqueClick

`func (o *CreateWebhookRequest) HasUniqueClick() bool`

HasUniqueClick returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


