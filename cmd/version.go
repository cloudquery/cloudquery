package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print full version info of cloudquery",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Version: %s\n", Version)
    fmt.Printf("Commit: %s\n", Commit)
    fmt.Printf("Date: %s\n", Date)
  },
}


func init() {
    rootCmd.AddCommand(versionCmd)
}
