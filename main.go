package main

import (
  "flag"
  "log"

  "github.com/gympi/image-primitive/ui/http_server"
)

var cfg = &http_server.HTTPServerConfig{}

func init() {
	flag.StringVar(&cfg.Host, "host", "localhost", "Host listen spec")
  flag.IntVar(&cfg.Port, "port", 9001, "Port listen spec")

	flag.StringVar(&cfg.AssetsPath, "assets-path", "public/static", "Path to assets dir")

  flag.Parse()
}

func main() {


  http_server.Run(cfg)

  if err := http_server.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
  }
}
