package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

// Task is used by pop to map your tasks database table to your go code.
type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Description string    `json:"description" db:"description"`
	ExecutedBy  string    `json:"executed_by" db:"executed_by"`
	RequestedBy string    `json:"requested_by" db:"requested_by"`
	ExecutedAt  time.Time `json:"executed_at" db:"executed_at"`
	IsDone      bool      `json:"is_done" db:"is_done"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Task) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}
