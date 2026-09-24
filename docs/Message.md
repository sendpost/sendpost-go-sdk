# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Unique identifier (UUID) for this email message | [optional] 
**SubAccountId** | Pointer to **int64** | ID of the sub-account that sent this email | [optional] 
**PublicIp** | Pointer to **string** | The public IP address used to send this email | [optional] 
**EmailType** | Pointer to **string** | Classification of the email, e.g. \&quot;transactional\&quot; or \&quot;marketing\&quot;.  | [optional] 
**SubmittedAt** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the email was submitted | [optional] 
**From** | Pointer to [**EmailAddress**](EmailAddress.md) | The sender&#39;s email address and display name | [optional] 
**ReplyTo** | Pointer to [**EmailAddress**](EmailAddress.md) | The Reply-To email address and display name | [optional] 
**To** | Pointer to [**Recipient**](Recipient.md) | The primary recipient, including any per-recipient CC/BCC and custom fields | [optional] 
**HeaderTo** | Pointer to [**Recipient**](Recipient.md) | The address rendered in the visible To header (may differ from the envelope recipient) | [optional] 
**HeaderCc** | Pointer to [**[]CopyTo**](CopyTo.md) | Addresses rendered in the visible Cc header | [optional] 
**HeaderBcc** | Pointer to [**[]CopyTo**](CopyTo.md) | Addresses rendered in the visible Bcc header | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | File attachments included with the email | [optional] 
**Groups** | Pointer to **[]string** | Tags/groups associated with this email | [optional] 
**IpPool** | Pointer to **string** | Name of the IP pool used for sending | [optional] 
**Headers** | Pointer to **map[string]string** | Custom SMTP headers set on the message | [optional] 
**Subject** | Pointer to **string** | The email subject line | [optional] 
**PreText** | Pointer to **string** | Preheader/preview text shown by many email clients after the subject | [optional] 
**HtmlBody** | Pointer to **string** | The HTML body of the email | [optional] 
**TextBody** | Pointer to **string** | The plain-text body of the email | [optional] 
**AmpBody** | Pointer to **string** | The AMP for Email body, if provided | [optional] 
**TrackOpens** | Pointer to **bool** | Whether open tracking was enabled for this email | [optional] 
**TrackClicks** | Pointer to **bool** | Whether click tracking was enabled for this email | [optional] 

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

### GetMessageId

`func (o *Message) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Message) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Message) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Message) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetSubAccountId

`func (o *Message) GetSubAccountId() int64`

GetSubAccountId returns the SubAccountId field if non-nil, zero value otherwise.

### GetSubAccountIdOk

`func (o *Message) GetSubAccountIdOk() (*int64, bool)`

GetSubAccountIdOk returns a tuple with the SubAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountId

`func (o *Message) SetSubAccountId(v int64)`

SetSubAccountId sets SubAccountId field to given value.

### HasSubAccountId

`func (o *Message) HasSubAccountId() bool`

HasSubAccountId returns a boolean if a field has been set.

### GetPublicIp

`func (o *Message) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *Message) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *Message) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *Message) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

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

`func (o *Message) GetSubmittedAt() int64`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *Message) GetSubmittedAtOk() (*int64, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *Message) SetSubmittedAt(v int64)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *Message) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetFrom

`func (o *Message) GetFrom() EmailAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Message) GetFromOk() (*EmailAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Message) SetFrom(v EmailAddress)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Message) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *Message) GetReplyTo() EmailAddress`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *Message) GetReplyToOk() (*EmailAddress, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *Message) SetReplyTo(v EmailAddress)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *Message) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTo

`func (o *Message) GetTo() Recipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Message) GetToOk() (*Recipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Message) SetTo(v Recipient)`

SetTo sets To field to given value.

### HasTo

`func (o *Message) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetHeaderTo

`func (o *Message) GetHeaderTo() Recipient`

GetHeaderTo returns the HeaderTo field if non-nil, zero value otherwise.

### GetHeaderToOk

`func (o *Message) GetHeaderToOk() (*Recipient, bool)`

GetHeaderToOk returns a tuple with the HeaderTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTo

`func (o *Message) SetHeaderTo(v Recipient)`

SetHeaderTo sets HeaderTo field to given value.

### HasHeaderTo

`func (o *Message) HasHeaderTo() bool`

HasHeaderTo returns a boolean if a field has been set.

### GetHeaderCc

`func (o *Message) GetHeaderCc() []CopyTo`

GetHeaderCc returns the HeaderCc field if non-nil, zero value otherwise.

### GetHeaderCcOk

`func (o *Message) GetHeaderCcOk() (*[]CopyTo, bool)`

GetHeaderCcOk returns a tuple with the HeaderCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderCc

`func (o *Message) SetHeaderCc(v []CopyTo)`

SetHeaderCc sets HeaderCc field to given value.

### HasHeaderCc

`func (o *Message) HasHeaderCc() bool`

HasHeaderCc returns a boolean if a field has been set.

### GetHeaderBcc

`func (o *Message) GetHeaderBcc() []CopyTo`

GetHeaderBcc returns the HeaderBcc field if non-nil, zero value otherwise.

### GetHeaderBccOk

`func (o *Message) GetHeaderBccOk() (*[]CopyTo, bool)`

GetHeaderBccOk returns a tuple with the HeaderBcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBcc

`func (o *Message) SetHeaderBcc(v []CopyTo)`

SetHeaderBcc sets HeaderBcc field to given value.

### HasHeaderBcc

`func (o *Message) HasHeaderBcc() bool`

HasHeaderBcc returns a boolean if a field has been set.

### GetAttachments

`func (o *Message) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Message) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Message) SetAttachments(v []Attachment)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


