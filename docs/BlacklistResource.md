# BlacklistResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the blacklist resource entry. | [optional] 
**Type** | Pointer to **string** | The kind of target being monitored (e.g. &#x60;domain&#x60; or &#x60;ip&#x60;). | [optional] 
**Target** | Pointer to **string** | The domain or IP address being monitored for blacklisting. | [optional] 
**AddDate** | Pointer to **int64** | UNIX epoch timestamp when the resource was added to monitoring. | [optional] 
**LastCheck** | Pointer to **int64** | UNIX epoch timestamp of the most recent blacklist check. | [optional] 
**Status** | Pointer to **string** | Current blacklist status of the target (e.g. &#x60;clean&#x60;, &#x60;listed&#x60;). | [optional] 
**Label** | Pointer to **string** | User-assigned label for the monitored resource. | [optional] 
**ContactListId** | Pointer to **string** | Identifier of the contact list associated with this resource, if any. | [optional] 
**BlacklistedCount** | Pointer to **string** | Number of RBLs the target is currently listed on (string-encoded). | [optional] 
**BlacklistedOn** | Pointer to [**[]BlacklistedOn**](BlacklistedOn.md) | The specific RBLs this target is currently listed on. | [optional] 
**Links** | Pointer to [**BlacklistLinks**](BlacklistLinks.md) |  | [optional] 

## Methods

### NewBlacklistResource

`func NewBlacklistResource() *BlacklistResource`

NewBlacklistResource instantiates a new BlacklistResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlacklistResourceWithDefaults

`func NewBlacklistResourceWithDefaults() *BlacklistResource`

NewBlacklistResourceWithDefaults instantiates a new BlacklistResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlacklistResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlacklistResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlacklistResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlacklistResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BlacklistResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlacklistResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlacklistResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlacklistResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTarget

`func (o *BlacklistResource) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BlacklistResource) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BlacklistResource) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BlacklistResource) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetAddDate

`func (o *BlacklistResource) GetAddDate() int64`

GetAddDate returns the AddDate field if non-nil, zero value otherwise.

### GetAddDateOk

`func (o *BlacklistResource) GetAddDateOk() (*int64, bool)`

GetAddDateOk returns a tuple with the AddDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDate

`func (o *BlacklistResource) SetAddDate(v int64)`

SetAddDate sets AddDate field to given value.

### HasAddDate

`func (o *BlacklistResource) HasAddDate() bool`

HasAddDate returns a boolean if a field has been set.

### GetLastCheck

`func (o *BlacklistResource) GetLastCheck() int64`

GetLastCheck returns the LastCheck field if non-nil, zero value otherwise.

### GetLastCheckOk

`func (o *BlacklistResource) GetLastCheckOk() (*int64, bool)`

GetLastCheckOk returns a tuple with the LastCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheck

`func (o *BlacklistResource) SetLastCheck(v int64)`

SetLastCheck sets LastCheck field to given value.

### HasLastCheck

`func (o *BlacklistResource) HasLastCheck() bool`

HasLastCheck returns a boolean if a field has been set.

### GetStatus

`func (o *BlacklistResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlacklistResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlacklistResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlacklistResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLabel

`func (o *BlacklistResource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BlacklistResource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BlacklistResource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BlacklistResource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetContactListId

`func (o *BlacklistResource) GetContactListId() string`

GetContactListId returns the ContactListId field if non-nil, zero value otherwise.

### GetContactListIdOk

`func (o *BlacklistResource) GetContactListIdOk() (*string, bool)`

GetContactListIdOk returns a tuple with the ContactListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactListId

`func (o *BlacklistResource) SetContactListId(v string)`

SetContactListId sets ContactListId field to given value.

### HasContactListId

`func (o *BlacklistResource) HasContactListId() bool`

HasContactListId returns a boolean if a field has been set.

### GetBlacklistedCount

`func (o *BlacklistResource) GetBlacklistedCount() string`

GetBlacklistedCount returns the BlacklistedCount field if non-nil, zero value otherwise.

### GetBlacklistedCountOk

`func (o *BlacklistResource) GetBlacklistedCountOk() (*string, bool)`

GetBlacklistedCountOk returns a tuple with the BlacklistedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedCount

`func (o *BlacklistResource) SetBlacklistedCount(v string)`

SetBlacklistedCount sets BlacklistedCount field to given value.

### HasBlacklistedCount

`func (o *BlacklistResource) HasBlacklistedCount() bool`

HasBlacklistedCount returns a boolean if a field has been set.

### GetBlacklistedOn

`func (o *BlacklistResource) GetBlacklistedOn() []BlacklistedOn`

GetBlacklistedOn returns the BlacklistedOn field if non-nil, zero value otherwise.

### GetBlacklistedOnOk

`func (o *BlacklistResource) GetBlacklistedOnOk() (*[]BlacklistedOn, bool)`

GetBlacklistedOnOk returns a tuple with the BlacklistedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedOn

`func (o *BlacklistResource) SetBlacklistedOn(v []BlacklistedOn)`

SetBlacklistedOn sets BlacklistedOn field to given value.

### HasBlacklistedOn

`func (o *BlacklistResource) HasBlacklistedOn() bool`

HasBlacklistedOn returns a boolean if a field has been set.

### GetLinks

`func (o *BlacklistResource) GetLinks() BlacklistLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BlacklistResource) GetLinksOk() (*BlacklistLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BlacklistResource) SetLinks(v BlacklistLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BlacklistResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


