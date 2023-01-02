package middlewares

import (
	"encoding/json"
	"fmt"
	"go-crud/config"
	"go-crud/models"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(ctx *fiber.Ctx) error {
	// get cookie from context
	tokenString := ctx.Cookies("AuthToken", "")

	if tokenString == "" {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration
		currentTime := float64(time.Now().Unix())
		if currentTime > claims["exp"].(float64) {
			return ctx.Status(http.StatusUnauthorized).SendString("Token expired")
		}

		// find the user with token sub
		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			return ctx.Status(http.StatusNotFound).SendString("User not found")
		}

		// atach to request
		userJSON, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		ctx.Set("user", string(userJSON))

		// continue
		return ctx.Next()

	} else {
		fmt.Println(err)
		return ctx.SendStatus(http.StatusUnauthorized)
	}

}
