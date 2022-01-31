package e2etest

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// Set this environment variable to any non empty value to enable the tests
	envVarSwitch = "E2E_TESTS"
	// Path to a Terraform state file
	envVarTFStatePath = "TFSTATE_PATH"

	testDSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
)

var providersToTestForDrift = []string{"aws"}

func TestDriftWithoutTheCloud(t *testing.T) {
	skipIfEnvVarNotSet(t, envVarSwitch)

	logger := hclog.New(&hclog.LoggerOptions{})
	ctx := context.Background()
	tmpdir := t.TempDir()
	tmpdir = "/tmp/cqtest"
	for _, providerName := range providersToTestForDrift {
		t.Run(providerName, func(t *testing.T) {
			c := prepareClient(t, ctx, providerName, tmpdir, logger)
			s, err := getProviderSchemaResponse(ctx, c, providerName)
			require.Nil(t, err)

			t.Run("successful run with all the tables", func(t *testing.T) {
				require.Nil(t, c.BuildProviderTables(ctx, providerName))
				mr, err := runDriftModule(t, ctx, c, "empty.tfstate", s)
				require.Nil(t, err)
				require.Nil(t, mr.Error)
				_, ok := mr.Result.(*drift.Results)
				require.True(t, ok, "result of module call is not *drift.Results")
			})

			t.Run("should fail with some tables missing", func(t *testing.T) {
				require.Nil(t, c.BuildProviderTables(ctx, providerName))
				// we'll keep deleting tables until we get an error
				// because some tables may not be used/ implemented in drift detection
				pool := setupDatabase(t)
				defer pool.Close()
				conn, err := pool.Acquire(ctx)
				require.Nil(t, err)
				defer conn.Release()
				_, err = conn.Exec(ctx, fmt.Sprintf("DROP TABLE cloudquery_%s_schema_migrations", providerName))
				require.Nil(t, err)
				for _, tab := range s.ResourceTables {
					_, err := conn.Exec(ctx, fmt.Sprintf("DROP TABLE %s CASCADE", tab.Name))
					require.Nil(t, err)
					mr, err := runDriftModule(t, ctx, c, "empty.tfstate", s)
					require.Nil(t, err)
					if mr.Error != nil {
						return
					}
				}

				t.Errorf("deleted all tables, but still no error from drift")
			})
		})
	}
}

func prepareClient(t *testing.T, ctx context.Context, providerName string, baseDir string, logger hclog.Logger) *client.Client {
	c, err := client.New(ctx, func(c *client.Client) {
		c.Providers = []*config.RequiredProvider{{Name: providerName, Version: "latest"}}
		c.PluginDirectory = path.Join(baseDir, "plugins")
		c.PolicyDirectory = path.Join(baseDir, "policies")
		c.DSN = testDSN
		c.Logger = logger
	})
	require.Nil(t, err)
	require.Nil(t, c.DownloadProviders(ctx))
	return c
}

func skipIfEnvVarNotSet(t *testing.T, name string) {
	if os.Getenv(name) == "" {
		t.Skipf("%s is not set or empty", name)
	}
}

func setupDatabase(t *testing.T) *pgxpool.Pool {
	poolCfg, err := pgxpool.ParseConfig(testDSN)
	require.NoError(t, err)
	poolCfg.LazyConnect = true
	pool, err := pgxpool.ConnectConfig(context.Background(), poolCfg)
	require.NoError(t, err)
	return pool
}

