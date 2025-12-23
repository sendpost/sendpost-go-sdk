# IPAllocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverflowPool** | **bool** | Determines whether emails should be sent over shared IP when the IP pool is full | 
**Ips** | **[]string** |  | 

## Methods

### NewIPAllocationRequest

`func NewIPAllocationRequest(overflowPool bool, ips []string, ) *IPAllocationRequest`

NewIPAllocationRequest instantiates a new IPAllocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAllocationRequestWithDefaults

`func NewIPAllocationRequestWithDefaults() *IPAllocationRequest`

NewIPAllocationRequestWithDefaults instantiates a new IPAllocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverflowPool

`func (o *IPAllocationRequest) GetOverflowPool() bool`

GetOverflowPool returns the OverflowPool field if non-nil, zero value otherwise.

### GetOverflowPoolOk

`func (o *IPAllocationRequest) GetOverflowPoolOk() (*bool, bool)`

GetOverflowPoolOk returns a tuple with the OverflowPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowPool

`func (o *IPAllocationRequest) SetOverflowPool(v bool)`

SetOverflowPool sets OverflowPool field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


