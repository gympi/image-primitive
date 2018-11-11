package handlers

import (
  "net/http"
  "log"
  "html/template"
  "path"
)

var (
    error_template =  template.Must(template.ParseFiles(path.Join("ui", "templates", "layout.html"), path.Join("ui", "templates", "error.html")))
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
 w.WriteHeader(status)
 if err := error_template.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
  log.Println(err.Error())
  http.Error(w, http.StatusText(500), 500)
  return
 }
}
