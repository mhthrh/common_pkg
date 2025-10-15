package account

import "github.com/mhthrh/common_pkg/pkg/model/visa"

type Validate struct {
	PayoutMethod      string                 `json:"payoutMethod"`
	TransactionDetail visa.TransactionDetail `json:"transactionDetail"`
	RecipientDetail   visa.Recipient         `json:"recipientDetail"`
	SenderDetail      visa.SenderDetail      `json:"senderDetail"`
}
