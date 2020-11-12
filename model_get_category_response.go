/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetCategoryResponse Successful response to get a single category and its ancestry. 
type GetCategoryResponse struct {
	// The category returned in this response. 
	Data CategoryResource `json:"data"`
}