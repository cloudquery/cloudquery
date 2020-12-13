package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

// Injected with at build time with -ldflags "-X github.com/cloudquery/cloudquery/cmd.Variable=Value"
var Version = "development"
var Commit = "development"
var Date = time.Now().String()

var rootCmd = &cobra.Command{
	Use:     "cloudquery",
	Short:   "cloudquery CLI",
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