func TestDriftWithTheCloud(t *testing.T) {
	skipIfEnvVarNotSet(t, envVarSwitch)

	// ensure that required envnironment variables are set
	tfStatePath := os.Getenv(envVarTFStatePath)
	require.NotEmptyf(t, tfStatePath, "%s is not set", envVarTFStatePath)

	tfStatePath = possiblySaveFromS3(t, tfStatePath)
	knownResources := loadKnownResources(t, tfStatePath)
	logger := hclog.New(&hclog.LoggerOptions{})
	ctx := context.Background()
	tmpdir := t.TempDir()
	tmpdir = "/tmp/cqtest"
	for _, providerName := range providersToTestForDrift {
		t.Run(providerName, func(t *testing.T) {
			c := prepareClient(t, ctx, providerName, tmpdir, logger)
			s, err := getProviderSchemaResponse(ctx, c, providerName)
			require.Nil(t, err)
			require.Nil(t, c.BuildProviderTables(ctx, providerName))

			t.Run("before a fetch all known resources are inside Missing", func(t *testing.T) {
				mr, err := runDriftModule(t, ctx, c, tfStatePath, s)
				require.Nil(t, err)
				require.Nil(t, mr.Error)
				result, ok := mr.Result.(*drift.Results)
				require.True(t, ok, "result of module call is not *drift.Results")
				assertContained(t, result, func(r *drift.Result) drift.ResourceList { return r.Missing }, knownResources)
			})

			_, err = c.Fetch(ctx, client.FetchRequest{
				Providers: []*config.Provider{{Name: providerName, Resources: []string{"*"}}},
			})
			require.Nil(t, err)

			t.Run("after a fetch with empty tf state all known resources are inside Extra", func(t *testing.T) {
				mr, err := runDriftModule(t, ctx, c, "empty.tfstate", s)
				require.Nil(t, err)
				require.Nil(t, mr.Error)
				result, ok := mr.Result.(*drift.Results)
				require.True(t, ok, "result of module call is not *drift.Results")
				assertContained(t, result, func(r *drift.Result) drift.ResourceList { return r.Extra }, knownResources)
			})

			t.Run("after a fetch with full tf state all known resources are inside Equal", func(t *testing.T) {
				mr, err := runDriftModule(t, ctx, c, tfStatePath, s)
				require.Nil(t, err)
				require.Nil(t, mr.Error)
				result, ok := mr.Result.(*drift.Results)
				require.True(t, ok, "result of module call is not *drift.Results")
				assertContained(t, result, func(r *drift.Result) drift.ResourceList { return r.Equal }, knownResources)
			})
		})
	}
}

func getProviderSchemaResponse(ctx context.Context, c *client.Client, providerName string) (*cqproto.GetProviderSchemaResponse, error) {
	s, err := c.GetProviderSchema(ctx, providerName)
	if err != nil {
		return nil, err
	}
	if s.Version == "" {
		d, err := c.Manager.GetPluginDetails(providerName)
		if err != nil {
			return nil, err
		}
		s.Version = d.Version
	}
	return s.GetProviderSchemaResponse, nil
}

func assertContained(t *testing.T, d *drift.Results, extractor func(*drift.Result) drift.ResourceList, known map[string][]string) {
	got := make(map[string][]string)
	for _, r := range d.Data {
		got[r.ResourceType] = append(got[r.ResourceType], extractor(r).IDs()...)
	}

	for resourceType, items := range known {
		t.Run(resourceType, func(t *testing.T) {
			diff := stringSetFromList(got[resourceType]).Sub(stringSetFromList(items))
			assert.Zerof(t, diff.Len(), "ids missing: %v", diff.ToList())
		})
	}
}

func runDriftModule(t *testing.T, ctx context.Context, c *client.Client, stateFilePath string, s *cqproto.GetProviderSchemaResponse) (*module.ExecutionResult, error) {
	driftParams := drift.RunParams{
		ListManaged:    true,
		TfMode:         "managed",
		StateFiles:     []string{stateFilePath},
		DisableFilters: true,
	}
	return c.ExecuteModule(ctx, client.ModuleRunRequest{
		Name:      "drift",
		Params:    driftParams,
		Providers: []*cqproto.GetProviderSchemaResponse{s},
	})
}

func parseS3URI(uri string) (string, string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", "", err
	}
	return u.Host, strings.TrimLeft(u.Path, "/"), nil
}
