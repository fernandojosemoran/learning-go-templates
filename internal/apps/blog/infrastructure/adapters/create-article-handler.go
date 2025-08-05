package blog

import (
	"html/template"
	"net/http"
	"net/url"

	blog "github.com/fernandojosemoran/go-templates/internal/apps/blog/infrastructure/dto/article"
)

func CreateArticleController(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var params url.Values = request.URL.Query()

	var title string = params.Get("title")
	var content string = params.Get("content")

	var dto blog.CreateArticleDto = blog.CreateArticleDto{
		Title:   title,
		Content: template.HTML(content),
	}

	newDto, err := dto.Create()

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	dto = *newDto

}
