package ui

import (
    "html/template"
    "net/http"
    "io"
    "log"
    "path"
    // "fmt"
    "os"
    "math/rand"
    "time"
    "path/filepath"
)

type Page struct {
 Title string
 Body  template.HTML
}

var (
    // компилируем шаблоны, если не удалось, то выходим
    primitive_view_template = template.Must(template.ParseFiles(path.Join("templates", "layout.html"), path.Join("templates", "view.html")))

    error_template =  template.Must(template.ParseFiles(path.Join("templates", "layout.html"), path.Join("templates", "error.html")))
)

func apiPrimitiveView(w http.ResponseWriter, r *http.Request) {
    // обработчик запросов

    body := ""

    page := Page{"Изменение картинки", template.HTML(body)}

    if err := primitive_view_template.ExecuteTemplate(w, "layout", page); err != nil {
        log.Println(err.Error())
        //http.Error(w, http.StatusText(500), 500)
        errorHandler(w, r, 500)
    }
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
 w.WriteHeader(status)
 if err := error_template.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
  log.Println(err.Error())
  http.Error(w, http.StatusText(500), 500)
  return
 }
}
