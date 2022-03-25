package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/pkg/client/meta_storage"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/spf13/afero"
)

var ErrPolicyOrQueryNotFound = errors.New("selected policy/query not found")

const testDBConnection = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"

type UpdateCallback func(update Update)

type Update struct {
	// PolicyID is the name of the policy that is being updated.
	PolicyName string
	// Version is the policy version.
	Version string
	// Source policy was fetched from
	Source string
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
	conn LowLevelQueryExecer
	log  hclog.Logger

	PolicyPath []string

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

	// ExecutionTime is when the policy has been started
	ExecutionTime time.Time

	// True if all policies have passed
	Passed bool

	// List of all query result sets
	Results []*QueryResult

	// Error is the reason the execution failed
	Error string

	// List of loaded Policies
	LoadedPolicies Policies
}

// ExecuteRequest is a request that triggers policy execution.
type ExecuteRequest struct {
	// Policy is the policy that should be executed.
	Policy *Policy

	// StopOnFailure if true policy execution will stop on first failure
	StopOnFailure bool

	// ProviderVersions describes current versions of providers in use.
	ProviderVersions map[string]*version.Version

	// UpdateCallback is the console ui update callback
	UpdateCallback UpdateCallback
}

// NewExecutor creates a new executor.
func NewExecutor(conn LowLevelQueryExecer, log hclog.Logger, progressUpdate UpdateCallback) *Executor {
	return &Executor{
		conn:           conn,
		log:            log,
		progressUpdate: progressUpdate,
		PolicyPath:     []string{},
	}
}

func (e *Executor) with(policy string, args ...interface{}) *Executor {
	policyPath := e.PolicyPath
	policyPath = append(policyPath, policy)
	return &Executor{
		conn:           e.conn,
		log:            e.log.With("policy", strings.Join(policyPath, "/")).With(args...),
		progressUpdate: e.progressUpdate,
		PolicyPath:     policyPath,
	}
}

// Execute executes given policy and the related sub queries/views.
func (e *Executor) Execute(ctx context.Context, req *ExecuteRequest, policy *Policy) (*ExecutionResult, error) {
	total := ExecutionResult{PolicyName: req.Policy.Name, Passed: true, Results: make([]*QueryResult, 0)}

	if !policy.HasChecks() {
		e.log.Warn("no checks or policies to execute")
		return &total, nil
	}

	e.log.Debug("Check policy versions", "versions", req.ProviderVersions)
	if err := e.checkVersions(policy.Config, req.ProviderVersions); err != nil {
		return nil, fmt.Errorf("%s: %w", policy.Name, err)
	}
	if err := e.checkFetches(ctx, policy.Config); err != nil {
		return nil, fmt.Errorf("%s: %w, please run `cloudquery fetch` before running policy", policy.Name, err)
	}

	for _, p := range policy.Policies {
		executor := e.with(p.Name)
		executor.log.Info("starting policy execution")
		r, err := executor.Execute(ctx, req, p)
		if err != nil {
			executor.log.Error("failed to execute policy", "err", err)
			return nil, fmt.Errorf("%s/%w", policy.Name, err)
		}
		total.Passed = total.Passed && r.Passed
		total.Results = append(total.Results, r.Results...)
		if !total.Passed && req.StopOnFailure {
			return &total, nil
		}
	}

	// TODO: A better idea here is to create a new session, create the views, run queries, and close the session.
	//       This will remove the need for 'deleteViews'.
	if err := e.createViews(ctx, policy); err != nil {
		return nil, err
	}
	defer e.deleteViews(ctx, policy)

	for _, q := range policy.Checks {
		e.log = e.log.With("query", q.Name)
		qr, err := e.executeQuery(ctx, q)
		if err != nil {
			e.log.Error("failed to execute query", "err", err)
			return nil, fmt.Errorf("%s/%w", policy.Name, err)
		}
		total.Passed = total.Passed && qr.Passed
		total.Results = append(total.Results, qr)
		e.log.Info("Check finished with result", "passed", qr.Passed)
		if e.progressUpdate != nil {
			e.progressUpdate(Update{
				FinishedQueries: 1,
			})
		}
		if !total.Passed && req.StopOnFailure {
			return &total, nil
		}
	}
	return &total, nil
}

