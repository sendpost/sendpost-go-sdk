# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the webhook. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the webhook is active or paused. | [optional] 
**Url** | Pointer to **string** | URL endpoint to which webhook calls need to be made. | [optional] 
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
**Created** | Pointer to **int64** | UNIX epoch nano timestamp when the webhook was created. | [optional] 
**CreatedBy** | Pointer to [**Member**](Member.md) | Member who created the webhook | [optional] 
**UpdatedBy** | Pointer to [**Member**](Member.md) | Member who updated the webhook | [optional] 

## Methods

### NewWebhook

`func NewWebhook() *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Webhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetProcessed

`func (o *Webhook) GetProcessed() bool`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *Webhook) GetProcessedOk() (*bool, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *Webhook) SetProcessed(v bool)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *Webhook) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetDelivered

`func (o *Webhook) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *Webhook) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *Webhook) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *Webhook) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDropped

`func (o *Webhook) GetDropped() bool`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *Webhook) GetDroppedOk() (*bool, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *Webhook) SetDropped(v bool)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *Webhook) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSoftBounced

`func (o *Webhook) GetSoftBounced() bool`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *Webhook) GetSoftBouncedOk() (*bool, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *Webhook) SetSoftBounced(v bool)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *Webhook) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetHardBounced

`func (o *Webhook) GetHardBounced() bool`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *Webhook) GetHardBouncedOk() (*bool, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *Webhook) SetHardBounced(v bool)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *Webhook) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetOpened

`func (o *Webhook) GetOpened() bool`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *Webhook) GetOpenedOk() (*bool, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *Webhook) SetOpened(v bool)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *Webhook) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *Webhook) GetClicked() bool`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *Webhook) GetClickedOk() (*bool, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *Webhook) SetClicked(v bool)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *Webhook) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *Webhook) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *Webhook) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *Webhook) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *Webhook) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *Webhook) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *Webhook) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *Webhook) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *Webhook) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetSent

`func (o *Webhook) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *Webhook) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *Webhook) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *Webhook) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *Webhook) GetSmtpDropped() bool`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *Webhook) GetSmtpDroppedOk() (*bool, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *Webhook) SetSmtpDropped(v bool)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *Webhook) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetUniqueOpen

`func (o *Webhook) GetUniqueOpen() bool`

GetUniqueOpen returns the UniqueOpen field if non-nil, zero value otherwise.

### GetUniqueOpenOk

`func (o *Webhook) GetUniqueOpenOk() (*bool, bool)`

GetUniqueOpenOk returns a tuple with the UniqueOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueOpen

`func (o *Webhook) SetUniqueOpen(v bool)`

SetUniqueOpen sets UniqueOpen field to given value.

### HasUniqueOpen

`func (o *Webhook) HasUniqueOpen() bool`

HasUniqueOpen returns a boolean if a field has been set.

### GetUniqueClick

`func (o *Webhook) GetUniqueClick() bool`

GetUniqueClick returns the UniqueClick field if non-nil, zero value otherwise.

### GetUniqueClickOk

`func (o *Webhook) GetUniqueClickOk() (*bool, bool)`

GetUniqueClickOk returns a tuple with the UniqueClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClick

`func (o *Webhook) SetUniqueClick(v bool)`

SetUniqueClick sets UniqueClick field to given value.

### HasUniqueClick

`func (o *Webhook) HasUniqueClick() bool`

HasUniqueClick returns a boolean if a field has been set.

### GetCreated

`func (o *Webhook) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Webhook) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Webhook) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Webhook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Webhook) GetCreatedBy() Member`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Webhook) GetCreatedByOk() (*Member, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Webhook) SetCreatedBy(v Member)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Webhook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Webhook) GetUpdatedBy() Member`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Webhook) GetUpdatedByOk() (*Member, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Webhook) SetUpdatedBy(v Member)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Webhook) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


