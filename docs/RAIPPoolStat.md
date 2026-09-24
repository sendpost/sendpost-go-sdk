# RAIPPoolStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date these IP-pool statistics apply to (YYYY-MM-DD, UTC). | [optional] 
**Stat** | Pointer to [**Stat**](Stat.md) |  | [optional] 

## Methods

### NewRAIPPoolStat

`func NewRAIPPoolStat() *RAIPPoolStat`

NewRAIPPoolStat instantiates a new RAIPPoolStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRAIPPoolStatWithDefaults

`func NewRAIPPoolStatWithDefaults() *RAIPPoolStat`

NewRAIPPoolStatWithDefaults instantiates a new RAIPPoolStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RAIPPoolStat) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RAIPPoolStat) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RAIPPoolStat) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RAIPPoolStat) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStat

`func (o *RAIPPoolStat) GetStat() Stat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *RAIPPoolStat) GetStatOk() (*Stat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *RAIPPoolStat) SetStat(v Stat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *RAIPPoolStat) HasStat() bool`

HasStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


