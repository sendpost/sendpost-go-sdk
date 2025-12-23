# \WebhookReferenceAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendPostWebhooksPost**](WebhookReferenceAPI.md#SendPostWebhooksPost) | **Post** /SendPostWebhooks | SendPost Webhook Object



## SendPostWebhooksPost

> SendPostWebhooksPost(ctx).WebhookObject(webhookObject).Execute()

SendPost Webhook Object

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
	webhookObject :=  // WebhookObject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookReferenceAPI.SendPostWebhooksPost(context.Background()).WebhookObject(webhookObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookReferenceAPI.SendPostWebhooksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPostWebhooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookObject** | [**WebhookObject**](WebhookObject.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

