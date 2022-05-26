package policy

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/cloudquery/cloudquery/pkg/core/state"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
)

func setupPolicyDatabase(t *testing.T, tableName string) (string, LowLevelQueryExecer, func(t *testing.T)) {
	baseDSN := os.Getenv("CQ_CLIENT_TEST_DSN")
	if baseDSN == "" {
		baseDSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}

	conn, err := sdkdb.New(context.Background(), hclog.NewNullLogger(), baseDSN)
	assert.NoError(t, err)

	// Setup test data
	err = conn.Exec(context.Background(), fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
	assert.NoError(t, err)
	err = conn.Exec(context.Background(), fmt.Sprintf("CREATE TABLE %s (id serial PRIMARY KEY, name VARCHAR(50) NOT NULL)", tableName))
	assert.NoError(t, err)
	err = conn.Exec(context.Background(), fmt.Sprintf("INSERT INTO %s VALUES (1, 'john')", tableName))
	assert.NoError(t, err)

	// Return conn and tear down func
	return baseDSN, conn, func(t *testing.T) {
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

	dsn, conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)

	sta, err := state.NewClient(context.Background(), dsn)
	if err != nil {
		assert.NoError(t, err)
	}
	defer sta.Close()

	executor := NewExecutor(conn, sta, nil)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			res, err := executor.executeQuery(context.Background(), &Check{
				Query:        tc.Query,
				ExpectOutput: tc.ExpectOutput,
			}, nil)
			assert.NoError(t, err)
			if tc.ShouldBeEmpty {
				assert.Empty(t, res.Rows)
			} else {
				assert.NotEmpty(t, res.Rows)
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
		ExpectedDiags []diag.FlatDiag
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
			ExpectedDiags: []diag.FlatDiag{{Err: "broken_policy_query/broken-query: ERROR: syntax error at or near \"SECT\" (SQLSTATE 42601)", Type: diag.DATABASE, Severity: diag.ERROR,
				Summary:     "broken_policy_query/broken-query: ERROR: syntax error at or near \"SECT\" (SQLSTATE 42601)",
				Description: diag.Description{Summary: "broken_policy_query/broken-query: ERROR: syntax error at or near \"SECT\" (SQLSTATE 42601)", Detail: ""}}},
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
			ExpectedDiags: []diag.FlatDiag{{Err: "broken_policy_view/failed to create view broken_policy_view/brokenview: ERROR: syntax error at or near \"INVALID\" (SQLSTATE 42601)", Type: diag.DATABASE, Severity: diag.ERROR,
				Summary:     "broken_policy_view/failed to create view broken_policy_view/brokenview: ERROR: syntax error at or near \"INVALID\" (SQLSTATE 42601)",
				Description: diag.Description{Summary: "broken_policy_view/failed to create view broken_policy_view/brokenview: ERROR: syntax error at or near \"INVALID\" (SQLSTATE 42601)", Detail: ""}}},
			ShouldBeEmpty: true,
			Pass:          true,
		},
	}

	dsn, conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)

	sta, err := state.NewClient(context.Background(), dsn)
	if err != nil {
		assert.NoError(t, err)
	}
	defer sta.Close()

	executor := NewExecutor(conn, sta, nil)

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

			res, diags := executor.Execute(context.Background(), execReq, p, nil)
			if tc.ExpectedDiags != nil {
				assert.ElementsMatch(t, tc.ExpectedDiags, diag.FlattenDiags(diags, false))
			} else {
				assert.Equal(t, []diag.FlatDiag{}, diag.FlattenDiags(diags, false))
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
		ExpectedDiags        []diag.FlatDiag
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
			Name:          "multilayer policy w/ using view inherited from parent",
			Policy:        multiLayerWithInheritedView,
			ShouldBeEmpty: true,
			ExpectedDiags: []diag.FlatDiag{{Err: "test/subpolicy/query-with-view: ERROR: relation \"testview\" does not exist (SQLSTATE 42P01)", Type: diag.DATABASE, Severity: diag.ERROR,
				Summary:     "test/subpolicy/query-with-view: ERROR: relation \"testview\" does not exist (SQLSTATE 42P01)",
				Description: diag.Description{Summary: "test/subpolicy/query-with-view: ERROR: relation \"testview\" does not exist (SQLSTATE 42P01)", Detail: ""}}},
		},
	}

	dsn, conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)

	sta, err := state.NewClient(context.Background(), dsn)
	if err != nil {
		assert.NoError(t, err)
	}
	defer sta.Close()

	executor := NewExecutor(conn, sta, nil)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			execReq := &ExecuteRequest{
				Policy:         tc.Policy,
				UpdateCallback: nil,
				StopOnFailure:  tc.StopOnFailure,
			}
			filtered := tc.Policy.Filter(tc.Selector)
			res, diags := executor.Execute(context.Background(), execReq, &filtered, nil)
			if tc.ExpectedDiags != nil {
				assert.ElementsMatch(t, tc.ExpectedDiags, diag.FlattenDiags(diags, false))
			} else {
				assert.Equal(t, []diag.FlatDiag{}, diag.FlattenDiags(diags, false))
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

func setupCheckFetchDatabase(db execution.QueryExecer, summary *state.FetchSummary, c *state.Client) (func(t *testing.T), error) {
	if summary == nil {
		return func(t *testing.T) {}, nil
	}
	summary.CqId = uuid.New()
	summary.FetchId = uuid.New()
	finish := time.Now().UTC()
	summary.Finish = finish
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

func TestExecutor_DisableFetchCheckFlag(t *testing.T) {
	dsn, conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)

	sta, err := state.NewClient(context.Background(), dsn)
	if err != nil {
		assert.NoError(t, err)
	}
	defer sta.Close()

	executor := NewExecutor(conn, sta, nil)

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
		ExpectedDiags     []diag.FlatDiag
	}{{
		Name:              "fetch_check_enabled",
		DisableFetchCheck: false,
		ExpectedDiags: []diag.FlatDiag{{Err: "failed to get fetch summary for provider testProvider: could not find a completed fetch for requested provider",
			Type: diag.USER, Severity: diag.ERROR, Summary: "failed to get fetch summary for provider testProvider: could not find a completed fetch for requested provider", Description: diag.Description{Resource: "", ResourceID: []string(nil), Summary: "failed to get fetch summary for provider testProvider: could not find a completed fetch for requested provider", Detail: "test: please run `cloudquery fetch` before running policy"}}},
	},
		{
			Name:              "fetch_check_disabled",
			DisableFetchCheck: true,
		},
	}

	executeRequest := &ExecuteRequest{
		Policy: policy,
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			defer viper.Reset()
			viper.Set("disable-fetch-check", tc.DisableFetchCheck)

			_, diags := executor.Execute(context.Background(), executeRequest, policy, nil)

			if tc.ExpectedDiags != nil {
				assert.ElementsMatch(t, tc.ExpectedDiags, diag.FlattenDiags(diags, false))
			} else {
				assert.Equal(t, []diag.FlatDiag{}, diag.FlattenDiags(diags, false))
			}
		})
	}

}

