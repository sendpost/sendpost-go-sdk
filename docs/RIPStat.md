# RIPStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date these IP statistics apply to (YYYY-MM-DD, UTC). | [optional] 
**Stat** | Pointer to [**IPStat**](IPStat.md) |  | [optional] 

## Methods

### NewRIPStat

`func NewRIPStat() *RIPStat`

NewRIPStat instantiates a new RIPStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRIPStatWithDefaults

`func NewRIPStatWithDefaults() *RIPStat`

NewRIPStatWithDefaults instantiates a new RIPStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RIPStat) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RIPStat) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RIPStat) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RIPStat) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStat

`func (o *RIPStat) GetStat() IPStat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *RIPStat) GetStatOk() (*IPStat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *RIPStat) SetStat(v IPStat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *RIPStat) HasStat() bool`

HasStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


