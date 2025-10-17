package account

import (
	"time"

	"github.com/mhthrh/common_pkg/pkg/model/visa"
)

type PayoutRequest struct {
	TransactionDetail visa.TransactionDetail `json:"transactionDetail"`
	PayoutMethod      string                 `json:"payoutMethod"`
	SenderDetail      visa.SenderDetail      `json:"senderDetail"`
	RecipientDetail   visa.Recipient         `json:"recipientDetail"`
}

type PayoutResponse struct {
	Error visa.ErrorResponse `json:"errorResponse"`
}

type TransactionDetail200 struct {
	Status                  string    `json:"status"`
	PayoutId                string    `json:"payoutId"`
	EndToEndId              string    `json:"endToEndId"`
	PayoutSpeed             string    `json:"payoutSpeed"`
	FxConversionRate        float64   `json:"fxConversionRate"`
	SettlementAmount        float64   `json:"settlementAmount"`
	ClientReferenceId       string    `json:"clientReferenceId"`
	DestinationAmount       float64   `json:"destinationAmount"`
	InitiatingPartyId       int       `json:"initiatingPartyId"`
	TransactionAmount       float64   `json:"transactionAmount"`
	ExpectedPostingDate     string    `json:"expectedPostingDate"`
	TransactionDateTime     time.Time `json:"transactionDateTime"`
	SettlementCurrencyCode  string    `json:"settlementCurrencyCode"`
	DestinationCurrencyCode string    `json:"destinationCurrencyCode"`
	TransactionCurrencyCode string    `json:"transactionCurrencyCode"`
}

type TransactionDetail202 struct {
	Status            string `json:"status"`
	ClientReferenceId string `json:"clientReferenceId"`
	InitiatingPartyId int    `json:"initiatingPartyId"`
}

type T400 struct {
	Reason  string `json:"reason"`
	Status  int    `json:"status"`
	Details []struct {
		Code     string `json:"code"`
		Message  string `json:"message"`
		Location string `json:"location"`
	} `json:"details"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Reason  string `json:"reason"`
	Status  int    `json:"status"`
	Details []struct {
		Code     string `json:"code"`
		Message  string `json:"message"`
		Location string `json:"location"`
	} `json:"details"`
	Message string `json:"message"`
}
