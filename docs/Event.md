# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Unique identifier for this specific event. Use this for idempotency - the same event may be delivered multiple times.  | [optional] 
**MessageId** | Pointer to **string** | Unique identifier of the email message this event belongs to. Use this to correlate events with the original send request.  | [optional] 
**Type** | Pointer to **int64** | Numeric event type code: - &#x60;0&#x60; &#x3D; processed (email accepted by API) - &#x60;1&#x60; &#x3D; dropped (not sent - suppression, invalid, etc.) - &#x60;2&#x60; &#x3D; delivered (accepted by recipient&#39;s mail server) - &#x60;3&#x60; &#x3D; softBounced (temporary failure, will retry) - &#x60;4&#x60; &#x3D; hardBounced (permanent failure) - &#x60;5&#x60; &#x3D; opened (tracking pixel loaded) - &#x60;6&#x60; &#x3D; clicked (link clicked) - &#x60;7&#x60; &#x3D; unsubscribed (clicked unsubscribe link) - &#x60;8&#x60; &#x3D; spam (marked as spam by recipient) - &#x60;9&#x60; &#x3D; sent (sent to mail server) - &#x60;10&#x60; &#x3D; smtpDropped (dropped at SMTP level)  | [optional] 
**TypeName** | Pointer to **string** | Human-readable event type name | [optional] 
**From** | Pointer to **string** | Sender email address | [optional] 
**To** | Pointer to **string** | Recipient email address | [optional] 
**Subject** | Pointer to **string** | Email subject line (useful for identifying the email) | [optional] 
**Groups** | Pointer to **[]string** | Tags/groups that were associated with the email | [optional] 
**SubmittedAt** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the email was originally submitted | [optional] 
**Timestamp** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when this event occurred | [optional] 
**EventMetadata** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *Event) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Event) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Event) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *Event) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetMessageId

`func (o *Event) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Event) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Event) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Event) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeName

`func (o *Event) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *Event) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *Event) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *Event) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetFrom

`func (o *Event) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Event) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Event) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Event) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Event) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Event) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Event) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Event) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSubject

`func (o *Event) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Event) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Event) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Event) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetGroups

`func (o *Event) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Event) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Event) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Event) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *Event) GetSubmittedAt() int64`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *Event) GetSubmittedAtOk() (*int64, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *Event) SetSubmittedAt(v int64)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *Event) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTimestamp

`func (o *Event) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Event) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Event) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Event) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEventMetadata

`func (o *Event) GetEventMetadata() EventMetadata`

GetEventMetadata returns the EventMetadata field if non-nil, zero value otherwise.

### GetEventMetadataOk

`func (o *Event) GetEventMetadataOk() (*EventMetadata, bool)`

GetEventMetadataOk returns a tuple with the EventMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMetadata

`func (o *Event) SetEventMetadata(v EventMetadata)`

SetEventMetadata sets EventMetadata field to given value.

### HasEventMetadata

`func (o *Event) HasEventMetadata() bool`

HasEventMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


