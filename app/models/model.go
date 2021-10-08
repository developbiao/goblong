package models

import "goblong/pkg/types"

// BaseModel
type BaseModel struct {
	ID uint64
}

// Get String ID
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
