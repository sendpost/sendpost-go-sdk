# AutoWarmupPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the auto-warmup plan | [optional] 
**Name** | Pointer to **string** | Name of the auto-warmup plan | [optional] 
**GmailWarmupPlan** | Pointer to **string** | Gmail warmup plan configuration in JSON format | [optional] 
**YahooWarmupPlan** | Pointer to **string** | Yahoo warmup plan configuration in JSON format | [optional] 
**AolWarmupPlan** | Pointer to **string** | AOL warmup plan configuration in JSON format | [optional] 
**MicrosoftWarmupPlan** | Pointer to **string** | Microsoft warmup plan configuration in JSON format | [optional] 
**ComcastWarmupPlan** | Pointer to **string** | Comcast warmup plan configuration in JSON format | [optional] 
**YandexWarmupPlan** | Pointer to **string** | Yandex warmup plan configuration in JSON format | [optional] 
**GmxWarmupPlan** | Pointer to **string** | GMX warmup plan configuration in JSON format | [optional] 
**MailruWarmupPlan** | Pointer to **string** | Mail.ru warmup plan configuration in JSON format | [optional] 
**IcloudWarmupPlan** | Pointer to **string** | iCloud warmup plan configuration in JSON format | [optional] 
**ZohoWarmupPlan** | Pointer to **string** | Zoho warmup plan configuration in JSON format | [optional] 
**QqWarmupPlan** | Pointer to **string** | QQ warmup plan configuration in JSON format | [optional] 
**DefaultWarmupPlan** | Pointer to **string** | Default warmup plan configuration in JSON format | [optional] 
**AttWarmupPlan** | Pointer to **string** | AT&amp;T warmup plan configuration in JSON format | [optional] 
**Office365WarmupPlan** | Pointer to **string** | Office365 warmup plan configuration in JSON format | [optional] 
**GoogleworkspaceWarmupPlan** | Pointer to **string** | Google Workspace warmup plan configuration in JSON format | [optional] 
**ProofpointWarmupPlan** | Pointer to **string** | Proofpoint warmup plan configuration in JSON format | [optional] 
**MimecastWarmupPlan** | Pointer to **string** | Mimecast warmup plan configuration in JSON format | [optional] 
**BarracudaWarmupPlan** | Pointer to **string** | Barracuda warmup plan configuration in JSON format | [optional] 
**CiscoironportWarmupPlan** | Pointer to **string** | Cisco IronPort warmup plan configuration in JSON format | [optional] 
**RackspaceWarmupPlan** | Pointer to **string** | Rackspace warmup plan configuration in JSON format | [optional] 
**ZohobusinessWarmupPlan** | Pointer to **string** | Zoho Business warmup plan configuration in JSON format | [optional] 
**AmazonworkmailWarmupPlan** | Pointer to **string** | Amazon WorkMail warmup plan configuration in JSON format | [optional] 
**SymantecWarmupPlan** | Pointer to **string** | Symantec warmup plan configuration in JSON format | [optional] 
**FortinetWarmupPlan** | Pointer to **string** | Fortinet warmup plan configuration in JSON format | [optional] 
**SophosWarmupPlan** | Pointer to **string** | Sophos warmup plan configuration in JSON format | [optional] 
**TrendmicroWarmupPlan** | Pointer to **string** | Trend Micro warmup plan configuration in JSON format | [optional] 
**CheckpointWarmupPlan** | Pointer to **string** | Checkpoint warmup plan configuration in JSON format | [optional] 
**Created** | Pointer to **int64** | UNIX epoch nano timestamp when the warmup plan was created | [optional] 
**Updated** | Pointer to **int64** | UNIX epoch nano timestamp when the warmup plan was last updated | [optional] 
**WarmupInterval** | Pointer to **int32** | Warmup interval in hours | [optional] 

## Methods

### NewAutoWarmupPlan

`func NewAutoWarmupPlan() *AutoWarmupPlan`

NewAutoWarmupPlan instantiates a new AutoWarmupPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoWarmupPlanWithDefaults

