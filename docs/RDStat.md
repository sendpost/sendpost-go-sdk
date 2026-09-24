# RDStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date these domain statistics apply to (YYYY-MM-DD, UTC). | [optional] 
**Stat** | Pointer to [**Stat**](Stat.md) |  | [optional] 

## Methods

### NewRDStat

`func NewRDStat() *RDStat`

NewRDStat instantiates a new RDStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDStatWithDefaults

`func NewRDStatWithDefaults() *RDStat`

NewRDStatWithDefaults instantiates a new RDStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RDStat) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RDStat) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RDStat) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RDStat) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStat

`func (o *RDStat) GetStat() Stat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *RDStat) GetStatOk() (*Stat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *RDStat) SetStat(v Stat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *RDStat) HasStat() bool`

HasStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


