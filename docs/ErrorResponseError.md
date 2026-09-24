# ErrorResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Machine-readable error code. Common codes: - &#x60;invalid_request&#x60; - Malformed request body or parameters - &#x60;authentication_failed&#x60; - Invalid or missing API key - &#x60;resource_not_found&#x60; - Requested resource doesn&#39;t exist - &#x60;resource_exists&#x60; - Resource with same identifier already exists - &#x60;validation_error&#x60; - Request validation failed - &#x60;rate_limit_exceeded&#x60; - Too many requests - &#x60;internal_error&#x60; - Server-side error  | [optional] 
**Message** | Pointer to **string** | Human-readable error description | [optional] 
**Param** | Pointer to **string** | The parameter that caused the error (if applicable) | [optional] 
**Type** | Pointer to **string** | Error category | [optional] 
**Details** | Pointer to [**[]ErrorResponseErrorDetailsInner**](ErrorResponseErrorDetailsInner.md) | Additional error details for validation errors | [optional] 

## Methods

### NewErrorResponseError

`func NewErrorResponseError() *ErrorResponseError`

NewErrorResponseError instantiates a new ErrorResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseErrorWithDefaults

`func NewErrorResponseErrorWithDefaults() *ErrorResponseError`

NewErrorResponseErrorWithDefaults instantiates a new ErrorResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParam

`func (o *ErrorResponseError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ErrorResponseError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ErrorResponseError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *ErrorResponseError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetType

`func (o *ErrorResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResponseError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorResponseError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorResponseError) GetDetails() []ErrorResponseErrorDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorResponseError) GetDetailsOk() (*[]ErrorResponseErrorDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorResponseError) SetDetails(v []ErrorResponseErrorDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorResponseError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


