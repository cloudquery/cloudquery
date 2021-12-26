package policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	CloudQueryOrg = "cloudquery-policies"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	// policyDirectory points to the local policy directory
	policyDirectory string

	// Instance of a database connection pool
	pool *pgxpool.Pool

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
}

// NewManager returns a new manager instance.
func NewManager(policyDir string, pool *pgxpool.Pool, logger hclog.Logger) *ManagerImpl {
	return &ManagerImpl{
		policyDirectory: policyDir,
		pool:            pool,
		logger:          logger,
	}
}

func (m *ManagerImpl) Load(ctx context.Context, policy *Policy) (*Policy, error) {
	var err error
	// if policy is configured with source we load it first
	if policy.Source != "" {
		policy, err = m.loadPolicyFromSource(ctx, policy.Name, policy.Source)
		if err != nil {
			return nil, err
		}
	}
	// load inner policies
	for i, p := range policy.Policies {
		policy.Policies[i], err = m.Load(ctx, p)
		if err != nil {
			return nil, err
		}
	}
	return policy, nil
}

func (m *ManagerImpl) Run(ctx context.Context, request *ExecuteRequest) (*ExecutionResult, error) {
	// Acquire connection from the connection pool
	conn, err := m.pool.Acquire(ctx)
	m.logger.Trace("acquired connection from the connection pool", "err", err)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection from the connection pool: %w", err)
	}
	defer conn.Release()

	var (
		totalQueriesToRun = request.Policy.TotalQueries()
		finishedQueries   = 0
	)

	m.logger.Info("policy Checks count", "total", totalQueriesToRun)
	// set the progress total queries to run
	if request.UpdateCallback != nil {
		request.UpdateCallback(Update{
			PolicyName:      request.Policy.Name,
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
				Version:         request.Policy.meta.Version,
				FinishedQueries: finishedQueries,
				QueriesCount:    totalQueriesToRun,
				Error:           "",
			})
		}
	}

	var selector []string
	if request.Policy.meta.SubPath != "" {
		selector = strings.Split(request.Policy.meta.SubPath, "/")
	}

	// execute the queries
	return NewExecutor(conn, m.logger, progressUpdate).ExecutePolicies(ctx, request, Policies{request.Policy}, selector)
}

func (m *ManagerImpl) loadPolicyFromSource(ctx context.Context, name, sourceURL string) (*Policy, error) {
	data, meta, err := LoadSource(ctx, "./.cq/policies", sourceURL)
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
	policy.Source = sourceURL
	return policy, nil
}
