# RStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The date these statistics apply to (YYYY-MM-DD, UTC). | [optional] 
**Stat** | Pointer to [**Stat**](Stat.md) |  | [optional] 

## Methods

### NewRStat

`func NewRStat() *RStat`

NewRStat instantiates a new RStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRStatWithDefaults

`func NewRStatWithDefaults() *RStat`

NewRStatWithDefaults instantiates a new RStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RStat) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RStat) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RStat) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RStat) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStat

`func (o *RStat) GetStat() Stat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *RStat) GetStatOk() (*Stat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *RStat) SetStat(v Stat)`

SetStat sets Stat field to given value.

### HasStat

`func (o *RStat) HasStat() bool`

HasStat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


