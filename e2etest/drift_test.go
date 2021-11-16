package e2etest

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

const (
	// Set this environment variable to any non empty value to enable the tests
	envVarSwitch = "E2E_TESTS"
	// Location of a YAML file with a list of known resources/ids
	envVarKnownResourcesPath = "KNOWN_RESOURCES"
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

	for _, providerName := range providersToTestForDrift {
		t.Run(providerName, func(t *testing.T) {
			c := prepareClient(t, ctx, providerName, tmpdir, logger)
			s, err := getProviderSchemaResponse(ctx, c, providerName)
			require.Nil(t, err)

			tests := []struct {
				name string
				test func(*testing.T)
			}{
				{
					"successful run with all the tables",
					func(t *testing.T) {
						require.Nil(t, c.BuildProviderTables(ctx, providerName))
						mr, err := runDriftModule(t, ctx, c, "empty.tfstate", s)
						require.Nil(t, err)
						require.Nil(t, mr.Error)
						_, ok := mr.Result.(*drift.Results)
						require.True(t, ok, "result of module call is not *drift.Results")
					},
				},

				{
					"should fail with some tables missing",
					func(t *testing.T) {
						require.Nil(t, c.BuildProviderTables(ctx, providerName))
						// we'll keep deleting tables until we get an error
						// because some tables may not be used/ implemented in drift detection
						pool := setupDatabase(t)
						defer pool.Close()
						conn, err := pool.Acquire(ctx)
						require.Nil(t, err)
						defer conn.Release()
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
					},
				},

				{
					"should fail with newly added table",
					func(t *testing.T) {
						t.Skip("currently drift module does not report required error")

						require.Nil(t, c.BuildProviderTables(ctx, providerName))

						altTables := make(map[string]*schema.Table, len(s.ResourceTables)+1)
						pool := setupDatabase(t)
						defer pool.Close()
						conn, err := pool.Acquire(ctx)
						require.Nil(t, err)
						defer conn.Release()
						// add extra copy of a resource to the new map
						for k, tab := range s.ResourceTables {
							t2 := *tab
							t2.Name += "2"
							t2.Relations = nil
							require.Nil(t, c.TableCreator.CreateTable(ctx, conn, &t2, nil))
							altTables[k+"2"] = &t2
							defer func() {
								_, err := conn.Exec(ctx, fmt.Sprintf("DROP TABLE %s CASCADE", t2.Name))
								if err != nil {
									t.Logf("error deleting a table: %v", err)
								}
							}()
							break // one is enough
						}

						// run a drift module with altered schema
						altS := *s
						altS.ResourceTables = altTables
						mr, err := runDriftModule(t, ctx, c, "empty.tfstate", &altS)
						require.Nil(t, err)
						require.NotNil(t, mr.Error)
					},
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, tt.test)
			}
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
	knownResourcesPath := os.Getenv(envVarKnownResourcesPath)
	require.NotEmptyf(t, knownResourcesPath, "%s is not set", envVarKnownResourcesPath)
	tfStatePath := os.Getenv(envVarTFStatePath)
	require.NotEmptyf(t, tfStatePath, "%s is not set", envVarTFStatePath)

	knownResources := readExpectedResources(t, knownResourcesPath)
	logger := hclog.New(&hclog.LoggerOptions{})
	ctx := context.Background()
	tmpdir := t.TempDir()
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
				assertEqual(t, result, func(r *drift.Result) drift.ResourceList { return r.Equal }, knownResources)
			})
		})
	}
}

func readExpectedResources(t *testing.T, fromPath string) map[string][]string {
	var r map[string][]string
	f, err := os.Open(fromPath)
	require.Nil(t, err)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	require.Nil(t, err)
	require.Nil(t, yaml.Unmarshal(b, &r))
	return r
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
	return s, nil
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

func assertEqual(t *testing.T, d *drift.Results, extractor func(*drift.Result) drift.ResourceList, known map[string][]string) {
	got := make(map[string][]string)
	for _, r := range d.Data {
		got[r.ResourceType] = append(got[r.ResourceType], extractor(r).IDs()...)
	}

	for resourceType, items := range known {
		t.Run(resourceType, func(t *testing.T) {
			got := stringSetFromList(got[resourceType])
			want := stringSetFromList(items)
			assert.Truef(t, got.Eq(want), "sets are different. first: %v, second: %v", got.ToList(), want.ToList())
		})
	}
}

func runDriftModule(t *testing.T, ctx context.Context, c *client.Client, stateFilePath string, s *cqproto.GetProviderSchemaResponse) (*module.ExecutionResult, error) {
	driftParams := drift.RunParams{
		ListManaged: true,
		TfMode:      "managed",
		StateFiles:  []string{stateFilePath},
	}
	return c.ExecuteModule(ctx, client.ModuleRunRequest{
		Name:      "drift",
		Params:    driftParams,
		Providers: []*cqproto.GetProviderSchemaResponse{s},
	})
}
