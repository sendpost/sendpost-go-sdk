# ValidationStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initiated** | Pointer to **int64** | Number of validation requests initiated. | [optional] 
**Processed** | Pointer to **int64** | Number of validation requests fully processed. | [optional] 
**Valid** | Pointer to **int64** | Number of addresses determined to be valid/deliverable. | [optional] 
**Invalid** | Pointer to **int64** | Number of addresses determined to be invalid/undeliverable. | [optional] 
**SoftBounced** | Pointer to **int64** | Number of addresses that soft-bounced during validation. | [optional] 
**HardBounced** | Pointer to **int64** | Number of addresses that hard-bounced during validation. | [optional] 
**CatchAll** | Pointer to **int64** | Number of addresses on catch-all (accept-all) domains. | [optional] 
**Unknown** | Pointer to **int64** | Number of addresses whose deliverability could not be determined. | [optional] 

## Methods

### NewValidationStat

`func NewValidationStat() *ValidationStat`

NewValidationStat instantiates a new ValidationStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationStatWithDefaults

`func NewValidationStatWithDefaults() *ValidationStat`

NewValidationStatWithDefaults instantiates a new ValidationStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiated

`func (o *ValidationStat) GetInitiated() int64`

GetInitiated returns the Initiated field if non-nil, zero value otherwise.

### GetInitiatedOk

`func (o *ValidationStat) GetInitiatedOk() (*int64, bool)`

GetInitiatedOk returns a tuple with the Initiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiated

`func (o *ValidationStat) SetInitiated(v int64)`

SetInitiated sets Initiated field to given value.

### HasInitiated

`func (o *ValidationStat) HasInitiated() bool`

HasInitiated returns a boolean if a field has been set.

### GetProcessed

`func (o *ValidationStat) GetProcessed() int64`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *ValidationStat) GetProcessedOk() (*int64, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *ValidationStat) SetProcessed(v int64)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *ValidationStat) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetValid

`func (o *ValidationStat) GetValid() int64`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidationStat) GetValidOk() (*int64, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidationStat) SetValid(v int64)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ValidationStat) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetInvalid

`func (o *ValidationStat) GetInvalid() int64`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *ValidationStat) GetInvalidOk() (*int64, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *ValidationStat) SetInvalid(v int64)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *ValidationStat) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetSoftBounced

`func (o *ValidationStat) GetSoftBounced() int64`

GetSoftBounced returns the SoftBounced field if non-nil, zero value otherwise.

### GetSoftBouncedOk

`func (o *ValidationStat) GetSoftBouncedOk() (*int64, bool)`

GetSoftBouncedOk returns a tuple with the SoftBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounced

`func (o *ValidationStat) SetSoftBounced(v int64)`

SetSoftBounced sets SoftBounced field to given value.

### HasSoftBounced

`func (o *ValidationStat) HasSoftBounced() bool`

HasSoftBounced returns a boolean if a field has been set.

### GetHardBounced

`func (o *ValidationStat) GetHardBounced() int64`

GetHardBounced returns the HardBounced field if non-nil, zero value otherwise.

### GetHardBouncedOk

`func (o *ValidationStat) GetHardBouncedOk() (*int64, bool)`

GetHardBouncedOk returns a tuple with the HardBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounced

`func (o *ValidationStat) SetHardBounced(v int64)`

SetHardBounced sets HardBounced field to given value.

### HasHardBounced

`func (o *ValidationStat) HasHardBounced() bool`

HasHardBounced returns a boolean if a field has been set.

### GetCatchAll

`func (o *ValidationStat) GetCatchAll() int64`

GetCatchAll returns the CatchAll field if non-nil, zero value otherwise.

### GetCatchAllOk

`func (o *ValidationStat) GetCatchAllOk() (*int64, bool)`

GetCatchAllOk returns a tuple with the CatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchAll

`func (o *ValidationStat) SetCatchAll(v int64)`

SetCatchAll sets CatchAll field to given value.

### HasCatchAll

`func (o *ValidationStat) HasCatchAll() bool`

HasCatchAll returns a boolean if a field has been set.

### GetUnknown

`func (o *ValidationStat) GetUnknown() int64`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *ValidationStat) GetUnknownOk() (*int64, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *ValidationStat) SetUnknown(v int64)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *ValidationStat) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


