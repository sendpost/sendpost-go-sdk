# IPPoolCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to [**[]EIP**](EIP.md) |  | [optional] 
**Tpsps** | Pointer to **[]int32** |  | [optional] 
**RoutingStrategy** | Pointer to **int32** |  | [optional] 
**RoutingMetaData** | Pointer to **string** |  | [optional] 
**OverflowPool** | Pointer to **bool** |  | [optional] 

## Methods

### NewIPPoolCreateRequest

`func NewIPPoolCreateRequest() *IPPoolCreateRequest`

NewIPPoolCreateRequest instantiates a new IPPoolCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPPoolCreateRequestWithDefaults

`func NewIPPoolCreateRequestWithDefaults() *IPPoolCreateRequest`

NewIPPoolCreateRequestWithDefaults instantiates a new IPPoolCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPPoolCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPPoolCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPPoolCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPPoolCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIps

`func (o *IPPoolCreateRequest) GetIps() []EIP`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IPPoolCreateRequest) GetIpsOk() (*[]EIP, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IPPoolCreateRequest) SetIps(v []EIP)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IPPoolCreateRequest) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetTpsps

`func (o *IPPoolCreateRequest) GetTpsps() []int32`

GetTpsps returns the Tpsps field if non-nil, zero value otherwise.

### GetTpspsOk

`func (o *IPPoolCreateRequest) GetTpspsOk() (*[]int32, bool)`

GetTpspsOk returns a tuple with the Tpsps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpsps

`func (o *IPPoolCreateRequest) SetTpsps(v []int32)`

SetTpsps sets Tpsps field to given value.

### HasTpsps

`func (o *IPPoolCreateRequest) HasTpsps() bool`

HasTpsps returns a boolean if a field has been set.

### GetRoutingStrategy

`func (o *IPPoolCreateRequest) GetRoutingStrategy() int32`

GetRoutingStrategy returns the RoutingStrategy field if non-nil, zero value otherwise.

### GetRoutingStrategyOk

`func (o *IPPoolCreateRequest) GetRoutingStrategyOk() (*int32, bool)`

GetRoutingStrategyOk returns a tuple with the RoutingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingStrategy

`func (o *IPPoolCreateRequest) SetRoutingStrategy(v int32)`

SetRoutingStrategy sets RoutingStrategy field to given value.

### HasRoutingStrategy

`func (o *IPPoolCreateRequest) HasRoutingStrategy() bool`

HasRoutingStrategy returns a boolean if a field has been set.

### GetRoutingMetaData

`func (o *IPPoolCreateRequest) GetRoutingMetaData() string`

GetRoutingMetaData returns the RoutingMetaData field if non-nil, zero value otherwise.

### GetRoutingMetaDataOk

`func (o *IPPoolCreateRequest) GetRoutingMetaDataOk() (*string, bool)`

GetRoutingMetaDataOk returns a tuple with the RoutingMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMetaData

`func (o *IPPoolCreateRequest) SetRoutingMetaData(v string)`

SetRoutingMetaData sets RoutingMetaData field to given value.

### HasRoutingMetaData

`func (o *IPPoolCreateRequest) HasRoutingMetaData() bool`

HasRoutingMetaData returns a boolean if a field has been set.

### GetOverflowPool

`func (o *IPPoolCreateRequest) GetOverflowPool() bool`

GetOverflowPool returns the OverflowPool field if non-nil, zero value otherwise.

### GetOverflowPoolOk

`func (o *IPPoolCreateRequest) GetOverflowPoolOk() (*bool, bool)`

GetOverflowPoolOk returns a tuple with the OverflowPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowPool

`func (o *IPPoolCreateRequest) SetOverflowPool(v bool)`

SetOverflowPool sets OverflowPool field to given value.

### HasOverflowPool

`func (o *IPPoolCreateRequest) HasOverflowPool() bool`

HasOverflowPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


