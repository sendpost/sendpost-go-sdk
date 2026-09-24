# \DomainAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubAccountDomain**](DomainAPI.md#CreateSubAccountDomain) | **Post** /subaccount/domain | Create Domain
[**DeleteSubAccountDomain**](DomainAPI.md#DeleteSubAccountDomain) | **Delete** /subaccount/domain/{domain_id} | Delete Domain
[**GetAllDomains**](DomainAPI.md#GetAllDomains) | **Get** /subaccount/domain | List Domains
[**GetSubAccountDomain**](DomainAPI.md#GetSubAccountDomain) | **Get** /subaccount/domain/{domain_id} | Get Domain



## CreateSubAccountDomain

> Domain CreateSubAccountDomain(ctx).CreateDomainRequest(createDomainRequest).Execute()

Create Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	createDomainRequest := *openapiclient.NewCreateDomainRequest("piedpiper.com") // CreateDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.CreateSubAccountDomain(context.Background()).CreateDomainRequest(createDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.CreateSubAccountDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountDomain`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.CreateSubAccountDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountDomainRequest struct via the builder pattern


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


## DeleteSubAccountDomain

> DeleteResponse DeleteSubAccountDomain(ctx, domainId).Execute()

Delete Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	domainId := "domainId_example" // string | The unique ID of the domain to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.DeleteSubAccountDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DeleteSubAccountDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSubAccountDomain`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DeleteSubAccountDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The unique ID of the domain to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubAccountDomainRequest struct via the builder pattern


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
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	limit := int32(20) // int32 | Number of records to return per request. Default is 20. (optional) (default to 20)
	offset := int32(0) // int32 | Number of initial records to skip for pagination. (optional) (default to 0)
	search := "mycompany" // string | Case insensitive search against domain names. Useful for finding specific domains in large lists. (optional)

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
 **limit** | **int32** | Number of records to return per request. Default is 20. | [default to 20]
 **offset** | **int32** | Number of initial records to skip for pagination. | [default to 0]
 **search** | **string** | Case insensitive search against domain names. Useful for finding specific domains in large lists. | 

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


## GetSubAccountDomain

> Domain GetSubAccountDomain(ctx, domainId).Execute()

Get Domain



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	domainId := "domainId_example" // string | The unique ID of the domain to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.GetSubAccountDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetSubAccountDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountDomain`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetSubAccountDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The unique ID of the domain to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountDomainRequest struct via the builder pattern


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

