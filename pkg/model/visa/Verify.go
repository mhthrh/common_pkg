package visa

type Verify struct {
	PayoutMethod                string    `json:"payoutMethod"`
	RecipientDetail             Recipient `json:"recipientDetail"`
	InitiatingPartyId           string    `json:"initiatingPartyId"`
	ClientVerificationReference string    `json:"clientVerificationReference"`
}
