package main

import (
    "net/http"
    "log"

    "github.com/gympi/image-primitive/ui/handlers"
)

func main() {
    // для отдачи сервером статичных файлов из папки public/static
    fs := http.FileServer(http.Dir("./public/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/admin/", handlers.ApiPrimitiveView)
    http.HandleFunc("/api/", handlers.ApiPrimitive)

    log.Println("Listening...")

    http.ListenAndServe(":9001", nil)
}
