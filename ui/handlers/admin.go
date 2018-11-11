package handlers

import (
    "html/template"
    "net/http"
    "log"
    "path"
)

var (
    primitive_view_template = template.Must(template.ParseFiles(path.Join("ui", "templates", "layout.html"), path.Join("ui", "templates", "view.html")))
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
