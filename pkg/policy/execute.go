package policy

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Executor implements the execution framework.
type Executor struct {
	// Connection to the database
	conn *pgxpool.Conn
	log  hclog.Logger
}

// QueryResult contains the result information from an executed query.
type QueryResult struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Columns     []string        `json:"result_headers"`
	Data        [][]interface{} `json:"result_rows"`
	Passed      bool            `json:"check_passed"`
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
func NewExecutor(conn *pgxpool.Conn, log hclog.Logger) *Executor {
	return &Executor{
		conn: conn,
		log:  log,
	}
}

// ExecutePolicy executes given policy and the related sub queries/views.
// The policy execution first creates all views that are defined in the top-level policy and any sub-policy it includes.
func (e *Executor) ExecutePolicy(ctx context.Context, execReq *ExecuteRequest, policy *config.Policy) (*ExecutionResult, error) {
	// Create temporary all Views recursively before executing any query
	if err := e.CreateViews(ctx, policy); err != nil {
		return nil, err
	}
	return e.executePolicy(ctx, policy, execReq)
}

// executePolicy executes a policy and collects all execution results for it's queries and sub-policies
// use the exported ExecutePolicy so views are created.
func (e *Executor) executePolicy(ctx context.Context, policy *config.Policy, execReq *ExecuteRequest) (*ExecutionResult, error) {
	execResults := &ExecutionResult{
		Passed:  true,
		Results: make(map[string]*QueryResult),
	}
	// Execute policies queries
	results, err := e.executePolicyQueries(ctx, policy, execReq)
	if err != nil {
		e.log.Error("failed to execute policy queries", "policy", policy.Name, "err", err)
		return nil, err
	}
	collectExecutionResults(execResults, policy.Name, results...)
	// Execute callback method
	if execReq.UpdateCallback != nil {
		for _, r := range results {
			execReq.UpdateCallback(r.Description, r.Passed)
		}
	}
	// Skip further execution if exit on failure is defined
	if execReq.StopOnFailure && !execResults.Passed {
		return nil, err
	}

	// Iterate over all given sub policies
	for _, subPolicy := range policy.Policies {
		e.log.Debug("executing policy", "policy", subPolicy.Name)
		// Execute policy
		execResult, err := e.executePolicy(ctx, subPolicy, execReq)
		if err != nil {
			e.log.Error("failed to execute policy", "policy", subPolicy.Name, "err", err)
			return nil, err
		}
		// If sub-policy didn't pass and we previous execution was okay so far, update passed to false
		if execResults.Passed && !execResult.Passed {
			execResults.Passed = execResult.Passed
		}
		for k, r := range execResult.Results {
			execResults.Results[policyPathJoin(policy.Name, k)] = r
		}
	}
	return execResults, nil
}

// executePolicyQueries executes the given policy's queries.
// Please use ExecutePolicy if possible.
func (e *Executor) executePolicyQueries(ctx context.Context, p *config.Policy, execReq *ExecuteRequest) ([]*QueryResult, error) {
	results := make([]*QueryResult, 0)
	// Execute queries
	for _, q := range p.Queries {
		res, err := e.ExecuteQuery(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("%s - %s: %w", p.Name, q.Name, err)
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
		Name:        q.Name,
		Description: q.Description,
		Columns:     make([]string, 0),
		Data:        make([][]interface{}, 0),
		Passed:      false,
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
	if data.Err() != nil {
		return nil, data.Err()
	}
	if (len(result.Data) == 0 && !q.ExpectOutput) || (q.ExpectOutput && len(result.Data) > 0) {
		result.Passed = true
	}
	return result, nil
}

// CreateViews creates temporary views for given config.Policy, and any views defined by sub-policies
func (e *Executor) CreateViews(ctx context.Context, policy *config.Policy) error {
	for _, v := range policy.Views {
		e.log.Debug("creating policy view", "policy", policy.Name, "view", v.Name)
		if err := e.CreateView(ctx, v); err != nil {
			return fmt.Errorf("%s - %s: %w", policy.Name, v.Name, err)
		}
	}
	for _, p := range policy.Policies {
		if err := e.CreateViews(ctx, p); err != nil {
			return err
		}
	}
	return nil
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
		execResult.Results[policyPathJoin(path, res.Name)] = res
	}
}
