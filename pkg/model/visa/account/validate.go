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

type ValidateR struct {
	PayoutMethod      string `json:"payoutMethod"`
	TransactionDetail struct {
		InitiatingPartyId             string `json:"initiatingPartyId"`
		BusinessApplicationId         string `json:"businessApplicationId"`
		TransactionAmount             string `json:"transactionAmount"`
		TransactionCurrencyCode       string `json:"transactionCurrencyCode"`
		EndToEndId                    string `json:"endToEndId"`
		SenderSourceOfFunds           string `json:"senderSourceOfFunds"`
		SenderBeneficiaryRelationship string `json:"senderBeneficiaryRelationship"`
		StatementNarrative            string `json:"statementNarrative"`
		PayoutSpeed                   string `json:"payoutSpeed"`
		SettlementCurrencyCode        string `json:"settlementCurrencyCode"`
	} `json:"transactionDetail"`
	RecipientDetail struct {
		Type    string `json:"type"`
		Address struct {
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
			City         string `json:"city"`
			PostalCode   string `json:"postalCode"`
			State        string `json:"state"`
			Country      string `json:"country"`
		} `json:"address"`
		Bank struct {
			BankName          string `json:"bankName"`
			AccountName       string `json:"accountName"`
			AccountNumber     string `json:"accountNumber"`
			AccountNumberType string `json:"accountNumberType"`
			CountryCode       string `json:"countryCode"`
			BankCode          string `json:"bankCode"`
			BankCodeType      string `json:"bankCodeType"`
			CurrencyCode      string `json:"currencyCode"`
		} `json:"bank"`
		IdentificationList []struct {
			IdType         string `json:"idType"`
			IdNumber       string `json:"idNumber"`
			IdIssueCountry string `json:"idIssueCountry"`
		} `json:"identificationList"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		CountryOfBirth string `json:"countryOfBirth"`
		CityOfBirth    string `json:"cityOfBirth"`
		ContactNumber  string `json:"contactNumber"`
		ContactEmail   string `json:"contactEmail"`
		DateOfBirth    string `json:"dateOfBirth"`
	} `json:"recipientDetail"`
	SenderDetail struct {
		Type    string `json:"type"`
		Address struct {
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
			City         string `json:"city"`
			PostalCode   string `json:"postalCode"`
			State        string `json:"state"`
			Country      string `json:"country"`
		} `json:"address"`
		IdentificationList []struct {
			IdType         string `json:"idType"`
			IdNumber       string `json:"idNumber"`
			IdIssueCountry string `json:"idIssueCountry"`
		} `json:"identificationList"`
		SenderAccountNumber string `json:"senderAccountNumber"`
		CountryOfBirth      string `json:"countryOfBirth"`
		CityOfBirth         string `json:"cityOfBirth"`
		ContactNumber       string `json:"contactNumber"`
		ContactNumberType   string `json:"contactNumberType"`
		ContactEmail        string `json:"contactEmail"`
		Name                string `json:"name"`
		DateOfBirth         string `json:"dateOfBirth"`
	} `json:"senderDetail"`
}

type Response200 struct {
	VerificationId    string `json:"verificationId"`
	ValidationDetails []struct {
	} `json:"validationDetails"`
	ExpectedPostingDate         string `json:"expectedPostingDate"`
	ValidationResultCode        string `json:"validationResultCode"`
	ClientVerificationReference string `json:"clientVerificationReference"`
}
type ResponseErr struct {
	ErrorResponse struct {
		Reason  string `json:"reason"`
		Status  int    `json:"status"`
		Details []struct {
			Code     string `json:"code"`
			Message  string `json:"message"`
			Location string `json:"location"`
		} `json:"details"`
		Message string `json:"message"`
	} `json:"errorResponse"`
}
