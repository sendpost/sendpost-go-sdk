# IPPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Ips** | Pointer to [**[]IP**](IP.md) |  | [optional] 
**ThirdPartySendingProviders** | Pointer to [**[]ThirdPartySendingProvider**](ThirdPartySendingProvider.md) |  | [optional] 
**RoutingStrategy** | Pointer to **int32** |  | [optional] 
**RoutingMetaData** | Pointer to **string** |  | [optional] 
**AutoWarmupEnabled** | Pointer to **bool** |  | [optional] 
**InfraMonitor** | Pointer to **bool** |  | [optional] 
**IpDomainWarmupStatus** | Pointer to **string** |  | [optional] 
**ShouldOverflow** | Pointer to **bool** | Indicates whether the IP should overflow, once email capacity of the IP Pool has been reached, should we send remaining emails over shared IP or not | [optional] 
**OverflowPoolName** | Pointer to **string** | The name of the overflow pool | [optional] 
**WarmupInterval** | Pointer to **int32** | The interval for the warmup | [optional] 

## Methods

### NewIPPool

`func NewIPPool() *IPPool`

NewIPPool instantiates a new IPPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPPoolWithDefaults

`func NewIPPoolWithDefaults() *IPPool`

NewIPPoolWithDefaults instantiates a new IPPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPPool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPPool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPPool) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IPPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IPPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *IPPool) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPPool) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPPool) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IPPool) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIps

`func (o *IPPool) GetIps() []IP`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IPPool) GetIpsOk() (*[]IP, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IPPool) SetIps(v []IP)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IPPool) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetThirdPartySendingProviders

`func (o *IPPool) GetThirdPartySendingProviders() []ThirdPartySendingProvider`

GetThirdPartySendingProviders returns the ThirdPartySendingProviders field if non-nil, zero value otherwise.

### GetThirdPartySendingProvidersOk

`func (o *IPPool) GetThirdPartySendingProvidersOk() (*[]ThirdPartySendingProvider, bool)`

GetThirdPartySendingProvidersOk returns a tuple with the ThirdPartySendingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartySendingProviders

`func (o *IPPool) SetThirdPartySendingProviders(v []ThirdPartySendingProvider)`

SetThirdPartySendingProviders sets ThirdPartySendingProviders field to given value.

### HasThirdPartySendingProviders

`func (o *IPPool) HasThirdPartySendingProviders() bool`

HasThirdPartySendingProviders returns a boolean if a field has been set.

### GetRoutingStrategy

`func (o *IPPool) GetRoutingStrategy() int32`

GetRoutingStrategy returns the RoutingStrategy field if non-nil, zero value otherwise.

### GetRoutingStrategyOk

`func (o *IPPool) GetRoutingStrategyOk() (*int32, bool)`

GetRoutingStrategyOk returns a tuple with the RoutingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingStrategy

`func (o *IPPool) SetRoutingStrategy(v int32)`

SetRoutingStrategy sets RoutingStrategy field to given value.

### HasRoutingStrategy

`func (o *IPPool) HasRoutingStrategy() bool`

HasRoutingStrategy returns a boolean if a field has been set.

### GetRoutingMetaData

`func (o *IPPool) GetRoutingMetaData() string`

GetRoutingMetaData returns the RoutingMetaData field if non-nil, zero value otherwise.

### GetRoutingMetaDataOk

`func (o *IPPool) GetRoutingMetaDataOk() (*string, bool)`

GetRoutingMetaDataOk returns a tuple with the RoutingMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMetaData

`func (o *IPPool) SetRoutingMetaData(v string)`

SetRoutingMetaData sets RoutingMetaData field to given value.

### HasRoutingMetaData

`func (o *IPPool) HasRoutingMetaData() bool`

HasRoutingMetaData returns a boolean if a field has been set.

### GetAutoWarmupEnabled

`func (o *IPPool) GetAutoWarmupEnabled() bool`

GetAutoWarmupEnabled returns the AutoWarmupEnabled field if non-nil, zero value otherwise.

