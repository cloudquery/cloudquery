package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/cloudquery/cloudquery/cli/internal/manifests"
	"github.com/cloudquery/cloudquery/cli/internal/packages"
	"github.com/spf13/cobra"
)

const (
	packageShort = "Package a CloudQuery plugin for distribution."
	packageLong  = `Package a CloudQuery plugin for distribution.

This creates a directory with the plugin binaries, package.json and documentation
required for the publish command.
`
	packageExample = `# Package a plugin with pre-built binary called my_plugin
cloudquery package ./my_plugin

# Package a plugin using gRPC server
cloudquery package --registry grpc 'localhost:7777'
`
)

// cloudquery package ./path [--docs ./docs] [--build ./build] [--registry local]
func newCmdPackage() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "package",
		Short:   packageShort,
		Long:    packageLong,
		Example: packageExample,
		Hidden:  true,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runPackage(ctx, cmd, args)
		},
	}

	cmd.Flags().String("docs", "docs", "Path to the plugin documentation directory.")
	cmd.Flags().String("build", "build", "Path to the plugin build directory.")
	cmd.Flags().String("registry", "local", "Registry to use when connecting to the plugin.")
	cmd.Flags().String("dist", "dist", "Path to the plugin distribution directory that will be created or updated.")
	return cmd
}

func runPackage(ctx context.Context, cmd *cobra.Command, args []string) error {
	fmt.Println("Hello from package command!")

	distDir := cmd.Flag("dist").Value.String()
	err := os.MkdirAll(distDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create dist directory: %w", err)
	}

	m, err := manifests.Read("manifest.yaml")
	if err != nil {
		return fmt.Errorf("failed to read manifest file: %w", err)
	}

	pj := packages.PackageJSON{
		SchemaVersion: 1,
		Kind:          m.Kind,
		Properties: packages.PluginProperties{
			Source:      m.Properties.Source,
			Destination: m.Properties.Destination,
		},
		Artifacts: cmd.Flag("build").Value.String(),
		Docs:      cmd.Flag("docs").Value.String(),
	}

	packageJSONPath := filepath.Join(distDir, "package.json")
	f, err := os.OpenFile(packageJSONPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0744)
	if err != nil {
		return fmt.Errorf("failed to create package.json: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(pj)
	if err != nil {
		return err
	}
	return nil
}
