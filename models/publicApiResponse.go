package models

type PublicApiResponse struct {
	Count   int `json:"count"`
	Entries []struct {
		Api         string `json:"API"`
		Description string `json:"Description"`
		Auth        string `json:"Auth"`
		Https       bool   `json:"HTTPS"`
		Cors        string `json:"Cors"`
		Link        string `json:"Link"`
		Category    string `json:"Category"`
	} `json:"entries"`
}
