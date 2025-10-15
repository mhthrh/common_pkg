package visa

type SenderDetail struct {
	Address               Address          `json:"address"`
	ContactEmail          string           `json:"contactEmail"`
	ContactNumber         string           `json:"contactNumber"`
	AdditionalData        []Additional     `json:"additionalData"`
	ContactNumberType     string           `json:"contactNumberType"`
	IdentificationList    []Identification `json:"identificationList"`
	SenderAccountNumber   string           `json:"senderAccountNumber"`
	SenderReferenceNumber string           `json:"senderReferenceNumber"`
	Name                  string           `json:"name"`
	Type                  string           `json:"type"`
	LastName              string           `json:"lastName"`
	FirstName             string           `json:"firstName"`
	CityOfBirth           string           `json:"cityOfBirth"`
	DateOfBirth           string           `json:"dateOfBirth"`
	CountryOfBirth        string           `json:"countryOfBirth"`
}
