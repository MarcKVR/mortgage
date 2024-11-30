package handler

import (
	"github.com/MarcKVR/mortgage/auth"
	"github.com/MarcKVR/mortgage/service"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthHandler struct {
		service service.AuthService
	}

	AuthResponse struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		// Meta   *meta.Meta  `json:"meta,omitempty"`
	}

	BodyLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body BodyLogin
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if body.Email == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email y contraseña son requeridos"})
	}

	logged, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Datos proprocionados incorrectos"})
	}

	token, err := auth.GenerateToken(body.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al generar token"})
	}

	return c.JSON(AuthResponse{
		Status: 200,
		Data:   fiber.Map{"token": token, "logged": logged},
	})
}
