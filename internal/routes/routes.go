package routes

import (
	user "github.com/fernandojosemoran/go-templates/internal/apps/user/infrastructure/adapters"
	"github.com/fernandojosemoran/go-templates/pkg/enums"
	"github.com/fernandojosemoran/go-templates/pkg/server"
)

const (
	get    enums.Method = 1
	post   enums.Method = 2
	put    enums.Method = 3
	delete enums.Method = 4
	patch  enums.Method = 5
)

var Handlers []server.Controller = []server.Controller{
	{
		Method:  enums.GetMethod(get),
		Path:    "/",
		Handler: user.RenderTemplateAdapter,
	},
	{
		Method:  enums.GetMethod(get),
		Path:    "/sentences",
		Handler: user.SentenceHandler,
	},
}
