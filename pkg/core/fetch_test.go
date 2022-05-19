package core

import (
	"context"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Fetch(t *testing.T) {
	latestVersion := getLatestVersion(t, "test")
	testCases := []struct {
		Name             string
		Options          FetchOptions
		ExpectedDiags    []diag.FlatDiag
		ExpectedResponse *FetchResponse
		Timeout          time.Duration
	}{
		{
			Name: "fetch-errors",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Alias:         "test_alias",
							Resources:     []string{"slow_resource", "panic_resource", "error_resource", "very_slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			ExpectedDiags: []diag.FlatDiag{
				{
					Err:         "error from provider",
					Resource:    "error_resource",
					Type:        diag.RESOLVING,
					Severity:    diag.ERROR,
					Summary:     "error from provider",
					Description: diag.Description{Resource: "error_resource", ResourceID: []string(nil), Summary: "error from provider", Detail: ""}},
				{
					Err:         "failed table panic_resource fetch. Error: resource with panic",
					Resource:    "panic_resource",
					Type:        diag.RESOLVING,
					Severity:    diag.ERROR,
					Summary:     "failed table panic_resource fetch. Error: resource with panic",
					Description: diag.Description{Resource: "panic_resource", ResourceID: []string(nil), Summary: "failed table panic_resource fetch. Error: resource with panic", Detail: ""},
				},
			},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test(test_alias)": {
				Name:                  "test",
				Alias:                 "test_alias",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFinished,
			}}},
		},
		{
			Name: "fetch-simple",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Resources:     []string{"slow_resource", "very_slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			ExpectedDiags: nil,
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFinished,
			}}},
		},
		{
			Name: "fetch-unique-fetch-id",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Resources:     []string{"slow_resource", "very_slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
				FetchId: uuid.New(),
			},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFinished,
			}}},
		},
		{
			Name: "fetch-timeout",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Resources:     []string{"slow_resource", "very_slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			Timeout: time.Second * 4,
			ExpectedDiags: []diag.FlatDiag{
				{
					Err:         "context deadline exceeded",
					Type:        diag.USER,
					Severity:    diag.ERROR,
					Summary:     "provider fetch was canceled by user / fetch deadline exceeded",
					Description: diag.Description{Resource: "", ResourceID: []string(nil), Summary: "provider fetch was canceled by user / fetch deadline exceeded", Detail: ""}},
			},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchCanceled,
			}}},
		},
		{
			Name: "fetch-default-config",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:      "test",
							Resources: []string{"slow_resource"},
						},
					},
				},
			},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFinished,
			}}},
		},
		{
			Name: "fetch-duplicates",
			Options: FetchOptions{
				ProvidersInfo: []ProviderInfo{
					{
						Provider: registry.Provider{
							Name:    "test",
							Version: registry.LatestVersion,
							Source:  registry.DefaultOrganization,
						},
						Config: &config.Provider{
							Name:          "test",
							Resources:     []string{"slow_resource", "slow_resource", "slow_resource"},
							Env:           nil,
							Configuration: nil,
						},
					},
				},
			},
			ExpectedDiags: []diag.FlatDiag{{Err: "resource \"slow_resource\" is duplicate", Type: 7, Severity: 2, Summary: "resource \"slow_resource\" is duplicate", Description: diag.Description{Summary: "resource \"slow_resource\" is duplicate", Detail: "configuration has duplicate resources"}}},
			ExpectedResponse: &FetchResponse{ProviderFetchSummary: map[string]*ProviderFetchSummary{"test": {
				Name:                  "test",
				Alias:                 "",
				Version:               registry.LatestVersion,
				TotalResourcesFetched: 0,
				Status:                FetchFailed,
			}}},
		},
	}

	pManager, err := plugin.NewManager(registry.NewRegistryHub(firebase.CloudQueryRegistryURL))
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			dsn := setupDB(t)
			storage := database.NewStorage(dsn, nil)
			rp := make([]registry.Provider, len(tc.Options.ProvidersInfo))
			for i, p := range tc.Options.ProvidersInfo {
				rp[i] = p.Provider
			}
			// download test provider if it doesn't already exist
			_, diags := Download(context.Background(), pManager, &DownloadOptions{
				Providers: rp,
				NoVerify:  false,
			})
			require.False(t, diags.HasDiags())

			for _, r := range rp {
				// Sync provider in table before fetch
				_, diags := Sync(context.Background(), storage, pManager, &SyncOptions{
					Provider:       r,
					DownloadLatest: false,
				})
				require.False(t, diags.HasDiags())
			}

			var (
				ctx    = context.Background()
				cancel context.CancelFunc
			)
			if tc.Timeout > 0 {
				ctx, cancel = context.WithTimeout(context.Background(), tc.Timeout)
				defer cancel()
			}
			resp, diags := Fetch(ctx, storage, pManager, &tc.Options)
			if tc.ExpectedDiags != nil {
				flattenedDiags := diag.FlattenDiags(diags, false)
				require.Len(t, flattenedDiags, len(tc.ExpectedDiags))
				for i, expected := range tc.ExpectedDiags {
					actual := flattenedDiags[i]
					assert.ElementsMatch(t, expected.Description.ResourceID, actual.Description.ResourceID)
					assert.Contains(t, actual.Description.Detail, expected.Description.Detail)
					assert.Contains(t, actual.Description.Resource, expected.Description.Resource)
					assert.Contains(t, actual.Description.Summary, expected.Description.Summary)
					assert.Contains(t, actual.Err, expected.Err)
					assert.Contains(t, actual.Summary, expected.Summary)

					assert.Equal(t, expected.Type, actual.Type)
					assert.Equal(t, expected.Severity, actual.Severity)
				}
			} else {
				assert.Equal(t, []diag.FlatDiag{}, diag.FlattenDiags(diags, false))
			}
			if tc.ExpectedResponse == nil {
				require.Nil(t, resp)
			} else {
				if tc.Options.FetchId != uuid.Nil {
					assert.Equal(t, tc.Options.FetchId, resp.FetchId)
				}
				for k, p := range tc.ExpectedResponse.ProviderFetchSummary {
					fetchSummary, ok := resp.ProviderFetchSummary[k]
					require.True(t, ok)
					assert.Equal(t, p.Name, fetchSummary.Name)
					if p.Version == registry.LatestVersion {
						assert.Equal(t, latestVersion, fetchSummary.Version)
					} else {
						assert.Equal(t, p.Version, fetchSummary.Version)
					}
					assert.Equal(t, p.Status, fetchSummary.Status)
					assert.Equal(t, p.TotalResourcesFetched, fetchSummary.TotalResourcesFetched)
				}
			}
		})
	}
}

