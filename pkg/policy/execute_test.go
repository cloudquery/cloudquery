package policy

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

func setupDatabase(t *testing.T, tableName string) (*pgxpool.Pool, func(t *testing.T)) {
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

	pool, tearDownFunc := setupDatabase(t, t.Name())
	defer tearDownFunc(t)
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)
	executor := NewExecutor(conn, hclog.Default())

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			res, err := executor.executeQuery(context.Background(), &config.Query{
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
		Queries       []*config.Query
		Views         []*config.View
		ShouldBeEmpty bool
		Pass          bool
		ErrorOutput   string
	}{
		{
			Name: "multiple_queries",
			Queries: []*config.Query{
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
			Views: []*config.View{
				{
					Name: "testview",
					Query: &config.Query{
						Name:  "get-john",
						Query: fmt.Sprintf("SELECT * FROM %s WHERE name LIKE 'john'", t.Name()),
					},
				},
			},
			Queries: []*config.Query{
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
			Queries: []*config.Query{
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
			Views: []*config.View{
				{
					Name: "brokenview",
					Query: &config.Query{
						Name:  "broken-query-view",
						Query: "TCELES * MOFR *",
					},
				},
			},
			ErrorOutput:   "broken_policy_view/brokenview/broken-query-view: ERROR: syntax error at or near \"TCELES\" (SQLSTATE 42601)",
			ShouldBeEmpty: true,
			Pass:          true,
		},
	}

	pool, tearDownFunc := setupDatabase(t, t.Name())
	defer tearDownFunc(t)
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)
	executor := NewExecutor(conn, hclog.Default())

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			p := &config.Policy{
				Name:    tc.Name,
				Queries: tc.Queries,
				Views:   tc.Views,
			}
			execReq := &ExecuteRequest{
				UpdateCallback: nil,
				StopOnFailure:  false,
			}

			res, err := executor.executePolicy(context.Background(), execReq, p, nil)
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
