package blogs

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Update blog
func Update(ctx *fiber.Ctx) error {
	// retrieve blog id from url
	id := ctx.Params("id")

	// retrieve data from context
	var body Request

	err := ctx.BodyParser(&body)

	if err != nil {
		return err
	}

	// find the blog by id
	var blog models.Blog
	config.DB.First(&blog, id)

	// update
	config.DB.Model(&blog).Updates(models.Blog{
		Title:    body.Title,
		Text:     body.Text,
		AuthorId: body.UserId,
	})

	// respond
	return ctx.Status(http.StatusOK).JSON(blog)
}
