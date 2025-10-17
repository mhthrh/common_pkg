package visa

type Recipient struct {
	Bank               Bank             `json:"bank"`
	ContactEmail       string           `json:"contactEmail"`
	ContactNumber      string           `json:"contactNumber"`
	AdditionalData     []Additional     `json:"additionalData"`
	ContactNumberType  string           `json:"contactNumberType"`
	IdentificationList []Identification `json:"identificationList"`
	Name               string           `json:"name"`
	Type               string           `json:"type"`
	Address            Address          `json:"address"`
	LastName           string           `json:"lastName"`
	FirstName          string           `json:"firstName"`
	CityOfBirth        string           `json:"cityOfBirth"`
	DateOfBirth        string           `json:"dateOfBirth"`
	CountryOfBirth     string           `json:"countryOfBirth"`
}
