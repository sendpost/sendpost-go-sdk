# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the team member | [optional] 
**Email** | Pointer to **string** | Email address of the team member (used for login) | [optional] 
**Name** | Pointer to **string** | Display name of the team member | [optional] 
**IsVerified** | Pointer to **bool** | Whether the member has verified their email address. Unverified members have limited access until verification is complete.  | [optional] 
**LogoUrl** | Pointer to **string** | URL of the member&#39;s profile picture/avatar | [optional] 
**CompanyName** | Pointer to **string** | Company or organization name | [optional] 
**OnboardQAnswered** | Pointer to **bool** | Whether the member has completed the onboarding questionnaire | [optional] 
**PhoneNumber** | Pointer to **string** | Contact phone number in E.164 format. Used for account recovery and important notifications.  | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the member was added | [optional] 

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

`func (o *Member) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Member) HasId() bool`

HasId returns a boolean if a field has been set.

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

### GetLogoUrl

`func (o *Member) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Member) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Member) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Member) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

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


