package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	apiAuth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/internal/api"
	"github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

const (
	testConnectionShort   = "Test plugin connections to sources and/or destinations"
	testConnectionExample = `# Test plugin connections to sources and/or destinations
cloudquery test-connection ./directory
# Test plugin connections from directories and files
cloudquery test-connection ./directory ./aws.yml ./pg.yml
`
)

func getSyncTestConnectionAPIClient() (*cloudquery_api.ClientWithResponses, error) {
	authClient := apiAuth.NewTokenClient()
	if authClient.GetTokenType() != apiAuth.SyncTestConnectionAPIKey {
		return nil, nil
	}

	token, err := authClient.GetToken()
	if err != nil {
		return nil, err
	}
	return api.NewClient(token.Value)
}

func updateSyncTestConnectionStatus(ctx context.Context, logger zerolog.Logger, status cloudquery_api.SyncTestConnectionStatus) {
	apiClient, err := getSyncTestConnectionAPIClient()
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get sync test connection API client")
		return
	}
	if apiClient == nil {
		return
	}
	teamName, syncTestConnectionId := os.Getenv("_CQ_TEAM_NAME"), os.Getenv("_CQ_SYNC_TEST_CONNECTION_ID")
	if teamName == "" || syncTestConnectionId == "" {
		log.Warn().Msg("Skipping sync test connection status update as environment variables are not set")
		return
	}
	syncTestConnectionUUID, err := uuid.Parse(syncTestConnectionId)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to parse sync test connection UUID")
		return
	}
	log.Info().Str("status", string(status)).Msg("Sending sync test connection to API")
	res, err := apiClient.UpdateSyncTestConnectionWithResponse(ctx, teamName, syncTestConnectionUUID, cloudquery_api.UpdateSyncTestConnectionJSONRequestBody{
		Status: status,
	})
	if err != nil {
		log.Warn().Err(err).Msg("Failed to send sync test connection to API")
		return
	}
	if res.StatusCode() != http.StatusOK {
		log.Warn().Str("status", res.Status()).Int("code", res.StatusCode()).Msg("Failed to send test connection to API")
	} else {
		log.Info().Str("status", string(status)).Msg("Sent sync test connection to API")
	}
}

func newCmdTestConnection() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "test-connection [files or directories]",
		Short:   testConnectionShort,
		Long:    testConnectionShort,
		Example: testConnectionExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    testConnection,
	}
	return cmd
}

func testConnection(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	updateSyncTestConnectionStatus(cmd.Context(), log.Logger, cloudquery_api.SyncTestConnectionStatusStarted)

	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewTestConnectionSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	sources := specReader.Sources
	destinations := specReader.Destinations

	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, destinations)
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	teamName, err := auth.GetTeamForToken(ctx, authToken)
	if err != nil {
		return fmt.Errorf("failed to get team name: %w", err)
	}
	opts := []managedplugin.Option{
		managedplugin.WithLogger(log.Logger),
		managedplugin.WithAuthToken(authToken.Value),
		managedplugin.WithTeamName(teamName),
	}
	if cqDir != "" {
		opts = append(opts, managedplugin.WithDirectory(cqDir))
	}
	if disableSentry {
		opts = append(opts, managedplugin.WithNoSentry())
	}

	sourcePluginConfigs := make([]managedplugin.Config, len(sources))
	sourceRegInferred := make([]bool, len(sources))
	for i, source := range sources {
		sourcePluginConfigs[i] = managedplugin.Config{
			Name:       source.Name,
			Version:    source.Version,
			Path:       source.Path,
			Registry:   SpecRegistryToPlugin(source.Registry),
			DockerAuth: source.DockerRegistryAuthToken,
		}
		sourceRegInferred[i] = source.RegistryInferred()
	}
	destinationPluginConfigs := make([]managedplugin.Config, len(destinations))
	destinationRegInferred := make([]bool, len(destinations))
	for i, destination := range destinations {
		destinationPluginConfigs[i] = managedplugin.Config{
			Name:       destination.Name,
			Version:    destination.Version,
			Path:       destination.Path,
			Registry:   SpecRegistryToPlugin(destination.Registry),
			DockerAuth: destination.DockerRegistryAuthToken,
		}
		destinationRegInferred[i] = destination.RegistryInferred()
	}

	sourceClients, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...)
	if err != nil {
		return enrichClientError(sourceClients, sourceRegInferred, err)
	}
	defer func() {
		if err := sourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	destinationClients, err := managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...)
	if err != nil {
		return enrichClientError(destinationClients, destinationRegInferred, err)
	}
	defer func() {
		if err := destinationClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()

	var initErrors []error
	var testConnectionResults []testConnectionResult
	for i, client := range sourceClients {
		pluginClient := plugin.NewPluginClient(client.Conn)
		log.Info().Str("source", sources[i].VersionString()).Msg("Initializing source")
		err := initPlugin(ctx, pluginClient, sources[i].Spec, false, invocationUUID.String())
		if err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to init source %v: %w", sources[i].VersionString(), err))
		} else {
			log.Info().Str("source", sources[i].VersionString()).Msg("Initialized source")
		}
		testResult, err := testPluginConnection(ctx, pluginClient, sources[i].Spec)
		if err != nil {
			return fmt.Errorf("failed to test source %v: %w", sources[i].VersionString(), err)
		}
		testConnectionResults = append(testConnectionResults, *testResult)
	}
	for i, client := range destinationClients {
		pluginClient := plugin.NewPluginClient(client.Conn)
		log.Info().Str("destination", destinations[i].VersionString()).Msg("Initializing destination")
		err := initPlugin(ctx, pluginClient, destinations[i].Spec, false, invocationUUID.String())
		if err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to init destination %v: %w", destinations[i].VersionString(), err))
		} else {
			log.Info().Str("destination", destinations[i].VersionString()).Msg("Initialized destination")
		}
		testResult, err := testPluginConnection(ctx, pluginClient, destinations[i].Spec)
		if err != nil {
			return fmt.Errorf("failed to test destination %v: %w", destinations[i].VersionString(), err)
		}
		testConnectionResults = append(testConnectionResults, *testResult)
	}

	allErrors := errors.Join(initErrors...)
	status := cloudquery_api.SyncTestConnectionStatusCompleted
	if allErrors != nil {
		status = cloudquery_api.SyncTestConnectionStatusFailed
	}
	updateSyncTestConnectionStatus(ctx, log.Logger, status)

	log.Info().Any("testresults", testConnectionResults).Msg("Test connection completed")

	return allErrors
}

type testConnectionResult struct {
	Success            bool
	FailureCode        string
	FailureDescription string
}

func testPluginConnection(ctx context.Context, pclient plugin.PluginClient, spec map[string]any) (*testConnectionResult, error) {
	var (
		specBytes []byte
		err       error
	)

	if len(spec) == 0 { // All nil or empty values to be marshaled as null
		specBytes = []byte(`null`)
	} else {
		specBytes, err = json.Marshal(spec)
		if err != nil {
			return nil, err
		}
	}

	in := &plugin.TestConnection_Request{
		Spec: specBytes,
	}

	resp, err := pclient.TestConnection(ctx, in)
	if err != nil {
		if gRPCErr, ok := grpcstatus.FromError(err); ok {
			if gRPCErr.Code() == codes.Unimplemented {
				return &testConnectionResult{}, nil
			}
		}
		return nil, err
	}

	return &testConnectionResult{
		Success:            resp.Success,
		FailureCode:        resp.FailureCode,
		FailureDescription: resp.FailureDescription,
	}, nil
}
