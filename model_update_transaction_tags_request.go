/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateTransactionTagsRequest Request to add or remove tags associated with a transaction. 
type UpdateTransactionTagsRequest struct {
	// The tags to add to or remove from the transaction. 
	Data []TagInputResourceIdentifier `json:"data"`
}
