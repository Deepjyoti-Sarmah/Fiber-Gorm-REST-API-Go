package main

import (
	"github.com/DeepjyotiSarmah/Fiber-REST-API/user"

	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Github with Go fiber and gorm")
}

func Routers(app *fiber.App){
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/users/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()

	app := fiber.New()
	app.Get("/", hello)

	Routers(app)

	app.Listen(":3000")
}
