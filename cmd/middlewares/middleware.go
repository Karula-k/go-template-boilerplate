package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func MiddleWare() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "missing token",
			})
		}
		id, err := VerifyToken(token)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "invalid token",
			})
		}
		c.Locals("userId", id)
		return c.Next()
	}
}
