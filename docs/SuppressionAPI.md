# \SuppressionAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSuppression**](SuppressionAPI.md#CreateSuppression) | **Post** /subaccount/suppression | Create Suppressions
[**DeleteSuppression**](SuppressionAPI.md#DeleteSuppression) | **Delete** /subaccount/suppression | Delete Suppressions
[**GetSuppressionList**](SuppressionAPI.md#GetSuppressionList) | **Get** /subaccount/suppression | List Suppressions



## CreateSuppression

> []Suppression CreateSuppression(ctx).CreateSuppressionRequest(createSuppressionRequest).Execute()

Create Suppressions



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
	createSuppressionRequest := *openapiclient.NewCreateSuppressionRequest() // CreateSuppressionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionAPI.CreateSuppression(context.Background()).CreateSuppressionRequest(createSuppressionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionAPI.CreateSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSuppression`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionAPI.CreateSuppression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSuppressionRequest** | [**CreateSuppressionRequest**](CreateSuppressionRequest.md) |  | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSuppression

> []DeleteSuppression200ResponseInner DeleteSuppression(ctx).DeleteSuppressionRequest(deleteSuppressionRequest).Execute()

Delete Suppressions



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
	deleteSuppressionRequest := *openapiclient.NewDeleteSuppressionRequest() // DeleteSuppressionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionAPI.DeleteSuppression(context.Background()).DeleteSuppressionRequest(deleteSuppressionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionAPI.DeleteSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSuppression`: []DeleteSuppression200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SuppressionAPI.DeleteSuppression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSuppressionRequest** | [**DeleteSuppressionRequest**](DeleteSuppressionRequest.md) |  | 

### Return type

[**[]DeleteSuppression200ResponseInner**](DeleteSuppression200ResponseInner.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressionList

> []Suppression GetSuppressionList(ctx).From(from).To(to).Limit(limit).Offset(offset).Search(search).Type_(type_).Execute()

List Suppressions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/sendpost/sendpost-go-sdk"
)

func main() {
	from := time.Now() // string | Start date for the suppression records
	to := time.Now() // string | End date for the suppression records (Note: `from` should be earlier than `to` and the date range should not exceed 60 days) 
	limit := int32(56) // int32 | Number of records to return per request (optional) (default to 20)
	offset := int32(56) // int32 | Number of initial records to skip (optional) (default to 0)
	search := "search_example" // string | Case-insensitive search against suppression email (optional)
	type_ := "type__example" // string | Type of suppression. Valid values: `hardBounce`, `manual`, `spamComplaint`, `unsubscribe`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionAPI.GetSuppressionList(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Search(search).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionAPI.GetSuppressionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuppressionList`: []Suppression
	fmt.Fprintf(os.Stdout, "Response from `SuppressionAPI.GetSuppressionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Start date for the suppression records | 
 **to** | **string** | End date for the suppression records (Note: &#x60;from&#x60; should be earlier than &#x60;to&#x60; and the date range should not exceed 60 days)  | 
 **limit** | **int32** | Number of records to return per request | [default to 20]
 **offset** | **int32** | Number of initial records to skip | [default to 0]
 **search** | **string** | Case-insensitive search against suppression email | 
 **type_** | **string** | Type of suppression. Valid values: &#x60;hardBounce&#x60;, &#x60;manual&#x60;, &#x60;spamComplaint&#x60;, &#x60;unsubscribe&#x60;  | 

### Return type

[**[]Suppression**](Suppression.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

