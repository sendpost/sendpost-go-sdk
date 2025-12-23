# DeleteSuppression200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the deleted suppression | [optional] 
**Message** | Pointer to **string** | A success message after the suppression(s) are deleted | [optional] 

## Methods

### NewDeleteSuppression200ResponseInner

`func NewDeleteSuppression200ResponseInner() *DeleteSuppression200ResponseInner`

NewDeleteSuppression200ResponseInner instantiates a new DeleteSuppression200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSuppression200ResponseInnerWithDefaults

`func NewDeleteSuppression200ResponseInnerWithDefaults() *DeleteSuppression200ResponseInner`

NewDeleteSuppression200ResponseInnerWithDefaults instantiates a new DeleteSuppression200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteSuppression200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteSuppression200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteSuppression200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteSuppression200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *DeleteSuppression200ResponseInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteSuppression200ResponseInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteSuppression200ResponseInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeleteSuppression200ResponseInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


