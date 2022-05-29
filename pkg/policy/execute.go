package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/cloudquery/cloudquery/internal"
	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/core/state"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

var ErrPolicyOrQueryNotFound = errors.New("selected policy/query not found")

const (
	statusError  = "error"
	statusFailed = "failed"
	statusPassed = "passed"
)

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

// Executor implements the execution framework.
type Executor struct {
	// Connection to the database
	conn         LowLevelQueryExecer
	stateManager *state.Client
	log          hclog.Logger

	PolicyPath []string

	// progressUpdate
	progressUpdate UpdateCallback
}

// QueryResult contains the result information from an executed query.
type QueryResult struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	QueryColumns []string  `json:"-"`
	Columns      []string  `json:"result_header"`
	Rows         Rows      `json:"result_rows"`
	Type         QueryType `json:"type"`
	Passed       bool      `json:"check_passed"`
}

type Row struct {
	// AdditionalData is any extra information that was returned from the result set
	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
	// Identifiers is a map of identifiers as defined by the policy
	Identifiers map[string]interface{} `json:"identifiers,omitempty"`
	// Reason is a user readable explanation returned by the query, or interpolated from check defined reason.
	Reason string `json:"reason,omitempty"`
	// Status is user defined status of the row i.e OK, ALERT etc'
	Status string `json:"status,omitempty"`
}
type Rows []Row

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
}

// ExecuteRequest is a request that triggers policy execution.
type ExecuteRequest struct {
	// Policy is the policy that should be executed.
	Policy *Policy
	// StopOnFailure if true policy execution will stop on first failure
	StopOnFailure bool
	// UpdateCallback is the console ui update callback
	UpdateCallback UpdateCallback
	// PolicyExecution represents the current policy execution
	PolicyExecution *state.PolicyExecution
	// DBPersistence defines weather or not to store run results
	DBPersistence bool
}

func (f Update) AllDone() bool {
	return f.FinishedQueries == f.QueriesCount
}

func (f Update) DoneCount() int {
	return f.FinishedQueries
}

func (r Rows) Len() int {
	return len(r)
}

func (r Rows) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Rows) Less(i, j int) bool {
	r1 := r[i]
	r2 := r[j]
	v1, v2 := make([]string, 0, len(r1.Identifiers)), make([]string, 0, len(r2.Identifiers))
	for _, v := range r1.Identifiers {
		v1 = append(v1, cast.ToString(v))
	}
	for _, v := range r2.Identifiers {
		v2 = append(v2, cast.ToString(v))
	}
	for l := 0; l < len(v1); l++ {
		if v1[l] < v2[l] {
			return true
		}
	}
	return false
}

// NewExecutor creates a new executor.
func NewExecutor(conn LowLevelQueryExecer, sta *state.Client, progressUpdate UpdateCallback) *Executor {
	return &Executor{
		conn:           conn,
		stateManager:   sta,
		log:            logging.NewZHcLog(&log.Logger, "policy"),
		progressUpdate: progressUpdate,
		PolicyPath:     []string{},
	}
}

func (e *Executor) with(policy string, args ...interface{}) *Executor {
	policyPath := e.PolicyPath
	policyPath = append(policyPath, policy)
	return &Executor{
		conn:           e.conn,
		stateManager:   e.stateManager,
		log:            e.log.With("policy", strings.Join(policyPath, "/")).With(args...),
		progressUpdate: e.progressUpdate,
		PolicyPath:     policyPath,
	}
}

