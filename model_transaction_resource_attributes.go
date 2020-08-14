/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It's new, it's exciting and it's just the beginning.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// TransactionResourceAttributes struct for TransactionResourceAttributes
type TransactionResourceAttributes struct {
	// The current processing status of this transaction, according to whether or not this transaction has settled or is still held.
	Status TransactionStatusEnum `json:"status"`
	// The original, unprocessed text of the transaction. This is often not a perfect indicator of the actual merchant, but it is useful for reconciliation purposes in some cases.
	RawText *string `json:"rawText"`
	// A short description for this transaction. Usually the merchant name for purchases.
	Description string `json:"description"`
	// Attached message for this transaction, such as a payment message, or a transfer note.
	Message *string `json:"message"`
	// If this transaction is currently in the `HELD` status, or was ever in the `HELD` status, the `amount` and `foreignAmount` of the transaction while `HELD`.
	HoldInfo *HoldInfoObject `json:"holdInfo"`
	// Details of how this transaction was rounded-up. If no Round Up was applied this field will be `null`.
	RoundUp *RoundUpObject `json:"roundUp"`
	// If all or part of this transaction was instantly reimbursed in the form of cashback, details of the reimbursement.
	Cashback *CashbackObject `json:"cashback"`
	// The amount of this transaction in Australian dollars. For transactions that were once `HELD` but are now `SETTLED`, refer to the `holdInfo` field for the original `amount` the transaction was `HELD` at.
	Amount MoneyObject `json:"amount"`
	// The foreign currency amount of this transaction. This field will be `null` for domestic transactions. The amount was converted to the AUD amount reflected in the `amount` of this transaction. Refer to the `holdInfo` field for the original `foreignAmount` the transaction was `HELD` at.
	ForeignAmount *MoneyObject `json:"foreignAmount"`
	// The date-time at which this transaction settled. This field will be `null` for transactions that are currently in the `HELD` status.
	SettledAt *time.Time `json:"settledAt"`
	// The date-time at which this transaction was first encountered.
	CreatedAt time.Time `json:"createdAt"`
}
