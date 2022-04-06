package policy

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/pkg/client/database"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	"github.com/cloudquery/cloudquery/pkg/client/meta_storage"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func setupPolicyDatabase(t *testing.T, tableName string) (LowLevelQueryExecer, func(t *testing.T)) {
	conn, err := sdkdb.New(context.Background(), hclog.NewNullLogger(), "postgres://postgres:pass@localhost:5432/postgres")
	assert.NoError(t, err)

	// Setup test data
	err = conn.Exec(context.Background(), fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
	assert.NoError(t, err)
	err = conn.Exec(context.Background(), fmt.Sprintf("CREATE TABLE %s (id serial PRIMARY KEY, name VARCHAR(50) NOT NULL)", tableName))
	assert.NoError(t, err)
	err = conn.Exec(context.Background(), fmt.Sprintf("INSERT INTO %s VALUES (1, 'john')", tableName))
	assert.NoError(t, err)

	// Return conn and tear down func
	return conn, func(t *testing.T) {
		err = conn.Exec(context.Background(), fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName))
		assert.NoError(t, err)
	}
}

func TestExecutor_executeQuery(t *testing.T) {
	cases := []struct {
		Name          string
		Query         string
		ExpectOutput  bool
		ShouldBeEmpty bool
	}{
		{
			Name:          "no output",
			Query:         fmt.Sprintf("SELECT * FROM %s WHERE name LIKE 'peter'", t.Name()),
			ExpectOutput:  false,
			ShouldBeEmpty: true,
		},
		{
			Name:          "output",
			Query:         fmt.Sprintf("SELECT * FROM %s WHERE name LIKE 'john'", t.Name()),
			ExpectOutput:  true,
			ShouldBeEmpty: false,
		},
	}

	conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)
	executor := NewExecutor(conn, hclog.Default(), nil)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			res, err := executor.executeQuery(context.Background(), &Check{
				Query:        tc.Query,
				ExpectOutput: tc.ExpectOutput,
			})
			assert.NoError(t, err)
			if tc.ShouldBeEmpty {
				assert.Empty(t, res.Data)
			} else {
				assert.NotEmpty(t, res.Data)
			}
		})
	}
}

func TestExecutor_executePolicy(t *testing.T) {
	cases := []struct {
		Name          string
		Queries       []*Check
		Views         []*View
		ShouldBeEmpty bool
		Pass          bool
		ErrorOutput   string
	}{
		{
			Name: "multiple_queries",
			Queries: []*Check{
				{
					Name:         "query-1",
					ExpectOutput: false,
					Query:        fmt.Sprintf("SELECT * from %s WHERE name LIKE 'peter'", t.Name()),
				},
				{
					Name:         "query-2",
					ExpectOutput: true,
					Query:        fmt.Sprintf("SELECT * from %s WHERE name LIKE 'john'", t.Name()),
				},
			},
			ShouldBeEmpty: false,
			Pass:          true,
		},
		{
			Name: "query_with_dependent_view",
			Views: []*View{
				{
					Name:  "testview",
					Query: fmt.Sprintf("SELECT * FROM %s WHERE name LIKE 'john'", t.Name()),
				},
			},
			Queries: []*Check{
				{
					Name:         "query-with-view",
					ExpectOutput: true,
					Query:        "SELECT * from testview",
				},
			},
			Pass:          true,
			ShouldBeEmpty: false,
		},
		{
			Name: "broken_policy_query",
			Queries: []*Check{
				{
					Name:  "broken-query",
					Query: "SECT * OM testview",
				},
			},
			ErrorOutput:   "broken_policy_query/broken-query: ERROR: syntax error at or near \"SECT\" (SQLSTATE 42601)",
			ShouldBeEmpty: true,
			Pass:          true,
		},
		{
			Name: "broken_policy_view",
			Views: []*View{
				{
					Name:  "brokenview",
					Query: "INVALID * MOFR *",
				},
			},
			Queries: []*Check{
				{
					Name:  "broken-query",
					Query: "SECT * OM testview",
				},
			},
			ErrorOutput:   "failed to create view broken_policy_view/brokenview: ERROR: syntax error at or near \"INVALID\" (SQLSTATE 42601)",
			ShouldBeEmpty: true,
			Pass:          true,
		},
	}

	conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)
	executor := NewExecutor(conn, hclog.Default(), nil)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			p := &Policy{
				Name:   tc.Name,
				Checks: tc.Queries,
				Views:  tc.Views,
			}
			execReq := &ExecuteRequest{
				Policy: &Policy{
					Name: tc.Name,
				},
				UpdateCallback: nil,
				StopOnFailure:  false,
			}

			res, err := executor.Execute(context.Background(), execReq, p)
			if tc.ErrorOutput != "" {
				assert.EqualError(t, err, tc.ErrorOutput)
			} else {
				assert.NoError(t, err)
			}
			if tc.ShouldBeEmpty {
				assert.Empty(t, res)
			} else {
				assert.NotEmpty(t, res)
			}
			if res != nil {
				assert.Equal(t, tc.Pass, res.Passed)
			}
		})
	}
}