// Execute executes given policy and the related sub queries/views.
func (e *Executor) Execute(ctx context.Context, req *ExecuteRequest, policy *Policy, identifiers []string) (*ExecutionResult, diag.Diagnostics) {
	total := ExecutionResult{PolicyName: req.Policy.String(), Passed: true, Results: make([]*QueryResult, 0), ExecutionTime: time.Now()}

	if !policy.HasChecks() {
		e.log.Warn("no checks or policies to execute")
		return &total, nil
	}

	if !viper.GetBool("disable-fetch-check") {
		if err := e.checkFetches(ctx, policy.Config); err != nil {
			return nil, diag.FromError(err, diag.USER, diag.WithDetails("%s: please run `cloudquery fetch` before running policy", policy.Name))
		}
	}
	if len(policy.Identifiers) > 0 {
		identifiers = policy.Identifiers
	}

	for _, p := range policy.Policies {
		executor := e.with(p.Name)
		executor.log.Info("starting policy execution")
		r, err := executor.Execute(ctx, req, p, identifiers)
		if err != nil {
			executor.log.Error("failed to execute policy", "err", err)
			return nil, diag.FromError(fmt.Errorf("%s/%w", policy.Name, err), diag.DATABASE)
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
		return nil, diag.FromError(fmt.Errorf("%s/%w", policy.Name, err), diag.DATABASE)
	}
	defer e.deleteViews(ctx, policy)

	for _, q := range policy.Checks {
		e.log = e.log.With("query", q.Name)
		qr, err := e.executeQuery(ctx, q, identifiers)
		if req.DBPersistence {
			if errStore := e.createCheckResult(ctx, req.PolicyExecution, q, qr, err); errStore != nil {
				e.log.Error("failed to create check result", "err", errStore)
			}
		}
		if err != nil {
			e.log.Error("failed to execute query", "err", err)
			return nil, diag.FromError(fmt.Errorf("%s/%w", policy.Name, err), diag.DATABASE)
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

func (e *Executor) createCheckResult(ctx context.Context, policyExecution *state.PolicyExecution, q *Check, qr *QueryResult, err error) error {
	if policyExecution == nil {
		return nil
	}
	checkResult := &state.CheckResult{
		ExecutionId:        policyExecution.Id,
		ExecutionTimestamp: policyExecution.Timestamp,
		Name:               q.Name,
		Selector:           normalizeCheckSelector(policyExecution, e.PolicyPath, q.Name),
		Description:        q.Title,
		Status:             statusFailed,
	}
	if err != nil {
		checkResult.Status = statusError
		checkResult.Error = err.Error()
		return e.stateManager.CreateCheckResult(ctx, checkResult)
	}
	if qr.Passed {
		checkResult.Status = statusPassed
	}

	rows := make([]map[string]interface{}, len(qr.Rows))
	for i, r := range qr.Rows {
		m := make(map[string]interface{})
		if len(r.AdditionalData) > 0 {
			m["data"] = r.AdditionalData
		}
		m = internal.FlattenRow(m)
		if len(r.Identifiers) > 0 {
			m["cq_identifiers"] = internal.FlattenRow(r.Identifiers)
		}
		if r.Reason != "" {
			m["cq_reason"] = r.Reason
		}
		if r.Status != "" {
			m["cq_status"] = r.Status
		}
		rows[i] = m
	}
	var byt []byte
	if byt, err = json.Marshal(rows); err != nil {
		return err
	}
	checkResult.RawResults = string(byt)

	return e.stateManager.CreateCheckResult(ctx, checkResult)
}

// checkFetches checks if there are fetch reports in database that satisfy providers from policy
func (e *Executor) checkFetches(ctx context.Context, policyConfig *Configuration) error {
	if policyConfig == nil {
		return nil
	}
	for _, p := range policyConfig.Providers {
		c, err := version.NewConstraint(p.Version)
		if err != nil {
			return fmt.Errorf("failed to parse version constraint for provider %s: %w", p.Type, err)
		}
		fetchSummary, err := e.stateManager.GetFetchSummaryForProvider(ctx, p.Type)
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
		if !c.Check(v.Core()) {
			return fmt.Errorf("the latest fetch for provider %s does not satisfy version requirement %s", p.Type, c)
		}
	}
	return nil
}

// executeQuery executes the given query and returns the result.
func (e *Executor) executeQuery(ctx context.Context, q *Check, identifiers []string) (*QueryResult, error) {
	e.log.Trace("query", q.Query)
	data, err := e.conn.Query(ctx, q.Query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", q.Name, err)
	}

	result := &QueryResult{
		Name:         q.Name,
		Description:  q.Title,
		QueryColumns: make([]string, 0),
		Columns:      []string{"status"},
		Rows:         make([]Row, 0),
		Type:         q.Type,
	}
	for _, fd := range data.FieldDescriptions() {
		result.QueryColumns = append(result.QueryColumns, string(fd.Name))
	}

	var rtpl *template.Template
	if q.Reason != "" {
		rtpl, err = template.New("query").Parse(q.Reason)
		if err != nil {
			log.Warn().Err(err).Msg("failed to to parse reason template")
		}
	}

	if len(identifiers) > 0 {
		result.Columns = append(result.Columns, identifiers...)
	}
	result.Columns = append(result.Columns, "reason")
	result.Columns = append(result.Columns, funk.SubtractString(result.QueryColumns, append([]string{"cq_status", "cq_reason"}, identifiers...))...)

	for data.Next() {
		values, err := data.Values()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", q.Name, err)
		}
		row, err := parseRow(result.QueryColumns, values, identifiers, rtpl)
		if err != nil {
			log.Warn().Err(err).Msg("failed to create reason for check")
		}
		result.Rows = append(result.Rows, row)
	}
	if data.Err() != nil {
		return nil, fmt.Errorf("%s: %w", q.Name, data.Err())
	}
	result.Passed = (len(result.Rows) == 0) == !q.ExpectOutput
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
			count++
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

	// result.PolicyName is the full selector for this policy run.
	// The name of the output file should just be the base policy.
	// e.g. for "aws//cis_v1.2.0", the output file should be "aws.json"
	basePolicyName, err := extractFirstPathComponent(result.PolicyName)
	if err != nil {
		return err
	}

	f, err := fs.Create(fmt.Sprintf("%s.json", filepath.Join(outputDir, basePolicyName)))
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

// extractFirstPathComponent extracts the first path component form a given string.
// e.g: "a/b/c" -> "a"
//      "a" -> "a"
//      "a//b" -> "a"
func extractFirstPathComponent(str string) (string, error) {
	regex := regexp.MustCompile(`^([^/]+)(?:/[^/]*)*`)

	matches := regex.FindSubmatch([]byte(str))

	if matches == nil {
		return "", fmt.Errorf("failed to extract first path component")
	}

	return string(matches[1]), nil
}

func parseRow(columns []string, values []interface{}, identifiers []string, reasonTpl *template.Template) (Row, error) {
	r := Row{
		AdditionalData: make(map[string]interface{}, len(values)),
		Identifiers:    make(map[string]interface{}, len(identifiers)),
		Reason:         "",
		Status:         statusFailed,
	}

	for i := 0; i < len(columns); i++ {
		switch {
		case columns[i] == "cq_reason":
			r.Reason = cast.ToString(values[i])
		case columns[i] == "cq_status":
			r.Status = cast.ToString(values[i])
		case funk.InStrings(identifiers, columns[i]):
			r.Identifiers[columns[i]] = values[i]
		default:
			r.AdditionalData[columns[i]] = values[i]
		}
	}

	if r.Reason == "" && reasonTpl != nil {
		var b strings.Builder
		if err := reasonTpl.Execute(&b, r.AdditionalData); err != nil {
			return r, err
		}
		r.Reason = b.String()
	}
	return r, nil
}

func normalizeCheckSelector(policyExecution *state.PolicyExecution, policyPath []string, checkName string) string {
	selector := []string{policyExecution.PolicyName}
	if !strings.HasPrefix(selector[0], policyExecution.Location+"//") {
		selector[0] = policyExecution.Location + "/"
	}
	selector = append(selector, policyPath...)
	return strings.Join(append(selector, checkName), "/")
}
