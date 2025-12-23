# EmailMessageObject

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

## Methods

### NewEmailMessageObject

`func NewEmailMessageObject() *EmailMessageObject`

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

### HasFrom

`func (o *EmailMessageObject) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

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

### HasTo

`func (o *EmailMessageObject) HasTo() bool`

HasTo returns a boolean if a field has been set.

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


