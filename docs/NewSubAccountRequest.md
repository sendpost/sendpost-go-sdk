# NewSubAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the new sub-account. Must be unique within your account. Use descriptive names like \&quot;Marketing - Production\&quot; or \&quot;Customer: Acme Corp\&quot;  | 

## Methods

### NewNewSubAccountRequest

`func NewNewSubAccountRequest(name string, ) *NewSubAccountRequest`

NewNewSubAccountRequest instantiates a new NewSubAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewSubAccountRequestWithDefaults

`func NewNewSubAccountRequestWithDefaults() *NewSubAccountRequest`

NewNewSubAccountRequestWithDefaults instantiates a new NewSubAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewSubAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewSubAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewSubAccountRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


