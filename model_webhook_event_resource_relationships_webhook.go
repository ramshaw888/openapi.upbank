/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It's new, it's exciting and it's just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WebhookEventResourceRelationshipsWebhook struct for WebhookEventResourceRelationshipsWebhook
type WebhookEventResourceRelationshipsWebhook struct {
	Data WebhookEventResourceRelationshipsWebhookData `json:"data"`
	Links AccountResourceRelationshipsTransactionsLinks `json:"links,omitempty"`
}