package models

type Question struct {
	Text        string `json:"text"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	DateAdded   string `json:"dateAdded"`
	Answers     []Answer `json:"answers"`
}

type Answer struct {
	Text string `json:"text"`
	IsCorrect bool   `json:"iscorrect"`
}
