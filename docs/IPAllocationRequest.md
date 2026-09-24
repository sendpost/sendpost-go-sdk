# IPAllocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **[]string** | List of IP addresses to allocate. These must be available IPs from SendPost&#39;s IP pool. Contact support to request IP allocation.  | 
**AutoWarmupEnabled** | Pointer to **bool** | Enable automatic IP warmup for newly allocated IPs. Recommended: true for new IPs to gradually build sender reputation.  | [optional] [default to true]

## Methods

### NewIPAllocationRequest

`func NewIPAllocationRequest(ips []string, ) *IPAllocationRequest`

NewIPAllocationRequest instantiates a new IPAllocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAllocationRequestWithDefaults

`func NewIPAllocationRequestWithDefaults() *IPAllocationRequest`

NewIPAllocationRequestWithDefaults instantiates a new IPAllocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *IPAllocationRequest) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IPAllocationRequest) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IPAllocationRequest) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetAutoWarmupEnabled

`func (o *IPAllocationRequest) GetAutoWarmupEnabled() bool`

GetAutoWarmupEnabled returns the AutoWarmupEnabled field if non-nil, zero value otherwise.

### GetAutoWarmupEnabledOk

`func (o *IPAllocationRequest) GetAutoWarmupEnabledOk() (*bool, bool)`

GetAutoWarmupEnabledOk returns a tuple with the AutoWarmupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWarmupEnabled

`func (o *IPAllocationRequest) SetAutoWarmupEnabled(v bool)`

SetAutoWarmupEnabled sets AutoWarmupEnabled field to given value.

### HasAutoWarmupEnabled

`func (o *IPAllocationRequest) HasAutoWarmupEnabled() bool`

HasAutoWarmupEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


