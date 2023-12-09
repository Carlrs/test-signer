package entity

import (
	"github.com/google/uuid"
	"time"
)

type Signature struct {
	ID        int       `gorm:"primary_key"`
	Uuid      uuid.UUID `json:"uuid" gorm:"type:uuid,unique"`
	User      string    `gorm:"type:varchar(128)"`
	Timestamp time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
}
