# CreateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The domain name to add for sending emails. Must be a valid domain you own and can configure DNS records for. Example: piedpiper.com (not subdomain like mail.piedpiper.com)  | 

## Methods

### NewCreateDomainRequest

`func NewCreateDomainRequest(name string, ) *CreateDomainRequest`

NewCreateDomainRequest instantiates a new CreateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainRequestWithDefaults

`func NewCreateDomainRequestWithDefaults() *CreateDomainRequest`

NewCreateDomainRequestWithDefaults instantiates a new CreateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDomainRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


