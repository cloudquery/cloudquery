package policy

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

func setupPolicyDatabase(t *testing.T, tableName string) (*pgxpool.Pool, func(t *testing.T)) {
	poolCfg, err := pgxpool.ParseConfig("postgres://postgres:pass@localhost:5432/postgres")
	assert.NoError(t, err)
	poolCfg.LazyConnect = true
	pool, err := pgxpool.ConnectConfig(context.Background(), poolCfg)
	assert.NoError(t, err)
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)

	// Setup test data
	_, err = conn.Exec(context.Background(), fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
	assert.NoError(t, err)
	_, err = conn.Exec(context.Background(), fmt.Sprintf("CREATE TABLE %s (id serial PRIMARY KEY, name VARCHAR(50) NOT NULL)", tableName))
	assert.NoError(t, err)
	_, err = conn.Exec(context.Background(), fmt.Sprintf("INSERT INTO %s VALUES (1, 'john')", tableName))
	assert.NoError(t, err)

	// Return conn and tear down func
	return pool, func(t *testing.T) {
		_, err = conn.Exec(context.Background(), fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName))
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
			Name:          "nooutput",
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

	pool, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)
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
					Query: "TCELES * MOFR *",
				},
			},
			ErrorOutput:   "failed to create view broken_policy_view/brokenview: ERROR: syntax error at or near \"TCELES\" (SQLSTATE 42601)",
			ShouldBeEmpty: true,
			Pass:          true,
		},
	}

	pool, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)
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
			Name:                 "multilayer policies \\w invalid selector",
			Policy:               multiLayerPolicy,
			Selector:             "invalidselector",
			Pass:                 true,
			ShouldBeEmpty:        true,
			TotalExpectedResults: 0,
			ErrorOutput:          "test//invalidselector: selected policy/query is not found",
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
	}

	pool, tearDownFunc := setupPolicyDatabase(t, t.Name())
	defer tearDownFunc(t)
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)
	executor := NewExecutor(conn, hclog.Default(), nil)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			execReq := &ExecuteRequest{
				Policy:         tc.Policy,
				UpdateCallback: nil,
				StopOnFailure:  tc.StopOnFailure,
			}
			filtered := tc.Policy.Filter(tc.Selector)
			log.Println(filtered.Checks)
			res, err := executor.Execute(context.Background(), execReq, &filtered)
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
				assert.Len(t, res.Results, tc.TotalExpectedResults)
			}
		})
	}
}
