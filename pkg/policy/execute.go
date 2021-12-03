package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/afero"

	"github.com/cloudquery/cloudquery/pkg/config"
)

var errPolicyOrQueryNotFound = errors.New("selected policy/query is not found")

type UpdateCallback func(update Update)

type Update struct {
	// PolicyID is the name of the policy that is being updated.
	PolicyName string
	// Version is the policy version.
	Version string
	// FinishedQueries is the number queries that have finished evaluating
	FinishedQueries int
	// QueriesCount is the amount of queries collected so far
	QueriesCount int
	// Error if any returned by the provider
	Error string
}

func (f Update) AllDone() bool {
	return f.FinishedQueries == f.QueriesCount
}

func (f Update) DoneCount() int {
	return f.FinishedQueries
}

// Executor implements the execution framework.
type Executor struct {
	// Connection to the database
	conn *pgxpool.Conn
	log  hclog.Logger

	// progressUpdate
	progressUpdate UpdateCallback
}

// QueryResult contains the result information from an executed query.
type QueryResult struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Columns     []string        `json:"result_headers"`
	Data        [][]interface{} `json:"result_rows"`
	Type        QueryType       `json:"type"`
	Passed      bool            `json:"check_passed"`
}

// ExecutionResult contains all policy execution results.
type ExecutionResult struct {
	// PolicyName is the running policy name
	PolicyName string

	// True if all policies have passed
	Passed bool

	// List of all query result sets
	Results []*QueryResult

	// Error is the reason the execution failed
	Error string
}

// ExecuteRequest is a request that triggers policy execution.
type ExecuteRequest struct {
	// Policy is the policy that should be executed.
	Policy *config.Policy

	// StopOnFailure if true policy execution will stop on first failure
	StopOnFailure bool

	// SkipVersioning if true policy will be executed without checking out the version of the policy repo using git tags
	SkipVersioning bool

	// ProviderVersions describes current versions of providers in use.
	ProviderVersions map[string]*version.Version

	// UpdateCallback is the console ui update callback
	UpdateCallback UpdateCallback
}

// NewExecutor creates a new executor.
func NewExecutor(conn *pgxpool.Conn, log hclog.Logger, progressUpdate UpdateCallback) *Executor {
	return &Executor{
		conn:           conn,
		log:            log,
		progressUpdate: progressUpdate,
	}
}

// executePolicy executes given policy and the related sub queries/views.
func (e *Executor) executePolicy(ctx context.Context, progressUpdate UpdateCallback, req *ExecuteRequest, policy *Policy, selector []string) (*ExecutionResult, error) {
	if err := e.checkVersions(policy.Config, req.ProviderVersions); err != nil {
		return nil, fmt.Errorf("%s: %w", policy.Name, err)
	}
	if err := e.createViews(ctx, policy); err != nil {
		return nil, err
	}
	var rest []string
	if len(selector) > 0 {
		rest = selector[1:]
	}
	var found bool
	total := ExecutionResult{PolicyName: req.Policy.Name, Passed: true, Results: make([]*QueryResult, 0)}
	for _, p := range policy.Policies {
		if len(selector) == 0 || p.Name == selector[0] {
			found = true
			r, err := e.executePolicy(ctx, progressUpdate, req, p, rest)
			if err != nil {
				return nil, fmt.Errorf("%s/%w", policy.Name, err)
			}
			total.Passed = total.Passed && r.Passed
			total.Results = append(total.Results, r.Results...)
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
				e.log.Error("failed to execute query", "policy", policy.Name, "err", err)
				return nil, fmt.Errorf("%s/%w", policy.Name, err)
			}
			total.Passed = total.Passed && qr.Passed
			total.Results = append(total.Results, qr)
			if progressUpdate != nil {
				progressUpdate(Update{
					FinishedQueries: 1,
				})
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

func (*Executor) checkVersions(policyConfig *Configuration, actual map[string]*version.Version) error {
	if policyConfig == nil {
		return nil
	}
	for _, p := range policyConfig.Providers {
		c, err := version.NewConstraint(p.Version)
		if err != nil {
			return fmt.Errorf("failed to parse version constraint for provider %s: %w", p.Type, err)
		}
		v, ok := actual[p.Type]
		if !ok {
			return fmt.Errorf("provider %s version is not defined in configuration", p.Type)
		}
		if !c.Check(v) {
			return fmt.Errorf("provider %s does not satisfy version requirement %s", p.Type, c)
		}
	}
	return nil
}

// executeQuery executes the given query and returns the result.
func (e *Executor) executeQuery(ctx context.Context, q *Query) (*QueryResult, error) {
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
func (e *Executor) createViews(ctx context.Context, policy *Policy) error {
	for _, v := range policy.Views {
		e.log.Debug("creating policy view", "policy", policy.Name, "view", v.Name)
		if err := e.createView(ctx, v); err != nil {
			return fmt.Errorf("%s/%s/%w", policy.Name, v.Name, err)
		}
	}
	return nil
}

// createView creates the given view temporary.
func (e *Executor) createView(ctx context.Context, v *View) error {
	// Add create view command
	v.Query.Query = fmt.Sprintf("CREATE OR REPLACE TEMPORARY VIEW %s AS %s", v.Name, v.Query.Query)

	// Create view and ignore the output
	_, err := e.executeQuery(ctx, v.Query)
	return err
}

func (e *Executor) ExecutePolicies(ctx context.Context, req *ExecuteRequest, policies Policies, selector []string) (*ExecutionResult, error) {
	var rest []string
	var pnames []string
	if len(selector) > 0 {
		rest = selector[1:]
	}
	var found bool
	total := ExecutionResult{PolicyName: req.Policy.Name, Passed: true, Results: make([]*QueryResult, 0)}
	for _, p := range policies {
		pnames = append(pnames, p.Name)
		if len(selector) == 0 || selector[0] == p.Name {
			found = true
			r, err := e.executePolicy(ctx, e.progressUpdate, req, p, rest)
			if err != nil {
				return nil, err
			}
			total.Passed = total.Passed && r.Passed
			total.Results = append(total.Results, r.Results...)
			if !total.Passed && req.StopOnFailure {
				return &total, nil
			}
		}
	}
	if !found && len(selector) > 0 {
		e.log.Error("policy not found with provided selector", "selector", selector, "policy names", pnames)
		return nil, fmt.Errorf("policy not found with provided selector: %s", selector)
	}
	return &total, nil
}

func GenerateExecutionResultFile(result *ExecutionResult, outputDir string) error {
	fs := afero.NewOsFs()

	if err := fs.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	f, err := fs.Create(fmt.Sprintf("%s.json", filepath.Join(outputDir, result.PolicyName)))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	data, err := json.Marshal(&result)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}
