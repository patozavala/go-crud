package users

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Delete an existent user
// @desc Delete user
// @Router /api/users/:id [delete]
func Delete(ctx *fiber.Ctx) error {
	// retrieve user id from url
	id := ctx.Params("id")

	// delete user
	config.DB.Delete(&models.User{}, id)
	// respond
	return ctx.Status(http.StatusOK).SendString("Successfully deleted")
}
