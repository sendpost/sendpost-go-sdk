# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the domain. | [optional] 
**Name** | Pointer to **string** | Name of the domain. | [optional] 
**Dkim** | Pointer to [**DomainDkim**](DomainDkim.md) |  | [optional] 
**ReturnPath** | Pointer to [**DomainReturnPath**](DomainReturnPath.md) |  | [optional] 
**Track** | Pointer to [**DomainTrack**](DomainTrack.md) |  | [optional] 
**Dmarc** | Pointer to [**DomainDmarc**](DomainDmarc.md) |  | [optional] 
**DkimConfig** | Pointer to **string** | DKIM configuration | [optional] 
**DkimVerified** | Pointer to **bool** | Status of DKIM verification ( true or false ) | [optional] 
**DmarcVerified** | Pointer to **bool** | Status of DMARC verification ( true or false) | [optional] 
**ReturnPathVerified** | Pointer to **bool** | Status of ReturnPath verification ( true or false ) | [optional] 
**TrackVerified** | Pointer to **bool** | Status of Track verification ( true or false ) | [optional] 
**Verified** | Pointer to **bool** | Overall verification status of the domain | [optional] 
**DomainRegisteredDate** | Pointer to **string** | Date when the domain was registered | [optional] 
**Created** | Pointer to **int64** | UNIX epoch timestamp in nanoseconds. | [optional] 
**GptVerified** | Pointer to **bool** | Status of GPT verification ( true or false ) | [optional] 
**Gpt** | Pointer to [**DomainGpt**](DomainGpt.md) |  | [optional] 
**DmarcFailureReason** | Pointer to **string** | Reason for DMARC verification failure | [optional] 
**DkimFailureReason** | Pointer to **string** | Reason for DKIM verification failure | [optional] 
**TrackFailureReason** | Pointer to **string** | Reason for Track verification failure | [optional] 
**ReturnPathFailureReason** | Pointer to **string** | Reason for ReturnPath verification failure | [optional] 

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

`func (o *Domain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v int32)`

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

### GetDkim

`func (o *Domain) GetDkim() DomainDkim`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *Domain) GetDkimOk() (*DomainDkim, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *Domain) SetDkim(v DomainDkim)`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *Domain) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetReturnPath

`func (o *Domain) GetReturnPath() DomainReturnPath`

GetReturnPath returns the ReturnPath field if non-nil, zero value otherwise.

### GetReturnPathOk

`func (o *Domain) GetReturnPathOk() (*DomainReturnPath, bool)`

GetReturnPathOk returns a tuple with the ReturnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPath

`func (o *Domain) SetReturnPath(v DomainReturnPath)`

SetReturnPath sets ReturnPath field to given value.

### HasReturnPath

`func (o *Domain) HasReturnPath() bool`

HasReturnPath returns a boolean if a field has been set.

### GetTrack

`func (o *Domain) GetTrack() DomainTrack`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *Domain) GetTrackOk() (*DomainTrack, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *Domain) SetTrack(v DomainTrack)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *Domain) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetDmarc

`func (o *Domain) GetDmarc() DomainDmarc`

GetDmarc returns the Dmarc field if non-nil, zero value otherwise.

### GetDmarcOk

`func (o *Domain) GetDmarcOk() (*DomainDmarc, bool)`

GetDmarcOk returns a tuple with the Dmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarc

`func (o *Domain) SetDmarc(v DomainDmarc)`

SetDmarc sets Dmarc field to given value.

### HasDmarc

`func (o *Domain) HasDmarc() bool`

HasDmarc returns a boolean if a field has been set.

### GetDkimConfig

`func (o *Domain) GetDkimConfig() string`

GetDkimConfig returns the DkimConfig field if non-nil, zero value otherwise.

### GetDkimConfigOk

`func (o *Domain) GetDkimConfigOk() (*string, bool)`

GetDkimConfigOk returns a tuple with the DkimConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimConfig

`func (o *Domain) SetDkimConfig(v string)`

SetDkimConfig sets DkimConfig field to given value.

### HasDkimConfig

`func (o *Domain) HasDkimConfig() bool`

HasDkimConfig returns a boolean if a field has been set.

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

### GetGptVerified

`func (o *Domain) GetGptVerified() bool`

GetGptVerified returns the GptVerified field if non-nil, zero value otherwise.

### GetGptVerifiedOk

`func (o *Domain) GetGptVerifiedOk() (*bool, bool)`

GetGptVerifiedOk returns a tuple with the GptVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGptVerified

`func (o *Domain) SetGptVerified(v bool)`

SetGptVerified sets GptVerified field to given value.

### HasGptVerified

`func (o *Domain) HasGptVerified() bool`

HasGptVerified returns a boolean if a field has been set.

### GetGpt

`func (o *Domain) GetGpt() DomainGpt`

GetGpt returns the Gpt field if non-nil, zero value otherwise.

### GetGptOk

`func (o *Domain) GetGptOk() (*DomainGpt, bool)`

GetGptOk returns a tuple with the Gpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpt

`func (o *Domain) SetGpt(v DomainGpt)`

SetGpt sets Gpt field to given value.

### HasGpt

`func (o *Domain) HasGpt() bool`

HasGpt returns a boolean if a field has been set.

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


