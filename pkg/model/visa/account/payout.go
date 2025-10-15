package account

import "github.com/mhthrh/common_pkg/pkg/model/visa"

type PayoutRequest struct {
	TransactionDetail visa.TransactionDetail `json:"transactionDetail"`
	PayoutMethod      string                 `json:"payoutMethod"`
	SenderDetail      visa.SenderDetail      `json:"senderDetail"`
	RecipientDetail   visa.Recipient         `json:"recipientDetail"`
}

type PayoutResponse struct {
	Status  int      `json:"status"`
	Reason  string   `json:"reason"`
	Details []Detail `json:"details"`
}

type Detail struct {
	Location string `json:"location"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
}
