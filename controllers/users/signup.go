package users

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Create a new user (SignUp)
// @desc Create user (SignUp)
// @Router /api/users/ [post]
func SignUp(ctx *fiber.Ctx) error {
	// retrieve data from context
	var body Request

	err := ctx.BodyParser(&body)

	if err != nil {
		return err
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return err
	}

	// create a new user
	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}
	result := config.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	// respond
	return ctx.Status(http.StatusCreated).JSON(user)

}
