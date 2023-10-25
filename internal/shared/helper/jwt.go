package helper

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/kharismajanuar/credit-app/infrastructure/auth/jwt"
)

func ExtractClaim(c *fiber.Ctx) *jwt.JwtClaims {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)

	extract := jwt.JwtClaims{
		UserID: uint(claims["userID"].(float64)),
		Email:  claims["email"].(string),
		Role:   claims["role"].(string),
	}

	return &extract
}
