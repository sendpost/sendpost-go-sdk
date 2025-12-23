# WebhookObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**Event**](Event.md) |  | [optional] 
**EmailMessage** | Pointer to [**EmailMessage**](EmailMessage.md) |  | [optional] 

## Methods

### NewWebhookObject

`func NewWebhookObject() *WebhookObject`

NewWebhookObject instantiates a new WebhookObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookObjectWithDefaults

`func NewWebhookObjectWithDefaults() *WebhookObject`

NewWebhookObjectWithDefaults instantiates a new WebhookObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookObject) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookObject) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookObject) SetEvent(v Event)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebhookObject) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEmailMessage

`func (o *WebhookObject) GetEmailMessage() EmailMessage`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *WebhookObject) GetEmailMessageOk() (*EmailMessage, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *WebhookObject) SetEmailMessage(v EmailMessage)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *WebhookObject) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


