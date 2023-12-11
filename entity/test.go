package entity

type Test struct {
	ID        int `json:"id" gorm:"primary_key"`
	Signature Signature
	Answers   []QuestionAnswer `json:"answers"`
}

func (t Test) Sign(signature Signature) {
	t.Signature = signature
}
