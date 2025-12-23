# AggregateStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processed** | Pointer to **int32** |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Delivered** | Pointer to **int32** |  | [optional] 
**Dropped** | Pointer to **int32** |  | [optional] 
**SmtpDropped** | Pointer to **int32** |  | [optional] 
**HardBounced** | Pointer to **int32** |  | [optional] 
**SoftBounced** | Pointer to **int32** |  | [optional] 
**Opened** | Pointer to **int32** |  | [optional] 
**Clicked** | Pointer to **int32** |  | [optional] 
**Unsubscribed** | Pointer to **int32** |  | [optional] 
**Spams** | Pointer to **int32** |  | [optional] 

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

`func (o *AggregateStats) GetProcessed() int32`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *AggregateStats) GetProcessedOk() (*int32, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *AggregateStats) SetProcessed(v int32)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *AggregateStats) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *AggregateStats) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *AggregateStats) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *AggregateStats) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *AggregateStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *AggregateStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *AggregateStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *AggregateStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *AggregateStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDropped

`func (o *AggregateStats) GetDropped() int32`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *AggregateStats) GetDroppedOk() (*int32, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *AggregateStats) SetDropped(v int32)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *AggregateStats) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *AggregateStats) GetSmtpDropped() int32`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *AggregateStats) GetSmtpDroppedOk() (*int32, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *AggregateStats) SetSmtpDropped(v int32)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *AggregateStats) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetHardBounced

`func (o *AggregateStats) GetHardBounced() int32`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *AggregateStats) GetHardBouncedOk() (*int32, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *AggregateStats) SetHardBounced(v int32)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *AggregateStats) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetSoftBounced

`func (o *AggregateStats) GetSoftBounced() int32`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *AggregateStats) GetSoftBouncedOk() (*int32, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *AggregateStats) SetSoftBounced(v int32)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *AggregateStats) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetOpened

`func (o *AggregateStats) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *AggregateStats) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *AggregateStats) SetOpened(v int32)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *AggregateStats) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *AggregateStats) GetClicked() int32`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *AggregateStats) GetClickedOk() (*int32, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *AggregateStats) SetClicked(v int32)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *AggregateStats) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *AggregateStats) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *AggregateStats) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *AggregateStats) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *AggregateStats) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpams

`func (o *AggregateStats) GetSpams() int32`

GetSpams returns the Spams field if non-nil, zero value otherwise.

### GetSpamsOk

`func (o *AggregateStats) GetSpamsOk() (*int32, bool)`

GetSpamsOk returns a tuple with the Spams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpams

`func (o *AggregateStats) SetSpams(v int32)`

SetSpams sets Spams field to given value.

### HasSpams

`func (o *AggregateStats) HasSpams() bool`

HasSpams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


