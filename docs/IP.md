# IP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for the IP | 
**Labels** | Pointer to **[]string** | Labels associated with the IP | [optional] 
**PublicIP** | **string** | The public IP address associated with the resource | 
**SystemDomain** | Pointer to [**Domain**](Domain.md) | Details of the system domain associated with the IP | [optional] 
**ReverseDNSHostname** | Pointer to **string** | The reverse DNS hostname for the IP | [optional] 
**Type** | Pointer to **int32** | Type of the IP | [optional] 
**GmailSettings** | Pointer to **string** | Configuration for Gmail delivery settings in JSON format | [optional] 
**YahooSettings** | Pointer to **string** | Configuration for Yahoo delivery settings in JSON format | [optional] 
**AolSettings** | Pointer to **string** | Configuration for AOL delivery settings in JSON format | [optional] 
**MicrosoftSettings** | Pointer to **string** | Configuration for Microsoft delivery settings in JSON format | [optional] 
**ComcastSettings** | Pointer to **string** | Configuration for Comcast delivery settings in JSON format | [optional] 
**YandexSettings** | Pointer to **string** | Configuration for Yandex delivery settings in JSON format | [optional] 
**GmxSettings** | Pointer to **string** | Configuration for GMX delivery settings in JSON format | [optional] 
**MailruSettings** | Pointer to **string** | Configuration for Mail.ru delivery settings in JSON format | [optional] 
**IcloudSettings** | Pointer to **string** | Configuration for iCloud delivery settings in JSON format | [optional] 
**ZohoSettings** | Pointer to **string** | Configuration for Zoho delivery settings in JSON format | [optional] 
**QqSettings** | Pointer to **string** | Configuration for QQ delivery settings in JSON format | [optional] 
**DefaultSettings** | Pointer to **string** | Default delivery settings in JSON format | [optional] 
**AttSettings** | Pointer to **string** | Configuration for AT&amp;T delivery settings in JSON format | [optional] 
**Office365Settings** | Pointer to **string** | Configuration for Office365 delivery settings in JSON format | [optional] 
**GoogleworkspaceSettings** | Pointer to **string** | Configuration for Google Workspace delivery settings in JSON format | [optional] 
**ProofpointSettings** | Pointer to **string** | Configuration for Proofpoint delivery settings in JSON format | [optional] 
**MimecastSettings** | Pointer to **string** | Configuration for Mimecast delivery settings in JSON format | [optional] 
**BarracudaSettings** | Pointer to **string** | Configuration for Barracuda delivery settings in JSON format | [optional] 
**CiscoironportSettings** | Pointer to **string** | Configuration for Cisco IronPort delivery settings in JSON format | [optional] 
**RackspaceSettings** | Pointer to **string** | Configuration for Rackspace delivery settings in JSON format | [optional] 
**ZohobusinessSettings** | Pointer to **string** | Configuration for Zoho Business delivery settings in JSON format | [optional] 
**AmazonworkmailSettings** | Pointer to **string** | Configuration for Amazon WorkMail delivery settings in JSON format | [optional] 
**SymantecSettings** | Pointer to **string** | Configuration for Symantec delivery settings in JSON format | [optional] 
**FortinetSettings** | Pointer to **string** | Configuration for Fortinet delivery settings in JSON format | [optional] 
**SophosSettings** | Pointer to **string** | Configuration for Sophos delivery settings in JSON format | [optional] 
**TrendmicroSettings** | Pointer to **string** | Configuration for Trend Micro delivery settings in JSON format | [optional] 
**CheckpointSettings** | Pointer to **string** | Configuration for Checkpoint delivery settings in JSON format | [optional] 
**Created** | **int64** | The timestamp (UNIX epoch) when the IP was created | 
**InfraClassification** | Pointer to **string** | Classification of the infrastructure | [optional] 
**InfraMonitor** | Pointer to **bool** | Indicates whether infrastructure monitoring is enabled | [optional] 
**State** | Pointer to **int32** | The state of the IP | [optional] 
**AutoWarmupPlan** | Pointer to **string** | The auto-warmup plan associated with the IP | [optional] 

