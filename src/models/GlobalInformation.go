package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GlobalInformation struct {
	gorm.Model            // Adds some metadata fields to the table
	ID          uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	CompanyName string
	Description string
}
