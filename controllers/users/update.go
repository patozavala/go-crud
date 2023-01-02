package users

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Update user
// @desc Update user
// @Router /api/users/:id [put]
func Update(ctx *fiber.Ctx) error {
	// retrieve user id from url
	id := ctx.Params("id")

	// retrieve data from context
	var body Request

	err := ctx.BodyParser(&body)

	if err != nil {
		return err
	}

	// find the user by id
	var user models.User
	config.DB.First(&user, id)

	// update
	config.DB.Model(&user).Updates(models.User{
		Username: body.Username,
		Email:    body.Email,
	})

	// respond
	return ctx.Status(http.StatusOK).JSON(user)
}
