# IPPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the IP pool | [optional] 
**Name** | Pointer to **string** | Display name for the IP pool. Must be unique within your account. Use descriptive names like \&quot;transactional\&quot;, \&quot;marketing\&quot;, \&quot;high-priority\&quot;.  | [optional] 
**Type** | Pointer to **int32** | Type of IP pool: - &#x60;0&#x60; &#x3D; Shared (uses shared IPs with pooled reputation) - &#x60;1&#x60; &#x3D; Dedicated (uses dedicated IPs exclusive to your account)  | [optional] 
**RoutingStrategy** | Pointer to **int32** | How emails are distributed across IPs/providers in this pool: - &#x60;0&#x60; &#x3D; Round Robin (equal distribution) - &#x60;1&#x60; &#x3D; Email Provider Strategy (route by recipient domain like Gmail, Yahoo) - &#x60;2&#x60; &#x3D; Volume Percentage Strategy (weighted distribution) - &#x60;3&#x60; &#x3D; Sending Domain Strategy (route by sender domain)  See the IPPools tag description for detailed routing configuration examples.  | [optional] 
**RoutingMetaData** | Pointer to **string** | JSON-encoded configuration for the selected routing strategy. Format depends on routingStrategy value. See IPPools documentation for examples.  For Round Robin (strategy 0): Use empty object &#x60;{}&#x60;  | [optional] 
**ShouldOverflow** | Pointer to **bool** | Whether to automatically overflow to a backup pool when this pool is unavailable (all IPs down) or at capacity (warmup limits reached).  | [optional] 
**OverflowPoolName** | Pointer to **string** | Name of the IP pool to overflow to when shouldOverflow is enabled. The overflow pool must exist. Common pattern: overflow to shared IP pool.  | [optional] 
**Ips** | Pointer to [**[]IP**](IP.md) | List of dedicated IPs assigned to this pool | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the IP pool was created | [optional] 

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

`func (o *IPPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPPool) SetId(v int64)`

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

### GetType

`func (o *IPPool) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPPool) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPPool) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *IPPool) HasType() bool`

HasType returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


