package policy

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var onceExec sync.Once

var executorInstance *Executor

type Executor struct {
	// Connection to the database
	conn *pgxpool.Conn
}

// QueryResult contains the result information from an executed query.
type QueryResult struct {
	Name    string          `json:"name"`
	Columns []string        `json:"result_headers"`
	Data    [][]interface{} `json:"result_rows"`
	Passed  bool            `json:"check_passed"`
}

// ExecutionResult contains all policy execution results.
type ExecutionResult struct {
	// True if all policies have passed
	Passed bool

	// Map of all query result sets
	Results map[string]*QueryResult
}

func NewExecutor(conn *pgxpool.Conn) *Executor {
	onceExec.Do(func() {
		executorInstance = &Executor{
			conn: conn,
		}
	})
	return executorInstance
}

// ExecutePolicy executes the given policy and its sub views/queries.
// Note: It does not execute sub policies that are attached to this policy.
// It is the callers responsibility to do that.
func (e *Executor) ExecutePolicy(ctx context.Context, p *config.Policy) ([]*QueryResult, error) {
	var results []*QueryResult

	// Create temporary Views
	for _, v := range p.Views {
		if err := e.CreateView(ctx, v); err != nil {
			return nil, err
		}
	}

	// Execute queries
	for _, q := range p.Queries {
		res, err := e.ExecuteQuery(ctx, q)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

// ExecuteQuery executes the given query and returns the result.
func (e *Executor) ExecuteQuery(ctx context.Context, q *config.Query) (*QueryResult, error) {
	data, err := e.conn.Query(ctx, q.Query)
	if err != nil {
		return nil, err
	}

	result := &QueryResult{
		Name:    q.Name,
		Columns: make([]string, 0),
		Data:    make([][]interface{}, 0),
		Passed:  false,
	}
	for _, fd := range data.FieldDescriptions() {
		result.Columns = append(result.Columns, string(fd.Name))
	}

	for data.Next() {
		values, err := data.Values()
		if err != nil {
			return nil, err
		}
		result.Data = append(result.Data, values)
	}
	if (len(result.Data) == 0 && !q.ExpectOutput) || (q.ExpectOutput && len(result.Data) > 0) {
		result.Passed = true
	}
	return result, nil
}

// CreateView creates the given view temporary.
func (e *Executor) CreateView(ctx context.Context, v *config.View) error {
	// Add create view command
	v.Query.Query = fmt.Sprintf("CREATE OR REPLACE TEMPORARY VIEW %s AS %s", v.Name, v.Query.Query)

	// Create view and ignore the output
	_, err := e.ExecuteQuery(ctx, v.Query)
	return err
}
