package account

import "github.com/mhthrh/common_pkg/pkg/model/visa"

type Payout struct {
	TransactionDetail visa.TransactionDetail `json:"transactionDetail"`
	PayoutMethod      string                 `json:"payoutMethod"`
	SenderDetail      visa.SenderDetail      `json:"senderDetail"`
	RecipientDetail   visa.Recipient         `json:"recipientDetail"`
}
