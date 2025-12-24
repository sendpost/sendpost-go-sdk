# Go API client for sendpost [![Go Reference](https://pkg.go.dev/badge/github.com/sendpost/sendpost-go-sdk.svg)](https://pkg.go.dev/github.com/sendpost/sendpost-go-sdk)

# Introduction

SendPost provides email API and SMTP relay which can be used not just to send & measure but also alert & optimised email sending.

You can use SendPost to:

* Send personalised emails to multiple recipients using email API 

* Track opens and clicks

* Analyse statistics around open, clicks, bounce, unsubscribe and spam 


At and advanced level you can use it to:

* Manage multiple sub-accounts which may map to your promotional or transactional sending, multiple product lines or multiple customers 

* Classify your emails using groups for better analysis

* Analyse and fix email sending at sub-account level, IP Pool level or group level

* Have automated alerts to notify disruptions regarding email sending

* Manage different dedicated IP Pools so to better control your email sending

* Automatically know when IP or domain is blacklisted or sender score is down

* Leverage pro deliverability tools to get significantly better email deliverability & inboxing


[<img src=\"https://run.pstmn.io/button.svg\" alt=\"Run In Postman\" style=\"width: 128px; height: 32px;\">](https://god.gw.postman.com/run-collection/33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D33476323-e6dbd27f-c4a7-4d49-bcac-94b0611b938b%26entityType%3Dcollection%26workspaceId%3D6b1e4f65-96a9-4136-9512-6266c852517e) 

# Overview

## REST API

SendPost API is built on REST API principles. Authenticated users can interact with any of the API endpoints to perform:

* **GET**- to get a resource

* **POST** - to create a resource

* **PUT** - to update an existing resource

* **DELETE** - to delete a resource


The API endpoint for all API calls is:
<code>https://api.sendpost.io/api/v1</code>


Some conventions that have been followed in the API design overall are following:


* All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header.

* All resource endpoints end with singular name and not plural. So we have <code>domain</code> instead of domains for domain resource endpoint. Likewise we have <code>sender</code> instead of senders for sender resource endpoint.

* Body submitted for POST / PUT API calls as well as JSON response from SendPost API follow camelcase convention

* All timestamps returned in response (created or submittedAt response fields) are UNIX nano epoch timestamp.


<aside class=\"success\">
All resources have either <code>/api/v1/subaccount</code> or <code>/api/v1/account</code> in their API call resource path based on who is authorised for the resource. All API calls with path <code>/api/v1/subaccount</code> use <code>X-SubAccount-ApiKey</code> in their request header. Likewise all API calls with path <code>/api/v1/account</code> use <code>X-Account-ApiKey</code> in their request header.
</aside>


SendPost uses conventional HTTP response codes to indicate the success or failure of an API request. 


* Codes in the <code>2xx</code> range indicate success. 

* Codes in the <code>4xx</code> range indicate an error owing due to unauthorize access, incorrect request parameters or body etc.

* Code in the <code>5xx</code> range indicate an eror with SendPost's servers ( internal service issue or maintenance )


<aside class=\"info\">
SendPost all responses return <code>created</code> in UNIX nano epoch timestamp. 
</aside>


## Authentication

SendPost uses API keys for authentication. You can register a new SendPost API key at our [developer portal](https://app.sendpost.io/register).


SendPost expects the API key to be included in all API requests to the server in a header that looks like the following:


`X-SubAccount-ApiKey: AHEZEP8192SEGH`


This API key is used for all Sub-Account level operations such as:

* Sending emails

* Retrieving stats regarding open, click, bounce, unsubscribe and spam

* Uploading suppressions list

* Verifying sending domains
and more

In addition to <code>X-SubAccount-ApiKey</code> you also have another API Key <code>X-Account-APIKey</code> which is used for Account level operations such as :

* Creating and managing sub-accounts

* Allocating IPs for your account

* Getting overall billing and usage information

* Email List validation

* Creating and managing alerts
and more


<aside class=\"notice\">
You must look at individual API reference page to look at whether <code>X-SubAccount-ApiKey</code> is required or <code>X-Account-ApiKey</code>
</aside>


In case an incorrect API Key header is specified or if it is missed you will get HTTP Response 401 ( Unauthorized ) response from SendPost.


## HTTP Response Headers


Code           | Reason                 | Details
---------------| -----------------------| -----------
200            | Success                | Everything went well
401            | Unauthorized           | Incorrect or missing API header either <code>X-SubAccount-ApiKey</code> or <code>X-Account-ApiKey</code>
403            | Forbidden              | Typically sent when resource with same name or details already exist
406            | Missing resource id    | Resource id specified is either missing or doesn't exist
422            | Unprocessable entity   | Request body is not in proper format
500            | Internal server error  | Some error happened at SendPost while processing API request
503            | Service Unavailable    | SendPost is offline for maintenance. Please try again later

# API SDKs

We have native SendPost SDKs in the following programming languages. You can integrate with them or create your own SDK with our API specification. In case you need any assistance with respect to API then do reachout to our team from website chat or email us at **hello@sendpost.io**


* [PHP](https://github.com/sendpost/sendpost_php_sdk)

* [Javascript](https://github.com/sendpost/sendpost_javascript_sdk)

* [Ruby](https://github.com/sendpost/sendpost_ruby_sdk)

* [Python](https://github.com/sendpost/sendpost_python_sdk)

* [Golang](https://github.com/sendpost/sendpost_go_sdk)


# API Reference

SendX REST API can be broken down into two major sub-sections:


* Sub-Account

* Account 


Sub-Account API operations enable common email sending API use-cases like sending bulk email, adding new domains or senders for email sending programmatically, retrieving stats, adding suppressions etc. All Sub-Account API operations need to pass <code>X-SubAccount-ApiKey</code> header with every API call.


The Account API operations allow users to manage multiple sub-accounts and manage IPs. A single parent SendPost account can have 100's of sub-accounts. You may want to create sub-accounts for different products your company is running or to segregate types of emails or for managing email sending across multiple customers of yours.


# Installation

## Prerequisites

Before installing the SendPost Go SDK, ensure you have the following installed:

1. **Go 1.18 or higher** - The SDK requires Go 1.18 or later
   - Check your Go version: `go version`
   - If you need to install or update Go, visit [https://golang.org/dl/](https://golang.org/dl/)
   - For macOS: `brew install go` or download from the official website
   - For Linux: Use your distribution's package manager or download from the official website
   - For Windows: Download the installer from the official website

2. **Git** - Required for cloning repositories (if installing from source)
   - Check if Git is installed: `git --version`
   - Install Git: [https://git-scm.com/downloads](https://git-scm.com/downloads)

3. **A SendPost account** - You'll need API keys to use the SDK
   - Sign up at [https://app.sendpost.io/register](https://app.sendpost.io/register)
   - Get your API keys from the dashboard

## Installing the SDK

### Option 1: Using go get (Recommended)

The easiest way to install the SendPost Go SDK is using `go get`:

```bash
go get github.com/sendpost/sendpost-go-sdk
```

This will download and install the SDK to your Go module cache.

### Option 2: Adding to your go.mod

If you're working in a Go module project, add the SDK to your `go.mod` file:

```bash
go mod init your-project-name
go get github.com/sendpost/sendpost-go-sdk
```

Or manually add it to your `go.mod`:

```go
module your-project-name

go 1.18

require (
    github.com/sendpost/sendpost-go-sdk v1.0.0
)
```

Then run:

```bash
go mod tidy
```

### Option 3: Installing from source

If you need to install from a specific branch or fork:

```bash
git clone https://github.com/sendpost/sendpost-go-sdk.git
cd sendpost-go-sdk
go mod download
```

## Verifying Installation

To verify that the SDK is installed correctly, create a simple test file:

```go
package main

import (
    "fmt"
    sendpost "github.com/sendpost/sendpost-go-sdk"
)

func main() {
    cfg := sendpost.NewConfiguration()
    fmt.Printf("SendPost SDK version: %s\n", cfg.UserAgent)
    fmt.Println("SDK installed successfully!")
}
```

Run it with:

```bash
go run main.go
```

If you see the success message, the SDK is installed correctly.

## Dependencies

The SendPost Go SDK uses only standard Go libraries and has minimal external dependencies. The SDK automatically handles:

- HTTP client operations (using `net/http`)
- JSON encoding/decoding (using `encoding/json`)
- Context management (using `context`)

No additional dependencies need to be installed manually.

## Quick Start

Once installed, you can start using the SDK. Here's a basic example:

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    sendpost "github.com/sendpost/sendpost-go-sdk"
)

func main() {
    // Get your API key from environment variable or set it directly
    apiKey := os.Getenv("SENDPOST_SUB_ACCOUNT_API_KEY")
    if apiKey == "" {
        apiKey = "YOUR_SUB_ACCOUNT_API_KEY_HERE"
    }

    // Create configuration
    cfg := sendpost.NewConfiguration()
    cfg.Servers = sendpost.ServerConfigurations{
        {
            URL: "https://api.sendpost.io/api/v1",
            Description: "SendPost API Server",
        },
    }

    // Create API client
    client := sendpost.NewAPIClient(cfg)

    // Set up authentication context
    ctx := context.WithValue(
        context.Background(),
        sendpost.ContextAPIKeys,
        map[string]sendpost.APIKey{
            "subAccountAuth": {Key: apiKey},
        },
    )

    // Example: Get all domains
    domains, resp, err := client.DomainAPI.SubaccountDomainGet(ctx).Execute()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Response status: %s\n", resp.Status)
    fmt.Printf("Number of domains: %d\n", len(domains))
}
```

## Configuration

### Setting up API Keys

You can set API keys in several ways:

1. **Environment Variables** (Recommended for production):
```bash
export SENDPOST_ACCOUNT_API_KEY="your_account_api_key"
export SENDPOST_SUB_ACCOUNT_API_KEY="your_sub_account_api_key"
```

2. **In your code**:
```go
ctx := context.WithValue(
    context.Background(),
    sendpost.ContextAPIKeys,
    map[string]sendpost.APIKey{
        "accountAuth": {Key: "your_account_api_key"},
        "subAccountAuth": {Key: "your_sub_account_api_key"},
    },
)
```

### Using a Proxy

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

Or configure it in your HTTP client:

```go
cfg := sendpost.NewConfiguration()
// Configure custom HTTP client with proxy if needed
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sendpost.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), sendpost.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sendpost.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), sendpost.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sendpost.ContextOperationServerIndices` and `sendpost.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), sendpost.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sendpost.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sendpost.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DomainAPI* | [**GetAllDomains**](docs/DomainAPI.md#getalldomains) | **Get** /subaccount/domain | List Domains
*DomainAPI* | [**SubaccountDomainDomainIdDelete**](docs/DomainAPI.md#subaccountdomaindomainiddelete) | **Delete** /subaccount/domain/{domain_id} | Delete Domain
*DomainAPI* | [**SubaccountDomainDomainIdGet**](docs/DomainAPI.md#subaccountdomaindomainidget) | **Get** /subaccount/domain/{domain_id} | Get Domain
*DomainAPI* | [**SubaccountDomainPost**](docs/DomainAPI.md#subaccountdomainpost) | **Post** /subaccount/domain | Create Domain
*EmailAPI* | [**SendEmail**](docs/EmailAPI.md#sendemail) | **Post** /subaccount/email/ | Send Email
*EmailAPI* | [**SendEmailWithTemplate**](docs/EmailAPI.md#sendemailwithtemplate) | **Post** /subaccount/email/template | Send Email With Template
*IPAPI* | [**AllocateNewIp**](docs/IPAPI.md#allocatenewip) | **Put** /account/ip/allocate | Allocate IP
*IPAPI* | [**DeleteIp**](docs/IPAPI.md#deleteip) | **Delete** /account/ip/{ip_id} | Delete IP
*IPAPI* | [**GetAllIps**](docs/IPAPI.md#getallips) | **Get** /account/ip/ | List IPs
*IPAPI* | [**GetSpecificIp**](docs/IPAPI.md#getspecificip) | **Get** /account/ip/{ip_id} | Get IP
*IPAPI* | [**UpdateIp**](docs/IPAPI.md#updateip) | **Put** /account/ip/{ip_id} | Update IP
*IPPoolsAPI* | [**CreateIPPool**](docs/IPPoolsAPI.md#createippool) | **Post** /account/ippool | Create IPPool
*IPPoolsAPI* | [**DeleteIPPool**](docs/IPPoolsAPI.md#deleteippool) | **Delete** /account/ippool/{ippool_id} | Delete IPPool
*IPPoolsAPI* | [**GetAllIPPools**](docs/IPPoolsAPI.md#getallippools) | **Get** /account/ippool | List IPPools
*IPPoolsAPI* | [**GetIPPoolById**](docs/IPPoolsAPI.md#getippoolbyid) | **Get** /account/ippool/{ippool_id} | Get IPPool
*IPPoolsAPI* | [**UpdateIPPool**](docs/IPPoolsAPI.md#updateippool) | **Put** /account/ippool/{ippool_id} | Update IPPool
*MessageAPI* | [**GetMessageById**](docs/MessageAPI.md#getmessagebyid) | **Get** /account/message/{message_id} | Get Message
*StatsAPI* | [**AccountSubaccountStatSubaccountIdAggregateGet**](docs/StatsAPI.md#accountsubaccountstatsubaccountidaggregateget) | **Get** /account/subaccount/stat/{subaccount_id}/aggregate | Get Aggregate Stats
*StatsAPI* | [**AccountSubaccountStatSubaccountIdGet**](docs/StatsAPI.md#accountsubaccountstatsubaccountidget) | **Get** /account/subaccount/stat/{subaccount_id} | List Stats
*StatsAPI* | [**GetAggregateStatsByGroup**](docs/StatsAPI.md#getaggregatestatsbygroup) | **Get** /account/subaccount/stat/{subaccount_id}/group | Get Group Aggregate Stats
*StatsAAPI* | [**GetAccountAggregateStats**](docs/StatsAAPI.md#getaccountaggregatestats) | **Get** /account/stat/aggregate | Get Account Aggregate Stats
*StatsAAPI* | [**GetAccountAggregateStatsByGroup**](docs/StatsAAPI.md#getaccountaggregatestatsbygroup) | **Get** /account/stat/aggregate/group | Get Account Group Aggregate Stats
*StatsAAPI* | [**GetAccountStatsByGroup**](docs/StatsAAPI.md#getaccountstatsbygroup) | **Get** /account/stat/group | List Account Group Stats
*StatsAAPI* | [**GetAllAccountStats**](docs/StatsAAPI.md#getallaccountstats) | **Get** /account/stat | List Account Stats
*SubAccountAPI* | [**CreateSubAccount**](docs/SubAccountAPI.md#createsubaccount) | **Post** /account/subaccount/ | Create Sub-Account
*SubAccountAPI* | [**DeleteSubAccount**](docs/SubAccountAPI.md#deletesubaccount) | **Delete** /account/subaccount/{subaccount_id} | Delete Sub-Account
*SubAccountAPI* | [**GetAllSubAccounts**](docs/SubAccountAPI.md#getallsubaccounts) | **Get** /account/subaccount/ | List Sub-Accounts
*SubAccountAPI* | [**GetSubAccount**](docs/SubAccountAPI.md#getsubaccount) | **Get** /account/subaccount/{subaccount_id} | Get Sub-Account
*SubAccountAPI* | [**UpdateSubAccount**](docs/SubAccountAPI.md#updatesubaccount) | **Put** /account/subaccount/{subaccount_id} | Update Sub-Account
*SuppressionAPI* | [**CreateSuppression**](docs/SuppressionAPI.md#createsuppression) | **Post** /subaccount/suppression | Create Suppressions
*SuppressionAPI* | [**DeleteSuppression**](docs/SuppressionAPI.md#deletesuppression) | **Delete** /subaccount/suppression | Delete Suppressions
*SuppressionAPI* | [**GetSuppressionList**](docs/SuppressionAPI.md#getsuppressionlist) | **Get** /subaccount/suppression | List Suppressions
*WebhookAPI* | [**CreateWebhook**](docs/WebhookAPI.md#createwebhook) | **Post** /account/webhook | Create Webhook
*WebhookAPI* | [**DeleteWebhook**](docs/WebhookAPI.md#deletewebhook) | **Delete** /account/webhook/{webhook_id} | Delete Webhook
*WebhookAPI* | [**GetAllWebhooks**](docs/WebhookAPI.md#getallwebhooks) | **Get** /account/webhook | List Webhooks
*WebhookAPI* | [**GetWebhook**](docs/WebhookAPI.md#getwebhook) | **Get** /account/webhook/{webhook_id} | Get Webhook
*WebhookAPI* | [**UpdateWebhook**](docs/WebhookAPI.md#updatewebhook) | **Put** /account/webhook/{webhook_id} | Update Webhook


## Documentation For Models

 - [AccountStats](docs/AccountStats.md)
 - [AccountStatsStat](docs/AccountStatsStat.md)
 - [AggregateStat](docs/AggregateStat.md)
 - [AggregateStats](docs/AggregateStats.md)
 - [AggregatedEmailStats](docs/AggregatedEmailStats.md)
 - [Attachment](docs/Attachment.md)
 - [CopyTo](docs/CopyTo.md)
 - [CreateDomainRequest](docs/CreateDomainRequest.md)
 - [CreateSubAccountRequest](docs/CreateSubAccountRequest.md)
 - [CreateSuppressionRequest](docs/CreateSuppressionRequest.md)
 - [CreateSuppressionRequestHardBounceInner](docs/CreateSuppressionRequestHardBounceInner.md)
 - [CreateSuppressionRequestManualInner](docs/CreateSuppressionRequestManualInner.md)
 - [CreateSuppressionRequestSpamComplaintInner](docs/CreateSuppressionRequestSpamComplaintInner.md)
 - [CreateSuppressionRequestUnsubscribeInner](docs/CreateSuppressionRequestUnsubscribeInner.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [DeleteSubAccountResponse](docs/DeleteSubAccountResponse.md)
 - [DeleteSuppression200ResponseInner](docs/DeleteSuppression200ResponseInner.md)
 - [DeleteSuppressionRequest](docs/DeleteSuppressionRequest.md)
 - [DeleteWebhookResponse](docs/DeleteWebhookResponse.md)
 - [Device](docs/Device.md)
 - [Domain](docs/Domain.md)
 - [DomainDkim](docs/DomainDkim.md)
 - [DomainDmarc](docs/DomainDmarc.md)
 - [DomainGpt](docs/DomainGpt.md)
 - [DomainReturnPath](docs/DomainReturnPath.md)
 - [DomainTrack](docs/DomainTrack.md)
 - [EIP](docs/EIP.md)
 - [EmailAddress](docs/EmailAddress.md)
 - [EmailMessage](docs/EmailMessage.md)
 - [EmailMessageFrom](docs/EmailMessageFrom.md)
 - [EmailMessageObject](docs/EmailMessageObject.md)
 - [EmailMessageReplyTo](docs/EmailMessageReplyTo.md)
 - [EmailMessageToInner](docs/EmailMessageToInner.md)
 - [EmailMessageToInnerBccInner](docs/EmailMessageToInnerBccInner.md)
 - [EmailMessageToInnerCcInner](docs/EmailMessageToInnerCcInner.md)
 - [EmailMessageWithTemplate](docs/EmailMessageWithTemplate.md)
 - [EmailResponse](docs/EmailResponse.md)
 - [EmailStats](docs/EmailStats.md)
 - [EmailStatsStats](docs/EmailStatsStats.md)
 - [Event](docs/Event.md)
 - [EventMetadata](docs/EventMetadata.md)
 - [GeoLocation](docs/GeoLocation.md)
 - [IP](docs/IP.md)
 - [IPAllocationRequest](docs/IPAllocationRequest.md)
 - [IPDeletionResponse](docs/IPDeletionResponse.md)
 - [IPPool](docs/IPPool.md)
 - [IPPoolCreateRequest](docs/IPPoolCreateRequest.md)
 - [IPPoolDeleteResponse](docs/IPPoolDeleteResponse.md)
 - [IPPoolUpdateRequest](docs/IPPoolUpdateRequest.md)
 - [IPUpdateRequest](docs/IPUpdateRequest.md)
 - [Member](docs/Member.md)
 - [Message](docs/Message.md)
 - [MessageHeaderTo](docs/MessageHeaderTo.md)
 - [MessageTo](docs/MessageTo.md)
 - [OperatingSystem](docs/OperatingSystem.md)
 - [Person](docs/Person.md)
 - [Recipient](docs/Recipient.md)
 - [SMTPAuth](docs/SMTPAuth.md)
 - [Stat](docs/Stat.md)
 - [StatStats](docs/StatStats.md)
 - [SubAccount](docs/SubAccount.md)
 - [Suppression](docs/Suppression.md)
 - [ThirdPartySendingProvider](docs/ThirdPartySendingProvider.md)
 - [UpdateSubAccount](docs/UpdateSubAccount.md)
 - [UpdateWebhook](docs/UpdateWebhook.md)
 - [UserAgent](docs/UserAgent.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookObject](docs/WebhookObject.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### accountAuth

- **Type**: API key
- **API key parameter name**: X-Account-ApiKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: accountAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sendpost.ContextAPIKeys,
		map[string]sendpost.APIKey{
			"accountAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### subAccountAuth

- **Type**: API key
- **API key parameter name**: X-SubAccount-ApiKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: subAccountAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sendpost.ContextAPIKeys,
		map[string]sendpost.APIKey{
			"subAccountAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Examples

For complete working examples, check out the [example-sdk-go](https://github.com/sendpost/go-esp-example) repository which demonstrates:

- Sending emails
- Managing domains
- Creating sub-accounts
- Managing IPs and IP pools
- Working with webhooks
- Retrieving statistics

## Troubleshooting

### Common Issues

1. **Module not found errors**: Make sure you've run `go mod tidy` after adding the SDK
2. **Authentication errors**: Verify your API keys are correct and set in the context
3. **Network errors**: Check your internet connection and firewall settings
4. **Version conflicts**: Ensure you're using Go 1.18 or higher

### Getting Help

- Check the [API documentation](https://docs.sendpost.io)
- Review the [example code](https://github.com/sendpost/example-sdk-go)
- Contact support: **hello@sendpost.io**
- Visit our [developer portal](https://app.sendpost.io)

## Author

SendPost Team

## License

See LICENSE file for details.