## Methods

### NewIP

`func NewIP(id int32, publicIP string, created int64, ) *IP`

NewIP instantiates a new IP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPWithDefaults

`func NewIPWithDefaults() *IP`

NewIPWithDefaults instantiates a new IP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IP) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IP) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IP) SetId(v int32)`

SetId sets Id field to given value.


### GetLabels

`func (o *IP) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IP) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IP) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IP) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPublicIP

`func (o *IP) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *IP) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *IP) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.


### GetSystemDomain

`func (o *IP) GetSystemDomain() Domain`

GetSystemDomain returns the SystemDomain field if non-nil, zero value otherwise.

### GetSystemDomainOk

`func (o *IP) GetSystemDomainOk() (*Domain, bool)`

GetSystemDomainOk returns a tuple with the SystemDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDomain

`func (o *IP) SetSystemDomain(v Domain)`

SetSystemDomain sets SystemDomain field to given value.

### HasSystemDomain

`func (o *IP) HasSystemDomain() bool`

HasSystemDomain returns a boolean if a field has been set.

### GetReverseDNSHostname

`func (o *IP) GetReverseDNSHostname() string`

GetReverseDNSHostname returns the ReverseDNSHostname field if non-nil, zero value otherwise.

### GetReverseDNSHostnameOk

`func (o *IP) GetReverseDNSHostnameOk() (*string, bool)`

GetReverseDNSHostnameOk returns a tuple with the ReverseDNSHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDNSHostname

`func (o *IP) SetReverseDNSHostname(v string)`

SetReverseDNSHostname sets ReverseDNSHostname field to given value.

### HasReverseDNSHostname

`func (o *IP) HasReverseDNSHostname() bool`

HasReverseDNSHostname returns a boolean if a field has been set.

### GetType

`func (o *IP) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IP) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IP) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *IP) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGmailSettings

`func (o *IP) GetGmailSettings() string`

GetGmailSettings returns the GmailSettings field if non-nil, zero value otherwise.

### GetGmailSettingsOk

`func (o *IP) GetGmailSettingsOk() (*string, bool)`

GetGmailSettingsOk returns a tuple with the GmailSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmailSettings

`func (o *IP) SetGmailSettings(v string)`

SetGmailSettings sets GmailSettings field to given value.

### HasGmailSettings

`func (o *IP) HasGmailSettings() bool`

HasGmailSettings returns a boolean if a field has been set.

### GetYahooSettings

`func (o *IP) GetYahooSettings() string`

GetYahooSettings returns the YahooSettings field if non-nil, zero value otherwise.

### GetYahooSettingsOk

`func (o *IP) GetYahooSettingsOk() (*string, bool)`

GetYahooSettingsOk returns a tuple with the YahooSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYahooSettings

`func (o *IP) SetYahooSettings(v string)`

SetYahooSettings sets YahooSettings field to given value.

### HasYahooSettings

`func (o *IP) HasYahooSettings() bool`

HasYahooSettings returns a boolean if a field has been set.

### GetAolSettings

`func (o *IP) GetAolSettings() string`

GetAolSettings returns the AolSettings field if non-nil, zero value otherwise.

### GetAolSettingsOk

`func (o *IP) GetAolSettingsOk() (*string, bool)`

GetAolSettingsOk returns a tuple with the AolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAolSettings

`func (o *IP) SetAolSettings(v string)`

SetAolSettings sets AolSettings field to given value.

### HasAolSettings

`func (o *IP) HasAolSettings() bool`

HasAolSettings returns a boolean if a field has been set.

### GetMicrosoftSettings

`func (o *IP) GetMicrosoftSettings() string`

GetMicrosoftSettings returns the MicrosoftSettings field if non-nil, zero value otherwise.

### GetMicrosoftSettingsOk

`func (o *IP) GetMicrosoftSettingsOk() (*string, bool)`

GetMicrosoftSettingsOk returns a tuple with the MicrosoftSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftSettings

`func (o *IP) SetMicrosoftSettings(v string)`

SetMicrosoftSettings sets MicrosoftSettings field to given value.

