# GeoLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CityId** | Pointer to **int32** | GeoNames city identifier | [optional] 
**CountryCode** | Pointer to **string** | Two-letter ISO 3166-1 alpha-2 country code | [optional] 
**ContinentCode** | Pointer to **string** | Two-letter continent code: AF (Africa), AN (Antarctica), AS (Asia), EU (Europe), NA (North America), OC (Oceania), SA (South America)  | [optional] 
**PostalCode** | Pointer to **string** | Postal/ZIP code | [optional] 
**TimeZone** | Pointer to **string** | IANA timezone identifier | [optional] 

## Methods

### NewGeoLocation

`func NewGeoLocation() *GeoLocation`

NewGeoLocation instantiates a new GeoLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoLocationWithDefaults

`func NewGeoLocationWithDefaults() *GeoLocation`

NewGeoLocationWithDefaults instantiates a new GeoLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCityId

`func (o *GeoLocation) GetCityId() int32`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *GeoLocation) GetCityIdOk() (*int32, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *GeoLocation) SetCityId(v int32)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *GeoLocation) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### GetCountryCode

`func (o *GeoLocation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GeoLocation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GeoLocation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GeoLocation) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetContinentCode

`func (o *GeoLocation) GetContinentCode() string`

GetContinentCode returns the ContinentCode field if non-nil, zero value otherwise.

### GetContinentCodeOk

`func (o *GeoLocation) GetContinentCodeOk() (*string, bool)`

GetContinentCodeOk returns a tuple with the ContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentCode

`func (o *GeoLocation) SetContinentCode(v string)`

SetContinentCode sets ContinentCode field to given value.

### HasContinentCode

`func (o *GeoLocation) HasContinentCode() bool`

HasContinentCode returns a boolean if a field has been set.

### GetPostalCode

`func (o *GeoLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *GeoLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *GeoLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *GeoLocation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetTimeZone

`func (o *GeoLocation) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GeoLocation) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GeoLocation) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GeoLocation) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


