package visa

type Recipient struct {
	Bank      Bank   `json:"bank"`
	Type      string `json:"type"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}
