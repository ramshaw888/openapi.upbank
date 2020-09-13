# \TagsApi

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsGet**](TagsApi.md#TagsGet) | **Get** /tags | List tags
[**TransactionsTransactionIdRelationshipsTagsDelete**](TagsApi.md#TransactionsTransactionIdRelationshipsTagsDelete) | **Delete** /transactions/{transactionId}/relationships/tags | Remove tags from transaction
[**TransactionsTransactionIdRelationshipsTagsPost**](TagsApi.md#TransactionsTransactionIdRelationshipsTagsPost) | **Post** /transactions/{transactionId}/relationships/tags | Add tags to transaction



## TagsGet

> ListTagsResponse TagsGet(ctx, optional)

List tags

Retrieve a list of all tags currently in use. The returned list is [paginated](#pagination) and can be scrolled by following the `next` and `prev` links where present. Results are ordered lexicographically. The `transactions` relationship for each tag exposes a link to get the transactions with the given tag. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TagsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of records to return in each page.  | 

### Return type

[**ListTagsResponse**](ListTagsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdRelationshipsTagsDelete

> TransactionsTransactionIdRelationshipsTagsDelete(ctx, transactionId, optional)

Remove tags from transaction

Disassociates one or more tags from a specific transaction. Tags that are not associated are silently ignored. An HTTP `204` is returned on success. The associated tags, along with this request URL, are also exposed via the `tags` relationship on the transaction resource returned from `/transactions/{id}`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string**| The unique identifier for the transaction.  | 
 **optional** | ***TransactionsTransactionIdRelationshipsTagsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TransactionsTransactionIdRelationshipsTagsDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransactionTagsRequest** | [**optional.Interface of UpdateTransactionTagsRequest**](UpdateTransactionTagsRequest.md)|  | 

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


## TransactionsTransactionIdRelationshipsTagsPost

> TransactionsTransactionIdRelationshipsTagsPost(ctx, transactionId, optional)

Add tags to transaction

Associates one or more tags with a specific transaction. No more than 6 tags may be present on any single transaction. Duplicate tags are silently ignored. An HTTP `204` is returned on success. The associated tags, along with this request URL, are also exposed via the `tags` relationship on the transaction resource returned from `/transactions/{id}`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string**| The unique identifier for the transaction.  | 
 **optional** | ***TransactionsTransactionIdRelationshipsTagsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TransactionsTransactionIdRelationshipsTagsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransactionTagsRequest** | [**optional.Interface of UpdateTransactionTagsRequest**](UpdateTransactionTagsRequest.md)|  | 

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

