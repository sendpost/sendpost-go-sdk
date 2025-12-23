# \StatsAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountSubaccountStatSubaccountIdAggregateGet**](StatsAPI.md#AccountSubaccountStatSubaccountIdAggregateGet) | **Get** /account/subaccount/stat/{subaccount_id}/aggregate | Get Aggregate Stats
[**AccountSubaccountStatSubaccountIdGet**](StatsAPI.md#AccountSubaccountStatSubaccountIdGet) | **Get** /account/subaccount/stat/{subaccount_id} | List Stats
[**GetAggregateStatsByGroup**](StatsAPI.md#GetAggregateStatsByGroup) | **Get** /account/subaccount/stat/{subaccount_id}/group | Get Group Aggregate Stats



## AccountSubaccountStatSubaccountIdAggregateGet

> AggregateStat AccountSubaccountStatSubaccountIdAggregateGet(ctx, subaccountId).From(from).To(to).Execute()

Get Aggregate Stats



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
	from := time.Now() // string | Start date for stats retrieval.
	to := time.Now() // string | Date to which stats should be retrieved ( Note than from date should be earlier than to date. Also the difference between from and to date shouldn't ne more than 60 days ) 
	subaccountId := int64(11) // int64 | The ID of the subaccount to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.AccountSubaccountStatSubaccountIdAggregateGet(context.Background(), subaccountId).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.AccountSubaccountStatSubaccountIdAggregateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountSubaccountStatSubaccountIdAggregateGet`: AggregateStat
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.AccountSubaccountStatSubaccountIdAggregateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subaccountId** | **int64** | The ID of the subaccount to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountSubaccountStatSubaccountIdAggregateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Start date for stats retrieval. | 
 **to** | **string** | Date to which stats should be retrieved ( Note than from date should be earlier than to date. Also the difference between from and to date shouldn&#39;t ne more than 60 days )  | 


### Return type

[**AggregateStat**](AggregateStat.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountSubaccountStatSubaccountIdGet

> []Stat AccountSubaccountStatSubaccountIdGet(ctx, subaccountId).From(from).To(to).Execute()

List Stats



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
	from := time.Now() // string | Start date for stats retrieval.
	to := time.Now() // string | Date to which stats should be retrieved ( Note than from date should be earlier than to date. Also the difference between from and to date shouldn't ne more than 60 days ) 
	subaccountId := int64(11) // int64 | The ID of the subaccount to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.AccountSubaccountStatSubaccountIdGet(context.Background(), subaccountId).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.AccountSubaccountStatSubaccountIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountSubaccountStatSubaccountIdGet`: []Stat
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.AccountSubaccountStatSubaccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subaccountId** | **int64** | The ID of the subaccount to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountSubaccountStatSubaccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Start date for stats retrieval. | 
 **to** | **string** | Date to which stats should be retrieved ( Note than from date should be earlier than to date. Also the difference between from and to date shouldn&#39;t ne more than 60 days )  | 


### Return type

[**[]Stat**](Stat.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAggregateStatsByGroup

> AggregateStat GetAggregateStatsByGroup(ctx, subaccountId).Group(group).From(from).To(to).Execute()

Get Group Aggregate Stats



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
	group := "group_example" // string | Group whose aggregated stats need to be retrieved
	from := time.Now() // string | The starting date for the aggregated stats
	to := time.Now() // string | The ending date for the aggregated stats (Note: `from` should be earlier than `to` and the date range should not exceed 366 days) 
	subaccountId := int64(11) // int64 | The ID of the subaccount to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetAggregateStatsByGroup(context.Background(), subaccountId).Group(group).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetAggregateStatsByGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAggregateStatsByGroup`: AggregateStat
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetAggregateStatsByGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subaccountId** | **int64** | The ID of the subaccount to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregateStatsByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string** | Group whose aggregated stats need to be retrieved | 
 **from** | **string** | The starting date for the aggregated stats | 
 **to** | **string** | The ending date for the aggregated stats (Note: &#x60;from&#x60; should be earlier than &#x60;to&#x60; and the date range should not exceed 366 days)  | 


### Return type

[**AggregateStat**](AggregateStat.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