`func NewAutoWarmupPlanWithDefaults() *AutoWarmupPlan`

NewAutoWarmupPlanWithDefaults instantiates a new AutoWarmupPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoWarmupPlan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoWarmupPlan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoWarmupPlan) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutoWarmupPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AutoWarmupPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoWarmupPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoWarmupPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoWarmupPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGmailWarmupPlan

`func (o *AutoWarmupPlan) GetGmailWarmupPlan() string`

GetGmailWarmupPlan returns the GmailWarmupPlan field if non-nil, zero value otherwise.

### GetGmailWarmupPlanOk

`func (o *AutoWarmupPlan) GetGmailWarmupPlanOk() (*string, bool)`

GetGmailWarmupPlanOk returns a tuple with the GmailWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmailWarmupPlan

`func (o *AutoWarmupPlan) SetGmailWarmupPlan(v string)`

SetGmailWarmupPlan sets GmailWarmupPlan field to given value.

### HasGmailWarmupPlan

`func (o *AutoWarmupPlan) HasGmailWarmupPlan() bool`

HasGmailWarmupPlan returns a boolean if a field has been set.

### GetYahooWarmupPlan

`func (o *AutoWarmupPlan) GetYahooWarmupPlan() string`

GetYahooWarmupPlan returns the YahooWarmupPlan field if non-nil, zero value otherwise.

### GetYahooWarmupPlanOk

`func (o *AutoWarmupPlan) GetYahooWarmupPlanOk() (*string, bool)`

GetYahooWarmupPlanOk returns a tuple with the YahooWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYahooWarmupPlan

`func (o *AutoWarmupPlan) SetYahooWarmupPlan(v string)`

SetYahooWarmupPlan sets YahooWarmupPlan field to given value.

### HasYahooWarmupPlan

`func (o *AutoWarmupPlan) HasYahooWarmupPlan() bool`

HasYahooWarmupPlan returns a boolean if a field has been set.

### GetAolWarmupPlan

`func (o *AutoWarmupPlan) GetAolWarmupPlan() string`

GetAolWarmupPlan returns the AolWarmupPlan field if non-nil, zero value otherwise.

### GetAolWarmupPlanOk

`func (o *AutoWarmupPlan) GetAolWarmupPlanOk() (*string, bool)`

GetAolWarmupPlanOk returns a tuple with the AolWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAolWarmupPlan

`func (o *AutoWarmupPlan) SetAolWarmupPlan(v string)`

SetAolWarmupPlan sets AolWarmupPlan field to given value.

### HasAolWarmupPlan

`func (o *AutoWarmupPlan) HasAolWarmupPlan() bool`

HasAolWarmupPlan returns a boolean if a field has been set.

### GetMicrosoftWarmupPlan

`func (o *AutoWarmupPlan) GetMicrosoftWarmupPlan() string`

GetMicrosoftWarmupPlan returns the MicrosoftWarmupPlan field if non-nil, zero value otherwise.

### GetMicrosoftWarmupPlanOk

`func (o *AutoWarmupPlan) GetMicrosoftWarmupPlanOk() (*string, bool)`

GetMicrosoftWarmupPlanOk returns a tuple with the MicrosoftWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftWarmupPlan

`func (o *AutoWarmupPlan) SetMicrosoftWarmupPlan(v string)`

SetMicrosoftWarmupPlan sets MicrosoftWarmupPlan field to given value.

### HasMicrosoftWarmupPlan

`func (o *AutoWarmupPlan) HasMicrosoftWarmupPlan() bool`

HasMicrosoftWarmupPlan returns a boolean if a field has been set.

### GetComcastWarmupPlan

`func (o *AutoWarmupPlan) GetComcastWarmupPlan() string`

GetComcastWarmupPlan returns the ComcastWarmupPlan field if non-nil, zero value otherwise.

### GetComcastWarmupPlanOk

`func (o *AutoWarmupPlan) GetComcastWarmupPlanOk() (*string, bool)`

