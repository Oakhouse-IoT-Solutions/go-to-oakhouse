// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Middleware template
const MiddlewareTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package middleware

import (
	"{{.ProjectName}}/util"
	"github.com/gofiber/fiber/v2"
)

func {{.MiddlewareName}}() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Add your middleware logic here
		
		// Example: Check some condition
		// if someCondition {
		//     return util.SendError(c, fiber.StatusForbidden, "Access denied", nil)
		// }
		
		return c.Next()
	}
}
`