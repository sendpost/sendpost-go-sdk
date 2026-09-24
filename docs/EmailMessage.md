# EmailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | Pointer to **string** | Unique identifier for this email message | [optional] 
**AccountID** | Pointer to **int64** | ID of the SendPost account that sent this email | [optional] 
**SubAccountID** | Pointer to **int64** | ID of the sub-account that sent this email | [optional] 
**IpID** | Pointer to **int64** | ID of the dedicated IP used for sending (0 if a shared IP was used) | [optional] 
**PublicIP** | Pointer to **string** | The public IP address used to send this email | [optional] 
**LocalIP** | Pointer to **string** | The internal/local IP address used to send this email | [optional] 
**EmailType** | Pointer to **string** | Classification of the email based on recipient domain: - &#x60;gmail&#x60; - Gmail recipient - &#x60;yahoo&#x60; - Yahoo recipient - &#x60;microsoft&#x60; - Outlook/Hotmail recipient - &#x60;default&#x60; - Other email providers  | [optional] 
**SubmittedAt** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the email was submitted | [optional] 
**From** | Pointer to [**EmailAddress**](EmailAddress.md) | The sender&#39;s email address | [optional] 
**ReplyTo** | Pointer to [**EmailAddress**](EmailAddress.md) | The reply-to address (if different from sender) | [optional] 
**To** | Pointer to [**Recipient**](Recipient.md) | The envelope recipient (actual delivery address) | [optional] 
**Groups** | Pointer to **[]string** | Tags/groups for categorization and analytics | [optional] 
**IpPool** | Pointer to **string** | Name of the IP pool used for sending | [optional] 
**Headers** | Pointer to **map[string]string** | Custom headers included in the email | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom fields sent with the email, available for personalization | [optional] 
**TrackOpens** | Pointer to **bool** | Whether open tracking was enabled | [optional] 
**TrackClicks** | Pointer to **bool** | Whether click tracking was enabled | [optional] 
**WebhookEndpoint** | Pointer to **string** | Custom webhook endpoint for this email (if specified) | [optional] 

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

`func (o *EmailMessage) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *EmailMessage) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *EmailMessage) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *EmailMessage) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetSubAccountID

`func (o *EmailMessage) GetSubAccountID() int64`

GetSubAccountID returns the SubAccountID field if non-nil, zero value otherwise.

### GetSubAccountIDOk

`func (o *EmailMessage) GetSubAccountIDOk() (*int64, bool)`

GetSubAccountIDOk returns a tuple with the SubAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountID

`func (o *EmailMessage) SetSubAccountID(v int64)`

SetSubAccountID sets SubAccountID field to given value.

### HasSubAccountID

`func (o *EmailMessage) HasSubAccountID() bool`

HasSubAccountID returns a boolean if a field has been set.

### GetIpID

`func (o *EmailMessage) GetIpID() int64`

GetIpID returns the IpID field if non-nil, zero value otherwise.

### GetIpIDOk

`func (o *EmailMessage) GetIpIDOk() (*int64, bool)`

GetIpIDOk returns a tuple with the IpID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpID

`func (o *EmailMessage) SetIpID(v int64)`

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

`func (o *EmailMessage) GetSubmittedAt() int64`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *EmailMessage) GetSubmittedAtOk() (*int64, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *EmailMessage) SetSubmittedAt(v int64)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *EmailMessage) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetFrom

`func (o *EmailMessage) GetFrom() EmailAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailMessage) GetFromOk() (*EmailAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailMessage) SetFrom(v EmailAddress)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailMessage) GetReplyTo() EmailAddress`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailMessage) GetReplyToOk() (*EmailAddress, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailMessage) SetReplyTo(v EmailAddress)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTo

`func (o *EmailMessage) GetTo() Recipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailMessage) GetToOk() (*Recipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailMessage) SetTo(v Recipient)`

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

`func (o *EmailMessage) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EmailMessage) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EmailMessage) SetCustomFields(v map[string]interface{})`

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


