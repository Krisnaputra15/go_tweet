package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"primary_key;default:UUID" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()

	if err != nil {
		err = errors.New("can't save invalid data")
	}

	return
}
