package helper

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

func GenerateToken(id int64) string {
	claims := jwt.MapClaims{
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(secretKey)
	return tokenString
}

func ParseToken(c *fiber.Ctx) int64 {
	authHeader := c.Get("Authorization")
	inputToken := strings.Split(authHeader, " ")[1]

	userID := 0
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		// Menggunakan secret key yang sama yang digunakan saat menandatangani token
		return secretKey, nil
	})
	if err != nil {
		return 0
	}
	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			userID = int(claims["userId"].(float64))
		} else {
			log.Fatal("Failed to get claims from token")
		}
	} else {
		log.Fatal("Invalid token")
	}
	return int64(userID)
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return fiber.ErrUnauthorized
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		c.Locals("user", token)
		return c.Next()
	}
}