var (
	multiLayerPolicy = &Policy{
		Name: "test",
		Policies: Policies{
			{
				Name: "subpolicy",
				Checks: []*Check{{
					Name:         "sub-query",
					Query:        "SELECT 1 as result;",
					ExpectOutput: true,
				},
					{
						Name:         "other-query",
						Query:        "SELECT 1 as result;",
						ExpectOutput: true,
					},
				},
			},
		},
		Checks: []*Check{{
			Query:        "SELECT 1 as result;",
			ExpectOutput: true,
		}},
	}
	failingPolicy = &Policy{
		Name: "test",
		Policies: Policies{
			{
				Name: "subpolicy",
				Checks: []*Check{{
					Name:         "sub-query",
					Query:        "SELECT 1 as result;",
					ExpectOutput: true,
				},
					{
						Name:  "other-query",
						Query: "SELECT 1 as result;",
					},
				},
			},
		},
		Checks: []*Check{{
			Query:        "SELECT 1 as result;",
			ExpectOutput: true,
		}},
	}
	multiLayerWithEmptySubPolicy = &Policy{
		Name: "test",
		Policies: Policies{
			{
				Name:   "subpolicy",
				Checks: []*Check{},
			},
		},
		Checks: []*Check{{
			Query:        "SELECT 1 as result;",
			ExpectOutput: true,
		}},
	}
	// views cannot be inherited from parent policies.
	multiLayerWithInheritedView = &Policy{
		Name: "test",
		Views: []*View{
			{
				Name:  "testview",
				Query: "SELECT 'something'",
			},
		},
		Policies: Policies{
			{
				Name: "subpolicy",
				Checks: []*Check{
					{
						Name:         "query-with-view",
						ExpectOutput: true,
						Query:        "SELECT * from testview",
					},
				},
			},
		},
	}
)

