package visa

type ErrorResponse struct {
	Status  int      `json:"status"`
	Reason  string   `json:"reason"`
	Details []Detail `json:"details"`
}

type Detail struct {
	Location string `json:"location"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
}
