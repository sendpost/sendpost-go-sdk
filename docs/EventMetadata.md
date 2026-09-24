# EventMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmtpCode** | Pointer to **int64** | SMTP response code from the receiving mail server. - 250: Success - 4xx: Temporary failure (soft bounce) - 5xx: Permanent failure (hard bounce)  | [optional] 
**SmtpDescription** | Pointer to **string** | Full SMTP response message from the receiving server. Useful for diagnosing delivery issues.  | [optional] 
**UserAgent** | Pointer to [**UserAgent**](UserAgent.md) | Parsed browser/email client information (for open/click events) | [optional] 
**Os** | Pointer to [**Os**](Os.md) | Parsed operating system information (for open/click events) | [optional] 
**Device** | Pointer to [**Device**](Device.md) | Device type information (for open/click events) | [optional] 
**Geo** | Pointer to [**GeoLocation**](GeoLocation.md) | Geographic location based on IP address (for open/click events) | [optional] 
**ClickedUrl** | Pointer to **string** | The original URL that was clicked (only for click events) | [optional] 
**TrackedIp** | Pointer to **string** | IP address of the user who triggered the event (open/click) | [optional] 
**RawUserAgent** | Pointer to **string** | Raw User-Agent header string from the HTTP request | [optional] 

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

`func (o *EventMetadata) GetSmtpCode() int64`

GetSmtpCode returns the SmtpCode field if non-nil, zero value otherwise.

### GetSmtpCodeOk

`func (o *EventMetadata) GetSmtpCodeOk() (*int64, bool)`

GetSmtpCodeOk returns a tuple with the SmtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpCode

`func (o *EventMetadata) SetSmtpCode(v int64)`

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

`func (o *EventMetadata) GetOs() Os`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *EventMetadata) GetOsOk() (*Os, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *EventMetadata) SetOs(v Os)`

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

### GetClickedUrl

`func (o *EventMetadata) GetClickedUrl() string`

GetClickedUrl returns the ClickedUrl field if non-nil, zero value otherwise.

### GetClickedUrlOk

`func (o *EventMetadata) GetClickedUrlOk() (*string, bool)`

GetClickedUrlOk returns a tuple with the ClickedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedUrl

`func (o *EventMetadata) SetClickedUrl(v string)`

SetClickedUrl sets ClickedUrl field to given value.

### HasClickedUrl

`func (o *EventMetadata) HasClickedUrl() bool`

HasClickedUrl returns a boolean if a field has been set.

### GetTrackedIp

`func (o *EventMetadata) GetTrackedIp() string`

GetTrackedIp returns the TrackedIp field if non-nil, zero value otherwise.

### GetTrackedIpOk

`func (o *EventMetadata) GetTrackedIpOk() (*string, bool)`

GetTrackedIpOk returns a tuple with the TrackedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedIp

`func (o *EventMetadata) SetTrackedIp(v string)`

SetTrackedIp sets TrackedIp field to given value.

### HasTrackedIp

`func (o *EventMetadata) HasTrackedIp() bool`

HasTrackedIp returns a boolean if a field has been set.

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


