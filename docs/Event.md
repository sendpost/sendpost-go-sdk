# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**IpID** | Pointer to **int32** |  | [optional] 
**IpPoolID** | Pointer to **int32** |  | [optional] 
**DomainID** | Pointer to **int32** |  | [optional] 
**TpspId** | Pointer to **int32** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**MessageSubject** | Pointer to **string** |  | [optional] 
**AccountID** | Pointer to **int32** |  | [optional] 
**SubAccountID** | Pointer to **int32** |  | [optional] 
**MessageID** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**FromName** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**ToName** | Pointer to **string** |  | [optional] 
**SubmittedAt** | Pointer to **int32** |  | [optional] 
**SmtpCode** | Pointer to **int32** |  | [optional] 
**SmtpDescription** | Pointer to **string** |  | [optional] 
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

### GetEventID

`func (o *Event) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *Event) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *Event) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *Event) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

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

### GetIpID

`func (o *Event) GetIpID() int32`

GetIpID returns the IpID field if non-nil, zero value otherwise.

### GetIpIDOk

`func (o *Event) GetIpIDOk() (*int32, bool)`

GetIpIDOk returns a tuple with the IpID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpID

`func (o *Event) SetIpID(v int32)`

SetIpID sets IpID field to given value.

### HasIpID

`func (o *Event) HasIpID() bool`

HasIpID returns a boolean if a field has been set.

### GetIpPoolID

`func (o *Event) GetIpPoolID() int32`

GetIpPoolID returns the IpPoolID field if non-nil, zero value otherwise.

### GetIpPoolIDOk

`func (o *Event) GetIpPoolIDOk() (*int32, bool)`

GetIpPoolIDOk returns a tuple with the IpPoolID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolID

`func (o *Event) SetIpPoolID(v int32)`

SetIpPoolID sets IpPoolID field to given value.

### HasIpPoolID

`func (o *Event) HasIpPoolID() bool`

HasIpPoolID returns a boolean if a field has been set.

### GetDomainID

`func (o *Event) GetDomainID() int32`

GetDomainID returns the DomainID field if non-nil, zero value otherwise.

### GetDomainIDOk

`func (o *Event) GetDomainIDOk() (*int32, bool)`

GetDomainIDOk returns a tuple with the DomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainID

`func (o *Event) SetDomainID(v int32)`

SetDomainID sets DomainID field to given value.

### HasDomainID

`func (o *Event) HasDomainID() bool`

HasDomainID returns a boolean if a field has been set.

### GetTpspId

`func (o *Event) GetTpspId() int32`

GetTpspId returns the TpspId field if non-nil, zero value otherwise.

### GetTpspIdOk

`func (o *Event) GetTpspIdOk() (*int32, bool)`

GetTpspIdOk returns a tuple with the TpspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpspId

`func (o *Event) SetTpspId(v int32)`

SetTpspId sets TpspId field to given value.

### HasTpspId

`func (o *Event) HasTpspId() bool`

HasTpspId returns a boolean if a field has been set.

### GetMessageType

`func (o *Event) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *Event) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *Event) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *Event) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetMessageSubject

`func (o *Event) GetMessageSubject() string`

GetMessageSubject returns the MessageSubject field if non-nil, zero value otherwise.

### GetMessageSubjectOk

`func (o *Event) GetMessageSubjectOk() (*string, bool)`

GetMessageSubjectOk returns a tuple with the MessageSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSubject

`func (o *Event) SetMessageSubject(v string)`

SetMessageSubject sets MessageSubject field to given value.

### HasMessageSubject

`func (o *Event) HasMessageSubject() bool`

HasMessageSubject returns a boolean if a field has been set.

### GetAccountID

`func (o *Event) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Event) GetAccountIDOk() (*int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Event) SetAccountID(v int32)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *Event) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetSubAccountID

`func (o *Event) GetSubAccountID() int32`

GetSubAccountID returns the SubAccountID field if non-nil, zero value otherwise.

### GetSubAccountIDOk

`func (o *Event) GetSubAccountIDOk() (*int32, bool)`

GetSubAccountIDOk returns a tuple with the SubAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountID

`func (o *Event) SetSubAccountID(v int32)`

SetSubAccountID sets SubAccountID field to given value.

### HasSubAccountID

`func (o *Event) HasSubAccountID() bool`

HasSubAccountID returns a boolean if a field has been set.

### GetMessageID

`func (o *Event) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *Event) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *Event) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *Event) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

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

### GetFromName

`func (o *Event) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *Event) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *Event) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *Event) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

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

### GetToName

`func (o *Event) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *Event) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *Event) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *Event) HasToName() bool`

HasToName returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *Event) GetSubmittedAt() int32`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *Event) GetSubmittedAtOk() (*int32, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *Event) SetSubmittedAt(v int32)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *Event) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetSmtpCode

`func (o *Event) GetSmtpCode() int32`

GetSmtpCode returns the SmtpCode field if non-nil, zero value otherwise.

### GetSmtpCodeOk

`func (o *Event) GetSmtpCodeOk() (*int32, bool)`

GetSmtpCodeOk returns a tuple with the SmtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpCode

`func (o *Event) SetSmtpCode(v int32)`

SetSmtpCode sets SmtpCode field to given value.

### HasSmtpCode

`func (o *Event) HasSmtpCode() bool`

HasSmtpCode returns a boolean if a field has been set.

### GetSmtpDescription

`func (o *Event) GetSmtpDescription() string`

GetSmtpDescription returns the SmtpDescription field if non-nil, zero value otherwise.

### GetSmtpDescriptionOk

`func (o *Event) GetSmtpDescriptionOk() (*string, bool)`

GetSmtpDescriptionOk returns a tuple with the SmtpDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDescription

`func (o *Event) SetSmtpDescription(v string)`

SetSmtpDescription sets SmtpDescription field to given value.

### HasSmtpDescription

`func (o *Event) HasSmtpDescription() bool`

HasSmtpDescription returns a boolean if a field has been set.

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


