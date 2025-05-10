package controllers

import (
	"context"

	"github.com/go-template-boilerplate/cmd/middlewares"
	"github.com/go-template-boilerplate/cmd/models"
	"github.com/go-template-boilerplate/cmd/utils"
	"github.com/go-template-boilerplate/generated"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login by username and password
//
//	@Summary		Login
//	@Description	Login by username and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.LoginRequest	true	"login body" extensions(x-order=1)
//	@Success		200		{string}	string
//	@Router			/auth/login [post]
func LoginController(ctx context.Context, queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		user, err := queries.GetUserByUsername(ctx, req.Username)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "username not found"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid password"})

		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "logged in"})
	}
}

// Login by username and password
//
//	@Summary		register
//	@Description	register by username and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RegisterRequest	true	"login body" extensions(x-order=1)
//	@Success		200		{string}	string
//	@Router			/auth/register [post]
func RegisterController(ctx context.Context, queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}
		hashedPassword, err := utils.HashedPassword(req.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
		}
		user, err := queries.CreateUsers(ctx, generated.CreateUsersParams{
			Username: req.Username,
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
