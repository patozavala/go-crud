package users

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// List all users
// @desc List all users
// @Router /api/users/ [get]
func GetAll(ctx *fiber.Ctx) error {
	// get all users
	var users []models.User
	config.DB.Find(&users)
	// respond
	return ctx.Status(http.StatusOK).JSON(users)
}
