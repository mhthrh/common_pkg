package visa

type Bank struct {
	BankName          string `json:"bankName"`
	AccountName       string `json:"accountName"`
	AccountNumber     string `json:"accountNumber"`
	AccountNumberType string `json:"accountNumberType"`
	CountryCode       string `json:"countryCode"`
	BankCode          string `json:"bankCode"`
	BankCodeType      string `json:"bankCodeType"`
	CurrencyCode      string `json:"currencyCode"`
}
