/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListTagsResponse Successful response to get all tags. This returns a paginated list of tags, which can be scrolled by following the `prev` and `next` links if present. 
type ListTagsResponse struct {
	// The list of tags returned in this response. 
	Data []TagResource `json:"data"`
	Links ListAccountsResponseLinks `json:"links"`
}
