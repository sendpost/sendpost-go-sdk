# EmailMessageWithTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**EmailAddress**](EmailAddress.md) |  | [optional] 
**ReplyTo** | Pointer to [**EmailAddress**](EmailAddress.md) |  | [optional] 
**To** | Pointer to [**[]Recipient**](Recipient.md) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**PreText** | Pointer to **string** |  | [optional] 
**HtmlBody** | Pointer to **string** |  | [optional] 
**TextBody** | Pointer to **string** |  | [optional] 
**AmpBody** | Pointer to **string** |  | [optional] 
**Ippool** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**TrackOpens** | Pointer to **bool** |  | [optional] 
**TrackClicks** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**WebhookEndpoint** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** | Template ID for the email template | [optional] 
**TemplateVariables** | Pointer to **map[string]string** | Key-Value pair of template variables | [optional] 

## Methods

### NewEmailMessageWithTemplate

`func NewEmailMessageWithTemplate() *EmailMessageWithTemplate`

NewEmailMessageWithTemplate instantiates a new EmailMessageWithTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageWithTemplateWithDefaults

`func NewEmailMessageWithTemplateWithDefaults() *EmailMessageWithTemplate`

NewEmailMessageWithTemplateWithDefaults instantiates a new EmailMessageWithTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *EmailMessageWithTemplate) GetFrom() EmailAddress`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailMessageWithTemplate) GetFromOk() (*EmailAddress, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailMessageWithTemplate) SetFrom(v EmailAddress)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailMessageWithTemplate) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailMessageWithTemplate) GetReplyTo() EmailAddress`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailMessageWithTemplate) GetReplyToOk() (*EmailAddress, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailMessageWithTemplate) SetReplyTo(v EmailAddress)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailMessageWithTemplate) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTo

`func (o *EmailMessageWithTemplate) GetTo() []Recipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailMessageWithTemplate) GetToOk() (*[]Recipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailMessageWithTemplate) SetTo(v []Recipient)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailMessageWithTemplate) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSubject

`func (o *EmailMessageWithTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailMessageWithTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailMessageWithTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailMessageWithTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetPreText

`func (o *EmailMessageWithTemplate) GetPreText() string`

GetPreText returns the PreText field if non-nil, zero value otherwise.

### GetPreTextOk

`func (o *EmailMessageWithTemplate) GetPreTextOk() (*string, bool)`

GetPreTextOk returns a tuple with the PreText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreText

`func (o *EmailMessageWithTemplate) SetPreText(v string)`

SetPreText sets PreText field to given value.

### HasPreText

`func (o *EmailMessageWithTemplate) HasPreText() bool`

HasPreText returns a boolean if a field has been set.

### GetHtmlBody

`func (o *EmailMessageWithTemplate) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *EmailMessageWithTemplate) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *EmailMessageWithTemplate) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *EmailMessageWithTemplate) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### GetTextBody

`func (o *EmailMessageWithTemplate) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *EmailMessageWithTemplate) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *EmailMessageWithTemplate) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *EmailMessageWithTemplate) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetAmpBody

`func (o *EmailMessageWithTemplate) GetAmpBody() string`

GetAmpBody returns the AmpBody field if non-nil, zero value otherwise.

### GetAmpBodyOk

`func (o *EmailMessageWithTemplate) GetAmpBodyOk() (*string, bool)`

GetAmpBodyOk returns a tuple with the AmpBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmpBody

`func (o *EmailMessageWithTemplate) SetAmpBody(v string)`

SetAmpBody sets AmpBody field to given value.

### HasAmpBody

`func (o *EmailMessageWithTemplate) HasAmpBody() bool`

HasAmpBody returns a boolean if a field has been set.

### GetIppool

`func (o *EmailMessageWithTemplate) GetIppool() string`

GetIppool returns the Ippool field if non-nil, zero value otherwise.

### GetIppoolOk

`func (o *EmailMessageWithTemplate) GetIppoolOk() (*string, bool)`

GetIppoolOk returns a tuple with the Ippool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIppool

`func (o *EmailMessageWithTemplate) SetIppool(v string)`

SetIppool sets Ippool field to given value.

### HasIppool

`func (o *EmailMessageWithTemplate) HasIppool() bool`

HasIppool returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailMessageWithTemplate) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailMessageWithTemplate) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailMessageWithTemplate) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailMessageWithTemplate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTrackOpens

`func (o *EmailMessageWithTemplate) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *EmailMessageWithTemplate) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *EmailMessageWithTemplate) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *EmailMessageWithTemplate) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *EmailMessageWithTemplate) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *EmailMessageWithTemplate) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *EmailMessageWithTemplate) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *EmailMessageWithTemplate) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetGroups

`func (o *EmailMessageWithTemplate) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmailMessageWithTemplate) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmailMessageWithTemplate) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EmailMessageWithTemplate) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAttachments

`func (o *EmailMessageWithTemplate) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EmailMessageWithTemplate) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EmailMessageWithTemplate) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EmailMessageWithTemplate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetWebhookEndpoint

`func (o *EmailMessageWithTemplate) GetWebhookEndpoint() string`

GetWebhookEndpoint returns the WebhookEndpoint field if non-nil, zero value otherwise.

### GetWebhookEndpointOk

`func (o *EmailMessageWithTemplate) GetWebhookEndpointOk() (*string, bool)`

GetWebhookEndpointOk returns a tuple with the WebhookEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpoint

`func (o *EmailMessageWithTemplate) SetWebhookEndpoint(v string)`

SetWebhookEndpoint sets WebhookEndpoint field to given value.

### HasWebhookEndpoint

`func (o *EmailMessageWithTemplate) HasWebhookEndpoint() bool`

HasWebhookEndpoint returns a boolean if a field has been set.

### GetTemplate

`func (o *EmailMessageWithTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailMessageWithTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailMessageWithTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailMessageWithTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTemplateId

`func (o *EmailMessageWithTemplate) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *EmailMessageWithTemplate) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *EmailMessageWithTemplate) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *EmailMessageWithTemplate) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVariables

`func (o *EmailMessageWithTemplate) GetTemplateVariables() map[string]string`

GetTemplateVariables returns the TemplateVariables field if non-nil, zero value otherwise.

### GetTemplateVariablesOk

`func (o *EmailMessageWithTemplate) GetTemplateVariablesOk() (*map[string]string, bool)`

GetTemplateVariablesOk returns a tuple with the TemplateVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVariables

`func (o *EmailMessageWithTemplate) SetTemplateVariables(v map[string]string)`

SetTemplateVariables sets TemplateVariables field to given value.

### HasTemplateVariables

`func (o *EmailMessageWithTemplate) HasTemplateVariables() bool`

HasTemplateVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


