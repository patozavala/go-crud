package blogs

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Get blog by id
// @desc Get blog by id
// @Router /api/users/:id [get]
func GetById(ctx *fiber.Ctx) error {
	// get blog id from the url
	id := ctx.Params("id")

	// find blog by id
	var blog models.Blog
	config.DB.First(&blog, id)

	// respond
	return ctx.Status(http.StatusOK).JSON(blog)
}
