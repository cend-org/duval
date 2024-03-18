package route

import (
	"duval/internal/route/api"
	"duval/internal/route/docs"
)

var RootRoutesGroup = []docs.RootDocumentation{
	{
		Group:       "/api",
		Description: "",
		Paths:       api.Routes,
		IsPublic:    true,
	},
	{
		Group:       "/translation",
		Description: "",
		Paths:       api.TranslationRoutes,
		IsPublic:    true,
	},
}
