package visa

type Response struct {
	VerificationId              string                    `json:"verificationId"`
	VerificationStatus          string                    `json:"verificationStatus"`
	ClientVerificationReference string                    `json:"clientVerificationReference"`
	AccountValidity             AccountVerificationDetail `json:"accountVerificationDetail"`
}

type AccountVerificationDetail struct {
	AccountIdentifierMatch string `json:"accountIdentifierMatch"`
	NameMatch              string `json:"nameMatch"`
	BankAccountName        string `json:"bankAccountName"`
}
