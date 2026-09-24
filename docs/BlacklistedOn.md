# BlacklistedOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rbl** | Pointer to **string** | Name/host of the real-time blacklist (RBL) the target is listed on. | [optional] 
**Delist** | Pointer to **string** | URL or instructions for requesting delisting from this RBL. | [optional] 

## Methods

### NewBlacklistedOn

`func NewBlacklistedOn() *BlacklistedOn`

NewBlacklistedOn instantiates a new BlacklistedOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlacklistedOnWithDefaults

`func NewBlacklistedOnWithDefaults() *BlacklistedOn`

NewBlacklistedOnWithDefaults instantiates a new BlacklistedOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRbl

`func (o *BlacklistedOn) GetRbl() string`

GetRbl returns the Rbl field if non-nil, zero value otherwise.

### GetRblOk

`func (o *BlacklistedOn) GetRblOk() (*string, bool)`

GetRblOk returns a tuple with the Rbl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbl

`func (o *BlacklistedOn) SetRbl(v string)`

SetRbl sets Rbl field to given value.

### HasRbl

`func (o *BlacklistedOn) HasRbl() bool`

HasRbl returns a boolean if a field has been set.

### GetDelist

`func (o *BlacklistedOn) GetDelist() string`

GetDelist returns the Delist field if non-nil, zero value otherwise.

### GetDelistOk

`func (o *BlacklistedOn) GetDelistOk() (*string, bool)`

GetDelistOk returns a tuple with the Delist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelist

`func (o *BlacklistedOn) SetDelist(v string)`

SetDelist sets Delist field to given value.

### HasDelist

`func (o *BlacklistedOn) HasDelist() bool`

HasDelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