### HasMicrosoftSettings

`func (o *IP) HasMicrosoftSettings() bool`

HasMicrosoftSettings returns a boolean if a field has been set.

### GetComcastSettings

`func (o *IP) GetComcastSettings() string`

GetComcastSettings returns the ComcastSettings field if non-nil, zero value otherwise.

### GetComcastSettingsOk

`func (o *IP) GetComcastSettingsOk() (*string, bool)`

GetComcastSettingsOk returns a tuple with the ComcastSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComcastSettings

`func (o *IP) SetComcastSettings(v string)`

SetComcastSettings sets ComcastSettings field to given value.

### HasComcastSettings

`func (o *IP) HasComcastSettings() bool`

HasComcastSettings returns a boolean if a field has been set.

### GetYandexSettings

`func (o *IP) GetYandexSettings() string`

GetYandexSettings returns the YandexSettings field if non-nil, zero value otherwise.

### GetYandexSettingsOk

`func (o *IP) GetYandexSettingsOk() (*string, bool)`

GetYandexSettingsOk returns a tuple with the YandexSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYandexSettings

`func (o *IP) SetYandexSettings(v string)`

SetYandexSettings sets YandexSettings field to given value.

### HasYandexSettings

`func (o *IP) HasYandexSettings() bool`

HasYandexSettings returns a boolean if a field has been set.

### GetGmxSettings

`func (o *IP) GetGmxSettings() string`

GetGmxSettings returns the GmxSettings field if non-nil, zero value otherwise.

### GetGmxSettingsOk

`func (o *IP) GetGmxSettingsOk() (*string, bool)`

GetGmxSettingsOk returns a tuple with the GmxSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmxSettings

`func (o *IP) SetGmxSettings(v string)`

SetGmxSettings sets GmxSettings field to given value.

### HasGmxSettings

`func (o *IP) HasGmxSettings() bool`

HasGmxSettings returns a boolean if a field has been set.

### GetMailruSettings

`func (o *IP) GetMailruSettings() string`

GetMailruSettings returns the MailruSettings field if non-nil, zero value otherwise.

### GetMailruSettingsOk

`func (o *IP) GetMailruSettingsOk() (*string, bool)`

GetMailruSettingsOk returns a tuple with the MailruSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailruSettings

`func (o *IP) SetMailruSettings(v string)`

SetMailruSettings sets MailruSettings field to given value.

### HasMailruSettings

`func (o *IP) HasMailruSettings() bool`

HasMailruSettings returns a boolean if a field has been set.

### GetIcloudSettings

`func (o *IP) GetIcloudSettings() string`

GetIcloudSettings returns the IcloudSettings field if non-nil, zero value otherwise.

### GetIcloudSettingsOk

`func (o *IP) GetIcloudSettingsOk() (*string, bool)`

GetIcloudSettingsOk returns a tuple with the IcloudSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcloudSettings

`func (o *IP) SetIcloudSettings(v string)`

SetIcloudSettings sets IcloudSettings field to given value.

### HasIcloudSettings

`func (o *IP) HasIcloudSettings() bool`

HasIcloudSettings returns a boolean if a field has been set.

### GetZohoSettings

`func (o *IP) GetZohoSettings() string`

GetZohoSettings returns the ZohoSettings field if non-nil, zero value otherwise.

### GetZohoSettingsOk

`func (o *IP) GetZohoSettingsOk() (*string, bool)`

GetZohoSettingsOk returns a tuple with the ZohoSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohoSettings

`func (o *IP) SetZohoSettings(v string)`

SetZohoSettings sets ZohoSettings field to given value.

### HasZohoSettings

`func (o *IP) HasZohoSettings() bool`

HasZohoSettings returns a boolean if a field has been set.

### GetQqSettings

`func (o *IP) GetQqSettings() string`

GetQqSettings returns the QqSettings field if non-nil, zero value otherwise.

### GetQqSettingsOk

`func (o *IP) GetQqSettingsOk() (*string, bool)`

GetQqSettingsOk returns a tuple with the QqSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQqSettings

`func (o *IP) SetQqSettings(v string)`

