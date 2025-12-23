# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | Pointer to **string** | Unique ID for the message. | [optional] 
**AccountID** | Pointer to **int32** | Account ID associated with the message. | [optional] 
**SubAccountID** | Pointer to **int32** | Sub-account ID associated with the message. | [optional] 
**IpID** | Pointer to **int32** | IP ID used for sending the message. | [optional] 
**AccountIPPoolID** | Pointer to **int32** | Account IP Pool ID associated with the message. | [optional] 
**PublicIP** | Pointer to **string** | Public IP address used for sending the message. | [optional] 
**LocalIP** | Pointer to **string** | Local IP address used for sending the message. | [optional] 
**EmailType** | Pointer to **string** | Type of email service used. | [optional] 
**SubmittedAt** | Pointer to **int32** | UNIX epoch nano timestamp when message was submitted. | [optional] 
**From** | Pointer to [**Person**](Person.md) | Object comprising name and email address of the sender | [optional] 
**ReplyTo** | Pointer to [**Person**](Person.md) | Object comprising name and email addresses to which email replies will go to | [optional] 
**To** | Pointer to [**MessageTo**](MessageTo.md) |  | [optional] 
**HeaderTo** | Pointer to [**MessageHeaderTo**](MessageHeaderTo.md) |  | [optional] 
**HeaderCc** | Pointer to **[]string** | List of CC recipients from email headers | [optional] 
**HeaderBcc** | Pointer to **[]string** | List of BCC recipients from email headers | [optional] 
**Attachments** | Pointer to **[]string** | List of attachments | [optional] 
**Groups** | Pointer to **[]string** | List of groups associated with the message | [optional] 
**IpPool** | Pointer to **string** | IP Pool from which emails will go out. Relevant only for customers on dedicated IP plans. | [optional] 
**Headers** | Pointer to **map[string]string** | Key-Value pair which are added to every email message being sent and also with webhooks triggered on events such as email delivered, open, click etc. They are useful to identify email, recipient etc. in your internal system | [optional] 
**CustomFields** | Pointer to **map[string]string** | Key-Value pair of custom fields at message level | [optional] 
**Subject** | Pointer to **string** | Email subject line. | [optional] 
**PreText** | Pointer to **string** | Text which appears on mobile right after email subject line. | [optional] 
**HtmlBody** | Pointer to **string** | HTML email content. | [optional] 
**TextBody** | Pointer to **string** | Text email content. | [optional] 
**AmpBody** | Pointer to **string** | AMP email content. | [optional] 
**TrackOpens** | Pointer to **bool** | Indicates if email opens need to be tracked. | [optional] 
**TrackClicks** | Pointer to **bool** | Indicates if email clicks need to be tracked. | [optional] 
**Attempt** | Pointer to **int32** | Number of delivery attempts made for the message. | [optional] 
**WebhookEndpoint** | Pointer to **string** | Webhook endpoint URL for the message. | [optional] 
**MxRecords** | Pointer to **[]string** | List of MX records for the recipient domain | [optional] 

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *Message) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *Message) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *Message) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *Message) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetAccountID

