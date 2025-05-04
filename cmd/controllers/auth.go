package controllers

import (
	"context"

	"github.com/go-template-boilerplate/cmd/middlewares"
	"github.com/go-template-boilerplate/cmd/utils"
	"github.com/go-template-boilerplate/generated"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(ctx context.Context, queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Params("username")
		password := c.Params("password")

		user, err := queries.GetUserByUsername(ctx, username)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "username not found"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid password"})

		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "logged in"})
	}
}

func RegisterController(ctx context.Context, queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Params("username")
		password := c.Params("password")

		hashedPassword, err := utils.HashedPassword(password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
		}
		user, err := queries.CreateUsers(ctx, generated.CreateUsersParams{
			Username: username,
			Password: hashedPassword,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
		}
		token, refreshToken, err := middlewares.GeneratedAccessAndRefreshTokens(&user)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create token"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "user created", "data": user,
			"access_token":  token,
			"refresh_token": refreshToken,
		})
	}
}

func RefreshToken(ctx context.Context, queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Params("refreshToken")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing refresh token"})
		}

		userID, error := middlewares.VerifyToken(authHeader)

		if error != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid jwt"})

		}

		user, err := queries.GetUsers(ctx, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get user"})
		}

		token, refreshToken, err := middlewares.GeneratedAccessAndRefreshTokens(&user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get refresh token"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true, "message": "token created",
			"token": token, "refreshToken": refreshToken,
		})

	}

}
