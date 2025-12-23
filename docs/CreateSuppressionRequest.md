# CreateSuppressionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardBounce** | Pointer to [**[]CreateSuppressionRequestHardBounceInner**](CreateSuppressionRequestHardBounceInner.md) | list of email addresses which you want to mark in hardBounce suppression list | [optional] 
**Manual** | Pointer to [**[]CreateSuppressionRequestManualInner**](CreateSuppressionRequestManualInner.md) | list of email addresses which you want to mark in manual suppression list | [optional] 
**Unsubscribe** | Pointer to [**[]CreateSuppressionRequestUnsubscribeInner**](CreateSuppressionRequestUnsubscribeInner.md) | list of email addresses which you want to mark in unsubscribe suppression list | [optional] 
**SpamComplaint** | Pointer to [**[]CreateSuppressionRequestSpamComplaintInner**](CreateSuppressionRequestSpamComplaintInner.md) | list of email addresses which you want to mark in spamComplaint suppression list | [optional] 

## Methods

### NewCreateSuppressionRequest

`func NewCreateSuppressionRequest() *CreateSuppressionRequest`

NewCreateSuppressionRequest instantiates a new CreateSuppressionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSuppressionRequestWithDefaults

`func NewCreateSuppressionRequestWithDefaults() *CreateSuppressionRequest`

NewCreateSuppressionRequestWithDefaults instantiates a new CreateSuppressionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardBounce

`func (o *CreateSuppressionRequest) GetHardBounce() []CreateSuppressionRequestHardBounceInner`

GetHardBounce returns the HardBounce field if non-nil, zero value otherwise.

### GetHardBounceOk

`func (o *CreateSuppressionRequest) GetHardBounceOk() (*[]CreateSuppressionRequestHardBounceInner, bool)`

GetHardBounceOk returns a tuple with the HardBounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounce

`func (o *CreateSuppressionRequest) SetHardBounce(v []CreateSuppressionRequestHardBounceInner)`

SetHardBounce sets HardBounce field to given value.

### HasHardBounce

`func (o *CreateSuppressionRequest) HasHardBounce() bool`

HasHardBounce returns a boolean if a field has been set.

### GetManual

`func (o *CreateSuppressionRequest) GetManual() []CreateSuppressionRequestManualInner`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *CreateSuppressionRequest) GetManualOk() (*[]CreateSuppressionRequestManualInner, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *CreateSuppressionRequest) SetManual(v []CreateSuppressionRequestManualInner)`

SetManual sets Manual field to given value.

### HasManual

`func (o *CreateSuppressionRequest) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *CreateSuppressionRequest) GetUnsubscribe() []CreateSuppressionRequestUnsubscribeInner`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *CreateSuppressionRequest) GetUnsubscribeOk() (*[]CreateSuppressionRequestUnsubscribeInner, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *CreateSuppressionRequest) SetUnsubscribe(v []CreateSuppressionRequestUnsubscribeInner)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *CreateSuppressionRequest) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetSpamComplaint

`func (o *CreateSuppressionRequest) GetSpamComplaint() []CreateSuppressionRequestSpamComplaintInner`

GetSpamComplaint returns the SpamComplaint field if non-nil, zero value otherwise.

### GetSpamComplaintOk

`func (o *CreateSuppressionRequest) GetSpamComplaintOk() (*[]CreateSuppressionRequestSpamComplaintInner, bool)`

GetSpamComplaintOk returns a tuple with the SpamComplaint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamComplaint

`func (o *CreateSuppressionRequest) SetSpamComplaint(v []CreateSuppressionRequestSpamComplaintInner)`

SetSpamComplaint sets SpamComplaint field to given value.

### HasSpamComplaint

`func (o *CreateSuppressionRequest) HasSpamComplaint() bool`

HasSpamComplaint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


