package entity

type Test struct {
	ID        int `json:"id" gorm:"primary_key"`
	Signature Signature
	Answers   []QuestionAnswer `json:"answers"`
}
