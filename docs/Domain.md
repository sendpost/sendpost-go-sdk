# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the domain | [optional] 
**Name** | Pointer to **string** | The domain name (e.g., \&quot;example.com\&quot;). This is the domain portion of your sending email addresses.  | [optional] 
**DnsProvider** | Pointer to **string** | Auto-detected DNS provider for this domain (e.g. \&quot;cloudflare\&quot;, \&quot;other\&quot;), used to tailor DNS-setup instructions. Read-only.  | [optional] 
**Dkim** | Pointer to [**DnsRecord**](DnsRecord.md) | DKIM (DomainKeys Identified Mail) DNS record configuration. DKIM cryptographically signs your emails to verify they haven&#39;t been tampered with. This is REQUIRED for sending emails.  | [optional] 
**ReturnPath** | Pointer to [**DnsRecord**](DnsRecord.md) | Return-Path (bounce handling) DNS record configuration. Configuring this allows bounce notifications to be properly routed through SendPost. RECOMMENDED for better deliverability.  | [optional] 
**Track** | Pointer to [**DnsRecord**](DnsRecord.md) | Tracking domain DNS record configuration. When configured, click tracking links use your domain instead of SendPost&#39;s domain. RECOMMENDED for brand consistency and improved click-through rates.  | [optional] 
**Dmarc** | Pointer to [**DnsRecord**](DnsRecord.md) | DMARC (Domain-based Message Authentication, Reporting &amp; Conformance) DNS record. DMARC builds on DKIM and SPF to provide email authentication and reporting. RECOMMENDED for enterprise senders.  | [optional] 
**DkimVerified** | Pointer to **bool** | Whether the DKIM DNS record has been verified successfully | [optional] 
**DmarcVerified** | Pointer to **bool** | Whether the DMARC DNS record has been verified successfully | [optional] 
**ReturnPathVerified** | Pointer to **bool** | Whether the Return-Path DNS record has been verified successfully | [optional] 
**TrackVerified** | Pointer to **bool** | Whether the tracking domain DNS record has been verified successfully | [optional] 
**Verified** | Pointer to **bool** | Overall verification status. True only if DKIM is verified (minimum requirement). For full verification, configure all DNS records.  | [optional] 
**DomainRegisteredDate** | Pointer to **string** | Date when this domain was originally registered (from WHOIS). Newer domains may have lower sender reputation initially.  | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds when the domain was added to SendPost | [optional] 
**DkimFailureReason** | Pointer to **string** | Detailed reason if DKIM verification failed (empty if verified or not attempted) | [optional] 
**DmarcFailureReason** | Pointer to **string** | Detailed reason if DMARC verification failed | [optional] 
**TrackFailureReason** | Pointer to **string** | Detailed reason if tracking domain verification failed | [optional] 
**ReturnPathFailureReason** | Pointer to **string** | Detailed reason if Return-Path verification failed | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Domain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDnsProvider

`func (o *Domain) GetDnsProvider() string`

GetDnsProvider returns the DnsProvider field if non-nil, zero value otherwise.

### GetDnsProviderOk

`func (o *Domain) GetDnsProviderOk() (*string, bool)`

GetDnsProviderOk returns a tuple with the DnsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProvider

`func (o *Domain) SetDnsProvider(v string)`

SetDnsProvider sets DnsProvider field to given value.

### HasDnsProvider

`func (o *Domain) HasDnsProvider() bool`

HasDnsProvider returns a boolean if a field has been set.

### GetDkim

`func (o *Domain) GetDkim() DnsRecord`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *Domain) GetDkimOk() (*DnsRecord, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *Domain) SetDkim(v DnsRecord)`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *Domain) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetReturnPath

`func (o *Domain) GetReturnPath() DnsRecord`

GetReturnPath returns the ReturnPath field if non-nil, zero value otherwise.

### GetReturnPathOk

`func (o *Domain) GetReturnPathOk() (*DnsRecord, bool)`

GetReturnPathOk returns a tuple with the ReturnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPath

`func (o *Domain) SetReturnPath(v DnsRecord)`

SetReturnPath sets ReturnPath field to given value.

### HasReturnPath

`func (o *Domain) HasReturnPath() bool`

HasReturnPath returns a boolean if a field has been set.

### GetTrack

`func (o *Domain) GetTrack() DnsRecord`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *Domain) GetTrackOk() (*DnsRecord, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *Domain) SetTrack(v DnsRecord)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *Domain) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetDmarc

`func (o *Domain) GetDmarc() DnsRecord`

GetDmarc returns the Dmarc field if non-nil, zero value otherwise.

### GetDmarcOk

`func (o *Domain) GetDmarcOk() (*DnsRecord, bool)`

GetDmarcOk returns a tuple with the Dmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarc

`func (o *Domain) SetDmarc(v DnsRecord)`

SetDmarc sets Dmarc field to given value.

### HasDmarc

`func (o *Domain) HasDmarc() bool`

HasDmarc returns a boolean if a field has been set.

### GetDkimVerified

`func (o *Domain) GetDkimVerified() bool`

GetDkimVerified returns the DkimVerified field if non-nil, zero value otherwise.

### GetDkimVerifiedOk

`func (o *Domain) GetDkimVerifiedOk() (*bool, bool)`

GetDkimVerifiedOk returns a tuple with the DkimVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimVerified

`func (o *Domain) SetDkimVerified(v bool)`

SetDkimVerified sets DkimVerified field to given value.

### HasDkimVerified

`func (o *Domain) HasDkimVerified() bool`

HasDkimVerified returns a boolean if a field has been set.

### GetDmarcVerified

`func (o *Domain) GetDmarcVerified() bool`

GetDmarcVerified returns the DmarcVerified field if non-nil, zero value otherwise.

### GetDmarcVerifiedOk

`func (o *Domain) GetDmarcVerifiedOk() (*bool, bool)`

GetDmarcVerifiedOk returns a tuple with the DmarcVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarcVerified

`func (o *Domain) SetDmarcVerified(v bool)`

SetDmarcVerified sets DmarcVerified field to given value.

### HasDmarcVerified

`func (o *Domain) HasDmarcVerified() bool`

HasDmarcVerified returns a boolean if a field has been set.

### GetReturnPathVerified

`func (o *Domain) GetReturnPathVerified() bool`

GetReturnPathVerified returns the ReturnPathVerified field if non-nil, zero value otherwise.

### GetReturnPathVerifiedOk

`func (o *Domain) GetReturnPathVerifiedOk() (*bool, bool)`

GetReturnPathVerifiedOk returns a tuple with the ReturnPathVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPathVerified

`func (o *Domain) SetReturnPathVerified(v bool)`

SetReturnPathVerified sets ReturnPathVerified field to given value.

### HasReturnPathVerified

`func (o *Domain) HasReturnPathVerified() bool`

HasReturnPathVerified returns a boolean if a field has been set.

### GetTrackVerified

`func (o *Domain) GetTrackVerified() bool`

GetTrackVerified returns the TrackVerified field if non-nil, zero value otherwise.

### GetTrackVerifiedOk

`func (o *Domain) GetTrackVerifiedOk() (*bool, bool)`

GetTrackVerifiedOk returns a tuple with the TrackVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackVerified

`func (o *Domain) SetTrackVerified(v bool)`

SetTrackVerified sets TrackVerified field to given value.

### HasTrackVerified

`func (o *Domain) HasTrackVerified() bool`

HasTrackVerified returns a boolean if a field has been set.

### GetVerified

`func (o *Domain) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Domain) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Domain) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *Domain) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetDomainRegisteredDate

