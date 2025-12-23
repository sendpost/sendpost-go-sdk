# SubAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the sub-account. | [optional] 
**ApiKey** | Pointer to **string** | API key for the sub-account. | [optional] 
**Name** | Pointer to **string** | Name of the sub-account. | [optional] 
**Labels** | Pointer to **[]string** | Labels associated with the sub-account | [optional] 
**SmtpAuths** | Pointer to [**[]SMTPAuth**](SMTPAuth.md) | SMTP Auths associated with the sub-account | [optional] 
**Type** | Pointer to **int32** | Type of the sub-account | [optional] 
**IsPlus** | Pointer to **bool** | Indicates whether the sub-account is a Plus sub-account | [optional] 
**Created** | Pointer to **int64** | UNIX epoch nano timestamp when the sub-account was created. | [optional] 
**CreatedBy** | Pointer to [**Member**](Member.md) | Member who created the sub-account | [optional] 
**UpdatedBy** | Pointer to [**Member**](Member.md) | Member who updated the sub-account | [optional] 
**Blocked** | Pointer to **bool** | Indicates whether the sub-account is blocked | [optional] 
**BlockedAt** | Pointer to **int32** | UNIX epoch nano timestamp when the sub-account was blocked (0 if not blocked) | [optional] 
**BlockReason** | Pointer to **string** | Reason for blocking the sub-account | [optional] 
**HbExempt** | Pointer to **bool** | Indicates whether the sub-account is exempt from hard bounce tracking | [optional] 
**GenerateWeeklyReport** | Pointer to **bool** | Indicates whether weekly reports are generated for this sub-account | [optional] 
**Handlers** | Pointer to **[]string** | Handlers associated with the sub-account | [optional] 

## Methods

### NewSubAccount

`func NewSubAccount() *SubAccount`

NewSubAccount instantiates a new SubAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountWithDefaults

`func NewSubAccountWithDefaults() *SubAccount`

NewSubAccountWithDefaults instantiates a new SubAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApiKey

`func (o *SubAccount) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *SubAccount) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *SubAccount) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *SubAccount) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetName

`func (o *SubAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *SubAccount) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SubAccount) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SubAccount) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SubAccount) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSmtpAuths

`func (o *SubAccount) GetSmtpAuths() []SMTPAuth`

GetSmtpAuths returns the SmtpAuths field if non-nil, zero value otherwise.

### GetSmtpAuthsOk

`func (o *SubAccount) GetSmtpAuthsOk() (*[]SMTPAuth, bool)`

GetSmtpAuthsOk returns a tuple with the SmtpAuths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAuths

`func (o *SubAccount) SetSmtpAuths(v []SMTPAuth)`

SetSmtpAuths sets SmtpAuths field to given value.

### HasSmtpAuths

`func (o *SubAccount) HasSmtpAuths() bool`

HasSmtpAuths returns a boolean if a field has been set.

### GetType

`func (o *SubAccount) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubAccount) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubAccount) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SubAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsPlus

`func (o *SubAccount) GetIsPlus() bool`

GetIsPlus returns the IsPlus field if non-nil, zero value otherwise.

### GetIsPlusOk

`func (o *SubAccount) GetIsPlusOk() (*bool, bool)`

GetIsPlusOk returns a tuple with the IsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlus

`func (o *SubAccount) SetIsPlus(v bool)`

SetIsPlus sets IsPlus field to given value.

### HasIsPlus

`func (o *SubAccount) HasIsPlus() bool`

HasIsPlus returns a boolean if a field has been set.

### GetCreated

`func (o *SubAccount) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubAccount) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubAccount) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SubAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SubAccount) GetCreatedBy() Member`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SubAccount) GetCreatedByOk() (*Member, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SubAccount) SetCreatedBy(v Member)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SubAccount) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SubAccount) GetUpdatedBy() Member`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SubAccount) GetUpdatedByOk() (*Member, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SubAccount) SetUpdatedBy(v Member)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SubAccount) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetBlocked

`func (o *SubAccount) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *SubAccount) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *SubAccount) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *SubAccount) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedAt

`func (o *SubAccount) GetBlockedAt() int32`

GetBlockedAt returns the BlockedAt field if non-nil, zero value otherwise.

### GetBlockedAtOk

`func (o *SubAccount) GetBlockedAtOk() (*int32, bool)`

GetBlockedAtOk returns a tuple with the BlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedAt

`func (o *SubAccount) SetBlockedAt(v int32)`

SetBlockedAt sets BlockedAt field to given value.

### HasBlockedAt

`func (o *SubAccount) HasBlockedAt() bool`

HasBlockedAt returns a boolean if a field has been set.

### GetBlockReason

`func (o *SubAccount) GetBlockReason() string`

GetBlockReason returns the BlockReason field if non-nil, zero value otherwise.

### GetBlockReasonOk

`func (o *SubAccount) GetBlockReasonOk() (*string, bool)`

GetBlockReasonOk returns a tuple with the BlockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReason

`func (o *SubAccount) SetBlockReason(v string)`

SetBlockReason sets BlockReason field to given value.

### HasBlockReason

`func (o *SubAccount) HasBlockReason() bool`

HasBlockReason returns a boolean if a field has been set.

### GetHbExempt

`func (o *SubAccount) GetHbExempt() bool`

GetHbExempt returns the HbExempt field if non-nil, zero value otherwise.

### GetHbExemptOk

`func (o *SubAccount) GetHbExemptOk() (*bool, bool)`

GetHbExemptOk returns a tuple with the HbExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbExempt

`func (o *SubAccount) SetHbExempt(v bool)`

SetHbExempt sets HbExempt field to given value.

### HasHbExempt

`func (o *SubAccount) HasHbExempt() bool`

HasHbExempt returns a boolean if a field has been set.

### GetGenerateWeeklyReport

`func (o *SubAccount) GetGenerateWeeklyReport() bool`

GetGenerateWeeklyReport returns the GenerateWeeklyReport field if non-nil, zero value otherwise.

### GetGenerateWeeklyReportOk

`func (o *SubAccount) GetGenerateWeeklyReportOk() (*bool, bool)`

GetGenerateWeeklyReportOk returns a tuple with the GenerateWeeklyReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateWeeklyReport

`func (o *SubAccount) SetGenerateWeeklyReport(v bool)`

SetGenerateWeeklyReport sets GenerateWeeklyReport field to given value.

### HasGenerateWeeklyReport

`func (o *SubAccount) HasGenerateWeeklyReport() bool`

HasGenerateWeeklyReport returns a boolean if a field has been set.

### GetHandlers

`func (o *SubAccount) GetHandlers() []string`

GetHandlers returns the Handlers field if non-nil, zero value otherwise.

### GetHandlersOk

`func (o *SubAccount) GetHandlersOk() (*[]string, bool)`

GetHandlersOk returns a tuple with the Handlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlers

`func (o *SubAccount) SetHandlers(v []string)`

SetHandlers sets Handlers field to given value.

### HasHandlers

`func (o *SubAccount) HasHandlers() bool`

HasHandlers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