GetComcastWarmupPlanOk returns a tuple with the ComcastWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComcastWarmupPlan

`func (o *AutoWarmupPlan) SetComcastWarmupPlan(v string)`

SetComcastWarmupPlan sets ComcastWarmupPlan field to given value.

### HasComcastWarmupPlan

`func (o *AutoWarmupPlan) HasComcastWarmupPlan() bool`

HasComcastWarmupPlan returns a boolean if a field has been set.

### GetYandexWarmupPlan

`func (o *AutoWarmupPlan) GetYandexWarmupPlan() string`

GetYandexWarmupPlan returns the YandexWarmupPlan field if non-nil, zero value otherwise.

### GetYandexWarmupPlanOk

`func (o *AutoWarmupPlan) GetYandexWarmupPlanOk() (*string, bool)`

GetYandexWarmupPlanOk returns a tuple with the YandexWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYandexWarmupPlan

`func (o *AutoWarmupPlan) SetYandexWarmupPlan(v string)`

SetYandexWarmupPlan sets YandexWarmupPlan field to given value.

### HasYandexWarmupPlan

`func (o *AutoWarmupPlan) HasYandexWarmupPlan() bool`

HasYandexWarmupPlan returns a boolean if a field has been set.

### GetGmxWarmupPlan

`func (o *AutoWarmupPlan) GetGmxWarmupPlan() string`

GetGmxWarmupPlan returns the GmxWarmupPlan field if non-nil, zero value otherwise.

### GetGmxWarmupPlanOk

`func (o *AutoWarmupPlan) GetGmxWarmupPlanOk() (*string, bool)`

GetGmxWarmupPlanOk returns a tuple with the GmxWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmxWarmupPlan

`func (o *AutoWarmupPlan) SetGmxWarmupPlan(v string)`

SetGmxWarmupPlan sets GmxWarmupPlan field to given value.

### HasGmxWarmupPlan

`func (o *AutoWarmupPlan) HasGmxWarmupPlan() bool`

HasGmxWarmupPlan returns a boolean if a field has been set.

### GetMailruWarmupPlan

`func (o *AutoWarmupPlan) GetMailruWarmupPlan() string`

GetMailruWarmupPlan returns the MailruWarmupPlan field if non-nil, zero value otherwise.

### GetMailruWarmupPlanOk

`func (o *AutoWarmupPlan) GetMailruWarmupPlanOk() (*string, bool)`

GetMailruWarmupPlanOk returns a tuple with the MailruWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailruWarmupPlan

`func (o *AutoWarmupPlan) SetMailruWarmupPlan(v string)`

SetMailruWarmupPlan sets MailruWarmupPlan field to given value.

### HasMailruWarmupPlan

`func (o *AutoWarmupPlan) HasMailruWarmupPlan() bool`

HasMailruWarmupPlan returns a boolean if a field has been set.

### GetIcloudWarmupPlan

`func (o *AutoWarmupPlan) GetIcloudWarmupPlan() string`

GetIcloudWarmupPlan returns the IcloudWarmupPlan field if non-nil, zero value otherwise.

### GetIcloudWarmupPlanOk

`func (o *AutoWarmupPlan) GetIcloudWarmupPlanOk() (*string, bool)`

GetIcloudWarmupPlanOk returns a tuple with the IcloudWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcloudWarmupPlan

`func (o *AutoWarmupPlan) SetIcloudWarmupPlan(v string)`

SetIcloudWarmupPlan sets IcloudWarmupPlan field to given value.

### HasIcloudWarmupPlan

`func (o *AutoWarmupPlan) HasIcloudWarmupPlan() bool`

HasIcloudWarmupPlan returns a boolean if a field has been set.

### GetZohoWarmupPlan

`func (o *AutoWarmupPlan) GetZohoWarmupPlan() string`

GetZohoWarmupPlan returns the ZohoWarmupPlan field if non-nil, zero value otherwise.

### GetZohoWarmupPlanOk

`func (o *AutoWarmupPlan) GetZohoWarmupPlanOk() (*string, bool)`

