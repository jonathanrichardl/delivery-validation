package models

type Orders struct {
	Id           int            `json:"id,omitempty"`
	Title        string         `json:"title"`
	Requirements []Requirements `json:"requirements"`
}

type Requirements struct {
	Requirementid   int    `json:"reqid,omitempty"`
	Request         string `json:"req"`
	ExpectedOutcome string `json:"outcome"`
	Status          bool   `json:"status"`
}

type ProgressForm struct {
	Fufillments []Fufillment `json:"fufillments"`
}

type Fufillment struct {
	Requirementid int    `json:"reqid"`
	Outcome       string `json:"outcome"`
}
