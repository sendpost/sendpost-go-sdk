# AggregateStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processed** | Pointer to **int64** | Total number of emails accepted by SendPost API for processing. This is the starting point - all emails submitted through the API.  | [optional] 
**Sent** | Pointer to **int64** | Number of emails sent to recipient mail servers. sent &#x3D; processed - dropped - smtpDropped  | [optional] 
**Dropped** | Pointer to **int64** | Number of emails dropped before sending. Common reasons: - Recipient email in suppression list (hard bounce, spam complaint, unsubscribe) - Invalid recipient email format - Sender domain not verified  | [optional] 
**SmtpDropped** | Pointer to **int64** | Number of emails dropped at SMTP level due to policy violations or rate limiting by the receiving server before delivery attempt completed.  | [optional] 
**Delivered** | Pointer to **int64** | Number of emails successfully delivered to recipient mail servers. Note: Delivered means accepted by the server, not necessarily in inbox.  | [optional] 
**SoftBounced** | Pointer to **int64** | Number of temporary delivery failures (soft bounces). Common causes: - Recipient mailbox full - Server temporarily unavailable - Message too large SendPost automatically retries soft bounces.  | [optional] 
**HardBounced** | Pointer to **int64** | Number of permanent delivery failures (hard bounces). Common causes: - Recipient email doesn&#39;t exist - Domain doesn&#39;t exist - Recipient has blocked sender Hard bounced addresses are automatically added to suppression list.  | [optional] 
**Opened** | Pointer to **int64** | Number of emails opened (tracking pixel loaded). Requires trackOpens&#x3D;true. Note: Some email clients block tracking pixels.  | [optional] 
**Clicked** | Pointer to **int64** | Number of emails with at least one link clicked. Requires trackClicks&#x3D;true.  | [optional] 
**Unsubscribed** | Pointer to **int64** | Number of recipients who clicked the unsubscribe link. Unsubscribed addresses are automatically added to suppression list.  | [optional] 
**Spam** | Pointer to **int64** | Number of spam complaints (recipient marked email as spam). High spam rates can severely impact your sender reputation. Target: Keep spam rate below 0.1%.  | [optional] 

## Methods

### NewAggregateStats

`func NewAggregateStats() *AggregateStats`

NewAggregateStats instantiates a new AggregateStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateStatsWithDefaults

`func NewAggregateStatsWithDefaults() *AggregateStats`

NewAggregateStatsWithDefaults instantiates a new AggregateStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *AggregateStats) GetProcessed() int64`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *AggregateStats) GetProcessedOk() (*int64, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *AggregateStats) SetProcessed(v int64)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *AggregateStats) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *AggregateStats) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *AggregateStats) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *AggregateStats) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *AggregateStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDropped

`func (o *AggregateStats) GetDropped() int64`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *AggregateStats) GetDroppedOk() (*int64, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *AggregateStats) SetDropped(v int64)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *AggregateStats) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *AggregateStats) GetSmtpDropped() int64`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *AggregateStats) GetSmtpDroppedOk() (*int64, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *AggregateStats) SetSmtpDropped(v int64)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *AggregateStats) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetDelivered

`func (o *AggregateStats) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *AggregateStats) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *AggregateStats) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *AggregateStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetSoftBounced

`func (o *AggregateStats) GetSoftBounced() int64`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *AggregateStats) GetSoftBouncedOk() (*int64, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *AggregateStats) SetSoftBounced(v int64)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *AggregateStats) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetHardBounced

`func (o *AggregateStats) GetHardBounced() int64`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *AggregateStats) GetHardBouncedOk() (*int64, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *AggregateStats) SetHardBounced(v int64)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *AggregateStats) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetOpened

`func (o *AggregateStats) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *AggregateStats) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *AggregateStats) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *AggregateStats) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *AggregateStats) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *AggregateStats) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *AggregateStats) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *AggregateStats) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *AggregateStats) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *AggregateStats) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *AggregateStats) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *AggregateStats) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *AggregateStats) GetSpam() int64`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *AggregateStats) GetSpamOk() (*int64, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *AggregateStats) SetSpam(v int64)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *AggregateStats) HasSpam() bool`

HasSpam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


