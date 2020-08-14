# \TagsApi

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsTransactionIdRelationshipsTagsDelete**](TagsApi.md#TransactionsTransactionIdRelationshipsTagsDelete) | **Delete** /transactions/{transactionId}/relationships/tags | Remove tags from transaction
[**TransactionsTransactionIdRelationshipsTagsPost**](TagsApi.md#TransactionsTransactionIdRelationshipsTagsPost) | **Post** /transactions/{transactionId}/relationships/tags | Add tags to transaction



## TransactionsTransactionIdRelationshipsTagsDelete

> TransactionsTransactionIdRelationshipsTagsDelete(ctx, transactionId, optional)

Remove tags from transaction

Delete a specific webhook by providing its unique identifier. Once deleted, webhook events will no longer be sent to the configured URL. 

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

