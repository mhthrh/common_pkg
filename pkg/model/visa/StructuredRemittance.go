package visa

type StructuredRemittance struct {
	Amount                      float64                       `json:"amount"`
	TaxAmount                   float64                       `json:"taxAmount"`
	TaxCurrencyCode             string                        `json:"taxCurrencyCode"`
	CreditorReference           CreditorReference             `json:"creditorReference"`
	AmountCurrencyCode          string                        `json:"amountCurrencyCode"`
	ReferredDocumentInformation []ReferredDocumentInformation `json:"referredDocumentInformation"`
}