### GetAutoWarmupEnabledOk

`func (o *IPPool) GetAutoWarmupEnabledOk() (*bool, bool)`

GetAutoWarmupEnabledOk returns a tuple with the AutoWarmupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWarmupEnabled

`func (o *IPPool) SetAutoWarmupEnabled(v bool)`

SetAutoWarmupEnabled sets AutoWarmupEnabled field to given value.

### HasAutoWarmupEnabled

`func (o *IPPool) HasAutoWarmupEnabled() bool`

HasAutoWarmupEnabled returns a boolean if a field has been set.

### GetInfraMonitor

`func (o *IPPool) GetInfraMonitor() bool`

GetInfraMonitor returns the InfraMonitor field if non-nil, zero value otherwise.

### GetInfraMonitorOk

`func (o *IPPool) GetInfraMonitorOk() (*bool, bool)`

GetInfraMonitorOk returns a tuple with the InfraMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraMonitor

`func (o *IPPool) SetInfraMonitor(v bool)`

SetInfraMonitor sets InfraMonitor field to given value.

### HasInfraMonitor

`func (o *IPPool) HasInfraMonitor() bool`

HasInfraMonitor returns a boolean if a field has been set.

### GetIpDomainWarmupStatus

`func (o *IPPool) GetIpDomainWarmupStatus() string`

GetIpDomainWarmupStatus returns the IpDomainWarmupStatus field if non-nil, zero value otherwise.

### GetIpDomainWarmupStatusOk

`func (o *IPPool) GetIpDomainWarmupStatusOk() (*string, bool)`

GetIpDomainWarmupStatusOk returns a tuple with the IpDomainWarmupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomainWarmupStatus

`func (o *IPPool) SetIpDomainWarmupStatus(v string)`

SetIpDomainWarmupStatus sets IpDomainWarmupStatus field to given value.

### HasIpDomainWarmupStatus

`func (o *IPPool) HasIpDomainWarmupStatus() bool`

HasIpDomainWarmupStatus returns a boolean if a field has been set.

### GetShouldOverflow

`func (o *IPPool) GetShouldOverflow() bool`

GetShouldOverflow returns the ShouldOverflow field if non-nil, zero value otherwise.

### GetShouldOverflowOk

`func (o *IPPool) GetShouldOverflowOk() (*bool, bool)`

GetShouldOverflowOk returns a tuple with the ShouldOverflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldOverflow

`func (o *IPPool) SetShouldOverflow(v bool)`

SetShouldOverflow sets ShouldOverflow field to given value.

### HasShouldOverflow

`func (o *IPPool) HasShouldOverflow() bool`

HasShouldOverflow returns a boolean if a field has been set.

### GetOverflowPoolName

`func (o *IPPool) GetOverflowPoolName() string`

GetOverflowPoolName returns the OverflowPoolName field if non-nil, zero value otherwise.

### GetOverflowPoolNameOk

`func (o *IPPool) GetOverflowPoolNameOk() (*string, bool)`

GetOverflowPoolNameOk returns a tuple with the OverflowPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowPoolName

`func (o *IPPool) SetOverflowPoolName(v string)`

SetOverflowPoolName sets OverflowPoolName field to given value.

### HasOverflowPoolName

`func (o *IPPool) HasOverflowPoolName() bool`

HasOverflowPoolName returns a boolean if a field has been set.

### GetWarmupInterval

`func (o *IPPool) GetWarmupInterval() int32`

GetWarmupInterval returns the WarmupInterval field if non-nil, zero value otherwise.

### GetWarmupIntervalOk

`func (o *IPPool) GetWarmupIntervalOk() (*int32, bool)`

GetWarmupIntervalOk returns a tuple with the WarmupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupInterval

`func (o *IPPool) SetWarmupInterval(v int32)`

SetWarmupInterval sets WarmupInterval field to given value.

### HasWarmupInterval

`func (o *IPPool) HasWarmupInterval() bool`

HasWarmupInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


