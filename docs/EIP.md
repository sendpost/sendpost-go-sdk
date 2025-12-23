# EIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIP** | **string** | list of IP resources which are a part of the IP Pool containing public IP information. Note that the IPs specified in the IPPool should have been allocated in advance for your account | 

## Methods

### NewEIP

`func NewEIP(publicIP string, ) *EIP`

NewEIP instantiates a new EIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEIPWithDefaults

`func NewEIPWithDefaults() *EIP`

NewEIPWithDefaults instantiates a new EIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIP

`func (o *EIP) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *EIP) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *EIP) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


