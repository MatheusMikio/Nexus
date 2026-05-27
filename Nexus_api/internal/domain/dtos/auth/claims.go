package auth

import(
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	ID   uuid.UUID   `json:"id"`
	jwt.RegisteredClaims
}