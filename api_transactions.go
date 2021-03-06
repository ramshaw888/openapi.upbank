/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	_bytes "bytes"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// TransactionsApiService TransactionsApi service
type TransactionsApiService service

// AccountsAccountIdTransactionsGetOpts Optional parameters for the method 'AccountsAccountIdTransactionsGet'
type AccountsAccountIdTransactionsGetOpts struct {
    PageSize optional.Int32
    FilterStatus optional.Interface
    FilterSince optional.Time
    FilterUntil optional.Time
    FilterCategory optional.String
    FilterTag optional.String
}

/*
AccountsAccountIdTransactionsGet List transactions by account
Retrieve a list of all transactions for a specific account. The returned list is [paginated](#pagination) and can be scrolled by following the &#x60;next&#x60; and &#x60;prev&#x60; links where present. To narrow the results to a specific date range pass one or both of &#x60;filter[since]&#x60; and &#x60;filter[until]&#x60; in the query string. These filter parameters **should not** be used for pagination. Results are ordered newest first to oldest last. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId The unique identifier for the account. 
 * @param optional nil or *AccountsAccountIdTransactionsGetOpts - Optional Parameters:
 * @param "PageSize" (optional.Int32) -  The number of records to return in each page. 
 * @param "FilterStatus" (optional.Interface of TransactionStatusEnum) -  The transaction status for which to return records. This can be used to filter `HELD` transactions from those that are `SETTLED`. 
 * @param "FilterSince" (optional.Time) -  The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes. 
 * @param "FilterUntil" (optional.Time) -  The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes. 
 * @param "FilterCategory" (optional.String) -  The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a `404` response. 
 * @param "FilterTag" (optional.String) -  A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given. 
@return ListTransactionsResponse
*/
func (a *TransactionsApiService) AccountsAccountIdTransactionsGet(ctx _context.Context, accountId string, localVarOptionals *AccountsAccountIdTransactionsGetOpts) (ListTransactionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListTransactionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/accounts/{accountId}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", _neturl.PathEscape(parameterToString(accountId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page[size]", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterStatus.IsSet() {
		localVarQueryParams.Add("filter[status]", parameterToString(localVarOptionals.FilterStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterSince.IsSet() {
		localVarQueryParams.Add("filter[since]", parameterToString(localVarOptionals.FilterSince.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterUntil.IsSet() {
		localVarQueryParams.Add("filter[until]", parameterToString(localVarOptionals.FilterUntil.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterCategory.IsSet() {
		localVarQueryParams.Add("filter[category]", parameterToString(localVarOptionals.FilterCategory.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterTag.IsSet() {
		localVarQueryParams.Add("filter[tag]", parameterToString(localVarOptionals.FilterTag.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// TransactionsGetOpts Optional parameters for the method 'TransactionsGet'
type TransactionsGetOpts struct {
    PageSize optional.Int32
    FilterStatus optional.Interface
    FilterSince optional.Time
    FilterUntil optional.Time
    FilterCategory optional.String
    FilterTag optional.String
}

/*
TransactionsGet List transactions
Retrieve a list of all transactions across all accounts for the currently authenticated user. The returned list is [paginated](#pagination) and can be scrolled by following the &#x60;next&#x60; and &#x60;prev&#x60; links where present. To narrow the results to a specific date range pass one or both of &#x60;filter[since]&#x60; and &#x60;filter[until]&#x60; in the query string. These filter parameters **should not** be used for pagination. Results are ordered newest first to oldest last. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TransactionsGetOpts - Optional Parameters:
 * @param "PageSize" (optional.Int32) -  The number of records to return in each page. 
 * @param "FilterStatus" (optional.Interface of TransactionStatusEnum) -  The transaction status for which to return records. This can be used to filter `HELD` transactions from those that are `SETTLED`. 
 * @param "FilterSince" (optional.Time) -  The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes. 
 * @param "FilterUntil" (optional.Time) -  The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes. 
 * @param "FilterCategory" (optional.String) -  The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a `404` response. 
 * @param "FilterTag" (optional.String) -  A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given. 
@return ListTransactionsResponse
*/
func (a *TransactionsApiService) TransactionsGet(ctx _context.Context, localVarOptionals *TransactionsGetOpts) (ListTransactionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListTransactionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transactions"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page[size]", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterStatus.IsSet() {
		localVarQueryParams.Add("filter[status]", parameterToString(localVarOptionals.FilterStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterSince.IsSet() {
		localVarQueryParams.Add("filter[since]", parameterToString(localVarOptionals.FilterSince.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterUntil.IsSet() {
		localVarQueryParams.Add("filter[until]", parameterToString(localVarOptionals.FilterUntil.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterCategory.IsSet() {
		localVarQueryParams.Add("filter[category]", parameterToString(localVarOptionals.FilterCategory.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterTag.IsSet() {
		localVarQueryParams.Add("filter[tag]", parameterToString(localVarOptionals.FilterTag.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
TransactionsIdGet Retrieve transaction
Retrieve a specific transaction by providing its unique identifier. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The unique identifier for the transaction. 
@return GetTransactionResponse
*/
func (a *TransactionsApiService) TransactionsIdGet(ctx _context.Context, id string) (GetTransactionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetTransactionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transactions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(id, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
