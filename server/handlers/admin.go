package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

var (
	primitive_view_template = template.Must(template.ParseFiles(path.Join("server", "templates", "layout.html"), path.Join("server", "templates", "view.html")))
)

func ApiPrimitiveView(w http.ResponseWriter, r *http.Request) {
	// обработчик запросов

	body := ""

	page := Page{"Изменение картинки", template.HTML(body)}

	if err := primitive_view_template.ExecuteTemplate(w, "layout", page); err != nil {
		log.Println(err.Error())
		//http.Error(w, http.StatusText(500), 500)
		errorHandler(w, r, 500)
	}
}
