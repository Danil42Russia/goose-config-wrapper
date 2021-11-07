package main

import (
	"log"
	"os"

	"github.com/Danil42Russia/goose-config-wrapper/pkg/config"
	"github.com/Danil42Russia/goose-config-wrapper/pkg/goose"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("Command not found")
	}
	command := args[0]

	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Can't read config: %v", err)
	}

	cmd := goose.New(cfg)
	log.Println(cmd.GenRunCmd(command))

	os.Exit(0)
}
