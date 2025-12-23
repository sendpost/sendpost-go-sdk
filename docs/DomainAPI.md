# \DomainAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllDomains**](DomainAPI.md#GetAllDomains) | **Get** /subaccount/domain | List Domains
[**SubaccountDomainDomainIdDelete**](DomainAPI.md#SubaccountDomainDomainIdDelete) | **Delete** /subaccount/domain/{domain_id} | Delete Domain
[**SubaccountDomainDomainIdGet**](DomainAPI.md#SubaccountDomainDomainIdGet) | **Get** /subaccount/domain/{domain_id} | Get Domain
[**SubaccountDomainPost**](DomainAPI.md#SubaccountDomainPost) | **Post** /subaccount/domain | Create Domain



## GetAllDomains

> []Domain GetAllDomains(ctx).Limit(limit).Offset(offset).Search(search).Execute()

List Domains



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
	search := "search_example" // string | Case insensitive search against domain names (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.GetAllDomains(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetAllDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDomains`: []Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetAllDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of records to return per request | 
 **offset** | **int32** | Number of initial records to skip | 
 **search** | **string** | Case insensitive search against domain names | 

### Return type

[**[]Domain**](Domain.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountDomainDomainIdDelete

> DeleteResponse SubaccountDomainDomainIdDelete(ctx, domainId).Execute()

Delete Domain



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
	domainId := "domainId_example" // string | The unique ID of the domain to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.SubaccountDomainDomainIdDelete(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.SubaccountDomainDomainIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountDomainDomainIdDelete`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.SubaccountDomainDomainIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The unique ID of the domain to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountDomainDomainIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountDomainDomainIdGet

> Domain SubaccountDomainDomainIdGet(ctx, domainId).Execute()

Get Domain



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
	domainId := "domainId_example" // string | The unique ID of the domain to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.SubaccountDomainDomainIdGet(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.SubaccountDomainDomainIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountDomainDomainIdGet`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.SubaccountDomainDomainIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The unique ID of the domain to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountDomainDomainIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountDomainPost

> Domain SubaccountDomainPost(ctx).CreateDomainRequest(createDomainRequest).Execute()

Create Domain



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
	createDomainRequest := *openapiclient.NewCreateDomainRequest() // CreateDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.SubaccountDomainPost(context.Background()).CreateDomainRequest(createDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.SubaccountDomainPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountDomainPost`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.SubaccountDomainPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountDomainPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainRequest** | [**CreateDomainRequest**](CreateDomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

