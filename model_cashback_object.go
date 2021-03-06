/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CashbackObject Provides information about an instant reimbursement in the form of cashback. 
type CashbackObject struct {
	// A brief description of why this cashback was paid. 
	Description string `json:"description"`
	// The total amount of cashback paid, represented as a positive value. 
	Amount MoneyObject `json:"amount"`
}
