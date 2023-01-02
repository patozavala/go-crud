package users

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Get user by id
// @desc Get user by id
// @Router /api/users/:id [get]
func GetById(ctx *fiber.Ctx) error {
	// get user id from the url
	id := ctx.Params("id")

	// find user by id
	var user models.User
	config.DB.First(&user, id)

	// respond
	return ctx.Status(http.StatusOK).JSON(user)
}
