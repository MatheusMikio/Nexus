package auth

import (
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	ID   uuid.UUID    `json:"id"`
	Role schemas.Role `json:"role"`
	jwt.RegisteredClaims
}
