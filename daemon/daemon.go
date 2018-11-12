package daemon

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gympi/image-primitive/server/http_server"
)

type Config struct {
	HTTPServer http_server.Config
}

func Run(cfg *Config) error {

  if err := http_server.Run(cfg.HTTPServer); err != nil {
		log.Printf("Error in http_server.Run(): %v", err)
  }

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}
