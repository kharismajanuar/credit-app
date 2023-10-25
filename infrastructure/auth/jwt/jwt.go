package jwt

import (
	"time"

	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/kharismajanuar/credit-app/config"
)

type AuthManager struct {
	config config.EnvConfig
}

func NewAuthManager(config config.EnvConfig) *AuthManager {
	return &AuthManager{
		config: config,
	}
}

func (a *AuthManager) CreateJwtToken(claim *JwtClaims) (string, error) {
	day := time.Hour * 24

	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"userID": claim.UserID,
		"email":  claim.Email,
		"role":   claim.Role,
		"exp":    time.Now().Add(day * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(a.config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	// Return the token
	return t, nil
}
