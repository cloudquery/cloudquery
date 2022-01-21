package policy

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	uuid "github.com/satori/go.uuid"

	"github.com/hashicorp/go-hclog"
)

const (
	CloudQueryOrg = "cloudquery-policies"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	// policyDirectory points to the local policy directory
	policyDirectory string

	// Instance of a database connection pool
	pool schema.QueryExecer

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
	Snapshot(ctx context.Context, policy *Policy) error
}

// NewManager returns a new manager instance.
func NewManager(policyDir string, pool schema.QueryExecer, logger hclog.Logger) *ManagerImpl {
	return &ManagerImpl{
		policyDirectory: policyDir,
		pool:            pool,
		logger:          logger,
	}
}

func createPath(path, queryName string) string {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	path = strings.TrimSuffix(path, "/")

	path = strings.TrimSuffix(path, "/query/"+queryName)

	u2 := uuid.NewV4()

	cleanedPath := filepath.Join("./database-data/", path, "/query"+"-"+queryName+"/", "tests", u2.String())

	// generate test name (uuid)
	err := os.MkdirAll(cleanedPath, os.ModePerm)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	return cleanedPath
}

func (m *ManagerImpl) Snapshot(ctx context.Context, policy *Policy) error {

	tableNames, _ := NewExecutor(m.pool, m.logger, nil).ExtractTableNames(ctx, policy.Checks[0].Query)

	snapShotPath := createPath("./", policy.Checks[0].Name)
	log.Println(snapShotPath)
	StoreSnapshot(snapShotPath, tableNames)
	return nil
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
	selector := strings.ReplaceAll(request.Policy.meta.subPolicy, "//", "/")
	filteredPolicy := request.Policy.Filter(selector)
	if len(filteredPolicy.Policies.All()) == 0 && len(filteredPolicy.Checks) == 0 {
		m.logger.Error("policy/query not found with provided sub-policy selector", "selector", selector, "available_policies", filteredPolicy.Policies.All())
		return nil, fmt.Errorf("%s//%s: %w", request.Policy.Name, selector, ErrPolicyOrQueryNotFound)
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
