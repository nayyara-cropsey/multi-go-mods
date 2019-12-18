package cmd

import (
	"github.com/nayyara-samuel/multi-go-mods/common/pkg/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	CLIName = "example"
)

var (
	BuildVersion = "unknown"
	verbose      bool
)

var RootCmd = &cobra.Command{
	Use:     CLIName,
	Version: BuildVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.WithFields(log.Fields{
			"context":   "share-command",
			"operation": "validation",
		})

		client := http.HttpClient{
			BaseUrl: "https://api.github.com",
		}
		client.Init()

		if verbose {
			logger.Debugf("Verbose mode set")
		}
		logger.Info("Running")
		response, err := client.Request("GET", "/user", []byte{}, map[string]string{})
		if err != nil {
			logger.Errorf("Got error response: %v: %v", string(response), err)
			return err
		}

		logger.Infof("Got response: %v", string(response))
		return nil
	},
}
