package policy

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/hashicorp/go-hclog"
)

const (
	CloudQueryOrg = "cloudquery-policies"
)

type LowLevelQueryExecer interface {
	execution.Copier
	execution.QueryExecer
}

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	// policyDirectory points to the local policy directory
	policyDirectory string

	// Instance of a database connection pool
	pool LowLevelQueryExecer

	// Logger instance
	logger hclog.Logger
}

// Manager is the interface that describes the interaction with the policy hub.
// Implemented by ManagerImpl.
type Manager interface {
	// Run the given policy.
	Run(ctx context.Context, request *ExecuteRequest) (*ExecutionResult, error)

	// Load the policy
	Load(ctx context.Context, policy *Policy) (*Policy, error)

	// Take a Snapshot of a policy
	Snapshot(ctx context.Context, policy *Policy, destination, selector string) error
}

// NewManager returns a new manager instance.
func NewManager(policyDir string, pool LowLevelQueryExecer, logger hclog.Logger) *ManagerImpl {
	return &ManagerImpl{
		policyDirectory: policyDir,
		pool:            pool,
		logger:          logger,
	}
}

//
func createSnapshotPath(directory, queryName string) (string, error) {
	path := strings.TrimSuffix(directory, "/")
	cleanedPath := filepath.Join(path, queryName, "tests", uuid.NewString())

	err := os.MkdirAll(cleanedPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return cleanedPath, nil
}

func (m *ManagerImpl) Snapshot(ctx context.Context, policy *Policy, outputPath, subpath string) error {
	e := NewExecutor(m.pool, m.logger, nil)

	if err := e.createViews(ctx, policy); err != nil {
		return err
	}

	tableNames, err := e.extractTableNames(ctx, policy.Checks[0].Query)
	if err != nil {
		return err
	}
	snapShotPath, err := createSnapshotPath(outputPath, subpath)
	if err != nil {
		return err
	}
	err = StoreSnapshot(ctx, e, snapShotPath, tableNames)
	if err != nil {
		return err
	}

	return StoreOutput(ctx, e, policy, snapShotPath)
}
func (m *ManagerImpl) Load(ctx context.Context, policy *Policy) (*Policy, error) {
	var err error
	// if policy is configured with source we load it first
	if policy.Source != "" {
		m.logger.Debug("loading policy from source", "policy", policy.Name, "source", policy.Source)
		policy, err = m.loadPolicyFromSource(ctx, policy.Name, policy.SubPolicy(), policy.Source)
		if err != nil {
			return nil, err
		}
	}
	// TODO: add recursive stop
	// load inner policies
	for i, p := range policy.Policies {
		m.logger.Debug("loading inner policy from source", "policy", policy.Name, "inner_policy", policy.Name)
		policy.Policies[i], err = m.Load(ctx, p)
		if err != nil {
			return nil, err
		}
	}
	return policy, nil
}

func (m *ManagerImpl) Run(ctx context.Context, request *ExecuteRequest) (*ExecutionResult, error) {
	var (
		totalQueriesToRun = request.Policy.TotalQueries()
		finishedQueries   = 0
	)
	filteredPolicy := request.Policy.Filter(request.Policy.meta.subPolicy)
	if !filteredPolicy.HasChecks() {
		m.logger.Error("policy/query not found with provided sub-policy selector", "selector", request.Policy.meta.subPolicy, "available_policies", filteredPolicy.Policies.All())
		return nil, fmt.Errorf("%s//%s: %w", request.Policy.Name, request.Policy.meta.subPolicy, ErrPolicyOrQueryNotFound)
	}
	totalQueriesToRun = filteredPolicy.TotalQueries()
	m.logger.Info("policy Checks count", "total", totalQueriesToRun)
	// set the progress total queries to run
	if request.UpdateCallback != nil {
		request.UpdateCallback(Update{
			PolicyName:      request.Policy.Name,
			Source:          request.Policy.Source,
			Version:         request.Policy.meta.Version,
			FinishedQueries: 0,
			QueriesCount:    totalQueriesToRun,
			Error:           "",
		})
	}

	// replace console update function to keep track the current status
	var progressUpdate = func(update Update) {
		finishedQueries += update.FinishedQueries
		if request.UpdateCallback != nil {
			request.UpdateCallback(Update{
				PolicyName:      request.Policy.Name,
				Source:          request.Policy.Source,
				Version:         request.Policy.meta.Version,
				FinishedQueries: finishedQueries,
				QueriesCount:    totalQueriesToRun,
				Error:           "",
			})
		}
	}
	// execute the queries
	return NewExecutor(m.pool, m.logger, progressUpdate).Execute(ctx, request, &filteredPolicy)
}

func (m *ManagerImpl) loadPolicyFromSource(ctx context.Context, name, subPolicy, sourceURL string) (*Policy, error) {
	data, meta, err := LoadSource(ctx, m.policyDirectory, sourceURL)
	if err != nil {
		return nil, err
	}
	f, dd := hclsyntax.ParseConfig(data, name, hcl.Pos{Byte: 0, Line: 1, Column: 1})
	if dd.HasErrors() {
		return nil, dd
	}
	policy, dd := DecodePolicy(f.Body, nil, meta.Directory)
	if dd.HasErrors() {
		return nil, dd
	}
	policy.meta = meta
	if subPolicy != "" {
		policy.meta.subPolicy = subPolicy
	}
	policy.Source = sourceURL
	return policy, nil
}
