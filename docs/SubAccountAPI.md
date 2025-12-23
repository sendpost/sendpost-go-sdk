# \SubAccountAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubAccount**](SubAccountAPI.md#CreateSubAccount) | **Post** /account/subaccount/ | Create Sub-Account
[**DeleteSubAccount**](SubAccountAPI.md#DeleteSubAccount) | **Delete** /account/subaccount/{subaccount_id} | Delete Sub-Account
[**GetAllSubAccounts**](SubAccountAPI.md#GetAllSubAccounts) | **Get** /account/subaccount/ | List Sub-Accounts
[**GetSubAccount**](SubAccountAPI.md#GetSubAccount) | **Get** /account/subaccount/{subaccount_id} | Get Sub-Account
[**UpdateSubAccount**](SubAccountAPI.md#UpdateSubAccount) | **Put** /account/subaccount/{subaccount_id} | Update Sub-Account



## CreateSubAccount

> SubAccount CreateSubAccount(ctx).CreateSubAccountRequest(createSubAccountRequest).Execute()

Create Sub-Account



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
	createSubAccountRequest := *openapiclient.NewCreateSubAccountRequest() // CreateSubAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccount(context.Background()).CreateSubAccountRequest(createSubAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccount`: SubAccount
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubAccountRequest** | [**CreateSubAccountRequest**](CreateSubAccountRequest.md) |  | 

### Return type

[**SubAccount**](SubAccount.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubAccount

> DeleteSubAccountResponse DeleteSubAccount(ctx, subaccountId).Execute()

Delete Sub-Account



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
	subaccountId := int32(12) // int32 | The ID of the sub-account to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.DeleteSubAccount(context.Background(), subaccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.DeleteSubAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSubAccount`: DeleteSubAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.DeleteSubAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subaccountId** | **int32** | The ID of the sub-account to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSubAccountResponse**](DeleteSubAccountResponse.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSubAccounts

> []SubAccount GetAllSubAccounts(ctx).Limit(limit).Offset(offset).Search(search).Execute()

List Sub-Accounts



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
	limit := int32(10) // int32 | Number of records to return per request. (optional)
	offset := int32(0) // int32 | Number of initial records to skip. (optional)
	search := "Hooli" // string | Case-insensitive search against the sub-account name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetAllSubAccounts(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetAllSubAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSubAccounts`: []SubAccount
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetAllSubAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSubAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of records to return per request. | 
 **offset** | **int32** | Number of initial records to skip. | 
 **search** | **string** | Case-insensitive search against the sub-account name. | 

### Return type

[**[]SubAccount**](SubAccount.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccount

> SubAccount GetSubAccount(ctx, subaccountId).Execute()

Get Sub-Account



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
	subaccountId := int32(11) // int32 | The ID of the sub-account to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccount(context.Background(), subaccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccount`: SubAccount
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subaccountId** | **int32** | The ID of the sub-account to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubAccount**](SubAccount.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubAccount

> SubAccount UpdateSubAccount(ctx, subaccountId).UpdateSubAccount(updateSubAccount).Execute()

Update Sub-Account



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
	updateSubAccount := *openapiclient.NewUpdateSubAccount() // UpdateSubAccount | 
	subaccountId := int32(12) // int32 | The ID of the sub-account to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.UpdateSubAccount(context.Background(), subaccountId).UpdateSubAccount(updateSubAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.UpdateSubAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubAccount`: SubAccount
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.UpdateSubAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subaccountId** | **int32** | The ID of the sub-account to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSubAccount** | [**UpdateSubAccount**](UpdateSubAccount.md) |  | 


### Return type

[**SubAccount**](SubAccount.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

