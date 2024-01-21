package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returns by the VerifyToken function
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains the payload data of token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	UserId    int64     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `jsom:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, user_id int64, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		UserId:    user_id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
