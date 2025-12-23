# IPPoolUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to [**[]IP**](IP.md) |  | [optional] 
**RoutingStrategy** | Pointer to **int32** |  | [optional] 
**RoutingMetaData** | Pointer to **string** |  | [optional] 

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

`func (o *IPPoolUpdateRequest) GetIps() []IP`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IPPoolUpdateRequest) GetIpsOk() (*[]IP, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IPPoolUpdateRequest) SetIps(v []IP)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IPPoolUpdateRequest) HasIps() bool`

HasIps returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


