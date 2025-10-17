package visa

type Identification struct {
	IdType         string `json:"idType"`
	IdNumber       string `json:"idNumber"`
	IdIssueCountry string `json:"idIssueCountry"`
}