// checkFetches checks if there are fetch reports in database that satisfy providers from policy
func (e *Executor) checkFetches(ctx context.Context, policyConfig *Configuration) error {
	if policyConfig == nil {
		return nil
	}
	metaStorage := meta_storage.NewClient(e.conn, e.log)
	for _, p := range policyConfig.Providers {
		c, err := version.NewConstraint(p.Version)
		if err != nil {
			return fmt.Errorf("failed to parse version constraint for provider %s: %w", p.Type, err)
		}
		fetchSummary, err := metaStorage.GetFetchSummaryForProvider(ctx, p.Type)
		if err != nil {
			return fmt.Errorf("failed to get fetch summary for provider %s: %w", p.Type, err)
		}
		if fetchSummary == nil {
			return fmt.Errorf("could not find finished fetches for provider %s", p.Type)
		}
		if !fetchSummary.IsSuccess {
			return fmt.Errorf("last fetch for provider %s wasn't successful", p.Type)
		}
		v, err := version.NewVersion(fetchSummary.ProviderVersion)
		if err != nil {
			return fmt.Errorf("failed to parse version for %s fetch summary: %w", p.Type, err)
		}
		if !c.Check(v) {
			return fmt.Errorf("the latest fetch for provider %s does not satisfy version requirement %s", p.Type, c)
		}
	}
	return nil
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
			return fmt.Errorf("provider %s version %s is not defined in configuration", p.Type, p.Version)
		}
		if !c.Check(v) {
			return fmt.Errorf("provider %s does not satisfy version requirement %s", p.Type, c)
		}
	}
	return nil
}

// executeQuery executes the given query and returns the result.
func (e *Executor) executeQuery(ctx context.Context, q *Check) (*QueryResult, error) {
	e.log.Trace("query", q.Query)
	data, err := e.conn.Query(ctx, q.Query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", q.Name, err)
	}

	result := &QueryResult{
		Name:        q.Name,
		Description: q.Title,
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

// createViews creates temporary views for the given policy (but not for its subpolicies)
func (e *Executor) createViews(ctx context.Context, policy *Policy) error {
	for _, v := range policy.Views {
		e.log.Info("creating policy view", "view", v.Name, "query", v.Query)
		if err := e.conn.Exec(ctx, fmt.Sprintf("CREATE TEMPORARY VIEW %s AS %s", v.Name, v.Query)); err != nil {
			return fmt.Errorf("failed to create view %s/%s: %w", policy.Name, v.Name, err)
		}
	}
	return nil
}

// deleteView deletes the temporary views for the given policy (but not for its subpolicies).
// This method should be executed in 'defer' statements, so it doesn't return an error.
func (e *Executor) deleteViews(ctx context.Context, policy *Policy) {
	for _, v := range policy.Views {

		// Validate that the view is actually a temp view
		data, err := e.conn.Query(ctx, fmt.Sprintf("SELECT table_name FROM INFORMATION_SCHEMA.VIEWS WHERE TABLE_NAME = '%s' and TABLE_SCHEMA LIKE 'pg_temp%%'", v.Name))
		if err != nil {
			e.log.Error("Failed to check if view is temporary", "policy", policy.Name, "view", v.Name, "err", err)
			continue
		}
		count := 0
		for data.Next() {
			count += 1
		}
		if data.Err() != nil {
			e.log.Error("Failed to check if view is temporary", "policy", policy.Name, "view", v.Name, "err", data.Err())
			continue
		}
		// If count is 0 then that means that no temp views with the correct name were found
		if count == 0 {
			continue
		}

		e.log.Info("deleting policy view", "view", v.Name)

		if err := e.conn.Exec(ctx, fmt.Sprintf("DROP VIEW %s", v.Name)); err != nil {
			e.log.Error("failed to drop view", "policy", policy.Name, "view", v.Name, "err", err)
			continue
		}
	}
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
