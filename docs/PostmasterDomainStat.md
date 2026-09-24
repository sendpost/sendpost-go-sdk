# PostmasterDomainStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The domain these Postmaster statistics belong to. | [optional] 
**Date** | Pointer to **string** | The date these statistics apply to (YYYY-MM-DD, UTC). | [optional] 
**DomainReputation** | Pointer to **string** | Google&#39;s reputation rating for the domain (e.g. &#x60;HIGH&#x60;, &#x60;MEDIUM&#x60;, &#x60;LOW&#x60;, &#x60;BAD&#x60;).  | [optional] 
**Spam** | Pointer to **string** | User-reported spam rate for the domain (fraction, string-encoded). | [optional] 
**DkimSuccess** | Pointer to **string** | Fraction of mail that passed DKIM authentication (string-encoded). | [optional] 
**SpfSuccess** | Pointer to **string** | Fraction of mail that passed SPF authentication (string-encoded). | [optional] 
**DmarcSuccess** | Pointer to **string** | Fraction of mail that passed DMARC authentication (string-encoded). | [optional] 

## Methods

### NewPostmasterDomainStat

`func NewPostmasterDomainStat() *PostmasterDomainStat`

NewPostmasterDomainStat instantiates a new PostmasterDomainStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostmasterDomainStatWithDefaults

`func NewPostmasterDomainStatWithDefaults() *PostmasterDomainStat`

NewPostmasterDomainStatWithDefaults instantiates a new PostmasterDomainStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostmasterDomainStat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostmasterDomainStat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostmasterDomainStat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostmasterDomainStat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDate

`func (o *PostmasterDomainStat) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PostmasterDomainStat) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PostmasterDomainStat) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PostmasterDomainStat) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDomainReputation

`func (o *PostmasterDomainStat) GetDomainReputation() string`

GetDomainReputation returns the DomainReputation field if non-nil, zero value otherwise.

### GetDomainReputationOk

`func (o *PostmasterDomainStat) GetDomainReputationOk() (*string, bool)`

GetDomainReputationOk returns a tuple with the DomainReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainReputation

`func (o *PostmasterDomainStat) SetDomainReputation(v string)`

SetDomainReputation sets DomainReputation field to given value.

### HasDomainReputation

`func (o *PostmasterDomainStat) HasDomainReputation() bool`

HasDomainReputation returns a boolean if a field has been set.

### GetSpam

`func (o *PostmasterDomainStat) GetSpam() string`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *PostmasterDomainStat) GetSpamOk() (*string, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *PostmasterDomainStat) SetSpam(v string)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *PostmasterDomainStat) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetDkimSuccess

`func (o *PostmasterDomainStat) GetDkimSuccess() string`

GetDkimSuccess returns the DkimSuccess field if non-nil, zero value otherwise.

### GetDkimSuccessOk

`func (o *PostmasterDomainStat) GetDkimSuccessOk() (*string, bool)`

GetDkimSuccessOk returns a tuple with the DkimSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimSuccess

`func (o *PostmasterDomainStat) SetDkimSuccess(v string)`

SetDkimSuccess sets DkimSuccess field to given value.

### HasDkimSuccess

`func (o *PostmasterDomainStat) HasDkimSuccess() bool`

HasDkimSuccess returns a boolean if a field has been set.

### GetSpfSuccess

`func (o *PostmasterDomainStat) GetSpfSuccess() string`

GetSpfSuccess returns the SpfSuccess field if non-nil, zero value otherwise.

### GetSpfSuccessOk

`func (o *PostmasterDomainStat) GetSpfSuccessOk() (*string, bool)`

GetSpfSuccessOk returns a tuple with the SpfSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpfSuccess

`func (o *PostmasterDomainStat) SetSpfSuccess(v string)`

SetSpfSuccess sets SpfSuccess field to given value.

### HasSpfSuccess

`func (o *PostmasterDomainStat) HasSpfSuccess() bool`

HasSpfSuccess returns a boolean if a field has been set.

### GetDmarcSuccess

`func (o *PostmasterDomainStat) GetDmarcSuccess() string`

GetDmarcSuccess returns the DmarcSuccess field if non-nil, zero value otherwise.

### GetDmarcSuccessOk

`func (o *PostmasterDomainStat) GetDmarcSuccessOk() (*string, bool)`

GetDmarcSuccessOk returns a tuple with the DmarcSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarcSuccess

`func (o *PostmasterDomainStat) SetDmarcSuccess(v string)`

SetDmarcSuccess sets DmarcSuccess field to given value.

### HasDmarcSuccess

`func (o *PostmasterDomainStat) HasDmarcSuccess() bool`

HasDmarcSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


