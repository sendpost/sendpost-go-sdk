# IP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the IP resource | [optional] 
**PublicIp** | Pointer to **string** | The public IPv4 address used for sending emails. This is the IP that receiving mail servers will see.  | [optional] 
**ReverseDnsHostname** | Pointer to **string** | The reverse DNS (PTR record) hostname for this IP. Properly configured rDNS is important for deliverability. Format: sp{id}.{region}.sendpost.email  | [optional] 
**Type** | Pointer to **int32** | Type of IP allocation: - &#x60;0&#x60; &#x3D; Shared IP (shared with other SendPost senders, pooled reputation) - &#x60;1&#x60; &#x3D; Dedicated IP (exclusive to your account, your own reputation)  | [optional] 
**AutoWarmupEnabled** | Pointer to **bool** | Whether automatic IP warmup is enabled. When enabled, SendPost automatically manages sending volume to gradually build reputation on this IP.  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) | Custom labels/tags for organizing IPs | [optional] 
**State** | Pointer to **int32** | Current state of the IP: - &#x60;0&#x60; &#x3D; Warmup (IP is in warmup phase, gradually building reputation) - &#x60;1&#x60; &#x3D; Normal (IP is fully warmed and ready for full sending volume)  | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the IP was allocated | [optional] 

## Methods

### NewIP

`func NewIP() *IP`

NewIP instantiates a new IP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPWithDefaults

`func NewIPWithDefaults() *IP`

NewIPWithDefaults instantiates a new IP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IP) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IP) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IP) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IP) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicIp

`func (o *IP) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *IP) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *IP) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *IP) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetReverseDnsHostname

`func (o *IP) GetReverseDnsHostname() string`

GetReverseDnsHostname returns the ReverseDnsHostname field if non-nil, zero value otherwise.

### GetReverseDnsHostnameOk

`func (o *IP) GetReverseDnsHostnameOk() (*string, bool)`

GetReverseDnsHostnameOk returns a tuple with the ReverseDnsHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDnsHostname

`func (o *IP) SetReverseDnsHostname(v string)`

SetReverseDnsHostname sets ReverseDnsHostname field to given value.

### HasReverseDnsHostname

`func (o *IP) HasReverseDnsHostname() bool`

HasReverseDnsHostname returns a boolean if a field has been set.

### GetType

`func (o *IP) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IP) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IP) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *IP) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoWarmupEnabled

`func (o *IP) GetAutoWarmupEnabled() bool`

GetAutoWarmupEnabled returns the AutoWarmupEnabled field if non-nil, zero value otherwise.

### GetAutoWarmupEnabledOk

`func (o *IP) GetAutoWarmupEnabledOk() (*bool, bool)`

GetAutoWarmupEnabledOk returns a tuple with the AutoWarmupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWarmupEnabled

`func (o *IP) SetAutoWarmupEnabled(v bool)`

SetAutoWarmupEnabled sets AutoWarmupEnabled field to given value.

### HasAutoWarmupEnabled

`func (o *IP) HasAutoWarmupEnabled() bool`

HasAutoWarmupEnabled returns a boolean if a field has been set.

### GetLabels

`func (o *IP) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IP) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IP) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IP) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetState

`func (o *IP) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IP) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IP) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *IP) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreated

`func (o *IP) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IP) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IP) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IP) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


