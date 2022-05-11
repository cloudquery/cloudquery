package policy

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/cloudquery/cloudquery/internal/logging"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/cloudquery/cloudquery/pkg/core/database"

	"github.com/rs/zerolog/log"

	"github.com/cloudquery/cq-provider-sdk/provider/execution"
)

const (
	CloudQueryOrg = "cloudquery-policies"
)

type LowLevelQueryExecer interface {
	execution.Copier
	execution.QueryExecer
}

func Snapshot(ctx context.Context, storage database.Storage, policy *Policy, outputPath, subpath string) error {
	db, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "executor-database"), storage.DSN())
	if err != nil {
		return err
	}
	e := NewExecutor(db, nil)

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
func Load(ctx context.Context, directory string, policy *Policy) (*Policy, error) {
	var err error
	// if policy is configured with source we load it first
	if policy.Source != "" {
		log.Debug().Str("policy", policy.Name).Str("source", policy.Source).Msg("loading policy from source")
		policy, err = loadPolicyFromSource(ctx, directory, policy.Name, policy.SubPolicy(), policy.Source)
		if err != nil {
			return nil, err
		}
	}
	// TODO: add recursive stop
	// load inner policies
	for i, p := range policy.Policies {
		log.Debug().Str("policy", policy.Name).Str("inner_policy", p.Name).Msg("loading inner policy from source")
		policy.Policies[i], err = Load(ctx, directory, p)
		if err != nil {
			return nil, err
		}
	}
	return policy, nil
}

// RunRequest is the request used to run one or more policy.
type RunRequest struct {
	// Policies to run
	Policies Policies
	// Directory to load / save policies to
	Directory string
	// OutputDir is the output dir for policy execution output.
	OutputDir string
	// RunCallBack is the callback method that is called after every policy execution.
	RunCallback UpdateCallback
}

type RunResponse struct {
	Policies   Policies
	Executions []*ExecutionResult
}

func Run(ctx context.Context, storage database.Storage, req *RunRequest) (*RunResponse, diag.Diagnostics) {
	var (
		diags diag.Diagnostics
		resp  = &RunResponse{
			Policies:   make(Policies, 0),
			Executions: make([]*ExecutionResult, 0),
		}
	)
	for _, p := range req.Policies {
		log.Info().Str("policy", p.Name).Str("version", p.Version()).Str("subPath", p.SubPolicy()).Msg("preparing to run policy")
		loadedPolicy, err := Load(ctx, req.Directory, p)
		if err != nil {
			return nil, diag.FromError(err, diag.INTERNAL)
		}
		resp.Policies = append(resp.Policies, loadedPolicy)
		log.Debug().Str("policy", p.Name).Str("version", p.Version()).Str("subPath", p.SubPolicy()).Msg("loaded policy successfully")
		result, dd := run(ctx, storage, &ExecuteRequest{
			Policy:         loadedPolicy,
			UpdateCallback: req.RunCallback,
		})
		diags = diags.Add(dd)
		if diags.HasErrors() {
			// this error means error in execution and not policy violation
			// we should exit immediately as this is a non-recoverable error
			// might mean schema is incorrect, provider version
			log.Error().Err(err).Msg("policy execution finished with error")
			return resp, diags
		}
		log.Info().Str("policy", p.Name).Msg("policy execution finished")
		resp.Executions = append(resp.Executions, result)
		if req.OutputDir == "" {
			continue
		}
		log.Info().Str("policy", p.Name).Str("version", p.Version()).Str("subPath", p.SubPolicy()).Msg("writing policy to output directory")
		if err := GenerateExecutionResultFile(result, req.OutputDir); err != nil {
			return nil, diags.Add(diag.FromError(err, diag.INTERNAL))
		}
	}
	return resp, diags
}

func run(ctx context.Context, storage database.Storage, request *ExecuteRequest) (*ExecutionResult, diag.Diagnostics) {
	var (
		totalQueriesToRun = request.Policy.TotalQueries()
		finishedQueries   = 0
	)
	filteredPolicy := request.Policy.Filter(request.Policy.meta.SubPolicy)
	if filteredPolicy.Config == nil {
		filteredPolicy.Config = request.Policy.Config
	}
	if !filteredPolicy.HasChecks() {
		log.Error().Str("selector", request.Policy.meta.SubPolicy).Strs("available_policies", filteredPolicy.Policies.All()).Msg("policy/query not found with provided sub-policy selector")
		return nil, diag.FromError(fmt.Errorf("%s//%s: %w", request.Policy.Name, request.Policy.meta.SubPolicy, ErrPolicyOrQueryNotFound),
			diag.USER, diag.WithDetails("%s//%s not found, run `cloudquery policy describe %s` to find all available policies", request.Policy.Name, request.Policy.meta.SubPolicy, request.Policy.Name))
	}
	totalQueriesToRun = filteredPolicy.TotalQueries()
	log.Info().Int("total", totalQueriesToRun).Msg("policy Checks count")
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
	db, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "executor-database"), storage.DSN())
	if err != nil {
		return nil, diag.FromError(err, diag.DATABASE)
	}
	// execute the queries
	return NewExecutor(db, progressUpdate).Execute(ctx, request, &filteredPolicy, nil)
}

func loadPolicyFromSource(ctx context.Context, directory, name, subPolicy, sourceURL string) (*Policy, error) {
	data, meta, err := LoadSource(ctx, directory, sourceURL)
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
		policy.meta.SubPolicy = subPolicy
	}
	policy.Source = sourceURL
	return policy, nil
}

func createSnapshotPath(directory, queryName string) (string, error) {
	path := strings.TrimSuffix(directory, "/")
	cleanedPath := filepath.Join(path, queryName, "tests", uuid.NewString())

	err := os.MkdirAll(cleanedPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return cleanedPath, nil
}
