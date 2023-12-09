package entity

import (
	"github.com/google/uuid"
	"time"
)

type Signature struct {
	ID        int       `gorm:"primary_key"`
	Uuid      uuid.UUID `json:"uuid" gorm:"type:uuid,unique"`
	User      string    `json:"user" gorm:"type:varchar(128)"`
	Timestamp time.Time `json:"timestamp" gorm:"DEFAULT:CURRENT_TIMESTAMP"`
}
