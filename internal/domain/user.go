package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user of the system.
// This is a pure domain model.
type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

	Email    string
	Username string
}