SetQqSettings sets QqSettings field to given value.

### HasQqSettings

`func (o *IP) HasQqSettings() bool`

HasQqSettings returns a boolean if a field has been set.

### GetDefaultSettings

`func (o *IP) GetDefaultSettings() string`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *IP) GetDefaultSettingsOk() (*string, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *IP) SetDefaultSettings(v string)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *IP) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetAttSettings

`func (o *IP) GetAttSettings() string`

GetAttSettings returns the AttSettings field if non-nil, zero value otherwise.

### GetAttSettingsOk

`func (o *IP) GetAttSettingsOk() (*string, bool)`

GetAttSettingsOk returns a tuple with the AttSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttSettings

`func (o *IP) SetAttSettings(v string)`

SetAttSettings sets AttSettings field to given value.

### HasAttSettings

`func (o *IP) HasAttSettings() bool`

HasAttSettings returns a boolean if a field has been set.

### GetOffice365Settings

`func (o *IP) GetOffice365Settings() string`

GetOffice365Settings returns the Office365Settings field if non-nil, zero value otherwise.

### GetOffice365SettingsOk

`func (o *IP) GetOffice365SettingsOk() (*string, bool)`

GetOffice365SettingsOk returns a tuple with the Office365Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Settings

`func (o *IP) SetOffice365Settings(v string)`

SetOffice365Settings sets Office365Settings field to given value.

### HasOffice365Settings

`func (o *IP) HasOffice365Settings() bool`

HasOffice365Settings returns a boolean if a field has been set.

### GetGoogleworkspaceSettings

`func (o *IP) GetGoogleworkspaceSettings() string`

GetGoogleworkspaceSettings returns the GoogleworkspaceSettings field if non-nil, zero value otherwise.

### GetGoogleworkspaceSettingsOk

`func (o *IP) GetGoogleworkspaceSettingsOk() (*string, bool)`

GetGoogleworkspaceSettingsOk returns a tuple with the GoogleworkspaceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleworkspaceSettings

`func (o *IP) SetGoogleworkspaceSettings(v string)`

SetGoogleworkspaceSettings sets GoogleworkspaceSettings field to given value.

### HasGoogleworkspaceSettings

`func (o *IP) HasGoogleworkspaceSettings() bool`

HasGoogleworkspaceSettings returns a boolean if a field has been set.

### GetProofpointSettings

`func (o *IP) GetProofpointSettings() string`

GetProofpointSettings returns the ProofpointSettings field if non-nil, zero value otherwise.

### GetProofpointSettingsOk

`func (o *IP) GetProofpointSettingsOk() (*string, bool)`

GetProofpointSettingsOk returns a tuple with the ProofpointSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofpointSettings

`func (o *IP) SetProofpointSettings(v string)`

SetProofpointSettings sets ProofpointSettings field to given value.

### HasProofpointSettings

`func (o *IP) HasProofpointSettings() bool`

HasProofpointSettings returns a boolean if a field has been set.

### GetMimecastSettings

`func (o *IP) GetMimecastSettings() string`

GetMimecastSettings returns the MimecastSettings field if non-nil, zero value otherwise.

### GetMimecastSettingsOk

`func (o *IP) GetMimecastSettingsOk() (*string, bool)`

GetMimecastSettingsOk returns a tuple with the MimecastSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimecastSettings

`func (o *IP) SetMimecastSettings(v string)`

SetMimecastSettings sets MimecastSettings field to given value.

### HasMimecastSettings

`func (o *IP) HasMimecastSettings() bool`

HasMimecastSettings returns a boolean if a field has been set.

### GetBarracudaSettings

`func (o *IP) GetBarracudaSettings() string`

GetBarracudaSettings returns the BarracudaSettings field if non-nil, zero value otherwise.

### GetBarracudaSettingsOk

`func (o *IP) GetBarracudaSettingsOk() (*string, bool)`

GetBarracudaSettingsOk returns a tuple with the BarracudaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarracudaSettings

`func (o *IP) SetBarracudaSettings(v string)`

SetBarracudaSettings sets BarracudaSettings field to given value.

