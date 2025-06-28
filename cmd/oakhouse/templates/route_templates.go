// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Resource route template
const ResourceRouteTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package route

import (
	"{{.ProjectName}}/handler"
	"github.com/gofiber/fiber/v2"
)

// Setup{{.Name}}Routes sets up routes for {{.Name}} resource
func Setup{{.Name}}Routes(api fiber.Router) {
	// Initialize handler
	{{.LowerName}}Handler := handler.New{{.Name}}Handler()

	// Setup routes
	{{.LowerName}}Group := api.Group("/{{.LowerName}}s")
	{{.LowerName}}Group.Get("/", {{.LowerName}}Handler.FindAll)
	{{.LowerName}}Group.Get("/:id", {{.LowerName}}Handler.FindByID)
	{{.LowerName}}Group.Post("/", {{.LowerName}}Handler.Create)
	{{.LowerName}}Group.Put("/:id", {{.LowerName}}Handler.Update)
	{{.LowerName}}Group.Delete("/:id", {{.LowerName}}Handler.Delete)
}
`