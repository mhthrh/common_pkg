package account

import "github.com/mhthrh/common_pkg/pkg/model/visa"

type VerifyRequest struct {
	PayoutMethod                string         `json:"payoutMethod"`
	RecipientDetail             visa.Recipient `json:"recipientDetail"`
	InitiatingPartyId           string         `json:"initiatingPartyId"`
	ClientVerificationReference string         `json:"clientVerificationReference"`
}

type VerifyResponse struct {
	VerificationId              string                  `json:"verificationId"`
	VerificationStatus          string                  `json:"verificationStatus"`
	ClientVerificationReference string                  `json:"clientVerificationReference"`
	AccountValidity             visa.VerificationDetail `json:"accountVerificationDetail"`
}
