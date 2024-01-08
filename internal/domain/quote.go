package domain

type Quote struct {
	Phrase      string `json:"phrase"`
	Author      string `json:"author"`
	Work        string `json:"work"`
	DateCreated string `json:"date_created"`
}
