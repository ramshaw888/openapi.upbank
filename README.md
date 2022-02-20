# Go API client for openapi

## Generated with

```
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.3.1 generate -i https://raw.githubusercontent.com/up-banking/api/master/v1/openapi.json -g go -o /local/out/go
```

The Up API gives you programmatic access to your balances and
transaction data. You can request past transactions or set up
webhooks to receive real-time events when new transactions hit your
account. It’s new, it’s exciting and it’s just the beginning.

## Overview

This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project. By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
  For more information, please visit [https://github.com/up-banking/api](https://github.com/up-banking/api)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.up.com.au/api/v1*

| Class                 | Method                                                                                                                               | HTTP request                                                   | Description                  |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------- | ---------------------------- |
| _AccountsApi_         | [**AccountsGet**](docs/AccountsApi.md#accountsget)                                                                                   | **Get** /accounts                                              | List accounts                |
| _AccountsApi_         | [**AccountsIdGet**](docs/AccountsApi.md#accountsidget)                                                                               | **Get** /accounts/{id}                                         | Retrieve account             |
| _CategoriesApi_       | [**CategoriesGet**](docs/CategoriesApi.md#categoriesget)                                                                             | **Get** /categories                                            | List categories              |
| _CategoriesApi_       | [**CategoriesIdGet**](docs/CategoriesApi.md#categoriesidget)                                                                         | **Get** /categories/{id}                                       | Retrieve category            |
| _CategoriesApi_       | [**TransactionsTransactionIdRelationshipsCategoryPatch**](docs/CategoriesApi.md#transactionstransactionidrelationshipscategorypatch) | **Patch** /transactions/{transactionId}/relationships/category | Categorize transaction       |
| _TagsApi_             | [**TagsGet**](docs/TagsApi.md#tagsget)                                                                                               | **Get** /tags                                                  | List tags                    |
| _TagsApi_             | [**TransactionsTransactionIdRelationshipsTagsDelete**](docs/TagsApi.md#transactionstransactionidrelationshipstagsdelete)             | **Delete** /transactions/{transactionId}/relationships/tags    | Remove tags from transaction |
| _TagsApi_             | [**TransactionsTransactionIdRelationshipsTagsPost**](docs/TagsApi.md#transactionstransactionidrelationshipstagspost)                 | **Post** /transactions/{transactionId}/relationships/tags      | Add tags to transaction      |
| _TransactionsApi_     | [**AccountsAccountIdTransactionsGet**](docs/TransactionsApi.md#accountsaccountidtransactionsget)                                     | **Get** /accounts/{accountId}/transactions                     | List transactions by account |
| _TransactionsApi_     | [**TransactionsGet**](docs/TransactionsApi.md#transactionsget)                                                                       | **Get** /transactions                                          | List transactions            |
| _TransactionsApi_     | [**TransactionsIdGet**](docs/TransactionsApi.md#transactionsidget)                                                                   | **Get** /transactions/{id}                                     | Retrieve transaction         |
| _UtilityEndpointsApi_ | [**UtilPingGet**](docs/UtilityEndpointsApi.md#utilpingget)                                                                           | **Get** /util/ping                                             | Ping                         |
| _WebhooksApi_         | [**WebhooksGet**](docs/WebhooksApi.md#webhooksget)                                                                                   | **Get** /webhooks                                              | List webhooks                |
| _WebhooksApi_         | [**WebhooksIdDelete**](docs/WebhooksApi.md#webhooksiddelete)                                                                         | **Delete** /webhooks/{id}                                      | Delete webhook               |
| _WebhooksApi_         | [**WebhooksIdGet**](docs/WebhooksApi.md#webhooksidget)                                                                               | **Get** /webhooks/{id}                                         | Retrieve webhook             |
| _WebhooksApi_         | [**WebhooksPost**](docs/WebhooksApi.md#webhookspost)                                                                                 | **Post** /webhooks                                             | Create webhook               |
| _WebhooksApi_         | [**WebhooksWebhookIdLogsGet**](docs/WebhooksApi.md#webhookswebhookidlogsget)                                                         | **Get** /webhooks/{webhookId}/logs                             | List webhook logs            |
| _WebhooksApi_         | [**WebhooksWebhookIdPingPost**](docs/WebhooksApi.md#webhookswebhookidpingpost)                                                       | **Post** /webhooks/{webhookId}/ping                            | Ping webhook                 |

## Documentation For Models

- [AccountResource](docs/AccountResource.md)
- [AccountResourceAttributes](docs/AccountResourceAttributes.md)
- [AccountResourceLinks](docs/AccountResourceLinks.md)
- [AccountResourceRelationships](docs/AccountResourceRelationships.md)
- [AccountResourceRelationshipsTransactions](docs/AccountResourceRelationshipsTransactions.md)
- [AccountResourceRelationshipsTransactionsLinks](docs/AccountResourceRelationshipsTransactionsLinks.md)
- [AccountTypeEnum](docs/AccountTypeEnum.md)
- [CashbackObject](docs/CashbackObject.md)
- [CategoryInputResourceIdentifier](docs/CategoryInputResourceIdentifier.md)
- [CategoryResource](docs/CategoryResource.md)
- [CategoryResourceAttributes](docs/CategoryResourceAttributes.md)
- [CategoryResourceRelationships](docs/CategoryResourceRelationships.md)
- [CategoryResourceRelationshipsChildren](docs/CategoryResourceRelationshipsChildren.md)
- [CategoryResourceRelationshipsChildrenData](docs/CategoryResourceRelationshipsChildrenData.md)
- [CategoryResourceRelationshipsParent](docs/CategoryResourceRelationshipsParent.md)
- [CategoryResourceRelationshipsParentData](docs/CategoryResourceRelationshipsParentData.md)
- [CreateWebhookRequest](docs/CreateWebhookRequest.md)
- [CreateWebhookResponse](docs/CreateWebhookResponse.md)
- [ErrorObject](docs/ErrorObject.md)
- [ErrorObjectSource](docs/ErrorObjectSource.md)
- [ErrorResponse](docs/ErrorResponse.md)
- [GetAccountResponse](docs/GetAccountResponse.md)
- [GetCategoryResponse](docs/GetCategoryResponse.md)
- [GetTransactionResponse](docs/GetTransactionResponse.md)
- [GetWebhookResponse](docs/GetWebhookResponse.md)
- [HoldInfoObject](docs/HoldInfoObject.md)
- [ListAccountsResponse](docs/ListAccountsResponse.md)
- [ListAccountsResponseLinks](docs/ListAccountsResponseLinks.md)
- [ListCategoriesResponse](docs/ListCategoriesResponse.md)
- [ListTagsResponse](docs/ListTagsResponse.md)
- [ListTransactionsResponse](docs/ListTransactionsResponse.md)
- [ListWebhookDeliveryLogsResponse](docs/ListWebhookDeliveryLogsResponse.md)
- [ListWebhooksResponse](docs/ListWebhooksResponse.md)
- [MoneyObject](docs/MoneyObject.md)
- [OwnershipTypeEnum](docs/OwnershipTypeEnum.md)
- [PingResponse](docs/PingResponse.md)
- [PingResponseMeta](docs/PingResponseMeta.md)
- [RoundUpObject](docs/RoundUpObject.md)
- [TagInputResourceIdentifier](docs/TagInputResourceIdentifier.md)
- [TagResource](docs/TagResource.md)
- [TransactionResource](docs/TransactionResource.md)
- [TransactionResourceAttributes](docs/TransactionResourceAttributes.md)
- [TransactionResourceRelationships](docs/TransactionResourceRelationships.md)
- [TransactionResourceRelationshipsAccount](docs/TransactionResourceRelationshipsAccount.md)
- [TransactionResourceRelationshipsAccountData](docs/TransactionResourceRelationshipsAccountData.md)
- [TransactionResourceRelationshipsCategory](docs/TransactionResourceRelationshipsCategory.md)
- [TransactionResourceRelationshipsCategoryLinks](docs/TransactionResourceRelationshipsCategoryLinks.md)
- [TransactionResourceRelationshipsTags](docs/TransactionResourceRelationshipsTags.md)
- [TransactionResourceRelationshipsTagsData](docs/TransactionResourceRelationshipsTagsData.md)
- [TransactionResourceRelationshipsTagsLinks](docs/TransactionResourceRelationshipsTagsLinks.md)
- [TransactionResourceRelationshipsTransferAccount](docs/TransactionResourceRelationshipsTransferAccount.md)
- [TransactionResourceRelationshipsTransferAccountData](docs/TransactionResourceRelationshipsTransferAccountData.md)
- [TransactionStatusEnum](docs/TransactionStatusEnum.md)
- [UpdateTransactionCategoryRequest](docs/UpdateTransactionCategoryRequest.md)
- [UpdateTransactionTagsRequest](docs/UpdateTransactionTagsRequest.md)
- [WebhookDeliveryLogResource](docs/WebhookDeliveryLogResource.md)
- [WebhookDeliveryLogResourceAttributes](docs/WebhookDeliveryLogResourceAttributes.md)
- [WebhookDeliveryLogResourceAttributesRequest](docs/WebhookDeliveryLogResourceAttributesRequest.md)
- [WebhookDeliveryLogResourceAttributesResponse](docs/WebhookDeliveryLogResourceAttributesResponse.md)
- [WebhookDeliveryLogResourceRelationships](docs/WebhookDeliveryLogResourceRelationships.md)
- [WebhookDeliveryLogResourceRelationshipsWebhookEvent](docs/WebhookDeliveryLogResourceRelationshipsWebhookEvent.md)
- [WebhookDeliveryLogResourceRelationshipsWebhookEventData](docs/WebhookDeliveryLogResourceRelationshipsWebhookEventData.md)
- [WebhookDeliveryStatusEnum](docs/WebhookDeliveryStatusEnum.md)
- [WebhookEventCallback](docs/WebhookEventCallback.md)
- [WebhookEventResource](docs/WebhookEventResource.md)
- [WebhookEventResourceAttributes](docs/WebhookEventResourceAttributes.md)
- [WebhookEventResourceRelationships](docs/WebhookEventResourceRelationships.md)
- [WebhookEventResourceRelationshipsTransaction](docs/WebhookEventResourceRelationshipsTransaction.md)
- [WebhookEventResourceRelationshipsTransactionData](docs/WebhookEventResourceRelationshipsTransactionData.md)
- [WebhookEventResourceRelationshipsWebhook](docs/WebhookEventResourceRelationshipsWebhook.md)
- [WebhookEventResourceRelationshipsWebhookData](docs/WebhookEventResourceRelationshipsWebhookData.md)
- [WebhookEventTypeEnum](docs/WebhookEventTypeEnum.md)
- [WebhookInputResource](docs/WebhookInputResource.md)
- [WebhookInputResourceAttributes](docs/WebhookInputResourceAttributes.md)
- [WebhookResource](docs/WebhookResource.md)
- [WebhookResourceAttributes](docs/WebhookResourceAttributes.md)
- [WebhookResourceRelationships](docs/WebhookResourceRelationships.md)

## Documentation For Authorization

### bearer_auth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

- `PtrBool`
- `PtrInt`
- `PtrInt32`
- `PtrInt64`
- `PtrFloat`
- `PtrFloat32`
- `PtrFloat64`
- `PtrString`
- `PtrTime`

## Author
