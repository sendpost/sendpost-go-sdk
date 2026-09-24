# SeedContactStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stat** | Pointer to [**Stat**](Stat.md) |  | [optional] 
**RStats** | Pointer to [**[]RStat**](RStat.md) | Per-day statistics for the seed-contact run. | [optional] 
**GroupStats** | Pointer to [**[]GroupStat**](GroupStat.md) | Statistics broken down by group/tag. | [optional] 
**DomainStats** | Pointer to [**[]DomainStat**](DomainStat.md) | Statistics broken down by sending domain. | [optional] 
**IpStats** | Pointer to [**[]IPStat**](IPStat.md) | Statistics broken down by sending IP address. | [optional] 
**ProviderStats** | Pointer to [**[]ProviderStat**](ProviderStat.md) | Statistics broken down by email provider. | [optional] 

## Methods

### NewSeedContactStats

`func NewSeedContactStats() *SeedContactStats`

NewSeedContactStats instantiates a new SeedContactStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeedContactStatsWithDefaults

`func NewSeedContactStatsWithDefaults() *SeedContactStats`

NewSeedContactStatsWithDefaults instantiates a new SeedContactStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStat

`func (o *SeedContactStats) GetStat() Stat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *SeedContactStats) GetStatOk() (*Stat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *SeedContactStats) SetStat(v Stat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *SeedContactStats) HasStat() bool`

HasStat returns a boolean if a field has been set.

### GetRStats

`func (o *SeedContactStats) GetRStats() []RStat`

GetRStats returns the RStats field if non-nil, zero value otherwise.

### GetRStatsOk

`func (o *SeedContactStats) GetRStatsOk() (*[]RStat, bool)`

GetRStatsOk returns a tuple with the RStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRStats

`func (o *SeedContactStats) SetRStats(v []RStat)`

SetRStats sets RStats field to given value.

### HasRStats

`func (o *SeedContactStats) HasRStats() bool`

HasRStats returns a boolean if a field has been set.

### GetGroupStats

`func (o *SeedContactStats) GetGroupStats() []GroupStat`

GetGroupStats returns the GroupStats field if non-nil, zero value otherwise.

### GetGroupStatsOk

`func (o *SeedContactStats) GetGroupStatsOk() (*[]GroupStat, bool)`

GetGroupStatsOk returns a tuple with the GroupStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupStats

`func (o *SeedContactStats) SetGroupStats(v []GroupStat)`

SetGroupStats sets GroupStats field to given value.

### HasGroupStats

`func (o *SeedContactStats) HasGroupStats() bool`

HasGroupStats returns a boolean if a field has been set.

### GetDomainStats

`func (o *SeedContactStats) GetDomainStats() []DomainStat`

GetDomainStats returns the DomainStats field if non-nil, zero value otherwise.

### GetDomainStatsOk

`func (o *SeedContactStats) GetDomainStatsOk() (*[]DomainStat, bool)`

GetDomainStatsOk returns a tuple with the DomainStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStats

`func (o *SeedContactStats) SetDomainStats(v []DomainStat)`

SetDomainStats sets DomainStats field to given value.

### HasDomainStats

`func (o *SeedContactStats) HasDomainStats() bool`

HasDomainStats returns a boolean if a field has been set.

### GetIpStats

`func (o *SeedContactStats) GetIpStats() []IPStat`

GetIpStats returns the IpStats field if non-nil, zero value otherwise.

### GetIpStatsOk

`func (o *SeedContactStats) GetIpStatsOk() (*[]IPStat, bool)`

GetIpStatsOk returns a tuple with the IpStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStats

`func (o *SeedContactStats) SetIpStats(v []IPStat)`

SetIpStats sets IpStats field to given value.

### HasIpStats

`func (o *SeedContactStats) HasIpStats() bool`

HasIpStats returns a boolean if a field has been set.

### GetProviderStats

`func (o *SeedContactStats) GetProviderStats() []ProviderStat`

GetProviderStats returns the ProviderStats field if non-nil, zero value otherwise.

### GetProviderStatsOk

`func (o *SeedContactStats) GetProviderStatsOk() (*[]ProviderStat, bool)`

GetProviderStatsOk returns a tuple with the ProviderStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderStats

`func (o *SeedContactStats) SetProviderStats(v []ProviderStat)`

SetProviderStats sets ProviderStats field to given value.

### HasProviderStats

`func (o *SeedContactStats) HasProviderStats() bool`

HasProviderStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


