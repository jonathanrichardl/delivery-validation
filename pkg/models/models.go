package models

type Orders struct {
	Id           int            `json:"id"`
	Title        string         `json:"title"`
	Requirements []Requirements `json:"requirements"`
}

type Requirements struct {
	Reqest          string `json:"req"`
	ExpectedOutcome string `json:"outcome"`
	Status          bool   `json:"status"`
}