func TestExecutor_Execute(t *testing.T) {
	cases := []struct {
		Name                 string
		Policy               *Policy
		Selector             string
		ShouldBeEmpty        bool
		Pass                 bool
		ErrorOutput          string
		TotalExpectedResults int
		StopOnFailure        bool
	}{
		{
			Name: "simple policy",
			Policy: &Policy{
				Name:     "test",
				Policies: nil,
				Checks: []*Check{{
					Query:        "SELECT 1 as result;",
					ExpectOutput: true,
				}},
			},
			Pass:                 true,
			TotalExpectedResults: 1,
		},
		{
			Name:                 "multilayer policies",
			Policy:               multiLayerPolicy,
			Pass:                 true,
			TotalExpectedResults: 3,
		},
		{
			Name:                 "multilayer policies \\w selector",
			Policy:               multiLayerPolicy,
			Selector:             "subpolicy",
			Pass:                 true,
			TotalExpectedResults: 2,
		},
		{
			Name:                 "multilayer policies \\w selector on query",
			Policy:               multiLayerPolicy,
			Selector:             "subpolicy/sub-query",
			Pass:                 true,
			TotalExpectedResults: 1,
		},
		{
			Name:                 "failing policy",
			Policy:               failingPolicy,
			Pass:                 false,
			TotalExpectedResults: 3,
		},
		{
			Name:                 "failing policy - stop on failure",
			Policy:               failingPolicy,
			Pass:                 false,
			TotalExpectedResults: 2,
			StopOnFailure:        true,
		},
		{
			Name:                 "failing policy \\w selector",
			Policy:               failingPolicy,
			Selector:             "subpolicy/sub-query",
			Pass:                 true,
			TotalExpectedResults: 1,
			StopOnFailure:        true,
		},
		{
			Name:                 "multilayer policy w/ empty subpolicy",
			Policy:               multiLayerWithEmptySubPolicy,
			Pass:                 true,
			TotalExpectedResults: 1,
		},
		{
			Name:        "multilayer policy w/ using view inherited from parent",
			Policy:      multiLayerWithInheritedView,
			ErrorOutput: "relation \"testview\" does not exist",
		},
	}

	conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)
	executor := NewExecutor(conn, hclog.Default(), nil)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			execReq := &ExecuteRequest{
				Policy:         tc.Policy,
				UpdateCallback: nil,
				StopOnFailure:  tc.StopOnFailure,
			}
			filtered := tc.Policy.Filter(tc.Selector)
			res, err := executor.Execute(context.Background(), execReq, &filtered)
			if tc.ErrorOutput != "" {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tc.ErrorOutput)
				}
				return
			} else {
				assert.NoError(t, err)
			}
			if tc.ShouldBeEmpty {
				assert.Empty(t, res)
			} else {
				assert.NotEmpty(t, res)
			}
			if res != nil {
				assert.Equal(t, tc.Pass, res.Passed)
				assert.Len(t, res.Results, tc.TotalExpectedResults)
			}
		})
	}
}

func setupCheckFetchDatabase(db execution.QueryExecer, summary *meta_storage.FetchSummary, c *meta_storage.Client) (func(t *testing.T), error) {
	if summary == nil {
		return func(t *testing.T) {}, nil
	}
	summary.CqId = uuid.New()
	summary.FetchId = uuid.New()
	finish := time.Now().UTC()
	summary.Finish = &finish
	err := c.SaveFetchSummary(context.Background(), summary)
	if err != nil {
		return nil, err
	}

	// Return conn and tear down func
	return func(t *testing.T) {
		err = db.Exec(context.Background(), fmt.Sprintf(`DELETE FROM "cloudquery"."fetches" WHERE "id" = '%s';`, summary.FetchId.String()))
		assert.NoError(t, err)
	}, nil
}

func TestExecuter_DisbleFetchCheckFlag(t *testing.T) {
	db, err := sdkdb.New(context.Background(), hclog.NewNullLogger(), testDBConnection)
	assert.NoError(t, err)

	metaStorage := meta_storage.NewClient(db, hclog.NewNullLogger())

	_, de, err := database.GetExecutor(hclog.NewNullLogger(), testDBConnection, &history.Config{})
	if err != nil {
		t.Fatal(fmt.Errorf("getExecutor: %w", err))
	}

	err = metaStorage.MigrateCore(context.Background(), de)
	assert.NoError(t, err)

	executor := NewExecutor(db, hclog.Default(), nil)

	policy := &Policy{
		Name:     "test",
		Policies: nil,
		Checks: []*Check{{
			Query:        "SELECT 1 as result;",
			ExpectOutput: true,
		}},
		Config: &Configuration{
			Providers: []*Provider{
				{
					Type:    "testProvider",
					Version: ">0.0.0",
				},
			},
		},
	}

	testCases := []struct {
		Name              string
		DisableFetchCheck bool
		ExpectedError     error
	}{{
		Name:              "fetch_check_enabled",
		DisableFetchCheck: false,
		ExpectedError:     errors.New("could not find a completed fetch for requested provider"),
	},
		{
			Name:              "fetch_check_disabled",
			DisableFetchCheck: true,
			ExpectedError:     nil,
		},
	}

	testProviderVersion, err := version.NewVersion("0.1.0")
	assert.NoError(t, err)

	executeRequest := &ExecuteRequest{
		Policy:           policy,
		ProviderVersions: map[string]*version.Version{"testProvider": testProviderVersion},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			defer viper.Reset()
			viper.Set("disable-fetch-check", tc.DisableFetchCheck)

			_, err = executor.Execute(context.Background(), executeRequest, policy)

			if tc.ExpectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.ExpectedError.Error())
			}
		})
	}

}

