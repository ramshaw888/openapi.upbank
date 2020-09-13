/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TagResource Provides information about a tag. 
type TagResource struct {
	// The type of this resource: `tags`
	Type string `json:"type"`
	// The label of the tag, which also acts as the tag’s unique identifier. 
	Id string `json:"id"`
	Relationships AccountResourceRelationships `json:"relationships"`
}
