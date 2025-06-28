// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Resource route template
const ResourceRouteTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package route

import (
	"{{.ProjectName}}/handler"
	"{{.ProjectName}}/repository"
	"{{.ProjectName}}/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Setup{{.Name}}Routes sets up routes for {{.Name}} resource
func Setup{{.Name}}Routes(api fiber.Router, db *gorm.DB) {
	// Initialize repository
	{{.LowerName}}Repo := repository.New{{.Name}}Repository(db)
	
	// Initialize service
	{{.LowerName}}Service := service.New{{.Name}}Service({{.LowerName}}Repo)
	
	// Initialize handler
	{{.LowerName}}Handler := handler.New{{.Name}}Handler({{.LowerName}}Service)

	// Setup routes
	{{.LowerName}}Group := api.Group("/{{.LowerName}}s")
	{{.LowerName}}Group.Get("/", {{.LowerName}}Handler.FindAll)
	{{.LowerName}}Group.Get("/:id", {{.LowerName}}Handler.FindById)
	{{.LowerName}}Group.Post("/", {{.LowerName}}Handler.Create)
	{{.LowerName}}Group.Put("/:id", {{.LowerName}}Handler.Update)
	{{.LowerName}}Group.Delete("/:id", {{.LowerName}}Handler.Delete)
}
`