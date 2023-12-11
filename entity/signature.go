package entity

import (
	"github.com/google/uuid"
	"time"
)

type Signature struct {
	ID        int       `gorm:"primary_key"`
	Uuid      uuid.UUID `json:"uuid" gorm:"type:uuid,unique"`
	User      uuid.UUID `gorm:"type:uuid"`
	Timestamp time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
}

type Signable interface {
	Sign(signature Signature)
}
