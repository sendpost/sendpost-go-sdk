# EmailStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The UTC date for which stats are retrieved | [optional] 
**Stats** | Pointer to [**EmailStatsStats**](EmailStatsStats.md) |  | [optional] 

## Methods

### NewEmailStats

`func NewEmailStats() *EmailStats`

NewEmailStats instantiates a new EmailStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStatsWithDefaults

`func NewEmailStatsWithDefaults() *EmailStats`

NewEmailStatsWithDefaults instantiates a new EmailStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EmailStats) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EmailStats) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EmailStats) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EmailStats) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStats

`func (o *EmailStats) GetStats() EmailStatsStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *EmailStats) GetStatsOk() (*EmailStatsStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *EmailStats) SetStats(v EmailStatsStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *EmailStats) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


