package routers

import (
	"go-crud/controllers/blogs"
	"go-crud/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {

	api := app.Group("/api")

	//http://localhost:3000/api/users
	usersGroup := api.Group("/users")
	usersGroup.Get("/", users.GetAll)
	usersGroup.Get("/:id", users.GetById)
	usersGroup.Post("/signup", users.SignUp)
	usersGroup.Post("/login", users.Login)
	usersGroup.Put("/:id", users.Update)
	usersGroup.Delete("/:id", users.Delete)

	//http://localhost:3000/api/blogs
	blogsGroup := api.Group("/blogs")
	blogsGroup.Get("/", blogs.GetAll)
	blogsGroup.Get("/:id", blogs.GetById)
	blogsGroup.Post("/", blogs.Create)
	blogsGroup.Put("/:id", blogs.Update)
	blogsGroup.Delete("/:id", blogs.Delete)

}