`func (o *Domain) GetDomainRegisteredDate() string`

GetDomainRegisteredDate returns the DomainRegisteredDate field if non-nil, zero value otherwise.

### GetDomainRegisteredDateOk

`func (o *Domain) GetDomainRegisteredDateOk() (*string, bool)`

GetDomainRegisteredDateOk returns a tuple with the DomainRegisteredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRegisteredDate

`func (o *Domain) SetDomainRegisteredDate(v string)`

SetDomainRegisteredDate sets DomainRegisteredDate field to given value.

### HasDomainRegisteredDate

`func (o *Domain) HasDomainRegisteredDate() bool`

HasDomainRegisteredDate returns a boolean if a field has been set.

### GetCreated

`func (o *Domain) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Domain) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Domain) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Domain) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDkimFailureReason

`func (o *Domain) GetDkimFailureReason() string`

GetDkimFailureReason returns the DkimFailureReason field if non-nil, zero value otherwise.

### GetDkimFailureReasonOk

`func (o *Domain) GetDkimFailureReasonOk() (*string, bool)`

GetDkimFailureReasonOk returns a tuple with the DkimFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimFailureReason

`func (o *Domain) SetDkimFailureReason(v string)`

SetDkimFailureReason sets DkimFailureReason field to given value.

### HasDkimFailureReason

`func (o *Domain) HasDkimFailureReason() bool`

HasDkimFailureReason returns a boolean if a field has been set.

### GetDmarcFailureReason

`func (o *Domain) GetDmarcFailureReason() string`

GetDmarcFailureReason returns the DmarcFailureReason field if non-nil, zero value otherwise.

### GetDmarcFailureReasonOk

`func (o *Domain) GetDmarcFailureReasonOk() (*string, bool)`

GetDmarcFailureReasonOk returns a tuple with the DmarcFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarcFailureReason

`func (o *Domain) SetDmarcFailureReason(v string)`

SetDmarcFailureReason sets DmarcFailureReason field to given value.

### HasDmarcFailureReason

`func (o *Domain) HasDmarcFailureReason() bool`

HasDmarcFailureReason returns a boolean if a field has been set.

### GetTrackFailureReason

`func (o *Domain) GetTrackFailureReason() string`

GetTrackFailureReason returns the TrackFailureReason field if non-nil, zero value otherwise.

### GetTrackFailureReasonOk

`func (o *Domain) GetTrackFailureReasonOk() (*string, bool)`

GetTrackFailureReasonOk returns a tuple with the TrackFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFailureReason

`func (o *Domain) SetTrackFailureReason(v string)`

SetTrackFailureReason sets TrackFailureReason field to given value.

### HasTrackFailureReason

`func (o *Domain) HasTrackFailureReason() bool`

HasTrackFailureReason returns a boolean if a field has been set.

### GetReturnPathFailureReason

`func (o *Domain) GetReturnPathFailureReason() string`

GetReturnPathFailureReason returns the ReturnPathFailureReason field if non-nil, zero value otherwise.

### GetReturnPathFailureReasonOk

`func (o *Domain) GetReturnPathFailureReasonOk() (*string, bool)`

GetReturnPathFailureReasonOk returns a tuple with the ReturnPathFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPathFailureReason

`func (o *Domain) SetReturnPathFailureReason(v string)`

SetReturnPathFailureReason sets ReturnPathFailureReason field to given value.

### HasReturnPathFailureReason

`func (o *Domain) HasReturnPathFailureReason() bool`

HasReturnPathFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


