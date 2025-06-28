// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// Model templates
const ModelTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{.ModelName}} struct {
	ID        uuid.UUID      ` + "`gorm:\"type:uuid;primary_key;default:gen_random_uuid()\" json:\"id\"`" + `
{{range .Fields}}	{{.Name}} {{.Type}} ` + "`gorm:\"{{.GormTag}}\" json:\"{{.JsonTag}}\"`" + `
{{end}}	CreatedAt time.Time      ` + "`gorm:\"autoCreateTime\" json:\"created_at\"`" + `
	UpdatedAt time.Time      ` + "`gorm:\"autoUpdateTime\" json:\"updated_at\"`" + `
	DeletedAt gorm.DeletedAt ` + "`gorm:\"index\" json:\"deleted_at,omitempty\"`" + `
}

func ({{.ModelName}}) TableName() string {
	return "{{.TableName}}"
}
`