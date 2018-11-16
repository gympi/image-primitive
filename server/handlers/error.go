package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

var (
	error_template = template.Must(template.ParseFiles(path.Join("server", "templates", "layout.html"), path.Join("server", "templates", "error.html")))
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if err := error_template.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
