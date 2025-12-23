# AggregatedEmailStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processed** | Pointer to **int32** | Total number of emails accepted by SendPost API | [optional] 
**Sent** | Pointer to **int32** | Total number of emails sent | [optional] 
**Delivered** | Pointer to **int32** | Total number of emails successfully delivered to SMTP | [optional] 
**Dropped** | Pointer to **int32** | Total number of emails dropped without delivery | [optional] 
**SmtpDropped** | Pointer to **int32** | Total number of emails dropped by SMTP | [optional] 
**HardBounced** | Pointer to **int32** | Total number of hard bounce errors | [optional] 
**SoftBounced** | Pointer to **int32** | Total number of soft bounce errors | [optional] 
**Opened** | Pointer to **int32** | Total number of emails opened by recipients | [optional] 
**Clicked** | Pointer to **int32** | Total number of links clicked by recipients | [optional] 
**Unsubscribed** | Pointer to **int32** | Total number of unsubscribed recipients | [optional] 
**Spam** | Pointer to **int32** | Total number of spams reported by recipients | [optional] 

## Methods

### NewAggregatedEmailStats

`func NewAggregatedEmailStats() *AggregatedEmailStats`

NewAggregatedEmailStats instantiates a new AggregatedEmailStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedEmailStatsWithDefaults

`func NewAggregatedEmailStatsWithDefaults() *AggregatedEmailStats`

NewAggregatedEmailStatsWithDefaults instantiates a new AggregatedEmailStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *AggregatedEmailStats) GetProcessed() int32`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *AggregatedEmailStats) GetProcessedOk() (*int32, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *AggregatedEmailStats) SetProcessed(v int32)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *AggregatedEmailStats) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *AggregatedEmailStats) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *AggregatedEmailStats) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *AggregatedEmailStats) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *AggregatedEmailStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *AggregatedEmailStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *AggregatedEmailStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *AggregatedEmailStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *AggregatedEmailStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDropped

`func (o *AggregatedEmailStats) GetDropped() int32`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *AggregatedEmailStats) GetDroppedOk() (*int32, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *AggregatedEmailStats) SetDropped(v int32)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *AggregatedEmailStats) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *AggregatedEmailStats) GetSmtpDropped() int32`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *AggregatedEmailStats) GetSmtpDroppedOk() (*int32, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *AggregatedEmailStats) SetSmtpDropped(v int32)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *AggregatedEmailStats) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetHardBounced

`func (o *AggregatedEmailStats) GetHardBounced() int32`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *AggregatedEmailStats) GetHardBouncedOk() (*int32, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *AggregatedEmailStats) SetHardBounced(v int32)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *AggregatedEmailStats) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetSoftBounced

`func (o *AggregatedEmailStats) GetSoftBounced() int32`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *AggregatedEmailStats) GetSoftBouncedOk() (*int32, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *AggregatedEmailStats) SetSoftBounced(v int32)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *AggregatedEmailStats) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetOpened

`func (o *AggregatedEmailStats) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *AggregatedEmailStats) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *AggregatedEmailStats) SetOpened(v int32)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *AggregatedEmailStats) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *AggregatedEmailStats) GetClicked() int32`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *AggregatedEmailStats) GetClickedOk() (*int32, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *AggregatedEmailStats) SetClicked(v int32)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *AggregatedEmailStats) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *AggregatedEmailStats) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *AggregatedEmailStats) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *AggregatedEmailStats) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *AggregatedEmailStats) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *AggregatedEmailStats) GetSpam() int32`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *AggregatedEmailStats) GetSpamOk() (*int32, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *AggregatedEmailStats) SetSpam(v int32)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *AggregatedEmailStats) HasSpam() bool`

HasSpam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


