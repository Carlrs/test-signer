package entity

type QuestionAnswer struct {
	ID       int    `gorm:"primary_key"`
	TestID   int    `gorm:"type:integer"`
	Question string `json:"question" gorm:"type:text"`
	Answer   string `json:"answer" gorm:"type:text"`
}