func TestExecutor_CheckFetches(t *testing.T) {
	dsn, conn, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)

	sta, err := state.NewClient(context.Background(), dsn)
	if err != nil {
		assert.NoError(t, err)
	}
	defer sta.Close()

	executor := NewExecutor(conn, sta, nil)

	finish := time.Now().UTC()
	assert.NoError(t, err)
	cases := []struct {
		Name   string
		Config Configuration
		f      *state.FetchSummary
		err    error
	}{
		{
			Name: "correct version",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test1", Version: "~> v0.2.0"},
				},
			},
			f:   &state.FetchSummary{ProviderName: "test1", ProviderVersion: "v0.2.3", Finish: finish, IsSuccess: true},
			err: nil,
		},
		{
			Name: "no finished fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "no_finish", Version: "~> v0.2.0"},
				},
			},
			f:   &state.FetchSummary{ProviderName: "test3", ProviderVersion: "v0.2.3", IsSuccess: false},
			err: errors.New("failed to get fetch summary for provider no_finish: could not find a completed fetch for requested provider"),
		},
		{
			Name: "no fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test3", Version: "~> v0.2.0"},
				},
			},
			f:   &state.FetchSummary{ProviderName: "test3", ProviderVersion: "v0.2.3", Finish: finish, IsSuccess: false},
			err: errors.New("last fetch for provider test3 wasn't successful"),
		},
		{
			Name: "no fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test4", Version: "~> v0.3.0"},
				},
			},
			f:   &state.FetchSummary{ProviderName: "test4", ProviderVersion: "v0.2.3", Finish: finish, IsSuccess: true},
			err: errors.New("the latest fetch for provider test4 does not satisfy version requirement ~> v0.3.0"),
		},
		{
			Name: "no fetches",
			Config: Configuration{
				Providers: []*Provider{
					{Type: "test4", Version: ""},
				},
			},
			f:   &state.FetchSummary{ProviderName: "test4", ProviderVersion: "v0.2.3", Finish: finish, IsSuccess: true},
			err: errors.New("failed to parse version constraint for provider test4: Malformed constraint: "),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			clear, err := setupCheckFetchDatabase(conn, tc.f, sta)
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

func TestInterpolate(t *testing.T) {
	tpl, err := template.New("query").Parse("{{.test}} {{.key}}")
	assert.Nil(t, err)
	var b strings.Builder

	err = tpl.Execute(&b, map[string]interface{}{"test": 1, "key": "lol"})
	assert.Nil(t, err)
	fmt.Print(b.String())
}

func TestRow_Sort(t *testing.T) {

	testCases := []struct {
		Name     string
		Data     Rows
		Expected Rows
	}{
		{
			Name: "simple",
			Data: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"a": "c"}},
			},
			Expected: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"a": "c"}},
			},
		},
		{
			Name: "same",
			Data: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"a": "b"}},
			},
			Expected: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"a": "b"}},
			},
		},
		{
			Name: "complex",
			Data: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"k": "b"}},
				{Identifiers: map[string]interface{}{"z": "b"}},
			},
			Expected: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"k": "b"}},
				{Identifiers: map[string]interface{}{"z": "b"}},
			},
		},
		{
			Name: "complex-2nd-level",
			Data: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"k": "c"}},
				{Identifiers: map[string]interface{}{"k": "b"}},
			},
			Expected: Rows{
				{Identifiers: map[string]interface{}{"a": "b"}},
				{Identifiers: map[string]interface{}{"k": "b"}},
				{Identifiers: map[string]interface{}{"k": "c"}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			sort.Sort(tc.Data)
			assert.Equal(t, tc.Expected, tc.Data)
		})
	}

}

