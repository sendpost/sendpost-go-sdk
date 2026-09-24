# \StatsAAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountAggregateStats**](StatsAAPI.md#GetAccountAggregateStats) | **Get** /account/stat/aggregate | Get Account Aggregate Stats
[**GetAccountAggregateStatsByGroup**](StatsAAPI.md#GetAccountAggregateStatsByGroup) | **Get** /account/stat/aggregate/group | Get Account Group Aggregate Stats
[**GetAccountStatsByGroup**](StatsAAPI.md#GetAccountStatsByGroup) | **Get** /account/stat/group | List Account Group Stats
[**GetAllAccountStats**](StatsAAPI.md#GetAllAccountStats) | **Get** /account/stat | List Account Stats



## GetAccountAggregateStats

> AggregateStats GetAccountAggregateStats(ctx).From(from).To(to).Execute()

Get Account Aggregate Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	from := time.Now() // string | Start date for aggregation (inclusive). Format YYYY-MM-DD.
	to := time.Now() // string | End date for aggregation (inclusive). Max 366 days from `from` date.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAAPI.GetAccountAggregateStats(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAAPI.GetAccountAggregateStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAggregateStats`: AggregateStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAAPI.GetAccountAggregateStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAggregateStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Start date for aggregation (inclusive). Format YYYY-MM-DD. | 
 **to** | **string** | End date for aggregation (inclusive). Max 366 days from &#x60;from&#x60; date. | 

### Return type

[**AggregateStats**](AggregateStats.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAggregateStatsByGroup

> AggregateStat GetAccountAggregateStatsByGroup(ctx).Group(group).From(from).To(to).Execute()

Get Account Group Aggregate Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	group := "shopify" // string | The group/tag name to filter and aggregate statistics by.
	from := time.Now() // string | Start date for aggregation (inclusive). Format YYYY-MM-DD.
	to := time.Now() // string | End date for aggregation (inclusive). Max 366 days from `from` date.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAAPI.GetAccountAggregateStatsByGroup(context.Background()).Group(group).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAAPI.GetAccountAggregateStatsByGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAggregateStatsByGroup`: AggregateStat
	fmt.Fprintf(os.Stdout, "Response from `StatsAAPI.GetAccountAggregateStatsByGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAggregateStatsByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string** | The group/tag name to filter and aggregate statistics by. | 
 **from** | **string** | Start date for aggregation (inclusive). Format YYYY-MM-DD. | 
 **to** | **string** | End date for aggregation (inclusive). Max 366 days from &#x60;from&#x60; date. | 

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


## GetAccountStatsByGroup

> []Stat GetAccountStatsByGroup(ctx).Group(group).From(from).To(to).Execute()

List Account Group Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	group := "shopify" // string | The group/tag name to filter statistics by.
	from := time.Now() // string | Start date for stats retrieval (inclusive). Format YYYY-MM-DD.
	to := time.Now() // string | End date for stats retrieval (inclusive). Max 31 days from `from` date.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAAPI.GetAccountStatsByGroup(context.Background()).Group(group).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAAPI.GetAccountStatsByGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountStatsByGroup`: []Stat
	fmt.Fprintf(os.Stdout, "Response from `StatsAAPI.GetAccountStatsByGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStatsByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string** | The group/tag name to filter statistics by. | 
 **from** | **string** | Start date for stats retrieval (inclusive). Format YYYY-MM-DD. | 
 **to** | **string** | End date for stats retrieval (inclusive). Max 31 days from &#x60;from&#x60; date. | 

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


## GetAllAccountStats

> []AccountStats GetAllAccountStats(ctx).From(from).To(to).Execute()

List Account Stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/sendpost/sendpost-go-sdk/v2"
)

func main() {
	from := time.Now() // string | Start date for stats retrieval (inclusive). Format YYYY-MM-DD.
	to := time.Now() // string | End date for stats retrieval (inclusive). Max 31 days from `from` date.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAAPI.GetAllAccountStats(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAAPI.GetAllAccountStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAccountStats`: []AccountStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAAPI.GetAllAccountStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccountStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Start date for stats retrieval (inclusive). Format YYYY-MM-DD. | 
 **to** | **string** | End date for stats retrieval (inclusive). Max 31 days from &#x60;from&#x60; date. | 

### Return type

[**[]AccountStats**](AccountStats.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

