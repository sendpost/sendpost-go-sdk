# \EmailAPI

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendEmail**](EmailAPI.md#SendEmail) | **Post** /subaccount/email/ | Send Email
[**SendEmailWithTemplate**](EmailAPI.md#SendEmailWithTemplate) | **Post** /subaccount/email/template | Send Email With Template



## SendEmail

> []EmailResponse SendEmail(ctx).EmailMessageObject(emailMessageObject).Execute()

Send Email



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
	emailMessageObject := *openapiclient.NewEmailMessageObject() // EmailMessageObject | Email message details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.SendEmail(context.Background()).EmailMessageObject(emailMessageObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.SendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmail`: []EmailResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.SendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailMessageObject** | [**EmailMessageObject**](EmailMessageObject.md) | Email message details | 

### Return type

[**[]EmailResponse**](EmailResponse.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailWithTemplate

> []EmailResponse SendEmailWithTemplate(ctx).EmailMessageWithTemplate(emailMessageWithTemplate).Execute()

Send Email With Template



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
	emailMessageWithTemplate := *openapiclient.NewEmailMessageWithTemplate() // EmailMessageWithTemplate | Email message details with template information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.SendEmailWithTemplate(context.Background()).EmailMessageWithTemplate(emailMessageWithTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.SendEmailWithTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmailWithTemplate`: []EmailResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.SendEmailWithTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailWithTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailMessageWithTemplate** | [**EmailMessageWithTemplate**](EmailMessageWithTemplate.md) | Email message details with template information | 

### Return type

[**[]EmailResponse**](EmailResponse.md)

### Authorization

[subAccountAuth](../README.md#subAccountAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

