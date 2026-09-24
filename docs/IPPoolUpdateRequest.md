# IPPoolUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New display name for the IP pool | [optional] 
**Ips** | Pointer to [**[]EIP**](EIP.md) | Updated list of IP addresses for this pool. This replaces the current IP list - include all IPs you want in the pool.  | [optional] 
**Tpsps** | Pointer to **[]int64** | Updated list of third-party sending provider IDs | [optional] 
**RoutingStrategy** | Pointer to **int32** | Updated routing strategy (see IPPoolCreateRequest for values) | [optional] 
**RoutingMetaData** | Pointer to **string** | Updated routing configuration (JSON) | [optional] 
**ShouldOverflow** | Pointer to **bool** | Whether to enable overflow to backup pool | [optional] 
**OverflowPoolName** | Pointer to **string** | Name of the overflow pool | [optional] 

## Methods

### NewIPPoolUpdateRequest

`func NewIPPoolUpdateRequest() *IPPoolUpdateRequest`

NewIPPoolUpdateRequest instantiates a new IPPoolUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPPoolUpdateRequestWithDefaults

`func NewIPPoolUpdateRequestWithDefaults() *IPPoolUpdateRequest`

NewIPPoolUpdateRequestWithDefaults instantiates a new IPPoolUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPPoolUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPPoolUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPPoolUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPPoolUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIps

`func (o *IPPoolUpdateRequest) GetIps() []EIP`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IPPoolUpdateRequest) GetIpsOk() (*[]EIP, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IPPoolUpdateRequest) SetIps(v []EIP)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IPPoolUpdateRequest) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetTpsps

`func (o *IPPoolUpdateRequest) GetTpsps() []int64`

GetTpsps returns the Tpsps field if non-nil, zero value otherwise.

### GetTpspsOk

`func (o *IPPoolUpdateRequest) GetTpspsOk() (*[]int64, bool)`

GetTpspsOk returns a tuple with the Tpsps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpsps

`func (o *IPPoolUpdateRequest) SetTpsps(v []int64)`

SetTpsps sets Tpsps field to given value.

### HasTpsps

`func (o *IPPoolUpdateRequest) HasTpsps() bool`

HasTpsps returns a boolean if a field has been set.

### GetRoutingStrategy

`func (o *IPPoolUpdateRequest) GetRoutingStrategy() int32`

GetRoutingStrategy returns the RoutingStrategy field if non-nil, zero value otherwise.

### GetRoutingStrategyOk

`func (o *IPPoolUpdateRequest) GetRoutingStrategyOk() (*int32, bool)`

GetRoutingStrategyOk returns a tuple with the RoutingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingStrategy

`func (o *IPPoolUpdateRequest) SetRoutingStrategy(v int32)`

SetRoutingStrategy sets RoutingStrategy field to given value.

### HasRoutingStrategy

`func (o *IPPoolUpdateRequest) HasRoutingStrategy() bool`

HasRoutingStrategy returns a boolean if a field has been set.

### GetRoutingMetaData

`func (o *IPPoolUpdateRequest) GetRoutingMetaData() string`

GetRoutingMetaData returns the RoutingMetaData field if non-nil, zero value otherwise.

### GetRoutingMetaDataOk

`func (o *IPPoolUpdateRequest) GetRoutingMetaDataOk() (*string, bool)`

GetRoutingMetaDataOk returns a tuple with the RoutingMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMetaData

`func (o *IPPoolUpdateRequest) SetRoutingMetaData(v string)`

SetRoutingMetaData sets RoutingMetaData field to given value.

### HasRoutingMetaData

`func (o *IPPoolUpdateRequest) HasRoutingMetaData() bool`

HasRoutingMetaData returns a boolean if a field has been set.

### GetShouldOverflow

`func (o *IPPoolUpdateRequest) GetShouldOverflow() bool`

GetShouldOverflow returns the ShouldOverflow field if non-nil, zero value otherwise.

### GetShouldOverflowOk

`func (o *IPPoolUpdateRequest) GetShouldOverflowOk() (*bool, bool)`

GetShouldOverflowOk returns a tuple with the ShouldOverflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldOverflow

`func (o *IPPoolUpdateRequest) SetShouldOverflow(v bool)`

SetShouldOverflow sets ShouldOverflow field to given value.

### HasShouldOverflow

`func (o *IPPoolUpdateRequest) HasShouldOverflow() bool`

HasShouldOverflow returns a boolean if a field has been set.

### GetOverflowPoolName

`func (o *IPPoolUpdateRequest) GetOverflowPoolName() string`

GetOverflowPoolName returns the OverflowPoolName field if non-nil, zero value otherwise.

### GetOverflowPoolNameOk

`func (o *IPPoolUpdateRequest) GetOverflowPoolNameOk() (*string, bool)`

GetOverflowPoolNameOk returns a tuple with the OverflowPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowPoolName

`func (o *IPPoolUpdateRequest) SetOverflowPoolName(v string)`

SetOverflowPoolName sets OverflowPoolName field to given value.

### HasOverflowPoolName

`func (o *IPPoolUpdateRequest) HasOverflowPoolName() bool`

HasOverflowPoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