GetZohoWarmupPlanOk returns a tuple with the ZohoWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohoWarmupPlan

`func (o *AutoWarmupPlan) SetZohoWarmupPlan(v string)`

SetZohoWarmupPlan sets ZohoWarmupPlan field to given value.

### HasZohoWarmupPlan

`func (o *AutoWarmupPlan) HasZohoWarmupPlan() bool`

HasZohoWarmupPlan returns a boolean if a field has been set.

### GetQqWarmupPlan

`func (o *AutoWarmupPlan) GetQqWarmupPlan() string`

GetQqWarmupPlan returns the QqWarmupPlan field if non-nil, zero value otherwise.

### GetQqWarmupPlanOk

`func (o *AutoWarmupPlan) GetQqWarmupPlanOk() (*string, bool)`

GetQqWarmupPlanOk returns a tuple with the QqWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQqWarmupPlan

`func (o *AutoWarmupPlan) SetQqWarmupPlan(v string)`

SetQqWarmupPlan sets QqWarmupPlan field to given value.

### HasQqWarmupPlan

`func (o *AutoWarmupPlan) HasQqWarmupPlan() bool`

HasQqWarmupPlan returns a boolean if a field has been set.

### GetDefaultWarmupPlan

`func (o *AutoWarmupPlan) GetDefaultWarmupPlan() string`

GetDefaultWarmupPlan returns the DefaultWarmupPlan field if non-nil, zero value otherwise.

### GetDefaultWarmupPlanOk

`func (o *AutoWarmupPlan) GetDefaultWarmupPlanOk() (*string, bool)`

GetDefaultWarmupPlanOk returns a tuple with the DefaultWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWarmupPlan

`func (o *AutoWarmupPlan) SetDefaultWarmupPlan(v string)`

SetDefaultWarmupPlan sets DefaultWarmupPlan field to given value.

### HasDefaultWarmupPlan

`func (o *AutoWarmupPlan) HasDefaultWarmupPlan() bool`

HasDefaultWarmupPlan returns a boolean if a field has been set.

### GetAttWarmupPlan

`func (o *AutoWarmupPlan) GetAttWarmupPlan() string`

GetAttWarmupPlan returns the AttWarmupPlan field if non-nil, zero value otherwise.

### GetAttWarmupPlanOk

`func (o *AutoWarmupPlan) GetAttWarmupPlanOk() (*string, bool)`

GetAttWarmupPlanOk returns a tuple with the AttWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttWarmupPlan

`func (o *AutoWarmupPlan) SetAttWarmupPlan(v string)`

SetAttWarmupPlan sets AttWarmupPlan field to given value.

### HasAttWarmupPlan

`func (o *AutoWarmupPlan) HasAttWarmupPlan() bool`

HasAttWarmupPlan returns a boolean if a field has been set.

### GetOffice365WarmupPlan

`func (o *AutoWarmupPlan) GetOffice365WarmupPlan() string`

GetOffice365WarmupPlan returns the Office365WarmupPlan field if non-nil, zero value otherwise.

### GetOffice365WarmupPlanOk

`func (o *AutoWarmupPlan) GetOffice365WarmupPlanOk() (*string, bool)`

GetOffice365WarmupPlanOk returns a tuple with the Office365WarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365WarmupPlan

`func (o *AutoWarmupPlan) SetOffice365WarmupPlan(v string)`

SetOffice365WarmupPlan sets Office365WarmupPlan field to given value.

### HasOffice365WarmupPlan

`func (o *AutoWarmupPlan) HasOffice365WarmupPlan() bool`

HasOffice365WarmupPlan returns a boolean if a field has been set.

### GetGoogleworkspaceWarmupPlan

`func (o *AutoWarmupPlan) GetGoogleworkspaceWarmupPlan() string`

GetGoogleworkspaceWarmupPlan returns the GoogleworkspaceWarmupPlan field if non-nil, zero value otherwise.

### GetGoogleworkspaceWarmupPlanOk

`func (o *AutoWarmupPlan) GetGoogleworkspaceWarmupPlanOk() (*string, bool)`

