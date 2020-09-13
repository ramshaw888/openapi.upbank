/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateWebhookResponse Successful response after creating a webhook. 
type CreateWebhookResponse struct {
	// The webhook that was created. 
	Data WebhookResource `json:"data"`
}
