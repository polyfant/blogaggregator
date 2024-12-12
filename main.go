package main

import (
	"fmt"
	"log"
	"github.com/polyfant/gator/internal/config"
)

func main() {
	// Read initial config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initial config:", cfg)

	// Update user and save
	err = cfg.SetUser("Jonas")
	if err != nil {
		log.Fatal(err)
	}

	// Read updated config
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated config:", updatedCfg)
}