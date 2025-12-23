# DeleteSubAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the deleted sub-account. | [optional] 
**Message** | Pointer to **string** | Message confirming the deletion. | [optional] 

## Methods

### NewDeleteSubAccountResponse

`func NewDeleteSubAccountResponse() *DeleteSubAccountResponse`

NewDeleteSubAccountResponse instantiates a new DeleteSubAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSubAccountResponseWithDefaults

`func NewDeleteSubAccountResponseWithDefaults() *DeleteSubAccountResponse`

NewDeleteSubAccountResponseWithDefaults instantiates a new DeleteSubAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteSubAccountResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteSubAccountResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteSubAccountResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteSubAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *DeleteSubAccountResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteSubAccountResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteSubAccountResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeleteSubAccountResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


