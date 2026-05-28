package auth

import (
	"errors"
	"time"

	"github.com/MatheusMikio/Nexus/config"
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/auth"
	"github.com/MatheusMikio/Nexus/internal/domain/schemas"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const accessTokenTTL = 2 * time.Hour

func GenerateAccessToken(userUuid uuid.UUID, userRole schemas.Role) (string, error) {
	claims := auth.Claims{
		ID: userUuid,
		Role: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.GetJwtSecret()
	if secret == "" {
		return "", errors.New("invalid signing method")
	}

	return token.SignedString([]byte(secret))
}

func ValidateAccessToken(tokenString string) (*auth.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&auth.Claims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			secret := config.GetJwtSecret()

			if secret == "" {
				return nil, errors.New("jwt secret is not configured")
			}

			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
