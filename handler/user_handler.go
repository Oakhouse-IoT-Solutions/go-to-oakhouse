package handler

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) FindAll(c *fiber.Ctx) error {
	return c.SendString("All Users retrieved successfully")
}

func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("User with ID " + idStr + " retrieved successfully")
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	return c.SendString("User created successfully")
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("User with ID " + idStr + " updated successfully")
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("User with ID " + idStr + " deleted successfully")
}
