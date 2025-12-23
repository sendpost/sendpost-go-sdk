# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the member | [optional] 
**IsVerified** | Pointer to **bool** | Indicates whether the member is verified | [optional] 
**IsForbidden** | Pointer to **bool** | Indicates whether the member is forbidden | [optional] 
**FirebaseUID** | Pointer to **string** | Firebase UID for the member | [optional] 
**Email** | Pointer to **string** | Email for the member | [optional] 
**Name** | Pointer to **string** | Name for the member | [optional] 
**Url** | Pointer to **string** | Logo URL for the member | [optional] 
**CompanyName** | Pointer to **string** | Company name for the member | [optional] 
**OnboardQAnswered** | Pointer to **bool** | Indicates whether the member has answered onboarding question | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number for the member | [optional] 
**NotesColor** | Pointer to **string** | Color for the member&#39;s notes | [optional] 
**Created** | Pointer to **int64** | UNIX epoch nano timestamp when the member was created | [optional] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Member) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Member) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsVerified

`func (o *Member) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *Member) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *Member) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *Member) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIsForbidden

`func (o *Member) GetIsForbidden() bool`

GetIsForbidden returns the IsForbidden field if non-nil, zero value otherwise.

### GetIsForbiddenOk

`func (o *Member) GetIsForbiddenOk() (*bool, bool)`

GetIsForbiddenOk returns a tuple with the IsForbidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForbidden

`func (o *Member) SetIsForbidden(v bool)`

SetIsForbidden sets IsForbidden field to given value.

### HasIsForbidden

`func (o *Member) HasIsForbidden() bool`

HasIsForbidden returns a boolean if a field has been set.

### GetFirebaseUID

`func (o *Member) GetFirebaseUID() string`

GetFirebaseUID returns the FirebaseUID field if non-nil, zero value otherwise.

### GetFirebaseUIDOk

`func (o *Member) GetFirebaseUIDOk() (*string, bool)`

GetFirebaseUIDOk returns a tuple with the FirebaseUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirebaseUID

`func (o *Member) SetFirebaseUID(v string)`

SetFirebaseUID sets FirebaseUID field to given value.

### HasFirebaseUID

`func (o *Member) HasFirebaseUID() bool`

HasFirebaseUID returns a boolean if a field has been set.

### GetEmail

`func (o *Member) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Member) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Member) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Member) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Member) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Member) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Member) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Member) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Member) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Member) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Member) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Member) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCompanyName

`func (o *Member) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Member) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Member) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Member) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetOnboardQAnswered

`func (o *Member) GetOnboardQAnswered() bool`

GetOnboardQAnswered returns the OnboardQAnswered field if non-nil, zero value otherwise.

### GetOnboardQAnsweredOk

`func (o *Member) GetOnboardQAnsweredOk() (*bool, bool)`

GetOnboardQAnsweredOk returns a tuple with the OnboardQAnswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardQAnswered

`func (o *Member) SetOnboardQAnswered(v bool)`

SetOnboardQAnswered sets OnboardQAnswered field to given value.

### HasOnboardQAnswered

`func (o *Member) HasOnboardQAnswered() bool`

HasOnboardQAnswered returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Member) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Member) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Member) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Member) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetNotesColor

`func (o *Member) GetNotesColor() string`

GetNotesColor returns the NotesColor field if non-nil, zero value otherwise.

### GetNotesColorOk

`func (o *Member) GetNotesColorOk() (*string, bool)`

GetNotesColorOk returns a tuple with the NotesColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesColor

`func (o *Member) SetNotesColor(v string)`

SetNotesColor sets NotesColor field to given value.

### HasNotesColor

`func (o *Member) HasNotesColor() bool`

HasNotesColor returns a boolean if a field has been set.

### GetCreated

`func (o *Member) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Member) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Member) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Member) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


