# \MessageAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMessages**](MessageAPI.md#GetAllMessages) | **Get** /account/message | List Messages
[**GetMessageById**](MessageAPI.md#GetMessageById) | **Get** /account/message/{message_id} | Get Message



## GetAllMessages

> []Message GetAllMessages(ctx).From(from).To(to).Limit(limit).Offset(offset).Execute()

List Messages



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
	from := time.Now() // time.Time | Start timestamp for message retrieval. ISO 8601 format.
	to := time.Now() // time.Time | End timestamp for message retrieval. Max 60 days from `from`.
	limit := int32(50) // int32 | Number of records to return per request. Default 50, max 100. (optional) (default to 50)
	offset := int32(0) // int32 | Number of initial records to skip for pagination. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.GetAllMessages(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.GetAllMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMessages`: []Message
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.GetAllMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Start timestamp for message retrieval. ISO 8601 format. | 
 **to** | **time.Time** | End timestamp for message retrieval. Max 60 days from &#x60;from&#x60;. | 
 **limit** | **int32** | Number of records to return per request. Default 50, max 100. | [default to 50]
 **offset** | **int32** | Number of initial records to skip for pagination. | [default to 0]

### Return type

[**[]Message**](Message.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageById

> Message GetMessageById(ctx, messageId).Execute()

Get Message



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
	messageId := "msg_01H2X3Y4Z5A6B7C8D9E0F1G2H3" // string | The unique message ID returned when the email was sent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.GetMessageById(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.GetMessageById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageById`: Message
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.GetMessageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | The unique message ID returned when the email was sent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[accountAuth](../README.md#accountAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