func Test_doNormalizeResources(t *testing.T) {
	tests := []struct {
		name      string
		requested []string
		skip      []string
		all       map[string]*schema.Table
		want      []string
		wantErr   bool
	}{
		{
			"wilcard",
			[]string{"*"},
			nil,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"1", "2", "3"},
			false,
		},
		{
			"wilcard with explicit",
			[]string{"*", "1"},
			nil,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"unknown resource",
			[]string{"1", "2", "x"},
			nil,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"duplicate resource",
			[]string{"1", "2", "1"},
			nil,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"ok, all explicit",
			[]string{"2", "1"},
			nil,
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"1", "2"},
			false,
		},
		{
			"ok, all explicit with ignores",
			[]string{"2", "1", "3"},
			[]string{"1"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"2", "3"},
			false,
		},
		{
			"ok, some globs",
			[]string{"c1.*", "c2.res4"},
			nil,
			map[string]*schema.Table{"c1.res1": nil, "c1.res2": nil, "c2.res3": nil, "c2.res4": nil, "c1a.res5": nil},
			[]string{"c1.res1", "c1.res2", "c2.res4"},
			false,
		},
		{
			"ok, some globs with skips",
			[]string{"c1.*", "c2.res4"},
			[]string{"c1.res1"},
			map[string]*schema.Table{"c1.res1": nil, "c1.res2": nil, "c2.res3": nil, "c2.res4": nil, "c1a.res5": nil},
			[]string{"c1.res2", "c2.res4"},
			false,
		},
		{
			"invalid glob 1",
			[]string{"c1*res1"},
			nil,
			map[string]*schema.Table{"c1.res1": nil, "c1.res2": nil, "c2.res3": nil, "c2.res4": nil},
			nil,
			true,
		},
		{
			"invalid glob 2",
			[]string{"c1.*res1"},
			nil,
			map[string]*schema.Table{"c1.res1": nil, "c1.res2": nil, "c2.res3": nil, "c2.res4": nil},
			nil,
			true,
		},
		{
			"invalid glob 3",
			[]string{"c1.res*"},
			nil,
			map[string]*schema.Table{"c1.res1": nil, "c1.res2": nil, "c2.res3": nil, "c2.res4": nil},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doNormalizeResources(tt.requested, tt.skip, tt.all)
			if (err != nil) != tt.wantErr {
				t.Errorf("doInterpolate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				tt.want = []string{}
			}
			if got == nil {
				got = []string{}
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