GetGoogleworkspaceWarmupPlanOk returns a tuple with the GoogleworkspaceWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleworkspaceWarmupPlan

`func (o *AutoWarmupPlan) SetGoogleworkspaceWarmupPlan(v string)`

SetGoogleworkspaceWarmupPlan sets GoogleworkspaceWarmupPlan field to given value.

### HasGoogleworkspaceWarmupPlan

`func (o *AutoWarmupPlan) HasGoogleworkspaceWarmupPlan() bool`

HasGoogleworkspaceWarmupPlan returns a boolean if a field has been set.

### GetProofpointWarmupPlan

`func (o *AutoWarmupPlan) GetProofpointWarmupPlan() string`

GetProofpointWarmupPlan returns the ProofpointWarmupPlan field if non-nil, zero value otherwise.

### GetProofpointWarmupPlanOk

`func (o *AutoWarmupPlan) GetProofpointWarmupPlanOk() (*string, bool)`

GetProofpointWarmupPlanOk returns a tuple with the ProofpointWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofpointWarmupPlan

`func (o *AutoWarmupPlan) SetProofpointWarmupPlan(v string)`

SetProofpointWarmupPlan sets ProofpointWarmupPlan field to given value.

### HasProofpointWarmupPlan

`func (o *AutoWarmupPlan) HasProofpointWarmupPlan() bool`

HasProofpointWarmupPlan returns a boolean if a field has been set.

### GetMimecastWarmupPlan

`func (o *AutoWarmupPlan) GetMimecastWarmupPlan() string`

GetMimecastWarmupPlan returns the MimecastWarmupPlan field if non-nil, zero value otherwise.

### GetMimecastWarmupPlanOk

`func (o *AutoWarmupPlan) GetMimecastWarmupPlanOk() (*string, bool)`

GetMimecastWarmupPlanOk returns a tuple with the MimecastWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimecastWarmupPlan

`func (o *AutoWarmupPlan) SetMimecastWarmupPlan(v string)`

SetMimecastWarmupPlan sets MimecastWarmupPlan field to given value.

### HasMimecastWarmupPlan

`func (o *AutoWarmupPlan) HasMimecastWarmupPlan() bool`

HasMimecastWarmupPlan returns a boolean if a field has been set.

### GetBarracudaWarmupPlan

`func (o *AutoWarmupPlan) GetBarracudaWarmupPlan() string`

GetBarracudaWarmupPlan returns the BarracudaWarmupPlan field if non-nil, zero value otherwise.

### GetBarracudaWarmupPlanOk

`func (o *AutoWarmupPlan) GetBarracudaWarmupPlanOk() (*string, bool)`

GetBarracudaWarmupPlanOk returns a tuple with the BarracudaWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarracudaWarmupPlan

`func (o *AutoWarmupPlan) SetBarracudaWarmupPlan(v string)`

SetBarracudaWarmupPlan sets BarracudaWarmupPlan field to given value.

### HasBarracudaWarmupPlan

`func (o *AutoWarmupPlan) HasBarracudaWarmupPlan() bool`

HasBarracudaWarmupPlan returns a boolean if a field has been set.

### GetCiscoironportWarmupPlan

`func (o *AutoWarmupPlan) GetCiscoironportWarmupPlan() string`

GetCiscoironportWarmupPlan returns the CiscoironportWarmupPlan field if non-nil, zero value otherwise.

### GetCiscoironportWarmupPlanOk

`func (o *AutoWarmupPlan) GetCiscoironportWarmupPlanOk() (*string, bool)`

GetCiscoironportWarmupPlanOk returns a tuple with the CiscoironportWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoironportWarmupPlan

`func (o *AutoWarmupPlan) SetCiscoironportWarmupPlan(v string)`

SetCiscoironportWarmupPlan sets CiscoironportWarmupPlan field to given value.

### HasCiscoironportWarmupPlan

`func (o *AutoWarmupPlan) HasCiscoironportWarmupPlan() bool`

