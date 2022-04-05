package client

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/cloudquery/cloudquery/internal/test/providertest"

	"github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hashicorp/go-hclog"

	"github.com/stretchr/testify/assert"
)

var (
	defaultProviderPath = filepath.Join(".", ".cq", "providers")
)

func TestPurgeProviderData(t *testing.T) {

	testCases := []struct {
		Name           string
		Options        *PurgeProviderDataOptions
		ExpectedResult *PurgeProviderDataResult
		ExpectedDiags  []diag.FlatDiag
		// Override plugin manager option
		PluginManagerCreator func() plugin.Manager
		Setup                func(t *testing.T, dsn string) func(t *testing.T)
	}{
		{
			Name: "no-providers-given",
			Options: &PurgeProviderDataOptions{
				Providers:  []string{},
				LastUpdate: 0,
				DryRun:     true,
			},
			ExpectedDiags: []diag.FlatDiag{
				{
					Err:      "no providers were given",
					Type:     diag.INTERNAL,
					Severity: diag.WARNING,
					Summary:  "no providers were given",
				},
			},
		},
		{
			Name: "bad-plugin-name",
			Options: &PurgeProviderDataOptions{
				Providers:  []string{"bad-plugin"},
				LastUpdate: 0,
				DryRun:     true,
			},
			ExpectedDiags: []diag.FlatDiag{
				{
					Err:      "no such provider bad-plugin. plugin might be missing from directory or wasn't downloaded",
					Type:     diag.INTERNAL,
					Severity: diag.ERROR,
					Summary:  "no such provider bad-plugin. plugin might be missing from directory or wasn't downloaded",
				},
			},
		},
		{
			Name: "dry-run-no-data",
			Options: &PurgeProviderDataOptions{
				Providers:  []string{"test"},
				LastUpdate: 0,
				DryRun:     true,
			},
			ExpectedResult: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
		},

		{
			Name: "dry-run-basic-data",
			Options: &PurgeProviderDataOptions{
				Providers:  []string{"test"},
				LastUpdate: 0,
				DryRun:     true,
			},
			Setup: func(t *testing.T, dsn string) func(t *testing.T) {
				tbl := providertest.Provider().ResourceMap["slow_resource"]
				r := schema.NewResourceData(schema.PostgresDialect{}, tbl, nil, nil, nil, time.Now())
				r.Set("cq_id", uuid.New())
				r.Set("cq_meta", schema.Meta{
					LastUpdate: time.Now().Add(time.Hour - 5),
					FetchId:    "",
				})

				insertData(t, dsn, tbl, schema.Resources{
					r,
				})
				return func(t *testing.T) {
					truncateTable(t, dsn, tbl.Name)
				}
			},
			ExpectedResult: &PurgeProviderDataResult{
				TotalAffected: 1,
				AffectedResources: map[string]int{
					"slow_resource": 1,
				},
			},
		},
	}

	dbDSN := setupDB(t)
	setupTestProvider(t, dbDSN)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			pm, err := plugin.NewManager(hclog.Default(), defaultProviderPath, registry.CloudQueryRegistryURL, nil)
			if !assert.Nil(t, err) {
				t.FailNow()
			}
			if tc.Setup != nil {
				teardown := tc.Setup(t, dbDSN)
				defer teardown(t)
			}

			result, diags := PurgeProviderData(context.TODO(), NewStorage(dbDSN), pm, tc.Options)
			if len(tc.ExpectedDiags) > 0 {
				assert.NotNil(t, diags)
				assert.Equal(t, tc.ExpectedDiags, diag.FlattenDiags(diags, true))

			} else {
				assert.Nil(t, diags)
			}
			if tc.ExpectedResult != nil {
				assert.Equal(t, tc.ExpectedResult, result)
			}

		})
	}
}

// INSERT INTO public.slow_resource(
//	cq_id, cq_meta, some_bool, upgrade_column, upgrade_column_2)
//	VALUES ('f6385024-e8d8-5961-823b-9feea81b34d9', '{"last_update": "2022-04-01T09:04:34.7637611Z"}', false, 1, 2);

func insertData(t *testing.T, dsn string, tbl *schema.Table, resources schema.Resources) {
	db, err := database.New(context.TODO(), hclog.Default(), dsn)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	defer db.Close()
	assert.Nil(t, db.Insert(context.TODO(), tbl, resources))
}

func truncateTable(t *testing.T, dsn, table string) {
	db, err := database.New(context.TODO(), hclog.Default(), dsn)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	defer db.Close()
	assert.Nil(t, db.Exec(context.TODO(), fmt.Sprintf("TRUNCATE %s", table)))
}

func setupTestProvider(t *testing.T, dsn string) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()
	// TODO: we set up client for now as not all commands are refactored
	c, err := New(context.TODO(), func(options *Client) {
		options.DSN = dsn
		options.Providers = requiredTestProviders()
	})
	assert.Nil(t, err)
	t.Cleanup(c.Close)

	if err := c.BuildProviderTables(context.TODO(), "test"); !assert.Nil(t, err) {
		t.FailNow()
	}
}
