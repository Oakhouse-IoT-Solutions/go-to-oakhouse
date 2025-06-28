// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Handler templates
const HandlerTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package handler

import (
	"github.com/gofiber/fiber/v2"
)

type {{.ModelName}}Handler struct {
}

func New{{.ModelName}}Handler() *{{.ModelName}}Handler {
	return &{{.ModelName}}Handler{}
}

func (h *{{.ModelName}}Handler) FindAll(c *fiber.Ctx) error {
	return c.SendString("All {{.ModelName}}s retrieved successfully")
}

func (h *{{.ModelName}}Handler) FindByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("{{.ModelName}} with ID " + idStr + " retrieved successfully")
}

func (h *{{.ModelName}}Handler) Create(c *fiber.Ctx) error {
	return c.SendString("{{.ModelName}} created successfully")
}

func (h *{{.ModelName}}Handler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("{{.ModelName}} with ID " + idStr + " updated successfully")
}

func (h *{{.ModelName}}Handler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	return c.SendString("{{.ModelName}} with ID " + idStr + " deleted successfully")
}
`