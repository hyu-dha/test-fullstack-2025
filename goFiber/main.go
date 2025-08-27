package main

import (
	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var user User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	return c.JSON(user)
}

func main() {
	app := fiber.New()

	app.Post("/", login)

	app.Listen(":3000")
}