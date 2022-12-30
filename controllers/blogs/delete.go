package blogs

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Delete an existent blog
func Delete(ctx *fiber.Ctx) error {
	// retrieve blog id from url
	id := ctx.Params("id")

	// delete blog
	config.DB.Delete(&models.Blog{}, id)
	// respond
	return ctx.Status(http.StatusOK).SendString("Successfully deleted")
}
