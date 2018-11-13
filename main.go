package main

import (
  "flag"
  "log"

  "github.com/gympi/image-primitive/daemon"
)

var cfg = &daemon.Config{}

func init() {
	flag.StringVar(&cfg.HTTPServer.Host, "host", "localhost", "Host listen spec")
  flag.IntVar(&cfg.HTTPServer.Port, "port", 9001, "Port listen spec")

  flag.IntVar(&cfg.GC.Timeout, "gc_timeout", 60, "Garbage collector timeout")

	flag.StringVar(&cfg.HTTPServer.AssetsPath, "assets-path", "public/static", "Path to assets dir")

  flag.Parse()
}

func main() {
  if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
  }
}
