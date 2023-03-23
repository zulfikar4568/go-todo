package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Todo struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;not null"`
	Description string    `gorm:"column:description"`
	// timestamp
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}
