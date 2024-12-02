package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"

	"github.com/spf13/cobra"
)

const (
	// logout command
	logoutShort = "Log out of CloudQuery Hub."
)

func newCmdLogout() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "logout",
		Short: logoutShort,
		Args:  cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runLogout(ctx, cmd)
		},
	}
	return loginCmd
}

func runLogout(_ context.Context, cmd *cobra.Command) error {
	err := auth.Logout()
	if err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}

	cmd.Println("CLI successfully logged out.")

	return nil
}
