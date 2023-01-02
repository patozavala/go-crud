package users

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login function for users
// @desc Login for users
// @Router /api/users/:id [post]
func Login(ctx *fiber.Ctx) error {
	// retrieve data from context
	var body Request

	if ctx.BodyParser(&body) != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Failed to read body")
	}

	// Look up the user
	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		return ctx.Status(http.StatusBadRequest).SendString("Invalid email")
	}

	// compare hashed password with the hash of the password provided by the user
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Invalid password")
	}
	// generate a jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Failed to create token")
	}

	// send token back to the user as a cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "AuthToken",
		Value:    tokenString,
		MaxAge:   3600 * 24,
		HTTPOnly: true,
		Secure:   false,
	})

	return ctx.Status(http.StatusOK).SendString("Now you are logged in 👋!")
}