`func (o *Message) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Message) GetAccountIDOk() (*int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Message) SetAccountID(v int32)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *Message) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetSubAccountID

`func (o *Message) GetSubAccountID() int32`

GetSubAccountID returns the SubAccountID field if non-nil, zero value otherwise.

### GetSubAccountIDOk

`func (o *Message) GetSubAccountIDOk() (*int32, bool)`

GetSubAccountIDOk returns a tuple with the SubAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountID

`func (o *Message) SetSubAccountID(v int32)`

SetSubAccountID sets SubAccountID field to given value.

### HasSubAccountID

`func (o *Message) HasSubAccountID() bool`

HasSubAccountID returns a boolean if a field has been set.

### GetIpID

`func (o *Message) GetIpID() int32`

GetIpID returns the IpID field if non-nil, zero value otherwise.

### GetIpIDOk

`func (o *Message) GetIpIDOk() (*int32, bool)`

GetIpIDOk returns a tuple with the IpID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpID

`func (o *Message) SetIpID(v int32)`

SetIpID sets IpID field to given value.

### HasIpID

`func (o *Message) HasIpID() bool`

HasIpID returns a boolean if a field has been set.

### GetAccountIPPoolID

`func (o *Message) GetAccountIPPoolID() int32`

GetAccountIPPoolID returns the AccountIPPoolID field if non-nil, zero value otherwise.

### GetAccountIPPoolIDOk

`func (o *Message) GetAccountIPPoolIDOk() (*int32, bool)`

GetAccountIPPoolIDOk returns a tuple with the AccountIPPoolID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIPPoolID

`func (o *Message) SetAccountIPPoolID(v int32)`

SetAccountIPPoolID sets AccountIPPoolID field to given value.

### HasAccountIPPoolID

`func (o *Message) HasAccountIPPoolID() bool`

HasAccountIPPoolID returns a boolean if a field has been set.

### GetPublicIP

`func (o *Message) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *Message) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *Message) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.

### HasPublicIP

`func (o *Message) HasPublicIP() bool`

HasPublicIP returns a boolean if a field has been set.

### GetLocalIP

`func (o *Message) GetLocalIP() string`

GetLocalIP returns the LocalIP field if non-nil, zero value otherwise.

### GetLocalIPOk

`func (o *Message) GetLocalIPOk() (*string, bool)`

GetLocalIPOk returns a tuple with the LocalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIP

`func (o *Message) SetLocalIP(v string)`

SetLocalIP sets LocalIP field to given value.

### HasLocalIP

`func (o *Message) HasLocalIP() bool`

HasLocalIP returns a boolean if a field has been set.

### GetEmailType

`func (o *Message) GetEmailType() string`

GetEmailType returns the EmailType field if non-nil, zero value otherwise.

### GetEmailTypeOk

`func (o *Message) GetEmailTypeOk() (*string, bool)`

GetEmailTypeOk returns a tuple with the EmailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailType

`func (o *Message) SetEmailType(v string)`

SetEmailType sets EmailType field to given value.

### HasEmailType

`func (o *Message) HasEmailType() bool`

HasEmailType returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *Message) GetSubmittedAt() int32`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *Message) GetSubmittedAtOk() (*int32, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *Message) SetSubmittedAt(v int32)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *Message) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetFrom

`func (o *Message) GetFrom() Person`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Message) GetFromOk() (*Person, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Message) SetFrom(v Person)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Message) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *Message) GetReplyTo() Person`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *Message) GetReplyToOk() (*Person, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *Message) SetReplyTo(v Person)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *Message) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTo

`func (o *Message) GetTo() MessageTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Message) GetToOk() (*MessageTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Message) SetTo(v MessageTo)`

SetTo sets To field to given value.

### HasTo

`func (o *Message) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetHeaderTo

`func (o *Message) GetHeaderTo() MessageHeaderTo`

GetHeaderTo returns the HeaderTo field if non-nil, zero value otherwise.

### GetHeaderToOk

`func (o *Message) GetHeaderToOk() (*MessageHeaderTo, bool)`

GetHeaderToOk returns a tuple with the HeaderTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTo

`func (o *Message) SetHeaderTo(v MessageHeaderTo)`

SetHeaderTo sets HeaderTo field to given value.

### HasHeaderTo

`func (o *Message) HasHeaderTo() bool`

HasHeaderTo returns a boolean if a field has been set.

### GetHeaderCc

`func (o *Message) GetHeaderCc() []string`

GetHeaderCc returns the HeaderCc field if non-nil, zero value otherwise.

### GetHeaderCcOk

`func (o *Message) GetHeaderCcOk() (*[]string, bool)`

GetHeaderCcOk returns a tuple with the HeaderCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderCc

`func (o *Message) SetHeaderCc(v []string)`

SetHeaderCc sets HeaderCc field to given value.

### HasHeaderCc

`func (o *Message) HasHeaderCc() bool`

HasHeaderCc returns a boolean if a field has been set.

### GetHeaderBcc

`func (o *Message) GetHeaderBcc() []string`

GetHeaderBcc returns the HeaderBcc field if non-nil, zero value otherwise.

### GetHeaderBccOk

`func (o *Message) GetHeaderBccOk() (*[]string, bool)`

GetHeaderBccOk returns a tuple with the HeaderBcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBcc

`func (o *Message) SetHeaderBcc(v []string)`

SetHeaderBcc sets HeaderBcc field to given value.

### HasHeaderBcc

`func (o *Message) HasHeaderBcc() bool`

HasHeaderBcc returns a boolean if a field has been set.

### GetAttachments

`func (o *Message) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Message) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Message) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Message) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetGroups

