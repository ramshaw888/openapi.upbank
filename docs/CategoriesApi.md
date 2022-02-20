# \CategoriesApi

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoriesGet**](CategoriesApi.md#CategoriesGet) | **Get** /categories | List categories
[**CategoriesIdGet**](CategoriesApi.md#CategoriesIdGet) | **Get** /categories/{id} | Retrieve category
[**TransactionsTransactionIdRelationshipsCategoryPatch**](CategoriesApi.md#TransactionsTransactionIdRelationshipsCategoryPatch) | **Patch** /transactions/{transactionId}/relationships/category | Categorize transaction



## CategoriesGet

> ListCategoriesResponse CategoriesGet(ctx).FilterParent(filterParent).Execute()

List categories



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
    filterParent := "good-life" // string | The unique identifier of a parent category for which to return only its children. Providing an invalid category identifier results in a `404` response.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.CategoriesGet(context.Background()).FilterParent(filterParent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.CategoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoriesGet`: ListCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.CategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterParent** | **string** | The unique identifier of a parent category for which to return only its children. Providing an invalid category identifier results in a &#x60;404&#x60; response.  | 

### Return type

[**ListCategoriesResponse**](ListCategoriesResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesIdGet

> GetCategoryResponse CategoriesIdGet(ctx, id).Execute()

Retrieve category



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
    id := "restaurants-and-cafes" // string | The unique identifier for the category. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.CategoriesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.CategoriesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoriesIdGet`: GetCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.CategoriesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the category.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCategoryResponse**](GetCategoryResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdRelationshipsCategoryPatch

> TransactionsTransactionIdRelationshipsCategoryPatch(ctx, transactionId).UpdateTransactionCategoryRequest(updateTransactionCategoryRequest).Execute()

Categorize transaction



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
    transactionId := "a572c7c3-b637-433c-a4ce-c0be5dcb0a5a" // string | The unique identifier for the transaction. 
    updateTransactionCategoryRequest := *openapiclient.NewUpdateTransactionCategoryRequest("TODO") // UpdateTransactionCategoryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.TransactionsTransactionIdRelationshipsCategoryPatch(context.Background(), transactionId).UpdateTransactionCategoryRequest(updateTransactionCategoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.TransactionsTransactionIdRelationshipsCategoryPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The unique identifier for the transaction.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdRelationshipsCategoryPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransactionCategoryRequest** | [**UpdateTransactionCategoryRequest**](UpdateTransactionCategoryRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