HasCiscoironportWarmupPlan returns a boolean if a field has been set.

### GetRackspaceWarmupPlan

`func (o *AutoWarmupPlan) GetRackspaceWarmupPlan() string`

GetRackspaceWarmupPlan returns the RackspaceWarmupPlan field if non-nil, zero value otherwise.

### GetRackspaceWarmupPlanOk

`func (o *AutoWarmupPlan) GetRackspaceWarmupPlanOk() (*string, bool)`

GetRackspaceWarmupPlanOk returns a tuple with the RackspaceWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackspaceWarmupPlan

`func (o *AutoWarmupPlan) SetRackspaceWarmupPlan(v string)`

SetRackspaceWarmupPlan sets RackspaceWarmupPlan field to given value.

### HasRackspaceWarmupPlan

`func (o *AutoWarmupPlan) HasRackspaceWarmupPlan() bool`

HasRackspaceWarmupPlan returns a boolean if a field has been set.

### GetZohobusinessWarmupPlan

`func (o *AutoWarmupPlan) GetZohobusinessWarmupPlan() string`

GetZohobusinessWarmupPlan returns the ZohobusinessWarmupPlan field if non-nil, zero value otherwise.

### GetZohobusinessWarmupPlanOk

`func (o *AutoWarmupPlan) GetZohobusinessWarmupPlanOk() (*string, bool)`

GetZohobusinessWarmupPlanOk returns a tuple with the ZohobusinessWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohobusinessWarmupPlan

`func (o *AutoWarmupPlan) SetZohobusinessWarmupPlan(v string)`

SetZohobusinessWarmupPlan sets ZohobusinessWarmupPlan field to given value.

### HasZohobusinessWarmupPlan

`func (o *AutoWarmupPlan) HasZohobusinessWarmupPlan() bool`

HasZohobusinessWarmupPlan returns a boolean if a field has been set.

### GetAmazonworkmailWarmupPlan

`func (o *AutoWarmupPlan) GetAmazonworkmailWarmupPlan() string`

GetAmazonworkmailWarmupPlan returns the AmazonworkmailWarmupPlan field if non-nil, zero value otherwise.

### GetAmazonworkmailWarmupPlanOk

`func (o *AutoWarmupPlan) GetAmazonworkmailWarmupPlanOk() (*string, bool)`

GetAmazonworkmailWarmupPlanOk returns a tuple with the AmazonworkmailWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonworkmailWarmupPlan

`func (o *AutoWarmupPlan) SetAmazonworkmailWarmupPlan(v string)`

SetAmazonworkmailWarmupPlan sets AmazonworkmailWarmupPlan field to given value.

### HasAmazonworkmailWarmupPlan

`func (o *AutoWarmupPlan) HasAmazonworkmailWarmupPlan() bool`

HasAmazonworkmailWarmupPlan returns a boolean if a field has been set.

### GetSymantecWarmupPlan

`func (o *AutoWarmupPlan) GetSymantecWarmupPlan() string`

GetSymantecWarmupPlan returns the SymantecWarmupPlan field if non-nil, zero value otherwise.

### GetSymantecWarmupPlanOk

`func (o *AutoWarmupPlan) GetSymantecWarmupPlanOk() (*string, bool)`

GetSymantecWarmupPlanOk returns a tuple with the SymantecWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymantecWarmupPlan

`func (o *AutoWarmupPlan) SetSymantecWarmupPlan(v string)`

SetSymantecWarmupPlan sets SymantecWarmupPlan field to given value.

### HasSymantecWarmupPlan

`func (o *AutoWarmupPlan) HasSymantecWarmupPlan() bool`

HasSymantecWarmupPlan returns a boolean if a field has been set.

### GetFortinetWarmupPlan

`func (o *AutoWarmupPlan) GetFortinetWarmupPlan() string`

GetFortinetWarmupPlan returns the FortinetWarmupPlan field if non-nil, zero value otherwise.

### GetFortinetWarmupPlanOk

`func (o *AutoWarmupPlan) GetFortinetWarmupPlanOk() (*string, bool)`

