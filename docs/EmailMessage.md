# EmailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | Pointer to **string** |  | [optional] 
**AccountID** | Pointer to **int32** |  | [optional] 
**SubAccountID** | Pointer to **int32** |  | [optional] 
**IpID** | Pointer to **int32** |  | [optional] 
**PublicIP** | Pointer to **string** |  | [optional] 
**LocalIP** | Pointer to **string** |  | [optional] 
**EmailType** | Pointer to **string** |  | [optional] 
**SubmittedAt** | Pointer to **int32** |  | [optional] 
**From** | Pointer to [**EmailMessageFrom**](EmailMessageFrom.md) |  | [optional] 
**ReplyTo** | Pointer to [**EmailMessageReplyTo**](EmailMessageReplyTo.md) |  | [optional] 
**To** | Pointer to [**[]EmailMessageToInner**](EmailMessageToInner.md) |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**IpPool** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]string** |  | [optional] 
**TrackOpens** | Pointer to **bool** |  | [optional] 
**TrackClicks** | Pointer to **bool** |  | [optional] 
**WebhookEndpoint** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailMessage

`func NewEmailMessage() *EmailMessage`

NewEmailMessage instantiates a new EmailMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageWithDefaults

`func NewEmailMessageWithDefaults() *EmailMessage`

NewEmailMessageWithDefaults instantiates a new EmailMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *EmailMessage) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *EmailMessage) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *EmailMessage) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *EmailMessage) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetAccountID

`func (o *EmailMessage) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *EmailMessage) GetAccountIDOk() (*int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *EmailMessage) SetAccountID(v int32)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *EmailMessage) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetSubAccountID

`func (o *EmailMessage) GetSubAccountID() int32`

GetSubAccountID returns the SubAccountID field if non-nil, zero value otherwise.

### GetSubAccountIDOk

`func (o *EmailMessage) GetSubAccountIDOk() (*int32, bool)`

GetSubAccountIDOk returns a tuple with the SubAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountID

`func (o *EmailMessage) SetSubAccountID(v int32)`

SetSubAccountID sets SubAccountID field to given value.

### HasSubAccountID

`func (o *EmailMessage) HasSubAccountID() bool`

HasSubAccountID returns a boolean if a field has been set.

### GetIpID

`func (o *EmailMessage) GetIpID() int32`

GetIpID returns the IpID field if non-nil, zero value otherwise.

### GetIpIDOk

`func (o *EmailMessage) GetIpIDOk() (*int32, bool)`

GetIpIDOk returns a tuple with the IpID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpID

`func (o *EmailMessage) SetIpID(v int32)`

SetIpID sets IpID field to given value.

### HasIpID

`func (o *EmailMessage) HasIpID() bool`

HasIpID returns a boolean if a field has been set.

### GetPublicIP

`func (o *EmailMessage) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *EmailMessage) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *EmailMessage) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.

### HasPublicIP

`func (o *EmailMessage) HasPublicIP() bool`

HasPublicIP returns a boolean if a field has been set.

### GetLocalIP

`func (o *EmailMessage) GetLocalIP() string`

GetLocalIP returns the LocalIP field if non-nil, zero value otherwise.

### GetLocalIPOk

`func (o *EmailMessage) GetLocalIPOk() (*string, bool)`

GetLocalIPOk returns a tuple with the LocalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIP

`func (o *EmailMessage) SetLocalIP(v string)`

SetLocalIP sets LocalIP field to given value.

### HasLocalIP

`func (o *EmailMessage) HasLocalIP() bool`

HasLocalIP returns a boolean if a field has been set.

### GetEmailType

`func (o *EmailMessage) GetEmailType() string`

GetEmailType returns the EmailType field if non-nil, zero value otherwise.

### GetEmailTypeOk

`func (o *EmailMessage) GetEmailTypeOk() (*string, bool)`

GetEmailTypeOk returns a tuple with the EmailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailType

`func (o *EmailMessage) SetEmailType(v string)`

SetEmailType sets EmailType field to given value.

### HasEmailType

`func (o *EmailMessage) HasEmailType() bool`

HasEmailType returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *EmailMessage) GetSubmittedAt() int32`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *EmailMessage) GetSubmittedAtOk() (*int32, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *EmailMessage) SetSubmittedAt(v int32)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *EmailMessage) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetFrom

`func (o *EmailMessage) GetFrom() EmailMessageFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailMessage) GetFromOk() (*EmailMessageFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailMessage) SetFrom(v EmailMessageFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailMessage) GetReplyTo() EmailMessageReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailMessage) GetReplyToOk() (*EmailMessageReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailMessage) SetReplyTo(v EmailMessageReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTo

`func (o *EmailMessage) GetTo() []EmailMessageToInner`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailMessage) GetToOk() (*[]EmailMessageToInner, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailMessage) SetTo(v []EmailMessageToInner)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGroups

`func (o *EmailMessage) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmailMessage) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmailMessage) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EmailMessage) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIpPool

`func (o *EmailMessage) GetIpPool() string`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *EmailMessage) GetIpPoolOk() (*string, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *EmailMessage) SetIpPool(v string)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *EmailMessage) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailMessage) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailMessage) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailMessage) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetCustomFields

`func (o *EmailMessage) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EmailMessage) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EmailMessage) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EmailMessage) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTrackOpens

`func (o *EmailMessage) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *EmailMessage) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *EmailMessage) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *EmailMessage) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *EmailMessage) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *EmailMessage) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *EmailMessage) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *EmailMessage) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetWebhookEndpoint

`func (o *EmailMessage) GetWebhookEndpoint() string`

GetWebhookEndpoint returns the WebhookEndpoint field if non-nil, zero value otherwise.

### GetWebhookEndpointOk

`func (o *EmailMessage) GetWebhookEndpointOk() (*string, bool)`

GetWebhookEndpointOk returns a tuple with the WebhookEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpoint

`func (o *EmailMessage) SetWebhookEndpoint(v string)`

SetWebhookEndpoint sets WebhookEndpoint field to given value.

### HasWebhookEndpoint

`func (o *EmailMessage) HasWebhookEndpoint() bool`

HasWebhookEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


