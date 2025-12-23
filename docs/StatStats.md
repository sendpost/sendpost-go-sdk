# StatStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processed** | Pointer to **int32** | Number of emails accepted by SendPost API. | [optional] 
**Sent** | Pointer to **int32** | Number of emails sent. | [optional] 
**Delivered** | Pointer to **int32** | Number of emails we were able to successfully deliver at SMTP without encountering any error | [optional] 
**Dropped** | Pointer to **int32** | Number of emails drop without attempting to deliver either because the email is invalid or email in in existing suppression list | [optional] 
**SmtpDropped** | Pointer to **int32** | Number of emails dropped by SMTP. | [optional] 
**HardBounced** | Pointer to **int32** | Number of emails where we got SMTP hard bounce error code by the recipient mail provider | [optional] 
**SoftBounced** | Pointer to **int32** | Number of emails where we got temporary soft bounce error by the recipent mail provider. Soft bounced emails are retried upto 5 times over 24 hour period before marking them as hardBounced. | [optional] 
**Opened** | Pointer to **int32** | Number of emails opened by recipients | [optional] 
**Clicked** | Pointer to **int32** | Number of email links clicked by recipients | [optional] 
**Unsubscribed** | Pointer to **int32** | Number of email recipients who unsubscribed from receiving further emails | [optional] 
**Spam** | Pointer to **int32** | Number of email recipients who marked emails as spam | [optional] 

## Methods

### NewStatStats

`func NewStatStats() *StatStats`

NewStatStats instantiates a new StatStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatStatsWithDefaults

`func NewStatStatsWithDefaults() *StatStats`

NewStatStatsWithDefaults instantiates a new StatStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *StatStats) GetProcessed() int32`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *StatStats) GetProcessedOk() (*int32, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *StatStats) SetProcessed(v int32)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *StatStats) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *StatStats) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *StatStats) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *StatStats) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *StatStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *StatStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *StatStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *StatStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *StatStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDropped

`func (o *StatStats) GetDropped() int32`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *StatStats) GetDroppedOk() (*int32, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *StatStats) SetDropped(v int32)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *StatStats) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *StatStats) GetSmtpDropped() int32`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *StatStats) GetSmtpDroppedOk() (*int32, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *StatStats) SetSmtpDropped(v int32)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *StatStats) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetHardBounced

`func (o *StatStats) GetHardBounced() int32`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *StatStats) GetHardBouncedOk() (*int32, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *StatStats) SetHardBounced(v int32)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *StatStats) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetSoftBounced

`func (o *StatStats) GetSoftBounced() int32`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *StatStats) GetSoftBouncedOk() (*int32, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *StatStats) SetSoftBounced(v int32)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *StatStats) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetOpened

`func (o *StatStats) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *StatStats) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *StatStats) SetOpened(v int32)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *StatStats) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *StatStats) GetClicked() int32`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *StatStats) GetClickedOk() (*int32, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *StatStats) SetClicked(v int32)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *StatStats) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *StatStats) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *StatStats) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *StatStats) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *StatStats) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *StatStats) GetSpam() int32`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *StatStats) GetSpamOk() (*int32, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *StatStats) SetSpam(v int32)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *StatStats) HasSpam() bool`

HasSpam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