GetFortinetWarmupPlanOk returns a tuple with the FortinetWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFortinetWarmupPlan

`func (o *AutoWarmupPlan) SetFortinetWarmupPlan(v string)`

SetFortinetWarmupPlan sets FortinetWarmupPlan field to given value.

### HasFortinetWarmupPlan

`func (o *AutoWarmupPlan) HasFortinetWarmupPlan() bool`

HasFortinetWarmupPlan returns a boolean if a field has been set.

### GetSophosWarmupPlan

`func (o *AutoWarmupPlan) GetSophosWarmupPlan() string`

GetSophosWarmupPlan returns the SophosWarmupPlan field if non-nil, zero value otherwise.

### GetSophosWarmupPlanOk

`func (o *AutoWarmupPlan) GetSophosWarmupPlanOk() (*string, bool)`

GetSophosWarmupPlanOk returns a tuple with the SophosWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSophosWarmupPlan

`func (o *AutoWarmupPlan) SetSophosWarmupPlan(v string)`

SetSophosWarmupPlan sets SophosWarmupPlan field to given value.

### HasSophosWarmupPlan

`func (o *AutoWarmupPlan) HasSophosWarmupPlan() bool`

HasSophosWarmupPlan returns a boolean if a field has been set.

### GetTrendmicroWarmupPlan

`func (o *AutoWarmupPlan) GetTrendmicroWarmupPlan() string`

GetTrendmicroWarmupPlan returns the TrendmicroWarmupPlan field if non-nil, zero value otherwise.

### GetTrendmicroWarmupPlanOk

`func (o *AutoWarmupPlan) GetTrendmicroWarmupPlanOk() (*string, bool)`

GetTrendmicroWarmupPlanOk returns a tuple with the TrendmicroWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendmicroWarmupPlan

`func (o *AutoWarmupPlan) SetTrendmicroWarmupPlan(v string)`

SetTrendmicroWarmupPlan sets TrendmicroWarmupPlan field to given value.

### HasTrendmicroWarmupPlan

`func (o *AutoWarmupPlan) HasTrendmicroWarmupPlan() bool`

HasTrendmicroWarmupPlan returns a boolean if a field has been set.

### GetCheckpointWarmupPlan

`func (o *AutoWarmupPlan) GetCheckpointWarmupPlan() string`

GetCheckpointWarmupPlan returns the CheckpointWarmupPlan field if non-nil, zero value otherwise.

### GetCheckpointWarmupPlanOk

`func (o *AutoWarmupPlan) GetCheckpointWarmupPlanOk() (*string, bool)`

GetCheckpointWarmupPlanOk returns a tuple with the CheckpointWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointWarmupPlan

`func (o *AutoWarmupPlan) SetCheckpointWarmupPlan(v string)`

SetCheckpointWarmupPlan sets CheckpointWarmupPlan field to given value.

### HasCheckpointWarmupPlan

`func (o *AutoWarmupPlan) HasCheckpointWarmupPlan() bool`

HasCheckpointWarmupPlan returns a boolean if a field has been set.

### GetCreated

`func (o *AutoWarmupPlan) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutoWarmupPlan) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutoWarmupPlan) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutoWarmupPlan) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AutoWarmupPlan) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutoWarmupPlan) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutoWarmupPlan) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutoWarmupPlan) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWarmupInterval

`func (o *AutoWarmupPlan) GetWarmupInterval() int32`

GetWarmupInterval returns the WarmupInterval field if non-nil, zero value otherwise.

### GetWarmupIntervalOk

`func (o *AutoWarmupPlan) GetWarmupIntervalOk() (*int32, bool)`

GetWarmupIntervalOk returns a tuple with the WarmupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupInterval

`func (o *AutoWarmupPlan) SetWarmupInterval(v int32)`

SetWarmupInterval sets WarmupInterval field to given value.

### HasWarmupInterval

`func (o *AutoWarmupPlan) HasWarmupInterval() bool`

HasWarmupInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


