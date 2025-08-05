package user

import (
	"bytes"
	"html/template"
	"net/http"

	user "github.com/fernandojosemoran/go-templates/internal/apps/user/infrastructure/dto"
)

func RenderTemplateAdapter(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl := "Hello {{ .Name }}"

	name := request.URL.Query().Get("name")

	dto, err := user.CreateUserDto(name)

	if err != nil {
		http.ServeFile(response, request, "web/404.html")
		return
	}

	t, err := template.New("test").Parse(tmpl)

	if err != nil {
		http.Error(response, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = t.Execute(response, dto)

	if err != nil {
		http.Error(response, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func SentenceHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := struct {
		IsLoggedIn bool
		Username   string
		Email      string
		Role       string
	}{
		IsLoggedIn: false,
		Username:   "fernando",
		Email:      "fernando@example.com",
		Role:       "admin",
	}

	tmpl, err := template.ParseFiles("web/about.html")

	if err != nil {
		http.Error(response, "Error parsing template", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		http.Error(response, "Error executing template", http.StatusInternalServerError)
		return
	}

	// Si todo salió bien, escribimos el contenido en la respuesta
	buf.WriteTo(response)
}
