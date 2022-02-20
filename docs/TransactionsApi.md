# \TransactionsApi

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAccountIdTransactionsGet**](TransactionsApi.md#AccountsAccountIdTransactionsGet) | **Get** /accounts/{accountId}/transactions | List transactions by account
[**TransactionsGet**](TransactionsApi.md#TransactionsGet) | **Get** /transactions | List transactions
[**TransactionsIdGet**](TransactionsApi.md#TransactionsIdGet) | **Get** /transactions/{id} | Retrieve transaction



## AccountsAccountIdTransactionsGet

> ListTransactionsResponse AccountsAccountIdTransactionsGet(ctx, accountId).PageSize(pageSize).FilterStatus(filterStatus).FilterSince(filterSince).FilterUntil(filterUntil).FilterCategory(filterCategory).FilterTag(filterTag).Execute()

List transactions by account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountId := "b5544658-4bbd-4eb1-8f63-a9909e0f564b" // string | The unique identifier for the account. 
    pageSize := int32(30) // int32 | The number of records to return in each page.  (optional)
    filterStatus := openapiclient.TransactionStatusEnum("HELD") // TransactionStatusEnum | The transaction status for which to return records. This can be used to filter `HELD` transactions from those that are `SETTLED`.  (optional)
    filterSince := time.Now() // time.Time | The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  (optional)
    filterUntil := time.Now() // time.Time | The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  (optional)
    filterCategory := "good-life" // string | The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a `404` response.  (optional)
    filterTag := "Holiday" // string | A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.AccountsAccountIdTransactionsGet(context.Background(), accountId).PageSize(pageSize).FilterStatus(filterStatus).FilterSince(filterSince).FilterUntil(filterUntil).FilterCategory(filterCategory).FilterTag(filterTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.AccountsAccountIdTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsAccountIdTransactionsGet`: ListTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.AccountsAccountIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique identifier for the account.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAccountIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of records to return in each page.  | 
 **filterStatus** | [**TransactionStatusEnum**](TransactionStatusEnum.md) | The transaction status for which to return records. This can be used to filter &#x60;HELD&#x60; transactions from those that are &#x60;SETTLED&#x60;.  | 
 **filterSince** | **time.Time** | The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 
 **filterUntil** | **time.Time** | The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 
 **filterCategory** | **string** | The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a &#x60;404&#x60; response.  | 
 **filterTag** | **string** | A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given.  | 

### Return type

[**ListTransactionsResponse**](ListTransactionsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsGet

> ListTransactionsResponse TransactionsGet(ctx).PageSize(pageSize).FilterStatus(filterStatus).FilterSince(filterSince).FilterUntil(filterUntil).FilterCategory(filterCategory).FilterTag(filterTag).Execute()

List transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    pageSize := int32(30) // int32 | The number of records to return in each page.  (optional)
    filterStatus := openapiclient.TransactionStatusEnum("HELD") // TransactionStatusEnum | The transaction status for which to return records. This can be used to filter `HELD` transactions from those that are `SETTLED`.  (optional)
    filterSince := time.Now() // time.Time | The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  (optional)
    filterUntil := time.Now() // time.Time | The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  (optional)
    filterCategory := "good-life" // string | The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a `404` response.  (optional)
    filterTag := "Holiday" // string | A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.TransactionsGet(context.Background()).PageSize(pageSize).FilterStatus(filterStatus).FilterSince(filterSince).FilterUntil(filterUntil).FilterCategory(filterCategory).FilterTag(filterTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.TransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsGet`: ListTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.TransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of records to return in each page.  | 
 **filterStatus** | [**TransactionStatusEnum**](TransactionStatusEnum.md) | The transaction status for which to return records. This can be used to filter &#x60;HELD&#x60; transactions from those that are &#x60;SETTLED&#x60;.  | 
 **filterSince** | **time.Time** | The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 
 **filterUntil** | **time.Time** | The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 
 **filterCategory** | **string** | The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a &#x60;404&#x60; response.  | 
 **filterTag** | **string** | A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given.  | 

### Return type

[**ListTransactionsResponse**](ListTransactionsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsIdGet

> GetTransactionResponse TransactionsIdGet(ctx, id).Execute()

Retrieve transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "7a9d19f9-106c-4e29-8591-52fc5d8f09c5" // string | The unique identifier for the transaction. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.TransactionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.TransactionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsIdGet`: GetTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.TransactionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transaction.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTransactionResponse**](GetTransactionResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

