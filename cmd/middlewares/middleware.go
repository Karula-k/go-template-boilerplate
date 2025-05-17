package middlewares

import (
	"github.com/go-template-boilerplate/cmd/models"
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(env *models.EnvModel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "missing token",
			})
		}
		id, err := VerifyToken(token, env)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "invalid token",
			})
		}
		c.Locals("userId", id)
		return c.Next()
	}
}