`func (o *Message) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Message) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Message) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Message) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIpPool

`func (o *Message) GetIpPool() string`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *Message) GetIpPoolOk() (*string, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *Message) SetIpPool(v string)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *Message) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetHeaders

`func (o *Message) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Message) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Message) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Message) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetCustomFields

`func (o *Message) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Message) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Message) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Message) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetSubject

`func (o *Message) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Message) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Message) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Message) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetPreText

`func (o *Message) GetPreText() string`

GetPreText returns the PreText field if non-nil, zero value otherwise.

### GetPreTextOk

`func (o *Message) GetPreTextOk() (*string, bool)`

GetPreTextOk returns a tuple with the PreText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreText

`func (o *Message) SetPreText(v string)`

SetPreText sets PreText field to given value.

### HasPreText

`func (o *Message) HasPreText() bool`

HasPreText returns a boolean if a field has been set.

### GetHtmlBody

`func (o *Message) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *Message) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *Message) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *Message) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### GetTextBody

`func (o *Message) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *Message) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *Message) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *Message) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetAmpBody

`func (o *Message) GetAmpBody() string`

GetAmpBody returns the AmpBody field if non-nil, zero value otherwise.

### GetAmpBodyOk

`func (o *Message) GetAmpBodyOk() (*string, bool)`

GetAmpBodyOk returns a tuple with the AmpBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmpBody

`func (o *Message) SetAmpBody(v string)`

SetAmpBody sets AmpBody field to given value.

### HasAmpBody

`func (o *Message) HasAmpBody() bool`

HasAmpBody returns a boolean if a field has been set.

### GetTrackOpens

`func (o *Message) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *Message) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *Message) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *Message) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *Message) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *Message) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *Message) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *Message) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetAttempt

`func (o *Message) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *Message) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *Message) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *Message) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetWebhookEndpoint

`func (o *Message) GetWebhookEndpoint() string`

GetWebhookEndpoint returns the WebhookEndpoint field if non-nil, zero value otherwise.

### GetWebhookEndpointOk

`func (o *Message) GetWebhookEndpointOk() (*string, bool)`

GetWebhookEndpointOk returns a tuple with the WebhookEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpoint

`func (o *Message) SetWebhookEndpoint(v string)`

SetWebhookEndpoint sets WebhookEndpoint field to given value.

### HasWebhookEndpoint

`func (o *Message) HasWebhookEndpoint() bool`

HasWebhookEndpoint returns a boolean if a field has been set.

### GetMxRecords

`func (o *Message) GetMxRecords() []string`

GetMxRecords returns the MxRecords field if non-nil, zero value otherwise.

### GetMxRecordsOk

`func (o *Message) GetMxRecordsOk() (*[]string, bool)`

GetMxRecordsOk returns a tuple with the MxRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxRecords

`func (o *Message) SetMxRecords(v []string)`

SetMxRecords sets MxRecords field to given value.

### HasMxRecords

`func (o *Message) HasMxRecords() bool`

HasMxRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


