package visa

type VerificationDetail struct {
	AccountIdentifierMatch string `json:"accountIdentifierMatch"`
	NameMatch              string `json:"nameMatch"`
	BankAccountName        string `json:"bankAccountName"`
}
