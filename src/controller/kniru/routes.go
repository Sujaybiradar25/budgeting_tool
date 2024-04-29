package kniru

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SimpleGet(c *fiber.Ctx) error {
	c.Context()
	msg := fmt.Sprintf("Hello World!")
	return c.SendString(msg)

}
