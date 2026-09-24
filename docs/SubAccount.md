# SubAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the sub-account | [optional] 
**AccountId** | Pointer to **int64** | Identifier of the parent account this sub-account belongs to | [optional] 
**Name** | Pointer to **string** | Display name for the sub-account. Must be unique within your account. Use descriptive names.  | [optional] 
**ApiKey** | Pointer to **string** | API key for this sub-account. Use this as the &#x60;X-SubAccount-ApiKey&#x60; header when making API calls for this sub-account (sending emails, managing domains, etc.).  **Security:** Treat this like a password. Rotate if compromised.  | [optional] 
**Type** | Pointer to **int32** | Type of sub-account: - &#x60;0&#x60; &#x3D; Default (the primary sub-account created with your account) - &#x60;1&#x60; &#x3D; Custom (additional sub-accounts you create)  Note: The default sub-account cannot be deleted.  | [optional] 
**IsPlus** | Pointer to **bool** | Whether this sub-account belongs to a SendX Plus customer. SendX Plus is a premium tier that provides enhanced features and support.  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) | Custom labels for organizing and filtering sub-accounts | [optional] 
**Blocked** | Pointer to **bool** | Whether the sub-account is blocked from sending. A blocked sub-account cannot send emails. Common reasons: - High bounce/spam rates - Billing issues - Policy violations - Manual suspension by administrator  | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the sub-account was created | [optional] 

## Methods

### NewSubAccount

`func NewSubAccount() *SubAccount`

NewSubAccount instantiates a new SubAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountWithDefaults

`func NewSubAccountWithDefaults() *SubAccount`

NewSubAccountWithDefaults instantiates a new SubAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SubAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *SubAccount) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubAccount) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubAccount) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SubAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *SubAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiKey

`func (o *SubAccount) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *SubAccount) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *SubAccount) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *SubAccount) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetType

`func (o *SubAccount) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubAccount) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubAccount) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SubAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsPlus

`func (o *SubAccount) GetIsPlus() bool`

GetIsPlus returns the IsPlus field if non-nil, zero value otherwise.

### GetIsPlusOk

`func (o *SubAccount) GetIsPlusOk() (*bool, bool)`

GetIsPlusOk returns a tuple with the IsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlus

`func (o *SubAccount) SetIsPlus(v bool)`

SetIsPlus sets IsPlus field to given value.

### HasIsPlus

`func (o *SubAccount) HasIsPlus() bool`

HasIsPlus returns a boolean if a field has been set.

### GetLabels

`func (o *SubAccount) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SubAccount) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SubAccount) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SubAccount) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetBlocked

`func (o *SubAccount) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *SubAccount) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *SubAccount) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *SubAccount) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCreated

`func (o *SubAccount) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubAccount) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubAccount) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SubAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


