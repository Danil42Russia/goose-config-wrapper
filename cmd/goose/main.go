package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Danil42Russia/goose-config-wrapper/pkg/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Println(fmt.Sprintf("%+v", cfg))

	os.Exit(0)
}
