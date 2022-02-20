/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AccountsApiService AccountsApi service
type AccountsApiService service

type ApiAccountsGetRequest struct {
	ctx _context.Context
	ApiService *AccountsApiService
	pageSize *int32
	filterAccountType *AccountTypeEnum
	filterOwnershipType *OwnershipTypeEnum
}

// The number of records to return in each page. 
func (r ApiAccountsGetRequest) PageSize(pageSize int32) ApiAccountsGetRequest {
	r.pageSize = &pageSize
	return r
}
// The type of account for which to return records. This can be used to filter Savers from spending accounts. 
func (r ApiAccountsGetRequest) FilterAccountType(filterAccountType AccountTypeEnum) ApiAccountsGetRequest {
	r.filterAccountType = &filterAccountType
	return r
}
// The account ownership structure for which to return records. This can be used to filter 2Up accounts from Up accounts. 
func (r ApiAccountsGetRequest) FilterOwnershipType(filterOwnershipType OwnershipTypeEnum) ApiAccountsGetRequest {
	r.filterOwnershipType = &filterOwnershipType
	return r
}

func (r ApiAccountsGetRequest) Execute() (ListAccountsResponse, *_nethttp.Response, error) {
	return r.ApiService.AccountsGetExecute(r)
}

/*
AccountsGet List accounts

Retrieve a paginated list of all accounts for the currently
authenticated user. The returned list is paginated and can be scrolled
by following the `prev` and `next` links where present.


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccountsGetRequest
*/
func (a *AccountsApiService) AccountsGet(ctx _context.Context) ApiAccountsGetRequest {
	return ApiAccountsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListAccountsResponse
func (a *AccountsApiService) AccountsGetExecute(r ApiAccountsGetRequest) (ListAccountsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListAccountsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.AccountsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.filterAccountType != nil {
		localVarQueryParams.Add("filter[accountType]", parameterToString(*r.filterAccountType, ""))
	}
	if r.filterOwnershipType != nil {
		localVarQueryParams.Add("filter[ownershipType]", parameterToString(*r.filterOwnershipType, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAccountsIdGetRequest struct {
	ctx _context.Context
	ApiService *AccountsApiService
	id string
}


func (r ApiAccountsIdGetRequest) Execute() (GetAccountResponse, *_nethttp.Response, error) {
	return r.ApiService.AccountsIdGetExecute(r)
}

/*
AccountsIdGet Retrieve account

Retrieve a specific account by providing its unique identifier.


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the account. 
 @return ApiAccountsIdGetRequest
*/
func (a *AccountsApiService) AccountsIdGet(ctx _context.Context, id string) ApiAccountsIdGetRequest {
	return ApiAccountsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetAccountResponse
func (a *AccountsApiService) AccountsIdGetExecute(r ApiAccountsIdGetRequest) (GetAccountResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetAccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.AccountsIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
