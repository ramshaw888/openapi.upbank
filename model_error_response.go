/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It's new, it's exciting and it's just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ErrorResponse Generic error response that returns one or more errors. 
type ErrorResponse struct {
	// The list of errors returned in this response. 
	Errors []ErrorObject `json:"errors"`
}
