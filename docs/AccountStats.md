# AccountStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Stat** | Pointer to [**AccountStatsStat**](AccountStatsStat.md) |  | [optional] 

## Methods

### NewAccountStats

`func NewAccountStats() *AccountStats`

NewAccountStats instantiates a new AccountStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatsWithDefaults

`func NewAccountStatsWithDefaults() *AccountStats`

NewAccountStatsWithDefaults instantiates a new AccountStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AccountStats) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AccountStats) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AccountStats) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AccountStats) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStat

`func (o *AccountStats) GetStat() AccountStatsStat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *AccountStats) GetStatOk() (*AccountStatsStat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *AccountStats) SetStat(v AccountStatsStat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *AccountStats) HasStat() bool`

HasStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


