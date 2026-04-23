package jwt

import (
	"time"

	"github.com/google/uuid"
)

// Maker is an interface for generating and verifying JWT tokens
type Maker interface {
	// GenerateToken creates a new token for a specific user email and duration
	GenerateAccessToken(userID uuid.UUID, email string, duration time.Duration) (string, error)
	// VerifyToken checks if the token is valid or not
	VerifyAccessToken(token string) (*Payload, error)
}