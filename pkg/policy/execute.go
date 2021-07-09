package policy

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Executor implements the execution framework.
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

// ExecutionCallback represents the format of the policy callback function.
type ExecutionCallback func(name string, passed bool)

// ExecuteRequest is a request that triggers policy execution.
type ExecuteRequest struct {
	// Policy is the policy that should be executed.
	Policy *Policy

	// UpdateCallback gets called when the client receives updates on policy execution (optional).
	UpdateCallback ExecutionCallback

	// StopOnFailure if true policy execution will stop on first failure
	StopOnFailure bool

	// SkipVersioning if true policy will be executed without checking out the version of the policy repo using git tags
	SkipVersioning bool
}

// NewExecutor creates a new executor.
func NewExecutor(conn *pgxpool.Conn) *Executor {
	return &Executor{
		conn: conn,
	}
}

// ExecutePolicies executes multiple given policies and the related sub queries/views.
// Note: It does not execute sub policies that are attached to the policies.
// Is is the callers responsibility to do that.
func (e *Executor) ExecutePolicies(ctx context.Context, execReq *ExecuteRequest, policyMap map[string]*config.Policy) (*ExecutionResult, error) {
	execResults := &ExecutionResult{
		Passed:  true,
		Results: make(map[string]*QueryResult),
	}

	// Iterate over all given policies
	for path, policy := range policyMap {
		// Execute policy
		results, err := e.executePolicy(ctx, policy, execReq)
		if err != nil {
			return nil, err
		}

		// Collect results
		collectExecutionResults(execResults, path, results...)

		// Execute callback method
		if execReq.UpdateCallback != nil {
			execReq.UpdateCallback(policy.Name, execResults.Passed)
		}

		// Skip further execution if exit on error is defined
		if execReq.StopOnFailure && !execResults.Passed {
			break
		}
	}
	return execResults, nil
}

// executePolicy executes the given policy and its sub views/queries.
// Please use ExecutePolicies if possible.
func (e *Executor) executePolicy(ctx context.Context, p *config.Policy, execReq *ExecuteRequest) ([]*QueryResult, error) {
	results := make([]*QueryResult, 0)

	// Create temporary Views
	for _, v := range p.Views {
		if err := e.CreateView(ctx, v); err != nil {
			return nil, fmt.Errorf("%s: %w", p.Name, err)
		}
	}

	// Execute queries
	for _, q := range p.Queries {
		res, err := e.ExecuteQuery(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p.Name, err)
		}
		results = append(results, res)

		// Stop execution if defined on error
		if execReq.StopOnFailure && !res.Passed {
			return results, nil
		}
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

// collectExecutionResults collects all query results and adds them to the
// execution results struct.
func collectExecutionResults(execResult *ExecutionResult, path string, results ...*QueryResult) {
	for _, res := range results {
		if !res.Passed {
			execResult.Passed = false
		}
		execResult.Results[filepath.Join(path, res.Name)] = res
	}
}
