# AccountStatsStat

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

### NewAccountStatsStat

`func NewAccountStatsStat() *AccountStatsStat`

NewAccountStatsStat instantiates a new AccountStatsStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatsStatWithDefaults

`func NewAccountStatsStatWithDefaults() *AccountStatsStat`

NewAccountStatsStatWithDefaults instantiates a new AccountStatsStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *AccountStatsStat) GetProcessed() int32`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *AccountStatsStat) GetProcessedOk() (*int32, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *AccountStatsStat) SetProcessed(v int32)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *AccountStatsStat) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetSent

`func (o *AccountStatsStat) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *AccountStatsStat) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *AccountStatsStat) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *AccountStatsStat) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *AccountStatsStat) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *AccountStatsStat) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *AccountStatsStat) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *AccountStatsStat) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDropped

`func (o *AccountStatsStat) GetDropped() int32`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *AccountStatsStat) GetDroppedOk() (*int32, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *AccountStatsStat) SetDropped(v int32)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *AccountStatsStat) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSmtpDropped

`func (o *AccountStatsStat) GetSmtpDropped() int32`

GetSmtpDropped returns the SmtpDropped field if non-nil, zero value otherwise.

### GetSmtpDroppedOk

`func (o *AccountStatsStat) GetSmtpDroppedOk() (*int32, bool)`

GetSmtpDroppedOk returns a tuple with the SmtpDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDropped

`func (o *AccountStatsStat) SetSmtpDropped(v int32)`

SetSmtpDropped sets SmtpDropped field to given value.

### HasSmtpDropped

`func (o *AccountStatsStat) HasSmtpDropped() bool`

HasSmtpDropped returns a boolean if a field has been set.

### GetHardBounced

`func (o *AccountStatsStat) GetHardBounced() int32`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *AccountStatsStat) GetHardBouncedOk() (*int32, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *AccountStatsStat) SetHardBounced(v int32)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *AccountStatsStat) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetSoftBounced

`func (o *AccountStatsStat) GetSoftBounced() int32`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *AccountStatsStat) GetSoftBouncedOk() (*int32, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *AccountStatsStat) SetSoftBounced(v int32)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *AccountStatsStat) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetOpened

`func (o *AccountStatsStat) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *AccountStatsStat) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *AccountStatsStat) SetOpened(v int32)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *AccountStatsStat) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *AccountStatsStat) GetClicked() int32`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *AccountStatsStat) GetClickedOk() (*int32, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *AccountStatsStat) SetClicked(v int32)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *AccountStatsStat) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *AccountStatsStat) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *AccountStatsStat) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *AccountStatsStat) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *AccountStatsStat) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetSpams

`func (o *AccountStatsStat) GetSpams() int32`

GetSpams returns the Spams field if non-nil, zero value otherwise.

### GetSpamsOk

`func (o *AccountStatsStat) GetSpamsOk() (*int32, bool)`

GetSpamsOk returns a tuple with the Spams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpams

`func (o *AccountStatsStat) SetSpams(v int32)`

SetSpams sets Spams field to given value.

### HasSpams

`func (o *AccountStatsStat) HasSpams() bool`

HasSpams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


