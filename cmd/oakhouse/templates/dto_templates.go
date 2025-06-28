// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// DTO templates
const GetDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import (
	"time"
)

type Get{{.ModelName}}Dto struct {
	Page     *int ` + "`json:\"page\" query:\"page\" validate:\"omitempty,gte=1\"`" + `
	PageSize *int ` + "`json:\"pageSize\" query:\"pageSize\" validate:\"omitempty,gte=1,lte=200\"`" + `
	
	// Date range filtering
	StartDate *time.Time ` + "`json:\"start_date\" query:\"start_date\" validate:\"omitempty\"`" + `
	EndDate   *time.Time ` + "`json:\"end_date\" query:\"end_date\" validate:\"omitempty\"`" + `
	
	// Add your filter fields here
{{range .Fields}}	{{.Name}} {{.QueryType}} ` + "`query:\"{{.QueryTag}}\" validate:\"omitempty\"`" + `
{{end}}
}

func (r *Get{{.ModelName}}Dto) SetDefaults() {
	if r.Page == nil || *r.Page < 1 {
		defaultPage := 1
		r.Page = &defaultPage
	}
	
	// Set default for PageSize
	if r.PageSize == nil || *r.PageSize < 1 {
		defaultPageSize := 20
		r.PageSize = &defaultPageSize
	} else if *r.PageSize > 200 {
		maxPageSize := 200
		r.PageSize = &maxPageSize
	}
}
`

const CreateDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

type Create{{.ModelName}}Dto struct {
{{range .Fields}}	{{.Name}} {{.Type}} ` + "`json:\"{{.JsonTag}}\" validate:\"required\"`" + `
{{end}}
}
`

const UpdateDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

type Update{{.ModelName}}Dto struct {
{{range .Fields}}	{{.Name}} *{{.Type}} ` + "`json:\"{{.JsonTag}}\" validate:\"omitempty\"`" + `
{{end}}
}
`
