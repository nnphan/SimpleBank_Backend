package jwt

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var ErrorExpiredToken = fmt.Errorf("token has expired")
var ErrorInvalidToken = fmt.Errorf("token is invalid")

// Payload defines the structure of the JWT payload
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

// Valid implements jwt.Claims.
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return ErrorExpiredToken
	}
	return nil
}

// NewPayload creates a new payload with a specific user ID, email, and role
func NewPayload(userID uuid.UUID, email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration), // Set expiration based on the provided duration
	}
	return payload, nil
}
