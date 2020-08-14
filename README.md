# Go API client for openapi

## Generated with

```
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i https://raw.githubusercontent.com/up-banking/api/master/v1/openapi.json -g go -o /local/out/go
```

## Summary

The Up API gives you programmatic access to your balances and
transaction data. You can request past transactions or set up
webhooks to receive real-time events when new transactions hit your
account. It's new, it's exciting and it's just the beginning.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

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
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.up.com.au/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**AccountsGet**](docs/AccountsApi.md#accountsget) | **Get** /accounts | List accounts
*AccountsApi* | [**AccountsIdGet**](docs/AccountsApi.md#accountsidget) | **Get** /accounts/{id} | Retrieve account
*TagsApi* | [**TransactionsTransactionIdRelationshipsTagsDelete**](docs/TagsApi.md#transactionstransactionidrelationshipstagsdelete) | **Delete** /transactions/{transactionId}/relationships/tags | Remove tags from transaction
*TagsApi* | [**TransactionsTransactionIdRelationshipsTagsPost**](docs/TagsApi.md#transactionstransactionidrelationshipstagspost) | **Post** /transactions/{transactionId}/relationships/tags | Add tags to transaction
*TransactionsApi* | [**AccountsAccountIdTransactionsGet**](docs/TransactionsApi.md#accountsaccountidtransactionsget) | **Get** /accounts/{accountId}/transactions | List transactions by account
*TransactionsApi* | [**TransactionsGet**](docs/TransactionsApi.md#transactionsget) | **Get** /transactions | List transactions
*TransactionsApi* | [**TransactionsIdGet**](docs/TransactionsApi.md#transactionsidget) | **Get** /transactions/{id} | Retrieve transaction
*UtilityEndpointsApi* | [**UtilPingGet**](docs/UtilityEndpointsApi.md#utilpingget) | **Get** /util/ping | Ping
*WebhooksApi* | [**WebhooksGet**](docs/WebhooksApi.md#webhooksget) | **Get** /webhooks | List webhooks
*WebhooksApi* | [**WebhooksIdDelete**](docs/WebhooksApi.md#webhooksiddelete) | **Delete** /webhooks/{id} | Delete webhook
*WebhooksApi* | [**WebhooksIdGet**](docs/WebhooksApi.md#webhooksidget) | **Get** /webhooks/{id} | Retrieve webhook
*WebhooksApi* | [**WebhooksPost**](docs/WebhooksApi.md#webhookspost) | **Post** /webhooks | Create webhook
*WebhooksApi* | [**WebhooksWebhookIdLogsGet**](docs/WebhooksApi.md#webhookswebhookidlogsget) | **Get** /webhooks/{webhookId}/logs | List webhook logs
*WebhooksApi* | [**WebhooksWebhookIdPingPost**](docs/WebhooksApi.md#webhookswebhookidpingpost) | **Post** /webhooks/{webhookId}/ping | Ping webhook


## Documentation For Models

 - [AccountResource](docs/AccountResource.md)
 - [AccountResourceAttributes](docs/AccountResourceAttributes.md)
 - [AccountResourceLinks](docs/AccountResourceLinks.md)
 - [AccountResourceRelationships](docs/AccountResourceRelationships.md)
 - [AccountResourceRelationshipsTransactions](docs/AccountResourceRelationshipsTransactions.md)
 - [AccountResourceRelationshipsTransactionsLinks](docs/AccountResourceRelationshipsTransactionsLinks.md)
 - [AccountTypeEnum](docs/AccountTypeEnum.md)
 - [CashbackObject](docs/CashbackObject.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [CreateWebhookResponse](docs/CreateWebhookResponse.md)
 - [ErrorObject](docs/ErrorObject.md)
 - [ErrorObjectSource](docs/ErrorObjectSource.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GetAccountResponse](docs/GetAccountResponse.md)
 - [GetTransactionResponse](docs/GetTransactionResponse.md)
 - [GetWebhookResponse](docs/GetWebhookResponse.md)
 - [HoldInfoObject](docs/HoldInfoObject.md)
 - [ListAccountsResponse](docs/ListAccountsResponse.md)
 - [ListAccountsResponseLinks](docs/ListAccountsResponseLinks.md)
 - [ListTransactionsResponse](docs/ListTransactionsResponse.md)
 - [ListWebhookDeliveryLogsResponse](docs/ListWebhookDeliveryLogsResponse.md)
 - [ListWebhooksResponse](docs/ListWebhooksResponse.md)
 - [MoneyObject](docs/MoneyObject.md)
 - [PingResponse](docs/PingResponse.md)
 - [PingResponseMeta](docs/PingResponseMeta.md)
 - [RoundUpObject](docs/RoundUpObject.md)
 - [TagInputResourceIdentifier](docs/TagInputResourceIdentifier.md)
 - [TransactionResource](docs/TransactionResource.md)
 - [TransactionResourceAttributes](docs/TransactionResourceAttributes.md)
 - [TransactionResourceRelationships](docs/TransactionResourceRelationships.md)
 - [TransactionResourceRelationshipsAccount](docs/TransactionResourceRelationshipsAccount.md)
 - [TransactionResourceRelationshipsAccountData](docs/TransactionResourceRelationshipsAccountData.md)
 - [TransactionResourceRelationshipsTags](docs/TransactionResourceRelationshipsTags.md)
 - [TransactionResourceRelationshipsTagsData](docs/TransactionResourceRelationshipsTagsData.md)
 - [TransactionResourceRelationshipsTagsLinks](docs/TransactionResourceRelationshipsTagsLinks.md)
 - [TransactionStatusEnum](docs/TransactionStatusEnum.md)
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



## bearer_auth

With your Up Token.

Example

```golang
ctx := context.WithValue(
  context.Background(),
  openapi.ContextAccessToken,
  os.Getenv("UP_TOKEN"),
)
r, err := client.Service.Operation(ctx, args)
```



## Author



