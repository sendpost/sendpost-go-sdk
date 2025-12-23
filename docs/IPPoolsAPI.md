# \IPPoolsAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPPool**](IPPoolsAPI.md#CreateIPPool) | **Post** /account/ippool | Create IPPool
[**DeleteIPPool**](IPPoolsAPI.md#DeleteIPPool) | **Delete** /account/ippool/{ippool_id} | Delete IPPool
[**GetAllIPPools**](IPPoolsAPI.md#GetAllIPPools) | **Get** /account/ippool | List IPPools
[**GetIPPoolById**](IPPoolsAPI.md#GetIPPoolById) | **Get** /account/ippool/{ippool_id} | Get IPPool
[**UpdateIPPool**](IPPoolsAPI.md#UpdateIPPool) | **Put** /account/ippool/{ippool_id} | Update IPPool



## CreateIPPool

> IPPool CreateIPPool(ctx).IPPoolCreateRequest(iPPoolCreateRequest).Execute()

Create IPPool



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
	iPPoolCreateRequest := *openapiclient.NewIPPoolCreateRequest() // IPPoolCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPPoolsAPI.CreateIPPool(context.Background()).IPPoolCreateRequest(iPPoolCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPPoolsAPI.CreateIPPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIPPool`: IPPool
	fmt.Fprintf(os.Stdout, "Response from `IPPoolsAPI.CreateIPPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iPPoolCreateRequest** | [**IPPoolCreateRequest**](IPPoolCreateRequest.md) |  | 

### Return type

[**IPPool**](IPPool.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIPPool

> IPPoolDeleteResponse DeleteIPPool(ctx, ippoolId).Execute()

Delete IPPool



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
	ippoolId := int32(756) // int32 | The ID of the IPPool to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPPoolsAPI.DeleteIPPool(context.Background(), ippoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPPoolsAPI.DeleteIPPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIPPool`: IPPoolDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `IPPoolsAPI.DeleteIPPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **int32** | The ID of the IPPool to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IPPoolDeleteResponse**](IPPoolDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIPPools

> []IPPool GetAllIPPools(ctx).Limit(limit).Offset(offset).Search(search).Execute()

List IPPools



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
	limit := int32(10) // int32 | Number of records to return per request (optional)
	offset := int32(0) // int32 | Number of initial records to skip (optional)
	search := "Transactional" // string | Case insensitive search against IPPool name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPPoolsAPI.GetAllIPPools(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPPoolsAPI.GetAllIPPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllIPPools`: []IPPool
	fmt.Fprintf(os.Stdout, "Response from `IPPoolsAPI.GetAllIPPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIPPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of records to return per request | 
 **offset** | **int32** | Number of initial records to skip | 
 **search** | **string** | Case insensitive search against IPPool name | 

### Return type

[**[]IPPool**](IPPool.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPPoolById

> IPPool GetIPPoolById(ctx, ippoolId).Execute()

Get IPPool



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
	ippoolId := int32(74) // int32 | The ID of the IPPool whose information you want to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPPoolsAPI.GetIPPoolById(context.Background(), ippoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPPoolsAPI.GetIPPoolById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPPoolById`: IPPool
	fmt.Fprintf(os.Stdout, "Response from `IPPoolsAPI.GetIPPoolById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **int32** | The ID of the IPPool whose information you want to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPPoolByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IPPool**](IPPool.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIPPool

> IPPool UpdateIPPool(ctx, ippoolId).IPPoolUpdateRequest(iPPoolUpdateRequest).Execute()

Update IPPool



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
	iPPoolUpdateRequest := *openapiclient.NewIPPoolUpdateRequest() // IPPoolUpdateRequest | 
	ippoolId := int32(756) // int32 | The ID of the IPPool to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPPoolsAPI.UpdateIPPool(context.Background(), ippoolId).IPPoolUpdateRequest(iPPoolUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPPoolsAPI.UpdateIPPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIPPool`: IPPool
	fmt.Fprintf(os.Stdout, "Response from `IPPoolsAPI.UpdateIPPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ippoolId** | **int32** | The ID of the IPPool to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iPPoolUpdateRequest** | [**IPPoolUpdateRequest**](IPPoolUpdateRequest.md) |  | 


### Return type

[**IPPool**](IPPool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

