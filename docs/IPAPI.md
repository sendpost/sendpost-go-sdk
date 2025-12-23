# \IPAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateNewIp**](IPAPI.md#AllocateNewIp) | **Put** /account/ip/allocate | Allocate IP
[**DeleteIp**](IPAPI.md#DeleteIp) | **Delete** /account/ip/{ip_id} | Delete IP
[**GetAllIps**](IPAPI.md#GetAllIps) | **Get** /account/ip/ | List IPs
[**GetSpecificIp**](IPAPI.md#GetSpecificIp) | **Get** /account/ip/{ip_id} | Get IP
[**UpdateIp**](IPAPI.md#UpdateIp) | **Put** /account/ip/{ip_id} | Update IP



## AllocateNewIp

> IP AllocateNewIp(ctx).IPAllocationRequest(iPAllocationRequest).Execute()

Allocate IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk"
)

func main() {
	iPAllocationRequest := *openapiclient.NewIPAllocationRequest(true, []string{"34.21.14.11"}) // IPAllocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAPI.AllocateNewIp(context.Background()).IPAllocationRequest(iPAllocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAPI.AllocateNewIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllocateNewIp`: IP
	fmt.Fprintf(os.Stdout, "Response from `IPAPI.AllocateNewIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllocateNewIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iPAllocationRequest** | [**IPAllocationRequest**](IPAllocationRequest.md) |  | 

### Return type

[**IP**](IP.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIp

> IPDeletionResponse DeleteIp(ctx, ipId).Execute()

Delete IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk"
)

func main() {
	ipId := int32(56) // int32 | The ID of the IP resource to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAPI.DeleteIp(context.Background(), ipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAPI.DeleteIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIp`: IPDeletionResponse
	fmt.Fprintf(os.Stdout, "Response from `IPAPI.DeleteIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipId** | **int32** | The ID of the IP resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IPDeletionResponse**](IPDeletionResponse.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIps

> []IP GetAllIps(ctx).Limit(limit).Offset(offset).Search(search).Execute()

List IPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk"
)

func main() {
	limit := int32(56) // int32 | Number of records to return per request (optional)
	offset := int32(56) // int32 | Number of initial records to skip (optional)
	search := "search_example" // string | Case insensitive search against IP's public IP address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAPI.GetAllIps(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAPI.GetAllIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllIps`: []IP
	fmt.Fprintf(os.Stdout, "Response from `IPAPI.GetAllIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of records to return per request | 
 **offset** | **int32** | Number of initial records to skip | 
 **search** | **string** | Case insensitive search against IP&#39;s public IP address | 

### Return type

[**[]IP**](IP.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificIp

> IP GetSpecificIp(ctx, ipId).Execute()

Get IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk"
)

func main() {
	ipId := int32(56) // int32 | The ID of the IP resource to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAPI.GetSpecificIp(context.Background(), ipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAPI.GetSpecificIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecificIp`: IP
	fmt.Fprintf(os.Stdout, "Response from `IPAPI.GetSpecificIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipId** | **int32** | The ID of the IP resource to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IP**](IP.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIp

> IP UpdateIp(ctx, ipId).IPUpdateRequest(iPUpdateRequest).Execute()

Update IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk"
)

func main() {
	iPUpdateRequest := *openapiclient.NewIPUpdateRequest(false) // IPUpdateRequest | 
	ipId := int32(56) // int32 | The ID of the IP resource to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAPI.UpdateIp(context.Background(), ipId).IPUpdateRequest(iPUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAPI.UpdateIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIp`: IP
	fmt.Fprintf(os.Stdout, "Response from `IPAPI.UpdateIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipId** | **int32** | The ID of the IP resource to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iPUpdateRequest** | [**IPUpdateRequest**](IPUpdateRequest.md) |  | 


### Return type

[**IP**](IP.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

