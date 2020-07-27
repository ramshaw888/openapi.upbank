/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// WebhookInputResourceAttributes struct for WebhookInputResourceAttributes
type WebhookInputResourceAttributes struct {
	// The URL that this webhook should post events to. This must be a valid HTTP or HTTPS URL that does not exceed 120 characters in length. 
	Url string `json:"url"`
}
