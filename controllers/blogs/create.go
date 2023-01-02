package blogs

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Create blog
// @desc Create blog
// @Router /api/users/ [post]
func Create(ctx *fiber.Ctx) error {
	// retrieve data from context
	var body Request

	err := ctx.BodyParser(&body)

	if err != nil {
		return err
	}

	// create a new blog
	blog := models.Blog{
		Title:    body.Title,
		Text:     body.Text,
		AuthorId: body.UserId,
	}
	result := config.DB.Create(&blog)

	if result.Error != nil {
		return result.Error
	}

	// respond
	return ctx.Status(http.StatusCreated).JSON(blog)

}
