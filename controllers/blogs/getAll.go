package blogs

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// List all blogs
// @desc List all blog
// @Router /api/users/ [get]
func GetAll(ctx *fiber.Ctx) error {
	// get all blogs
	var blogs []models.Blog
	config.DB.Find(&blogs)
	// respond
	return ctx.Status(http.StatusOK).JSON(blogs)
}
