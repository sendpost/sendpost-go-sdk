# Label

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the label | [optional] 
**Name** | Pointer to **string** | Display name for the label (max 50 characters) | [optional] 
**Color** | Pointer to **string** | Hex color code for visual identification in the dashboard. Format: 6-character hex without # prefix.  | [optional] 
**Type** | Pointer to **int32** | Resource type this label applies to: - &#x60;0&#x60; &#x3D; IP label - &#x60;1&#x60; &#x3D; Sub-account label - &#x60;2&#x60; &#x3D; IP Pool label  | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the label was created | [optional] 

## Methods

### NewLabel

`func NewLabel() *Label`

NewLabel instantiates a new Label object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelWithDefaults

`func NewLabelWithDefaults() *Label`

NewLabelWithDefaults instantiates a new Label object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Label) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Label) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Label) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Label) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Label) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Label) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Label) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Label) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *Label) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Label) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Label) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Label) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetType

`func (o *Label) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Label) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Label) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Label) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *Label) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Label) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Label) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Label) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


