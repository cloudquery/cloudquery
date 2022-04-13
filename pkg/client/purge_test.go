package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-test/resources"

	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hashicorp/go-hclog"

	"github.com/stretchr/testify/assert"
)

func TestPurgeProviderData(t *testing.T) {

	testCases := []struct {
		Name    string
		Options *PurgeProviderDataOptions

		// Expected results and diags from first dry run
		ExpectedDryRunResult *PurgeProviderDataResult
		ExpectedDryRunDiags  []diag.FlatDiag

		// Expected results diags from normal run, if RunResults is nil, normal purge is not called.
		ExpectedRunResults *PurgeProviderDataResult
		ExpectedRunDiags   []diag.FlatDiag

		// Expected results and diags from secondary dry run
		SecondaryDryRunUpdate          time.Duration
		ExpectedSecondaryDryRunResults *PurgeProviderDataResult
		ExpectedSecondaryRunDiags      []diag.FlatDiag

		// Override plugin manager option
		PluginManagerCreator func() plugin.Manager
		Setup                func(t *testing.T, dsn string) func(t *testing.T)
	}{
		{
			Name: "no-providers-given",
			Options: &PurgeProviderDataOptions{
				Providers:  []registry.Provider{},
				LastUpdate: 0,
				DryRun:     true,
			},
			ExpectedDryRunDiags: []diag.FlatDiag{
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
				Providers:  []registry.Provider{{Name: "bad-plugin", Version: registry.LatestVersion, Source: registry.DefaultOrganization}},
				LastUpdate: 0,
				DryRun:     true,
			},
			ExpectedDryRunDiags: []diag.FlatDiag{
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
				Providers:  []registry.Provider{{Name: "test", Version: registry.LatestVersion, Source: registry.DefaultOrganization}},
				LastUpdate: 0,
				DryRun:     true,
			},
			ExpectedDryRunResult: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
		},
		{
			Name: "basic-data-purge",
			Options: &PurgeProviderDataOptions{
				Providers:  []registry.Provider{{Name: "test", Version: registry.LatestVersion, Source: registry.DefaultOrganization}},
				LastUpdate: 0,
			},
			Setup: func(t *testing.T, dsn string) func(t *testing.T) {
				tbl := resources.Provider().ResourceMap["slow_resource"]
				r := schema.NewResourceData(schema.PostgresDialect{}, tbl, nil, nil, nil, time.Now())
				_ = r.Set("cq_id", uuid.New())
				_ = r.Set("cq_meta", schema.Meta{
					LastUpdate: time.Now().Add(-time.Hour * 5),
					FetchId:    "",
				})
				insertData(t, dsn, tbl, schema.Resources{
					r,
				})
				return func(t *testing.T) {
					truncateTable(t, dsn, tbl.Name)
				}
			},
			ExpectedDryRunResult: &PurgeProviderDataResult{
				TotalAffected: 1,
				AffectedResources: map[string]int{
					"slow_resource": 1,
				},
			},
			ExpectedRunResults: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
			ExpectedSecondaryDryRunResults: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
		},
		{
			Name: "no-data-purge",
			Options: &PurgeProviderDataOptions{
				Providers:  []registry.Provider{{Name: "test", Version: registry.LatestVersion, Source: registry.DefaultOrganization}},
				LastUpdate: time.Hour * 10,
			},
			Setup: func(t *testing.T, dsn string) func(t *testing.T) {
				tbl := resources.Provider().ResourceMap["slow_resource"]
				r := schema.NewResourceData(schema.PostgresDialect{}, tbl, nil, nil, nil, time.Now())
				_ = r.Set("cq_id", uuid.New())
				_ = r.Set("cq_meta", schema.Meta{
					LastUpdate: time.Now().Add(-time.Hour * 5),
					FetchId:    "",
				})
				insertData(t, dsn, tbl, schema.Resources{
					r,
				})
				return func(t *testing.T) {
					truncateTable(t, dsn, tbl.Name)
				}
			},
			ExpectedDryRunResult: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
			ExpectedRunResults: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
			// We update time to verify data is still there and wasn't purged
			SecondaryDryRunUpdate: 1,
			ExpectedSecondaryDryRunResults: &PurgeProviderDataResult{
				TotalAffected: 1,
				AffectedResources: map[string]int{
					"slow_resource": 1,
				},
			},
		},
		{
			Name: "single-data-purge",
			Options: &PurgeProviderDataOptions{
				Providers:  []registry.Provider{{Name: "test", Version: registry.LatestVersion, Source: registry.DefaultOrganization}},
				LastUpdate: time.Hour * 6,
			},
			Setup: func(t *testing.T, dsn string) func(t *testing.T) {
				tbl := resources.Provider().ResourceMap["slow_resource"]
				r := schema.NewResourceData(schema.PostgresDialect{}, tbl, nil, nil, nil, time.Now())
				_ = r.Set("cq_id", uuid.New())
				_ = r.Set("cq_meta", schema.Meta{
					LastUpdate: time.Now().UTC().Add(-time.Hour * 5),
					FetchId:    "",
				})
				r2 := schema.NewResourceData(schema.PostgresDialect{}, tbl, nil, nil, nil, time.Now())
				_ = r2.Set("cq_id", uuid.New())
				_ = r2.Set("cq_meta", schema.Meta{
					LastUpdate: time.Now().UTC().Add(-time.Hour * 15),
					FetchId:    "",
				})

				insertData(t, dsn, tbl, schema.Resources{
					r,
					r2,
				})
				return func(t *testing.T) {
					truncateTable(t, dsn, tbl.Name)
				}
			},
			ExpectedDryRunResult: &PurgeProviderDataResult{
				TotalAffected: 1,
				AffectedResources: map[string]int{
					"slow_resource": 1,
				},
			},
			ExpectedRunResults: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
			// We update time to verify data is still there and wasn't purged
			SecondaryDryRunUpdate: time.Hour * 4,
			ExpectedSecondaryDryRunResults: &PurgeProviderDataResult{
				TotalAffected: 1,
				AffectedResources: map[string]int{
					"slow_resource": 1,
				},
			},
		},
		{
			Name: "data-purge-different-times",
			Options: &PurgeProviderDataOptions{
				Providers:  []registry.Provider{{Name: "test", Version: registry.LatestVersion, Source: registry.DefaultOrganization}},
				LastUpdate: time.Hour * 4,
			},
			Setup: func(t *testing.T, dsn string) func(t *testing.T) {
				tbl := resources.Provider().ResourceMap["slow_resource"]
				r := schema.NewResourceData(schema.PostgresDialect{}, tbl, nil, nil, nil, time.Now())
				_ = r.Set("cq_id", uuid.New())
				_ = r.Set("cq_meta", schema.Meta{
					LastUpdate: time.Now().UTC().Add(-time.Hour * 5),
					FetchId:    "",
				})
				r2 := schema.NewResourceData(schema.PostgresDialect{}, tbl, nil, nil, nil, time.Now())
				_ = r2.Set("cq_id", uuid.New())
				_ = r2.Set("cq_meta", schema.Meta{
					LastUpdate: time.Now().UTC().Add(-time.Hour * 15),
					FetchId:    "",
				})

				insertData(t, dsn, tbl, schema.Resources{
					r,
					r2,
				})
				return func(t *testing.T) {
					truncateTable(t, dsn, tbl.Name)
				}
			},
			ExpectedDryRunResult: &PurgeProviderDataResult{
				TotalAffected: 2,
				AffectedResources: map[string]int{
					"slow_resource": 2,
				},
			},
			ExpectedRunResults: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
			// We update time to verify data is still there and wasn't purged
			SecondaryDryRunUpdate: 0,
			ExpectedSecondaryDryRunResults: &PurgeProviderDataResult{
				TotalAffected:     0,
				AffectedResources: make(map[string]int),
			},
		},
	}

	dbDSN := setupDB(t)
	setupTestProvider(t, dbDSN)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			pm, err := plugin.NewManager(registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(".")), plugin.WithAllowReattach())
			if !assert.Nil(t, err) {
				t.FailNow()
			}
			if tc.Setup != nil {
				teardown := tc.Setup(t, dbDSN)
				defer teardown(t)
			}

			if len(tc.ExpectedDryRunDiags) > 0 || tc.ExpectedDryRunResult != nil {
				tc.Options.DryRun = true
				result, diags := PurgeProviderData(context.TODO(), NewStorage(dbDSN, nil), pm, tc.Options)
				checkPurgeOutput(t, tc.ExpectedDryRunResult, result, tc.ExpectedDryRunDiags, diag.FlattenDiags(diags, true))
			}

			if len(tc.ExpectedRunDiags) > 0 || tc.ExpectedRunResults != nil {
				tc.Options.DryRun = false
				result, diags := PurgeProviderData(context.TODO(), NewStorage(dbDSN, nil), pm, tc.Options)
				checkPurgeOutput(t, tc.ExpectedRunResults, result, tc.ExpectedRunDiags, diag.FlattenDiags(diags, true))
			}

			if len(tc.ExpectedSecondaryRunDiags) > 0 || tc.ExpectedSecondaryDryRunResults != nil {
				tc.Options.DryRun = true
				if tc.SecondaryDryRunUpdate > 0 {
					tc.Options.LastUpdate = tc.SecondaryDryRunUpdate
				}
				result, diags := PurgeProviderData(context.TODO(), NewStorage(dbDSN, nil), pm, tc.Options)
				checkPurgeOutput(t, tc.ExpectedSecondaryDryRunResults, result, tc.ExpectedSecondaryRunDiags, diag.FlattenDiags(diags, true))
			}

		})
	}
}

func checkPurgeOutput(t *testing.T, expectedResult, actualResult *PurgeProviderDataResult, expectedDiags, actualDiags []diag.FlatDiag) {
	if len(expectedDiags) > 0 {
		assert.Equal(t, expectedDiags, actualDiags)
	} else {
		assert.Len(t, actualDiags, 0)
	}
	if expectedResult != nil {
		assert.Equal(t, expectedResult, actualResult)
	}

}

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

	provider := registry.Provider{
		Name:    "test",
		Source:  "cloudquery",
		Version: "v0.0.11",
	}
	pm, err := plugin.NewManager(registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(".")), plugin.WithAllowReattach())
	assert.Nil(t, err)
	_, diags := Download(context.TODO(), pm, &DownloadOptions{
		Providers: []registry.Provider{provider},
		NoVerify:  false,
	})
	assert.Nil(t, diags)

	if _, err := Sync(context.TODO(), NewStorage(dsn, nil), pm, &SyncOptions{provider, true}); !assert.Nil(t, err) {
		t.FailNow()
	}
}
