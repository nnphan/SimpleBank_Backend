package jwt

import (
	"errors"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const minSecretKeySize = 32

// JWTMaker is an interface for generating and verifying JWT tokens
type JWTMaker struct {
	secretKey string
}

// GenerateAccessToken implements Maker.
func (maker *JWTMaker) GenerateAccessToken(userID uuid.UUID, email string, duration time.Duration) (string, error) {
	p, err := NewPayload(userID, email, duration)
	if err != nil {
		return "", fmt.Errorf("failed to create the payload: %w", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	return token.SignedString([]byte(maker.secretKey))
}

// VerifyAccessToken implements Maker.
func (maker *JWTMaker) VerifyAccessToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	// Parse the token with the claims
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrorExpiredToken) {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}
	if !jwtToken.Valid {
		return nil, ErrorInvalidToken
	}
	return jwtToken.Claims.(*Payload), nil
}

// NewJWTMaker creates a new JWTMaker with the given secret key
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}


