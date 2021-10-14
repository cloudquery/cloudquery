package policy

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/go-hclog"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var errPolicyOrQueryNotFound = errors.New("selected policy/query is not found")

// Executor implements the execution framework.
type Executor struct {
	// Connection to the database
	conn *pgxpool.Conn
	log  hclog.Logger
}

// QueryResult contains the result information from an executed query.
type QueryResult struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Columns     []string         `json:"result_headers"`
	Data        [][]interface{}  `json:"result_rows"`
	Type        config.QueryType `json:"type"`
	Passed      bool             `json:"check_passed"`
}

// ExecutionResult contains all policy execution results.
type ExecutionResult struct {
	// True if all policies have passed
	Passed bool

	// Map of all query result sets
	Results map[string]*QueryResult
}

// ExecutionCallback represents the format of the policy callback function.
type ExecutionCallback func(name string, queryType config.QueryType, passed bool)

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

// executePolicy executes given policy and the related sub queries/views.
func (e *Executor) executePolicy(ctx context.Context, req *ExecuteRequest, policy *config.Policy, selector []string) (*ExecutionResult, error) {
	if err := e.createViews(ctx, policy); err != nil {
		return nil, err
	}
	var rest []string
	if len(selector) > 0 {
		rest = selector[1:]
	}
	var found bool
	total := ExecutionResult{Passed: true, Results: make(map[string]*QueryResult)}
	for _, p := range policy.Policies {
		if len(selector) == 0 || p.Name == selector[0] {
			found = true
			r, err := e.executePolicy(ctx, req, p, rest)
			if err != nil {
				return nil, fmt.Errorf("%s/%w", policy.Name, err)
			}
			total.Passed = total.Passed && r.Passed
			for k, v := range r.Results {
				total.Results[policyPathJoin(policy.Name, k)] = v
			}
			if !total.Passed && req.StopOnFailure {
				return &total, nil
			}

		}
	}
	for _, q := range policy.Queries {
		if len(selector) == 0 || q.Name == selector[0] {
			found = true
			qr, err := e.executeQuery(ctx, q)
			if err != nil {
				return nil, fmt.Errorf("%s/%w", policy.Name, err)
			}
			total.Passed = total.Passed && qr.Passed
			total.Results[policyPathJoin(policy.Name, q.Name)] = qr
			if req.UpdateCallback != nil {
				req.UpdateCallback(q.Name, qr.Type, qr.Passed)
			}
			if !total.Passed && req.StopOnFailure {
				return &total, nil
			}
		}
	}
	if !found && len(selector) > 0 {
		return nil, fmt.Errorf("%s: %w", policy.Name, errPolicyOrQueryNotFound)
	}
	return &total, nil
}

// executeQuery executes the given query and returns the result.
func (e *Executor) executeQuery(ctx context.Context, q *config.Query) (*QueryResult, error) {
	data, err := e.conn.Query(ctx, q.Query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", q.Name, err)
	}

	result := &QueryResult{
		Name:        q.Name,
		Description: q.Description,
		Columns:     make([]string, 0),
		Data:        make([][]interface{}, 0),
		Type:        q.Type,
	}
	for _, fd := range data.FieldDescriptions() {
		result.Columns = append(result.Columns, string(fd.Name))
	}

	for data.Next() {
		values, err := data.Values()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", q.Name, err)
		}
		result.Data = append(result.Data, values)
	}
	if data.Err() != nil {
		return nil, fmt.Errorf("%s: %w", q.Name, data.Err())
	}
	result.Passed = (len(result.Data) == 0) == !q.ExpectOutput
	return result, nil
}

// createViews creates temporary views for given config.Policy, and any views defined by sub-policies
func (e *Executor) createViews(ctx context.Context, policy *config.Policy) error {
	for _, v := range policy.Views {
		e.log.Debug("creating policy view", "policy", policy.Name, "view", v.Name)
		if err := e.createView(ctx, v); err != nil {
			return fmt.Errorf("%s/%s/%w", policy.Name, v.Name, err)
		}
	}
	return nil
}

// createView creates the given view temporary.
func (e *Executor) createView(ctx context.Context, v *config.View) error {
	// Add create view command
	v.Query.Query = fmt.Sprintf("CREATE OR REPLACE TEMPORARY VIEW %s AS %s", v.Name, v.Query.Query)

	// Create view and ignore the output
	_, err := e.executeQuery(ctx, v.Query)
	return err
}

func (e *Executor) ExecutePolicies(ctx context.Context, req *ExecuteRequest, policies []*config.Policy, selector []string) (*ExecutionResult, error) {
	var rest []string
	if len(selector) > 0 {
		rest = selector[1:]
	}
	var found bool
	total := ExecutionResult{Passed: true, Results: make(map[string]*QueryResult)}
	for _, p := range policies {
		if len(selector) == 0 || selector[0] == p.Name {
			found = true
			r, err := e.executePolicy(ctx, req, p, rest)
			if err != nil {
				return nil, err
			}
			total.Passed = total.Passed && r.Passed
			for k, v := range r.Results {
				total.Results[k] = v
			}
			if !total.Passed && req.StopOnFailure {
				return &total, nil
			}
		}
	}
	if !found && len(selector) > 0 {
		return nil, errPolicyOrQueryNotFound
	}
	return &total, nil
}
