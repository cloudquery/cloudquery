package cmd

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	cqauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/internal/api"
	"github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

const (
	initShort   = `Generate a configuration file for a sync`
	initExample = ``
)

var (
	sourcesOrder      = []string{"aws", "azure", "gcp"}
	destinationsOrder = []string{"postgresql", "bigquery"}
)

func newCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   initShort,
		Long:    initShort,
		Example: initExample,
		Args:    cobra.ExactArgs(0),
		RunE:    initCmd,
	}
	cmd.Flags().String("source", "", "Source plugin name or path")
	cmd.Flags().String("destination", "", "Destination plugin name or path")
	cmd.Flags().String("spec-path", "", "Output spec file path")
	return cmd
}

func normalizePluginPath(pluginNameOrPath string) (string, error) {
	parts := strings.Split(pluginNameOrPath, "/")
	if len(parts) == 1 {
		return "cloudquery/" + pluginNameOrPath, nil
	}
	if len(parts) != 2 {
		return "", errors.New("invalid plugin path")
	}
	return pluginNameOrPath, nil
}

func parseFlags(cmd *cobra.Command) (source, destination, specPath string, allErrors error) {
	source, err := cmd.Flags().GetString("source")
	allErrors = errors.Join(allErrors, err)
	if source != "" {
		source, err = normalizePluginPath(source)
		allErrors = errors.Join(allErrors, err)
	}
	destination, err = cmd.Flags().GetString("destination")
	allErrors = errors.Join(allErrors, err)
	if destination != "" {
		destination, err = normalizePluginPath(destination)
		allErrors = errors.Join(allErrors, err)
	}
	specPath, err = cmd.Flags().GetString("spec-path")
	allErrors = errors.Join(allErrors, err)
	return source, destination, specPath, allErrors
}

func pluginFilter(pluginPath string, kind cqapi.PluginKind) func(plugin cqapi.ListPlugin) bool {
	return func(plugin cqapi.ListPlugin) bool {
		return plugin.TeamName+"/"+plugin.Name == pluginPath && plugin.Kind == kind
	}
}

func pluginName(plugin cqapi.ListPlugin, _ int) string {
	return plugin.Name
}

func officialReleasedPluginsByKind(kind cqapi.PluginKind) func(plugin cqapi.ListPlugin, _ int) bool {
	return func(plugin cqapi.ListPlugin, _ int) bool {
		return plugin.Kind == kind && plugin.Official && plugin.ReleaseStage != cqapi.PluginReleaseStageComingSoon
	}
}

func pluginsSorter(plugins []cqapi.ListPlugin, prioritySlice []string) func(a, b int) bool {
	return func(a, b int) bool {
		indexOfA := lo.IndexOf(prioritySlice, plugins[a].Name)
		indexOfB := lo.IndexOf(prioritySlice, plugins[b].Name)
		if indexOfA == -1 && indexOfB != -1 {
			return false
		}
		if indexOfA != -1 && indexOfB == -1 {
			return true
		}
		if indexOfA == -1 && indexOfB == -1 {
			return plugins[a].Name < plugins[b].Name
		}
		return indexOfA < indexOfB
	}
}

func initCmd(cmd *cobra.Command, args []string) error {
	source, destination, _, err := parseFlags(cmd)
	if err != nil {
		return err
	}

	authClient := cqauth.NewTokenClient()
	token, err := authClient.GetToken()
	var user *cqapi.User
	if err == nil {
		user, _ = auth.GetUser(cmd.Context(), token)
	}

	apiClient, err := api.NewAnonymousClient()
	if err != nil {
		return err
	}
	if user != nil {
		apiClient, err = api.NewClient(token.Value)
		if err != nil {
			return err
		}
	}

	allPlugins, err := api.ListAllPlugins(apiClient)
	if err != nil {
		return err
	}

	var notFoundPlugins error
	if source != "" {
		sourcePluginFilter := pluginFilter(source, cqapi.PluginKindSource)
		sourceFound := lo.SomeBy(allPlugins, sourcePluginFilter)
		if !sourceFound {
			notFoundPlugins = errors.Join(notFoundPlugins, fmt.Errorf("source plugin %q not found", source))
		}
	}
	if destination != "" {
		destinationPluginFilter := pluginFilter(destination, cqapi.PluginKindDestination)
		destinationFound := lo.SomeBy(allPlugins, destinationPluginFilter)
		if !destinationFound {
			notFoundPlugins = errors.Join(notFoundPlugins, fmt.Errorf("destination plugin %q not found", destination))
		}
	}

	if notFoundPlugins != nil {
		return notFoundPlugins
	}

	officialDestinations := lo.Filter(allPlugins, officialReleasedPluginsByKind(cqapi.PluginKindDestination))

	if source == "" {
		officialSources := lo.Filter(allPlugins, officialReleasedPluginsByKind(cqapi.PluginKindSource))
		sort.SliceStable(officialSources, pluginsSorter(officialSources, sourcesOrder))
		prompt := promptui.Select{
			Label: "Select Source Plugin",
			Items: lo.Map(officialSources, pluginName),
		}

		_, source, err = prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed %w", err)
		}
		source, _ = normalizePluginPath(source)
	}
	_, sourceIndex, _ := lo.FindIndexOf(allPlugins, pluginFilter(source, cqapi.PluginKindSource))

	if destination == "" {
		sort.SliceStable(officialDestinations, pluginsSorter(officialDestinations, destinationsOrder))
		prompt := promptui.Select{
			Label: "Select Destination Plugin",
			Items: lo.Map(officialDestinations, pluginName),
		}

		_, destination, err = prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed %w", err)
		}
		destination, _ = normalizePluginPath(destination)
	}
	_, destinationIndex, _ := lo.FindIndexOf(allPlugins, pluginFilter(destination, cqapi.PluginKindDestination))

	fmt.Println("Selected Source Plugin:", source, sourceIndex)
	fmt.Println("Selected Destination Plugin:", destination, destinationIndex)

	return nil
}
