# IPUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoWarmupEnabled** | Pointer to **bool** | Toggle automatic IP warmup on or off. - &#x60;true&#x60;: SendPost automatically manages daily sending limits - &#x60;false&#x60;: You manage sending volume manually (advanced users)  Warning: Disabling warmup and sending high volume on a new IP can damage sender reputation.  | [optional] 

## Methods

### NewIPUpdateRequest

`func NewIPUpdateRequest() *IPUpdateRequest`

NewIPUpdateRequest instantiates a new IPUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPUpdateRequestWithDefaults

`func NewIPUpdateRequestWithDefaults() *IPUpdateRequest`

NewIPUpdateRequestWithDefaults instantiates a new IPUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoWarmupEnabled

`func (o *IPUpdateRequest) GetAutoWarmupEnabled() bool`

GetAutoWarmupEnabled returns the AutoWarmupEnabled field if non-nil, zero value otherwise.

### GetAutoWarmupEnabledOk

`func (o *IPUpdateRequest) GetAutoWarmupEnabledOk() (*bool, bool)`

GetAutoWarmupEnabledOk returns a tuple with the AutoWarmupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWarmupEnabled

`func (o *IPUpdateRequest) SetAutoWarmupEnabled(v bool)`

SetAutoWarmupEnabled sets AutoWarmupEnabled field to given value.

### HasAutoWarmupEnabled

`func (o *IPUpdateRequest) HasAutoWarmupEnabled() bool`

HasAutoWarmupEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


