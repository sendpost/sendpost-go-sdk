# Stat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Date for which stats are retrieved (UTC). | [optional] 
**Stats** | Pointer to [**StatStats**](StatStats.md) |  | [optional] 

## Methods

### NewStat

`func NewStat() *Stat`

NewStat instantiates a new Stat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatWithDefaults

`func NewStatWithDefaults() *Stat`

NewStatWithDefaults instantiates a new Stat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *Stat) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Stat) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Stat) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Stat) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStats

`func (o *Stat) GetStats() StatStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Stat) GetStatsOk() (*StatStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Stat) SetStats(v StatStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Stat) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


