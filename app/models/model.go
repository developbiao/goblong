package models

import (
	"goblong/pkg/types"
	"time"
)

// BaseModel
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"created_at;index"`
	UpdatedAt time.Time `gorm:"updated_at;index"`
}

// Get String ID
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