### HasBarracudaSettings

`func (o *IP) HasBarracudaSettings() bool`

HasBarracudaSettings returns a boolean if a field has been set.

### GetCiscoironportSettings

`func (o *IP) GetCiscoironportSettings() string`

GetCiscoironportSettings returns the CiscoironportSettings field if non-nil, zero value otherwise.

### GetCiscoironportSettingsOk

`func (o *IP) GetCiscoironportSettingsOk() (*string, bool)`

GetCiscoironportSettingsOk returns a tuple with the CiscoironportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoironportSettings

`func (o *IP) SetCiscoironportSettings(v string)`

SetCiscoironportSettings sets CiscoironportSettings field to given value.

### HasCiscoironportSettings

`func (o *IP) HasCiscoironportSettings() bool`

HasCiscoironportSettings returns a boolean if a field has been set.

### GetRackspaceSettings

`func (o *IP) GetRackspaceSettings() string`

GetRackspaceSettings returns the RackspaceSettings field if non-nil, zero value otherwise.

### GetRackspaceSettingsOk

`func (o *IP) GetRackspaceSettingsOk() (*string, bool)`

GetRackspaceSettingsOk returns a tuple with the RackspaceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackspaceSettings

`func (o *IP) SetRackspaceSettings(v string)`

SetRackspaceSettings sets RackspaceSettings field to given value.

### HasRackspaceSettings

`func (o *IP) HasRackspaceSettings() bool`

HasRackspaceSettings returns a boolean if a field has been set.

### GetZohobusinessSettings

`func (o *IP) GetZohobusinessSettings() string`

GetZohobusinessSettings returns the ZohobusinessSettings field if non-nil, zero value otherwise.

### GetZohobusinessSettingsOk

`func (o *IP) GetZohobusinessSettingsOk() (*string, bool)`

GetZohobusinessSettingsOk returns a tuple with the ZohobusinessSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohobusinessSettings

`func (o *IP) SetZohobusinessSettings(v string)`

SetZohobusinessSettings sets ZohobusinessSettings field to given value.

### HasZohobusinessSettings

`func (o *IP) HasZohobusinessSettings() bool`

HasZohobusinessSettings returns a boolean if a field has been set.

### GetAmazonworkmailSettings

`func (o *IP) GetAmazonworkmailSettings() string`

GetAmazonworkmailSettings returns the AmazonworkmailSettings field if non-nil, zero value otherwise.

### GetAmazonworkmailSettingsOk

`func (o *IP) GetAmazonworkmailSettingsOk() (*string, bool)`

GetAmazonworkmailSettingsOk returns a tuple with the AmazonworkmailSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonworkmailSettings

`func (o *IP) SetAmazonworkmailSettings(v string)`

SetAmazonworkmailSettings sets AmazonworkmailSettings field to given value.

### HasAmazonworkmailSettings

`func (o *IP) HasAmazonworkmailSettings() bool`

HasAmazonworkmailSettings returns a boolean if a field has been set.

### GetSymantecSettings

`func (o *IP) GetSymantecSettings() string`

GetSymantecSettings returns the SymantecSettings field if non-nil, zero value otherwise.

### GetSymantecSettingsOk

`func (o *IP) GetSymantecSettingsOk() (*string, bool)`

GetSymantecSettingsOk returns a tuple with the SymantecSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymantecSettings

`func (o *IP) SetSymantecSettings(v string)`

SetSymantecSettings sets SymantecSettings field to given value.

### HasSymantecSettings

`func (o *IP) HasSymantecSettings() bool`

HasSymantecSettings returns a boolean if a field has been set.

### GetFortinetSettings

`func (o *IP) GetFortinetSettings() string`

GetFortinetSettings returns the FortinetSettings field if non-nil, zero value otherwise.

### GetFortinetSettingsOk

`func (o *IP) GetFortinetSettingsOk() (*string, bool)`

GetFortinetSettingsOk returns a tuple with the FortinetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFortinetSettings

`func (o *IP) SetFortinetSettings(v string)`

SetFortinetSettings sets FortinetSettings field to given value.

