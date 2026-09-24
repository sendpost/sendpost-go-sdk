# NewWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the webhook is active immediately after creation. Set to false to configure and test before activating.  | [optional] [default to true]
**Url** | **string** | HTTPS URL endpoint to receive webhook POST requests. Must: - Use HTTPS (HTTP not allowed for security) - Be publicly accessible - Return 2xx status within 10 seconds - Handle duplicate deliveries (use eventId for idempotency)  | 
**Processed** | Pointer to **bool** | Fire when email is accepted by SendPost API | [optional] [default to false]
**Sent** | Pointer to **bool** | Fire when email is sent to recipient&#39;s mail server | [optional] [default to false]
**Delivered** | Pointer to **bool** | Fire when email is accepted by recipient&#39;s mail server | [optional] [default to true]
**Dropped** | Pointer to **bool** | Fire when email is not sent (suppression, invalid, etc.) | [optional] [default to true]
**SmtpDropped** | Pointer to **bool** | Fire when email is rejected at SMTP level | [optional] [default to false]
**SoftBounced** | Pointer to **bool** | Fire on temporary delivery failure (will retry) | [optional] [default to true]
**HardBounced** | Pointer to **bool** | Fire on permanent delivery failure | [optional] [default to true]
**Opened** | Pointer to **bool** | Fire when email is opened. Fires on EVERY open. Consider using &#x60;uniqueOpen&#x60; instead to reduce volume.  | [optional] [default to true]
**Clicked** | Pointer to **bool** | Fire when a link is clicked. Fires on EVERY click. Consider using &#x60;uniqueClick&#x60; instead to reduce volume.  | [optional] [default to true]
**Unsubscribed** | Pointer to **bool** | Fire when recipient clicks unsubscribe link | [optional] [default to true]
**Spam** | Pointer to **bool** | Fire when recipient marks email as spam | [optional] [default to true]
**UniqueOpen** | Pointer to **bool** | Fire only on FIRST open of each email (unique opens). More efficient than &#x60;opened&#x60; if you only need engagement metrics.  | [optional] [default to false]
**UniqueClick** | Pointer to **bool** | Fire only on FIRST click of each email (unique clicks). More efficient than &#x60;clicked&#x60; if you only need engagement metrics.  | [optional] [default to false]

## Methods

### NewNewWebhookRequest

`func NewNewWebhookRequest(url string, ) *NewWebhookRequest`

NewNewWebhookRequest instantiates a new NewWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewWebhookRequestWithDefaults

`func NewNewWebhookRequestWithDefaults() *NewWebhookRequest`

NewNewWebhookRequestWithDefaults instantiates a new NewWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NewWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NewWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *NewWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NewWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProcessed

`func (o *NewWebhookRequest) GetProcessed() bool`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *NewWebhookRequest) GetProcessedOk() (*bool, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *NewWebhookRequest) SetProcessed(v bool)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *NewWebhookRequest) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *NewWebhookRequest) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *NewWebhookRequest) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *NewWebhookRequest) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *NewWebhookRequest) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *NewWebhookRequest) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *NewWebhookRequest) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *NewWebhookRequest) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *NewWebhookRequest) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDropped

`func (o *NewWebhookRequest) GetDropped() bool`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *NewWebhookRequest) GetDroppedOk() (*bool, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *NewWebhookRequest) SetDropped(v bool)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *NewWebhookRequest) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *NewWebhookRequest) GetSmtpDropped() bool`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *NewWebhookRequest) GetSmtpDroppedOk() (*bool, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *NewWebhookRequest) SetSmtpDropped(v bool)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *NewWebhookRequest) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetSoftBounced

`func (o *NewWebhookRequest) GetSoftBounced() bool`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *NewWebhookRequest) GetSoftBouncedOk() (*bool, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *NewWebhookRequest) SetSoftBounced(v bool)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *NewWebhookRequest) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetHardBounced

`func (o *NewWebhookRequest) GetHardBounced() bool`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *NewWebhookRequest) GetHardBouncedOk() (*bool, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *NewWebhookRequest) SetHardBounced(v bool)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *NewWebhookRequest) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetOpened

`func (o *NewWebhookRequest) GetOpened() bool`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *NewWebhookRequest) GetOpenedOk() (*bool, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *NewWebhookRequest) SetOpened(v bool)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *NewWebhookRequest) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *NewWebhookRequest) GetClicked() bool`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *NewWebhookRequest) GetClickedOk() (*bool, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *NewWebhookRequest) SetClicked(v bool)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *NewWebhookRequest) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *NewWebhookRequest) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *NewWebhookRequest) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *NewWebhookRequest) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *NewWebhookRequest) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *NewWebhookRequest) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *NewWebhookRequest) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *NewWebhookRequest) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *NewWebhookRequest) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetUniqueOpen

`func (o *NewWebhookRequest) GetUniqueOpen() bool`

GetUniqueOpen returns the UniqueOpen field if non-nil, zero value otherwise.

### GetUniqueOpenOk

`func (o *NewWebhookRequest) GetUniqueOpenOk() (*bool, bool)`

GetUniqueOpenOk returns a tuple with the UniqueOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueOpen

`func (o *NewWebhookRequest) SetUniqueOpen(v bool)`

SetUniqueOpen sets UniqueOpen field to given value.

### HasUniqueOpen

`func (o *NewWebhookRequest) HasUniqueOpen() bool`

HasUniqueOpen returns a boolean if a field has been set.

### GetUniqueClick

`func (o *NewWebhookRequest) GetUniqueClick() bool`

GetUniqueClick returns the UniqueClick field if non-nil, zero value otherwise.

### GetUniqueClickOk

`func (o *NewWebhookRequest) GetUniqueClickOk() (*bool, bool)`

GetUniqueClickOk returns a tuple with the UniqueClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueClick

`func (o *NewWebhookRequest) SetUniqueClick(v bool)`

SetUniqueClick sets UniqueClick field to given value.

### HasUniqueClick

`func (o *NewWebhookRequest) HasUniqueClick() bool`

HasUniqueClick returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


