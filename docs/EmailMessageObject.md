# EmailMessageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**EmailAddress**](EmailAddress.md) | The sender&#39;s email address and optional display name | 
**ReplyTo** | Pointer to [**EmailAddress**](EmailAddress.md) | The reply-to email address. If not specified, replies will go to the &#x60;from&#x60; address | [optional] 
**To** | [**[]Recipient**](Recipient.md) | List of recipients. Each recipient can have their own CC, BCC, and custom fields for personalization. Maximum 1000 recipients per API call.  | 
**Subject** | Pointer to **string** | Email subject line. Supports Handlebars templating for personalization. Example: \&quot;Hello, {{firstName}}! Your order is ready\&quot;  | [optional] 
**PreText** | Pointer to **string** | Preview text (preheader) shown in email clients before opening the email. This text appears after the subject line in most email clients&#39; inbox view.  | [optional] 
**HtmlBody** | Pointer to **string** | HTML content of the email. Supports Handlebars templating for personalization. Use {{customFieldName}} to insert recipient-specific values.  | [optional] 
**TextBody** | Pointer to **string** | Plain text content of the email. Used as fallback when HTML cannot be rendered. Also improves deliverability as some spam filters prefer multipart emails.  | [optional] 
**AmpBody** | Pointer to **string** | AMP HTML content for supported email clients (Gmail, Yahoo). Enables interactive email experiences like carousels, forms, and real-time content. See https://amp.dev/about/email/ for more details.  | [optional] 
**Template** | Pointer to **string** | Name of a pre-defined template to use for this email. When specified, the template&#39;s subject, htmlBody, and textBody will be used unless explicitly overridden in this request.  | [optional] 
**Ippool** | Pointer to **string** | Name of the IP pool to use for sending this email. If not specified, the default IP pool for the sub-account will be used.  | [optional] 
**Headers** | Pointer to **map[string]string** | Custom email headers to include in the message. Common uses: adding List-Unsubscribe headers, custom tracking IDs, or priority flags. Note: Some headers like From, To, Subject are set automatically and cannot be overridden.  | [optional] 
**TrackOpens** | Pointer to **bool** | Whether to track email opens using a tracking pixel. When enabled, a 1x1 transparent image is inserted into the HTML body. Default: true (if not specified)  | [optional] [default to true]
**TrackClicks** | Pointer to **bool** | Whether to track link clicks by rewriting URLs through SendPost&#39;s tracking domain. When enabled, all links in htmlBody are replaced with tracking URLs. Default: true (if not specified)  | [optional] [default to true]
**Groups** | Pointer to **[]string** | Tags/groups to categorize this email for analytics and reporting. Use groups to segment your email statistics (e.g., by campaign, email type, or customer segment).  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | File attachments to include with the email. Maximum total attachment size: 25MB. Supported formats: PDF, images, documents, etc.  | [optional] 
**WebhookEndpoint** | Pointer to **string** | Custom webhook URL to receive events for this specific email. Overrides the default webhook configured at the account level. Useful for per-email or per-customer webhook routing.  | [optional] 

## Methods

### NewEmailMessageObject

`func NewEmailMessageObject(from EmailAddress, to []Recipient, ) *EmailMessageObject`

NewEmailMessageObject instantiates a new EmailMessageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageObjectWithDefaults

`func NewEmailMessageObjectWithDefaults() *EmailMessageObject`

NewEmailMessageObjectWithDefaults instantiates a new EmailMessageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *EmailMessageObject) GetFrom() EmailAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailMessageObject) GetFromOk() (*EmailAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailMessageObject) SetFrom(v EmailAddress)`

SetFrom sets From field to given value.


### GetReplyTo

`func (o *EmailMessageObject) GetReplyTo() EmailAddress`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailMessageObject) GetReplyToOk() (*EmailAddress, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailMessageObject) SetReplyTo(v EmailAddress)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailMessageObject) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTo

`func (o *EmailMessageObject) GetTo() []Recipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailMessageObject) GetToOk() (*[]Recipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailMessageObject) SetTo(v []Recipient)`

SetTo sets To field to given value.


### GetSubject

`func (o *EmailMessageObject) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailMessageObject) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailMessageObject) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailMessageObject) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetPreText

`func (o *EmailMessageObject) GetPreText() string`

GetPreText returns the PreText field if non-nil, zero value otherwise.

### GetPreTextOk

`func (o *EmailMessageObject) GetPreTextOk() (*string, bool)`

GetPreTextOk returns a tuple with the PreText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreText

`func (o *EmailMessageObject) SetPreText(v string)`

SetPreText sets PreText field to given value.

### HasPreText

`func (o *EmailMessageObject) HasPreText() bool`

HasPreText returns a boolean if a field has been set.

### GetHtmlBody

`func (o *EmailMessageObject) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *EmailMessageObject) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *EmailMessageObject) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *EmailMessageObject) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### GetTextBody

`func (o *EmailMessageObject) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *EmailMessageObject) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *EmailMessageObject) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *EmailMessageObject) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetAmpBody

`func (o *EmailMessageObject) GetAmpBody() string`

GetAmpBody returns the AmpBody field if non-nil, zero value otherwise.

### GetAmpBodyOk

`func (o *EmailMessageObject) GetAmpBodyOk() (*string, bool)`

GetAmpBodyOk returns a tuple with the AmpBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmpBody

`func (o *EmailMessageObject) SetAmpBody(v string)`

SetAmpBody sets AmpBody field to given value.

### HasAmpBody

`func (o *EmailMessageObject) HasAmpBody() bool`

HasAmpBody returns a boolean if a field has been set.

### GetTemplate

`func (o *EmailMessageObject) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailMessageObject) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailMessageObject) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailMessageObject) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetIppool

`func (o *EmailMessageObject) GetIppool() string`

GetIppool returns the Ippool field if non-nil, zero value otherwise.

### GetIppoolOk

`func (o *EmailMessageObject) GetIppoolOk() (*string, bool)`

GetIppoolOk returns a tuple with the Ippool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIppool

`func (o *EmailMessageObject) SetIppool(v string)`

SetIppool sets Ippool field to given value.

### HasIppool

`func (o *EmailMessageObject) HasIppool() bool`

HasIppool returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailMessageObject) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailMessageObject) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailMessageObject) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailMessageObject) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTrackOpens

`func (o *EmailMessageObject) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *EmailMessageObject) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *EmailMessageObject) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *EmailMessageObject) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *EmailMessageObject) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *EmailMessageObject) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *EmailMessageObject) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *EmailMessageObject) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetGroups

`func (o *EmailMessageObject) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmailMessageObject) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmailMessageObject) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EmailMessageObject) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAttachments

`func (o *EmailMessageObject) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EmailMessageObject) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EmailMessageObject) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EmailMessageObject) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetWebhookEndpoint

`func (o *EmailMessageObject) GetWebhookEndpoint() string`

GetWebhookEndpoint returns the WebhookEndpoint field if non-nil, zero value otherwise.

### GetWebhookEndpointOk

`func (o *EmailMessageObject) GetWebhookEndpointOk() (*string, bool)`

GetWebhookEndpointOk returns a tuple with the WebhookEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpoint

`func (o *EmailMessageObject) SetWebhookEndpoint(v string)`

SetWebhookEndpoint sets WebhookEndpoint field to given value.

### HasWebhookEndpoint

`func (o *EmailMessageObject) HasWebhookEndpoint() bool`

HasWebhookEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


