package handler

import (
	"os"
	"strconv"

	"github.com/MarcKVR/mortgage/domain"
	"github.com/MarcKVR/mortgage/packages/meta"
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
		Meta   *meta.Meta  `json:"meta,omitempty"`
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

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	filters := service.Filters{
		Name:  c.Query("name", ""),
		Email: c.Query("email", ""),
	}

	limit, _ := strconv.Atoi(c.Query("limit", os.Getenv("PAGINATOR_LIMIT_DEFAULT")))
	page, _ := strconv.Atoi(c.Query("page", os.Getenv("DEFAULT_PAGE")))

	total, err := h.service.Count(service.Filters{
		Name:  filters.Name,
		Email: filters.Email,
	})
	if err != nil {
		c.Status(500)
		return c.JSON(Response{Status: 500, Err: err.Error()})
	}

	meta, err := meta.New(page, limit, int(total))
	if err != nil {
		c.Status(400)
		return c.JSON(Response{Status: 400, Err: err.Error()})
	}

	users, err := h.service.GetUsers(filters, meta.Limit(), meta.Offset())
	if err != nil {
		c.Status(500)
		return c.JSON(Response{Status: 500, Err: err.Error()})
	}

	return c.JSON(UserResponse{
		Status: 200,
		Data:   users,
		Meta:   meta,
	})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		c.Status(400)
		return c.JSON(Response{Status: fiber.StatusBadRequest, Err: "id is required"})
	}

	var userInput struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := domain.User{
		Name:  userInput.Name,
		Email: userInput.Email,
		ID:    id,
	}

	if err := h.service.Update(id, &user); err != nil {
		c.Status(422)
		return c.JSON(Response{Status: fiber.StatusUnprocessableEntity, Err: err.Error()})
	}

	return c.JSON(Response{
		Status: 200,
		Data:   map[string]string{"id": user.ID, "name": user.Name, "email": user.Email},
	})
}
