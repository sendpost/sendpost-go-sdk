# IPPoolStat

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
**IpPoolId** | Pointer to **int64** | Unique identifier of the IP pool. | [optional] 
**IpPoolName** | Pointer to **string** | Human-readable name of the IP pool. | [optional] 
**Pending** | Pointer to **int64** | Number of emails queued on this pool but not yet processed. | [optional] 

## Methods

### NewIPPoolStat

`func NewIPPoolStat() *IPPoolStat`

NewIPPoolStat instantiates a new IPPoolStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPPoolStatWithDefaults

`func NewIPPoolStatWithDefaults() *IPPoolStat`

NewIPPoolStatWithDefaults instantiates a new IPPoolStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *IPPoolStat) GetProcessed() int64`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *IPPoolStat) GetProcessedOk() (*int64, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *IPPoolStat) SetProcessed(v int64)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *IPPoolStat) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *IPPoolStat) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *IPPoolStat) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *IPPoolStat) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *IPPoolStat) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDropped

`func (o *IPPoolStat) GetDropped() int64`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *IPPoolStat) GetDroppedOk() (*int64, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *IPPoolStat) SetDropped(v int64)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *IPPoolStat) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *IPPoolStat) GetSmtpDropped() int64`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *IPPoolStat) GetSmtpDroppedOk() (*int64, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *IPPoolStat) SetSmtpDropped(v int64)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *IPPoolStat) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetDelivered

`func (o *IPPoolStat) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *IPPoolStat) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *IPPoolStat) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *IPPoolStat) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetSoftBounced

`func (o *IPPoolStat) GetSoftBounced() int64`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *IPPoolStat) GetSoftBouncedOk() (*int64, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *IPPoolStat) SetSoftBounced(v int64)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *IPPoolStat) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetHardBounced

`func (o *IPPoolStat) GetHardBounced() int64`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *IPPoolStat) GetHardBouncedOk() (*int64, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *IPPoolStat) SetHardBounced(v int64)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *IPPoolStat) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetOpened

`func (o *IPPoolStat) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *IPPoolStat) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *IPPoolStat) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *IPPoolStat) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *IPPoolStat) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *IPPoolStat) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *IPPoolStat) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *IPPoolStat) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *IPPoolStat) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *IPPoolStat) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *IPPoolStat) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *IPPoolStat) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *IPPoolStat) GetSpam() int64`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *IPPoolStat) GetSpamOk() (*int64, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *IPPoolStat) SetSpam(v int64)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *IPPoolStat) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetIpPoolId

`func (o *IPPoolStat) GetIpPoolId() int64`

GetIpPoolId returns the IpPoolId field if non-nil, zero value otherwise.

### GetIpPoolIdOk

`func (o *IPPoolStat) GetIpPoolIdOk() (*int64, bool)`

GetIpPoolIdOk returns a tuple with the IpPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolId

`func (o *IPPoolStat) SetIpPoolId(v int64)`

SetIpPoolId sets IpPoolId field to given value.

### HasIpPoolId

`func (o *IPPoolStat) HasIpPoolId() bool`

HasIpPoolId returns a boolean if a field has been set.

### GetIpPoolName

`func (o *IPPoolStat) GetIpPoolName() string`

GetIpPoolName returns the IpPoolName field if non-nil, zero value otherwise.

### GetIpPoolNameOk

`func (o *IPPoolStat) GetIpPoolNameOk() (*string, bool)`

GetIpPoolNameOk returns a tuple with the IpPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolName

`func (o *IPPoolStat) SetIpPoolName(v string)`

SetIpPoolName sets IpPoolName field to given value.

### HasIpPoolName

`func (o *IPPoolStat) HasIpPoolName() bool`

HasIpPoolName returns a boolean if a field has been set.

### GetPending

`func (o *IPPoolStat) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *IPPoolStat) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *IPPoolStat) SetPending(v int64)`

SetPending sets Pending field to given value.

### HasPending

`func (o *IPPoolStat) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