### HasFortinetSettings

`func (o *IP) HasFortinetSettings() bool`

HasFortinetSettings returns a boolean if a field has been set.

### GetSophosSettings

`func (o *IP) GetSophosSettings() string`

GetSophosSettings returns the SophosSettings field if non-nil, zero value otherwise.

### GetSophosSettingsOk

`func (o *IP) GetSophosSettingsOk() (*string, bool)`

GetSophosSettingsOk returns a tuple with the SophosSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSophosSettings

`func (o *IP) SetSophosSettings(v string)`

SetSophosSettings sets SophosSettings field to given value.

### HasSophosSettings

`func (o *IP) HasSophosSettings() bool`

HasSophosSettings returns a boolean if a field has been set.

### GetTrendmicroSettings

`func (o *IP) GetTrendmicroSettings() string`

GetTrendmicroSettings returns the TrendmicroSettings field if non-nil, zero value otherwise.

### GetTrendmicroSettingsOk

`func (o *IP) GetTrendmicroSettingsOk() (*string, bool)`

GetTrendmicroSettingsOk returns a tuple with the TrendmicroSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendmicroSettings

`func (o *IP) SetTrendmicroSettings(v string)`

SetTrendmicroSettings sets TrendmicroSettings field to given value.

### HasTrendmicroSettings

`func (o *IP) HasTrendmicroSettings() bool`

HasTrendmicroSettings returns a boolean if a field has been set.

### GetCheckpointSettings

`func (o *IP) GetCheckpointSettings() string`

GetCheckpointSettings returns the CheckpointSettings field if non-nil, zero value otherwise.

### GetCheckpointSettingsOk

`func (o *IP) GetCheckpointSettingsOk() (*string, bool)`

GetCheckpointSettingsOk returns a tuple with the CheckpointSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointSettings

`func (o *IP) SetCheckpointSettings(v string)`

SetCheckpointSettings sets CheckpointSettings field to given value.

### HasCheckpointSettings

`func (o *IP) HasCheckpointSettings() bool`

HasCheckpointSettings returns a boolean if a field has been set.

### GetCreated

`func (o *IP) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IP) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IP) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetInfraClassification

`func (o *IP) GetInfraClassification() string`

GetInfraClassification returns the InfraClassification field if non-nil, zero value otherwise.

### GetInfraClassificationOk

`func (o *IP) GetInfraClassificationOk() (*string, bool)`

GetInfraClassificationOk returns a tuple with the InfraClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraClassification

`func (o *IP) SetInfraClassification(v string)`

SetInfraClassification sets InfraClassification field to given value.

### HasInfraClassification

`func (o *IP) HasInfraClassification() bool`

HasInfraClassification returns a boolean if a field has been set.

### GetInfraMonitor

`func (o *IP) GetInfraMonitor() bool`

GetInfraMonitor returns the InfraMonitor field if non-nil, zero value otherwise.

### GetInfraMonitorOk

`func (o *IP) GetInfraMonitorOk() (*bool, bool)`

GetInfraMonitorOk returns a tuple with the InfraMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraMonitor

`func (o *IP) SetInfraMonitor(v bool)`

SetInfraMonitor sets InfraMonitor field to given value.

### HasInfraMonitor

`func (o *IP) HasInfraMonitor() bool`

HasInfraMonitor returns a boolean if a field has been set.

### GetState

`func (o *IP) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IP) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IP) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *IP) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAutoWarmupPlan

`func (o *IP) GetAutoWarmupPlan() string`

GetAutoWarmupPlan returns the AutoWarmupPlan field if non-nil, zero value otherwise.

### GetAutoWarmupPlanOk

`func (o *IP) GetAutoWarmupPlanOk() (*string, bool)`

GetAutoWarmupPlanOk returns a tuple with the AutoWarmupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWarmupPlan

`func (o *IP) SetAutoWarmupPlan(v string)`

SetAutoWarmupPlan sets AutoWarmupPlan field to given value.

### HasAutoWarmupPlan

`func (o *IP) HasAutoWarmupPlan() bool`

HasAutoWarmupPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


