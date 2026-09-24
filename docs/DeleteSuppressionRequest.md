# DeleteSuppressionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suppressions** | Pointer to [**[]DeleteSuppressionRequestSuppressionsInner**](DeleteSuppressionRequestSuppressionsInner.md) | List of email addresses to remove from suppression. Each email will be removed regardless of suppression type. | [optional] 

## Methods

### NewDeleteSuppressionRequest

`func NewDeleteSuppressionRequest() *DeleteSuppressionRequest`

NewDeleteSuppressionRequest instantiates a new DeleteSuppressionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSuppressionRequestWithDefaults

`func NewDeleteSuppressionRequestWithDefaults() *DeleteSuppressionRequest`

NewDeleteSuppressionRequestWithDefaults instantiates a new DeleteSuppressionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppressions

`func (o *DeleteSuppressionRequest) GetSuppressions() []DeleteSuppressionRequestSuppressionsInner`

GetSuppressions returns the Suppressions field if non-nil, zero value otherwise.

### GetSuppressionsOk

`func (o *DeleteSuppressionRequest) GetSuppressionsOk() (*[]DeleteSuppressionRequestSuppressionsInner, bool)`

GetSuppressionsOk returns a tuple with the Suppressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressions

`func (o *DeleteSuppressionRequest) SetSuppressions(v []DeleteSuppressionRequestSuppressionsInner)`

SetSuppressions sets Suppressions field to given value.

### HasSuppressions

`func (o *DeleteSuppressionRequest) HasSuppressions() bool`

HasSuppressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


