package main

import (
	"github.com/nayyara-samuel/multi-go-mods/example/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalf("Failed with error: %v", err)
	}
}
