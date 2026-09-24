# IPPoolCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the IP pool. Must be unique within your account. Use descriptive names like \&quot;transactional\&quot;, \&quot;marketing-bulk\&quot;, \&quot;high-priority\&quot;  | 
**Ips** | Pointer to [**[]EIP**](EIP.md) | List of dedicated IP addresses to include in this pool. IPs must already be allocated to your account.  | [optional] 
**Tpsps** | Pointer to **[]int64** | List of third-party sending provider IDs to include in this pool. TPSPs must be pre-configured in your account.  | [optional] 
**RoutingStrategy** | Pointer to **int32** | Email routing strategy: - &#x60;0&#x60; &#x3D; Round Robin (equal distribution) - &#x60;1&#x60; &#x3D; Email Provider Strategy (route by recipient domain) - &#x60;2&#x60; &#x3D; Volume Percentage Strategy (weighted distribution) - &#x60;3&#x60; &#x3D; Sending Domain Strategy (route by sender domain)  | [optional] [default to 0]
**RoutingMetaData** | Pointer to **string** | JSON-encoded routing configuration. See IPPools documentation for format. Use &#x60;{}&#x60; for round-robin strategy.  | [optional] [default to "{}"]
**ShouldOverflow** | Pointer to **bool** | Whether to overflow to shared pool when this pool is unavailable | [optional] [default to false]
**OverflowPoolName** | Pointer to **string** | Name of the IP pool to overflow to (if shouldOverflow is true) | [optional] 

## Methods

### NewIPPoolCreateRequest

`func NewIPPoolCreateRequest(name string, ) *IPPoolCreateRequest`

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

`func (o *IPPoolCreateRequest) GetTpsps() []int64`

GetTpsps returns the Tpsps field if non-nil, zero value otherwise.

### GetTpspsOk

`func (o *IPPoolCreateRequest) GetTpspsOk() (*[]int64, bool)`

GetTpspsOk returns a tuple with the Tpsps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpsps

`func (o *IPPoolCreateRequest) SetTpsps(v []int64)`

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

### GetShouldOverflow

`func (o *IPPoolCreateRequest) GetShouldOverflow() bool`

GetShouldOverflow returns the ShouldOverflow field if non-nil, zero value otherwise.

### GetShouldOverflowOk

`func (o *IPPoolCreateRequest) GetShouldOverflowOk() (*bool, bool)`

GetShouldOverflowOk returns a tuple with the ShouldOverflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldOverflow

`func (o *IPPoolCreateRequest) SetShouldOverflow(v bool)`

SetShouldOverflow sets ShouldOverflow field to given value.

### HasShouldOverflow

`func (o *IPPoolCreateRequest) HasShouldOverflow() bool`

HasShouldOverflow returns a boolean if a field has been set.

### GetOverflowPoolName

`func (o *IPPoolCreateRequest) GetOverflowPoolName() string`

GetOverflowPoolName returns the OverflowPoolName field if non-nil, zero value otherwise.

### GetOverflowPoolNameOk

`func (o *IPPoolCreateRequest) GetOverflowPoolNameOk() (*string, bool)`

GetOverflowPoolNameOk returns a tuple with the OverflowPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowPoolName

`func (o *IPPoolCreateRequest) SetOverflowPoolName(v string)`

SetOverflowPoolName sets OverflowPoolName field to given value.

### HasOverflowPoolName

`func (o *IPPoolCreateRequest) HasOverflowPoolName() bool`

HasOverflowPoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


