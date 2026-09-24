# DailyStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processed** | Pointer to **int64** | Total emails accepted by the API | [optional] 
**Sent** | Pointer to **int64** | Total emails sent to recipient mail servers | [optional] 
**Dropped** | Pointer to **int64** | Total emails dropped before sending | [optional] 
**SmtpDropped** | Pointer to **int64** | Total emails dropped at the SMTP level before delivery | [optional] 
**Delivered** | Pointer to **int64** | Total emails delivered successfully | [optional] 
**SoftBounced** | Pointer to **int64** | Total temporary delivery failures | [optional] 
**HardBounced** | Pointer to **int64** | Total permanent delivery failures | [optional] 
**Opened** | Pointer to **int64** | Total email opens | [optional] 
**Clicked** | Pointer to **int64** | Total link clicks | [optional] 
**Unsubscribed** | Pointer to **int64** | Total unsubscribes | [optional] 
**Spam** | Pointer to **int64** | Total spam complaints | [optional] 

## Methods

### NewDailyStatistics

`func NewDailyStatistics() *DailyStatistics`

NewDailyStatistics instantiates a new DailyStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyStatisticsWithDefaults

`func NewDailyStatisticsWithDefaults() *DailyStatistics`

NewDailyStatisticsWithDefaults instantiates a new DailyStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *DailyStatistics) GetProcessed() int64`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *DailyStatistics) GetProcessedOk() (*int64, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *DailyStatistics) SetProcessed(v int64)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *DailyStatistics) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *DailyStatistics) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *DailyStatistics) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *DailyStatistics) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *DailyStatistics) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDropped

`func (o *DailyStatistics) GetDropped() int64`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *DailyStatistics) GetDroppedOk() (*int64, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *DailyStatistics) SetDropped(v int64)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *DailyStatistics) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *DailyStatistics) GetSmtpDropped() int64`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *DailyStatistics) GetSmtpDroppedOk() (*int64, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *DailyStatistics) SetSmtpDropped(v int64)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *DailyStatistics) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetDelivered

`func (o *DailyStatistics) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *DailyStatistics) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *DailyStatistics) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *DailyStatistics) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetSoftBounced

`func (o *DailyStatistics) GetSoftBounced() int64`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *DailyStatistics) GetSoftBouncedOk() (*int64, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *DailyStatistics) SetSoftBounced(v int64)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *DailyStatistics) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetHardBounced

`func (o *DailyStatistics) GetHardBounced() int64`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *DailyStatistics) GetHardBouncedOk() (*int64, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *DailyStatistics) SetHardBounced(v int64)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *DailyStatistics) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetOpened

`func (o *DailyStatistics) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *DailyStatistics) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *DailyStatistics) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *DailyStatistics) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *DailyStatistics) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *DailyStatistics) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *DailyStatistics) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *DailyStatistics) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *DailyStatistics) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *DailyStatistics) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *DailyStatistics) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *DailyStatistics) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpam

`func (o *DailyStatistics) GetSpam() int64`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *DailyStatistics) GetSpamOk() (*int64, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *DailyStatistics) SetSpam(v int64)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *DailyStatistics) HasSpam() bool`

HasSpam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


