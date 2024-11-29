package handler

import (
	"github.com/MarcKVR/mortgage/domain"
	"github.com/MarcKVR/mortgage/service"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler struct {
		service service.UserService
	}

	UserResponse struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		// Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newUser, err := h.service.Create(&user)
	if err != nil {
		c.Status(422)
		return c.JSON(Response{Status: 422, Err: err.Error()})
	}

	return c.JSON(UserResponse{
		Status: 201,
		Data:   map[string]string{"id": newUser.ID, "name": newUser.Name, "email": newUser.Email},
	})
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.Get(id)

	if err != nil {
		c.Status(404)
		return c.JSON(Response{Status: 404, Err: err.Error()})
	}

	return c.JSON(Response{
		Status: 200,
		Data:   map[string]string{"id": user.ID, "name": user.Name, "email": user.Email},
	})
}
