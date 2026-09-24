# DnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The DNS hostname where this record should be created | [optional] 
**Type** | Pointer to **string** | The DNS record type (TXT or CNAME) | [optional] 
**TextValue** | Pointer to **string** | The value to set for this DNS record | [optional] 

## Methods

### NewDnsRecord

`func NewDnsRecord() *DnsRecord`

NewDnsRecord instantiates a new DnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordWithDefaults

`func NewDnsRecordWithDefaults() *DnsRecord`

NewDnsRecordWithDefaults instantiates a new DnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DnsRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DnsRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DnsRecord) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DnsRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetType

`func (o *DnsRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DnsRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTextValue

`func (o *DnsRecord) GetTextValue() string`

GetTextValue returns the TextValue field if non-nil, zero value otherwise.

### GetTextValueOk

`func (o *DnsRecord) GetTextValueOk() (*string, bool)`

GetTextValueOk returns a tuple with the TextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextValue

`func (o *DnsRecord) SetTextValue(v string)`

SetTextValue sets TextValue field to given value.

### HasTextValue

`func (o *DnsRecord) HasTextValue() bool`

HasTextValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


