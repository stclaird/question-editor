package models

type Question struct {
	Text        string `json:"text" form:"text" binding:"required"`
	Type        string `json:"type" form:"type" binding:"required"`
	Category    string `json:"category" form:"category" binding:"required,alphanum"`
	Subcategory string `json:"subcategory" form:"subcategory" binding:"required,alphanum"`
	DateAdded   string `json:"dateAdded" form:"dateadded"`
	Certification string `json:"certification" form:"certification"`
	AnswerSupport string `json:"answerSupport" form:"answersupport"`
	Answers     []Answer `json:"answers" form:"answers" binding:"required"`
}

type Answer struct {
	Text string `json:"text"`
	IsCorrect bool   `json:"iscorrect"`
}
