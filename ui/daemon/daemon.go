package daemon

import (
	"log"
	// "net"
	"os"
	"os/signal"
	"syscall"

  "github.com/gympi/image-primitive/ui/http_server"
)

type Config struct {
	ListenSpec string

  HttpServer http_server.Config
}
//
// func Run(cfg *Config) error {
// 	log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)
//
//   l, err := net.Listen("tcp", cfg.ListenSpec)
// 	if err != nil {
// 		log.Printf("Error creating listener: %v\n", err)
// 		return err
// 	}
//
// 	http_server.Start(cfg.UI, l)
// l, err := net.Listen("tcp", cfg.ListenSpec)
// 	if err != nil {
// 		log.Printf("Error creating listener: %v\n", err)
// 		return err
// 	}
// 	waitForSignal()
//
// 	return nil
// }

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}
