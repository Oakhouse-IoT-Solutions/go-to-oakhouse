// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package templates

// DTO templates
const GetDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

import "time"

type Get{{.ModelName}}Dto struct {
	Page     int ` + "`query:\"page\" validate:\"min=1\"`" + `
	PageSize int ` + "`query:\"page_size\" validate:\"min=1,max=100\"`" + `
	
	// Add your filter fields here
{{range .Fields}}	{{.Name}} {{.QueryType}} ` + "`query:\"{{.QueryTag}}\"`" + `
{{end}}	
	// Date range filters
	StartDate time.Time ` + "`query:\"start_date\"`" + `
	EndDate   time.Time ` + "`query:\"end_date\"`" + `
}

func (dto *Get{{.ModelName}}Dto) SetDefaults() {
	if dto.Page <= 0 {
		dto.Page = 1
	}
	if dto.PageSize <= 0 {
		dto.PageSize = 10
	}
	if dto.PageSize > 100 {
		dto.PageSize = 100
	}
}
`

const CreateDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

type Create{{.ModelName}}Dto struct {
{{range .Fields}}	{{.Name}} {{.Type}} ` + "`json:\"{{.JsonTag}}\"`" + `
{{end}}
}
`

const UpdateDtoTemplate = `// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package {{.PackageName}}

type Update{{.ModelName}}Dto struct {
{{range .Fields}}	{{.Name}} *{{.Type}} ` + "`json:\"{{.JsonTag}}\"`" + `
{{end}}
}
`