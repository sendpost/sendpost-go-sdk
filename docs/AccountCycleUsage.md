# AccountCycleUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processed** | Pointer to **int64** | Total emails processed by the account since the last billing cycle reset. | [optional] 

## Methods

### NewAccountCycleUsage

`func NewAccountCycleUsage() *AccountCycleUsage`

NewAccountCycleUsage instantiates a new AccountCycleUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCycleUsageWithDefaults

`func NewAccountCycleUsageWithDefaults() *AccountCycleUsage`

NewAccountCycleUsageWithDefaults instantiates a new AccountCycleUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *AccountCycleUsage) GetProcessed() int64`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *AccountCycleUsage) GetProcessedOk() (*int64, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *AccountCycleUsage) SetProcessed(v int64)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *AccountCycleUsage) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


