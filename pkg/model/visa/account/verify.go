package account

import "github.com/mhthrh/common_pkg/pkg/model/visa"

type Verify struct {
	PayoutMethod                string         `json:"payoutMethod"`
	RecipientDetail             visa.Recipient `json:"recipientDetail"`
	InitiatingPartyId           string         `json:"initiatingPartyId"`
	ClientVerificationReference string         `json:"clientVerificationReference"`
}
