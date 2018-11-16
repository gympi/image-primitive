package gc

import (
	"log"
	"math/rand"
	"time"
)

type Config struct {
	Timeout int
}

func Run(cfg Config) error {
	log.Printf("Starting garbage collector images on")

	for true {
		log.Printf("Run garbage collector...")
		amt := time.Duration(rand.Intn(cfg.Timeout))
		time.Sleep(time.Second * amt)
	}

	return nil
}