// Tests the path of the policy execution result file is correct.
// It should always be named as the root policy. i.e. for policy run "aws//cis_v1.2.0", it should be named
// aws.json
func TestGenerateExecutionResultFile(t *testing.T) {
	tempDirPath, err := ioutil.TempDir("", "cqtest")
	if !assert.NoError(t, err) {
		return
	}
	defer os.RemoveAll(tempDirPath)

	testCases := []struct {
		PolicyName   string
		OutputDir    string
		ExpectedFile string
	}{
		{
			"aws//cis_v1.2.0",
			filepath.Join(tempDirPath, "1"),
			filepath.Join(tempDirPath, "1", "aws.json"),
		},
		{
			"aws//cis_v1.2.0",
			filepath.Join(tempDirPath, ".", "2"),
			filepath.Join(tempDirPath, "2", "aws.json"),
		},
		{
			"aws//cis_v1.2.0",
			filepath.Join(tempDirPath, "3", "..", "3"),
			filepath.Join(tempDirPath, "3", "aws.json"),
		},
		{
			"aws//cis_v1.2.0/somecheck",
			filepath.Join(tempDirPath, "4"),
			filepath.Join(tempDirPath, "4", "aws.json"),
		},
		{
			"aws",
			filepath.Join(tempDirPath, "5", "aws.json"),
			filepath.Join(tempDirPath, "5"),
		},
	}

	for _, tc := range testCases {
		executiionResult := ExecutionResult{PolicyName: tc.PolicyName}
		err := GenerateExecutionResultFile(&executiionResult, tc.OutputDir)
		if !assert.NoError(t, err) {
			continue
		}

		_, err = os.Stat(tc.ExpectedFile) // os.Stat will return an error if file doesn't exist
		assert.NoError(t, err)
	}

}

func TestNormalizeCheckSelector(t *testing.T) {

	testCases := []struct {
		pe               *state.PolicyExecution
		PolicyPath       []string
		PolicyName       string
		ExpectedSelector string
	}{
		{
			&state.PolicyExecution{
				Location:   "aws",
				PolicyName: "aws",
			},
			[]string{"pci_dss_v3.2.1", "cloudtrail"},
			"4",
			"aws//pci_dss_v3.2.1/cloudtrail/4",
		}, {
			&state.PolicyExecution{
				Location:   "aws",
				PolicyName: "aws//pci_dss_v3.2.1",
			},
			[]string{"cloudtrail"},
			"4",
			"aws//pci_dss_v3.2.1/cloudtrail/4",
		}, {
			&state.PolicyExecution{
				Location:   "aws",
				PolicyName: "aws//pci_dss_v3.2.1/cloudtrail",
			},
			[]string{},
			"4",
			"aws//pci_dss_v3.2.1/cloudtrail/4",
		},
	}

	for _, tc := range testCases {
		computedSelector := normalizeCheckSelector(tc.pe, tc.PolicyPath, tc.PolicyName)
		assert.Equal(t, computedSelector, tc.ExpectedSelector)

	}

}
