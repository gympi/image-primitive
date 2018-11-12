package http_server

import (
    "net/http"
    "log"
    "bytes"
    "strconv"

    "github.com/gympi/image-primitive/server/handlers"
)

type Config struct {
  Host string
  Port int
	AssetsPath string
}

func Run(cfg Config) error {
    // для отдачи сервером статичных файлов из папки public/static
    fs := http.FileServer(http.Dir(cfg.AssetsPath))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/admin/", handlers.ApiPrimitiveView)
    http.HandleFunc("/api/", handlers.ApiPrimitive)

    log.Println("Listening...")

    var listen_spec bytes.Buffer
  	listen_spec.WriteString(cfg.Host)
  	listen_spec.WriteString(":")
    listen_spec.WriteString(strconv.Itoa(cfg.Port))

    log.Printf("Starting, HTTP on: %s\n", listen_spec.String())

    go http.ListenAndServe(listen_spec.String(), nil)

    return nil
}
