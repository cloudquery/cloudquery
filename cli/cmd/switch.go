package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/internal/team"
	"github.com/spf13/cobra"
)

const (
	switchShort   = "Switches between teams."
	switchLong    = `Switches between teams.`
	switchExample = `
# Switch to a different team
cloudquery switch my-team
`
)

func newCmdSwitch() *cobra.Command {
	switchCmd := &cobra.Command{
		Use:     "switch",
		Short:   switchShort,
		Long:    switchLong,
		Example: switchExample,
		Args:    cobra.MaximumNArgs(1),
		Hidden:  true,
		RunE:    runSwitch,
	}
	return switchCmd
}

func runSwitch(cmd *cobra.Command, args []string) error {
	apiURL := getEnvOrDefault(envAPIURL, defaultAPIURL)

	tc := auth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	cl, err := team.NewClient(apiURL, token)
	if err != nil {
		return fmt.Errorf("failed to create API client: %w", err)
	}

	if len(args) == 0 {
		// Print the current team context
		currentTeam, err := config.GetValue("team")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to get current team: %w", err)
		}

		allTeams, err := cl.ListAllTeams(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list all teams: %w", err)
		}

		if currentTeam == "" {
			cmd.Println("Your team is not set.")
			if len(allTeams) == 1 {
				cmd.Println("As you are currently a member of only one team, this will be used as your default team.")
			}
		} else {
			cmd.Printf("Your current team is set to %v.\n\n", currentTeam)
		}
		cmd.Println("Teams available to you:", strings.Join(allTeams, ", ")+"\n")
		cmd.Println("To switch teams, run `cloudquery switch <team>`")
		return nil
	}
	selectedTeam := args[0]
	err = cl.ValidateTeam(cmd.Context(), selectedTeam)
	if err != nil {
		return fmt.Errorf("failed to switch teams: %w", err)
	}
	err = config.SetValue("team", selectedTeam)
	if err != nil {
		return fmt.Errorf("failed to set team value: %w", err)
	}
	cmd.Printf("Successfully switched teams to %v.\n", selectedTeam)
	return nil
}
