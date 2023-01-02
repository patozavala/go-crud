# GO-CRUD

I want to learn [Go](https://go.dev/), so I implemented a basic API with [FIBER](https://gofiber.io/) and [GORM](https://github.com/go-gorm/gorm) which implements:

- CRUD operations for a simple Users model
- Users login where [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt) is used
- CRUD operations for a simple blog model (connecting to users by foreign key)
- Connection to a database (I used [postgresql](https://www.elephantsql.com/))
- Environment variables with [godotenv](https://github.com/joho/godotenv)


# Disclaimer

This is an introductory project to look how Go works. This project is not production ready and it has some bad practices. All improvements and comments are very welcome.

# Future work

Some improvements need to be implemented before the app is ready for production

- Implement middlewares
- Implement testing
- Accelerate queries
- Add descriptions and metadata to functions
