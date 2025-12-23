# ThirdPartySendingProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**OauthToken** | Pointer to **string** |  | [optional] 
**RetryTime** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 

## Methods

### NewThirdPartySendingProvider

`func NewThirdPartySendingProvider() *ThirdPartySendingProvider`

NewThirdPartySendingProvider instantiates a new ThirdPartySendingProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartySendingProviderWithDefaults

`func NewThirdPartySendingProviderWithDefaults() *ThirdPartySendingProvider`

NewThirdPartySendingProviderWithDefaults instantiates a new ThirdPartySendingProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartySendingProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartySendingProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartySendingProvider) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ThirdPartySendingProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ThirdPartySendingProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThirdPartySendingProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThirdPartySendingProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThirdPartySendingProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ThirdPartySendingProvider) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdPartySendingProvider) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdPartySendingProvider) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ThirdPartySendingProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomain

`func (o *ThirdPartySendingProvider) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ThirdPartySendingProvider) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ThirdPartySendingProvider) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ThirdPartySendingProvider) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEndpoint

`func (o *ThirdPartySendingProvider) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ThirdPartySendingProvider) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ThirdPartySendingProvider) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ThirdPartySendingProvider) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetKey

`func (o *ThirdPartySendingProvider) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ThirdPartySendingProvider) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ThirdPartySendingProvider) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ThirdPartySendingProvider) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecret

`func (o *ThirdPartySendingProvider) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ThirdPartySendingProvider) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ThirdPartySendingProvider) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ThirdPartySendingProvider) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPort

`func (o *ThirdPartySendingProvider) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ThirdPartySendingProvider) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ThirdPartySendingProvider) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ThirdPartySendingProvider) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetOauthToken

`func (o *ThirdPartySendingProvider) GetOauthToken() string`

GetOauthToken returns the OauthToken field if non-nil, zero value otherwise.

### GetOauthTokenOk

`func (o *ThirdPartySendingProvider) GetOauthTokenOk() (*string, bool)`

GetOauthTokenOk returns a tuple with the OauthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthToken

`func (o *ThirdPartySendingProvider) SetOauthToken(v string)`

SetOauthToken sets OauthToken field to given value.

### HasOauthToken

`func (o *ThirdPartySendingProvider) HasOauthToken() bool`

HasOauthToken returns a boolean if a field has been set.

### GetRetryTime

`func (o *ThirdPartySendingProvider) GetRetryTime() int32`

GetRetryTime returns the RetryTime field if non-nil, zero value otherwise.

### GetRetryTimeOk

`func (o *ThirdPartySendingProvider) GetRetryTimeOk() (*int32, bool)`

GetRetryTimeOk returns a tuple with the RetryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTime

`func (o *ThirdPartySendingProvider) SetRetryTime(v int32)`

SetRetryTime sets RetryTime field to given value.

### HasRetryTime

`func (o *ThirdPartySendingProvider) HasRetryTime() bool`

HasRetryTime returns a boolean if a field has been set.

### GetCreated

`func (o *ThirdPartySendingProvider) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ThirdPartySendingProvider) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ThirdPartySendingProvider) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ThirdPartySendingProvider) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


