package visa

type TransactionDetail struct {
	QuoteId                       int                    `json:"quoteId"`
	EndToEndId                    string                 `json:"endToEndId"`
	PaymentRail                   string                 `json:"paymentRail"`
	PayoutSpeed                   string                 `json:"payoutSpeed"`
	AdditionalData                []Additional           `json:"additionalData"`
	PurposeOfPayment              string                 `json:"purposeOfPayment"`
	ClientReferenceId             string                 `json:"clientReferenceId"`
	InitiatingPartyId             int                    `json:"initiatingPartyId"`
	TransactionAmount             float64                `json:"transactionAmount"`
	StatementNarrative            string                 `json:"statementNarrative"`
	SenderSourceOfFunds           string                 `json:"senderSourceOfFunds"`
	StructuredRemittance          []StructuredRemittance `json:"structuredRemittance"`
	BusinessApplicationId         string                 `json:"businessApplicationId"`
	SettlementCurrencyCode        string                 `json:"settlementCurrencyCode"`
	TransactionCurrencyCode       string                 `json:"transactionCurrencyCode"`
	SenderBeneficiaryRelationship string                 `json:"senderBeneficiaryRelationship"`
}
