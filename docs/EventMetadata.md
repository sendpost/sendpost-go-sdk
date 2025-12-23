# EventMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmtpCode** | Pointer to **int32** |  | [optional] 
**SmtpDescription** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to [**UserAgent**](UserAgent.md) |  | [optional] 
**Os** | Pointer to [**OperatingSystem**](OperatingSystem.md) |  | [optional] 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**Geo** | Pointer to [**GeoLocation**](GeoLocation.md) |  | [optional] 
**ClickedURL** | Pointer to **string** |  | [optional] 
**TrackedIP** | Pointer to **string** |  | [optional] 
**RawUserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewEventMetadata

`func NewEventMetadata() *EventMetadata`

NewEventMetadata instantiates a new EventMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMetadataWithDefaults

`func NewEventMetadataWithDefaults() *EventMetadata`

NewEventMetadataWithDefaults instantiates a new EventMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmtpCode

`func (o *EventMetadata) GetSmtpCode() int32`

GetSmtpCode returns the SmtpCode field if non-nil, zero value otherwise.

### GetSmtpCodeOk

`func (o *EventMetadata) GetSmtpCodeOk() (*int32, bool)`

GetSmtpCodeOk returns a tuple with the SmtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpCode

`func (o *EventMetadata) SetSmtpCode(v int32)`

SetSmtpCode sets SmtpCode field to given value.

### HasSmtpCode

`func (o *EventMetadata) HasSmtpCode() bool`

HasSmtpCode returns a boolean if a field has been set.

### GetSmtpDescription

`func (o *EventMetadata) GetSmtpDescription() string`

GetSmtpDescription returns the SmtpDescription field if non-nil, zero value otherwise.

### GetSmtpDescriptionOk

`func (o *EventMetadata) GetSmtpDescriptionOk() (*string, bool)`

GetSmtpDescriptionOk returns a tuple with the SmtpDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpDescription

`func (o *EventMetadata) SetSmtpDescription(v string)`

SetSmtpDescription sets SmtpDescription field to given value.

### HasSmtpDescription

`func (o *EventMetadata) HasSmtpDescription() bool`

HasSmtpDescription returns a boolean if a field has been set.

### GetUserAgent

`func (o *EventMetadata) GetUserAgent() UserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *EventMetadata) GetUserAgentOk() (*UserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *EventMetadata) SetUserAgent(v UserAgent)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *EventMetadata) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetOs

`func (o *EventMetadata) GetOs() OperatingSystem`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *EventMetadata) GetOsOk() (*OperatingSystem, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *EventMetadata) SetOs(v OperatingSystem)`

SetOs sets Os field to given value.

### HasOs

`func (o *EventMetadata) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetDevice

`func (o *EventMetadata) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *EventMetadata) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *EventMetadata) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *EventMetadata) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetGeo

`func (o *EventMetadata) GetGeo() GeoLocation`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *EventMetadata) GetGeoOk() (*GeoLocation, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *EventMetadata) SetGeo(v GeoLocation)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *EventMetadata) HasGeo() bool`

HasGeo returns a boolean if a field has been set.

### GetClickedURL

`func (o *EventMetadata) GetClickedURL() string`

GetClickedURL returns the ClickedURL field if non-nil, zero value otherwise.

### GetClickedURLOk

`func (o *EventMetadata) GetClickedURLOk() (*string, bool)`

GetClickedURLOk returns a tuple with the ClickedURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedURL

`func (o *EventMetadata) SetClickedURL(v string)`

SetClickedURL sets ClickedURL field to given value.

### HasClickedURL

`func (o *EventMetadata) HasClickedURL() bool`

HasClickedURL returns a boolean if a field has been set.

### GetTrackedIP

`func (o *EventMetadata) GetTrackedIP() string`

GetTrackedIP returns the TrackedIP field if non-nil, zero value otherwise.

### GetTrackedIPOk

`func (o *EventMetadata) GetTrackedIPOk() (*string, bool)`

GetTrackedIPOk returns a tuple with the TrackedIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedIP

`func (o *EventMetadata) SetTrackedIP(v string)`

SetTrackedIP sets TrackedIP field to given value.

### HasTrackedIP

`func (o *EventMetadata) HasTrackedIP() bool`

HasTrackedIP returns a boolean if a field has been set.

### GetRawUserAgent

`func (o *EventMetadata) GetRawUserAgent() string`

GetRawUserAgent returns the RawUserAgent field if non-nil, zero value otherwise.

### GetRawUserAgentOk

`func (o *EventMetadata) GetRawUserAgentOk() (*string, bool)`

GetRawUserAgentOk returns a tuple with the RawUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUserAgent

`func (o *EventMetadata) SetRawUserAgent(v string)`

SetRawUserAgent sets RawUserAgent field to given value.

### HasRawUserAgent

`func (o *EventMetadata) HasRawUserAgent() bool`

HasRawUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


