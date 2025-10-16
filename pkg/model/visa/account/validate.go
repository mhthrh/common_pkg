package account

import "github.com/mhthrh/common_pkg/pkg/model/visa"

type ValidateRequest struct {
	PayoutMethod      string                 `json:"payoutMethod"`
	TransactionDetail visa.TransactionDetail `json:"transactionDetail"`
	RecipientDetail   visa.Recipient         `json:"recipientDetail"`
	SenderDetail      visa.SenderDetail      `json:"senderDetail"`
}

type ValidateResponse struct {
	ValidationResultCode string             `json:"validationResultCode"`
	ExpectedPostingDate  string             `json:"expectedPostingDate"`
	Error                visa.ErrorResponse `json:"errorResponse"`
}
