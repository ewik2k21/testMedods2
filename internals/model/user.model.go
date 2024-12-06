package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email        string    `gorm:"unique;not_null"`
	UserName     string    `gorm:"not_null"`
	PasswordHash string    `gorm:"not_null"`
}
