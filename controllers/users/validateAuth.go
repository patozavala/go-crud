package users

import (
	"encoding/json"
	"go-crud/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Validate that a user is authorized (logged in)
// @desc Validate that a user is logged in
// @Router /api/users/:id/validate [get]
func ValidateAuth(ctx *fiber.Ctx) error {
	userString := ctx.GetRespHeader("User", "")

	if userString == "" {
		return ctx.Status(http.StatusNotFound).SendString("User not found")
	}

	var user models.User

	err := json.Unmarshal([]byte(userString), &user)
	if err != nil {
		panic(err)
	}

	return ctx.Status(http.StatusOK).JSON(user)
}