func TestExecutor_CheckFetches(t *testing.T) {
	// create database connection
	db, err := sdkdb.New(context.Background(), hclog.NewNullLogger(), testDBConnection)
	assert.NoError(t, err)

	metaStorage := meta_storage.NewClient(db, hclog.NewNullLogger())

	_, de, err := database.GetExecutor(hclog.NewNullLogger(), testDBConnection, &history.Config{})
	if err != nil {
		t.Fatal(fmt.Errorf("getExecutor: %w", err))
	}

	err = metaStorage.MigrateCore(context.Background(), de)
	assert.NoError(t, err)

	executor := NewExecutor(db, hclog.Default(), nil)

	finish := time.Now().UTC()
	assert.NoError(t, err)
	cases := []struct {
		Name   string
		Config Configuration
		f      *meta_storage.FetchSummary
		err    error
	}{
		{
			Name: "correct version",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test1", Version: "~> v0.2.0"},
				},
			},
			f:   &meta_storage.FetchSummary{ProviderName: "test1", ProviderVersion: "v0.2.3", Finish: &finish, IsSuccess: true},
			err: nil,
		},
		{
			Name: "no fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test2", Version: "~> v0.2.0"},
				},
			},
			err: errors.New("failed to get fetch summary for provider test2: could not find a completed fetch for requested provider"),
		}, {
			Name: "no finished fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "no_finish", Version: "~> v0.2.0"},
				},
			},
			f:   &meta_storage.FetchSummary{ProviderName: "test3", ProviderVersion: "v0.2.3", IsSuccess: false},
			err: errors.New("failed to get fetch summary for provider no_finish: could not find a completed fetch for requested provider"),
		},
		{
			Name: "no fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test3", Version: "~> v0.2.0"},
				},
			},
			f:   &meta_storage.FetchSummary{ProviderName: "test3", ProviderVersion: "v0.2.3", Finish: &finish, IsSuccess: false},
			err: errors.New("last fetch for provider test3 wasn't successful"),
		},
		{
			Name: "no fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test4", Version: "~> v0.3.0"},
				},
			},
			f:   &meta_storage.FetchSummary{ProviderName: "test4", ProviderVersion: "v0.2.3", Finish: &finish, IsSuccess: true},
			err: errors.New("the latest fetch for provider test4 does not satisfy version requirement ~> v0.3.0"),
		},
		{
			Name: "no fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test4", Version: ""},
				},
			},
			f:   &meta_storage.FetchSummary{ProviderName: "test4", ProviderVersion: "v0.2.3", Finish: &finish, IsSuccess: true},
			err: errors.New("failed to parse version constraint for provider test4: Malformed constraint: "),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			clear, err := setupCheckFetchDatabase(db, tc.f, metaStorage)
			assert.NoError(t, err)

			err = executor.checkFetches(context.Background(), &tc.Config)
			if tc.err != nil {
				assert.Equal(t, tc.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
			clear(t)
		})

	}
}
