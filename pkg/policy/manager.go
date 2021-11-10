package policy

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"net/url"
	"path/filepath"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	pathDelimiter    = "/"
	versionDelimiter = "@"

	cloudQueryOrg = "cloudquery-policies"
	gitHubUrl     = "https://github.com/"

	defaultPolicyFileName = "policy"
)

var defaultSupportedPolicyExtensions = []string{"hcl", "json"}

type RemotePolicy struct {
	// SourceControl is the source control which the policy hosted on. Github / Gitlab
	SourceControl string

	// Organization is the organization / user which own the policy.
	Organization string

	// Repository is the policy repository name.
	Repository string

	// Repository is the policy repository version.
	Version string
}

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
	Run(ctx context.Context, req *ExecuteRequest, policyWrapper *PolicyWrapper) (*ExecutionResult, error)

	// Load the policy from local / remote location
	Load(ctx context.Context, p *config.Policy, execReq *ExecuteRequest) (*PolicyWrapper, error)

	// DownloadPolicy downloads the policy into the manager path.
	DownloadPolicy(ctx context.Context, p *RemotePolicy) error

	// ParsePolicySource transform config.Policy into RemotePolicy
	ParsePolicySource(policy *config.Policy) (*RemotePolicy, error)

	// ParsePolicyFromArgs transform cmd args into RemotePolicy
	ParsePolicyFromArgs(args []string) (*RemotePolicy, error)
}

// NewManager returns a new manager instance.
func NewManager(policyDir string, pool *pgxpool.Pool, logger hclog.Logger) Manager {
	return &ManagerImpl{
		policyDirectory: policyDir,
		pool:            pool,
		logger:          logger,
	}
}

func (m *ManagerImpl) Run(ctx context.Context, execReq *ExecuteRequest, policyWrapper *PolicyWrapper) (*ExecutionResult, error) {
	// Acquire connection from the connection pool
	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection from the connection pool: %s", err.Error())
	}
	defer conn.Release()

	var totalQueriesToRun = policyWrapper.getQueriesCount()
	var finishedQueries = 0

	// set the progress total queries to run
	execReq.UpdateCallback(Update{
		PolicyName:      execReq.Policy.Name,
		Version:         execReq.Policy.Version,
		FinishedQueries: 0,
		QueriesCount:    totalQueriesToRun,
		Error:           "",
	})

	// replace console update function to keep track the current status
	var progressUpdate = func(update Update) {
		finishedQueries += update.FinishedQueries
		execReq.UpdateCallback(Update{
			PolicyName:      execReq.Policy.Name,
			Version:         execReq.Policy.Version,
			FinishedQueries: finishedQueries,
			QueriesCount:    totalQueriesToRun,
			Error:           "",
		})
	}

	executor := NewExecutor(conn, m.logger)

	var selector []string
	if execReq.Policy.SubPath != "" {
		selector = strings.Split(execReq.Policy.SubPath, "/")
	}

	return executor.ExecutePolicies(ctx, progressUpdate, execReq, policyWrapper.Policies, selector)
}

// DownloadPolicy downloads the given policy from GitHub and stores it in the local policy directory.
func (m *ManagerImpl) DownloadPolicy(ctx context.Context, p *RemotePolicy) error {
	// Make sure that the local policy organization folder exists
	osFs := file.NewOsFs()
	policyOrgFolder := filepath.Join(m.policyDirectory, p.Organization)
	if err := osFs.MkdirAll(policyOrgFolder, 0744); err != nil {
		return fmt.Errorf("failed to create organization policy directory: %s", policyOrgFolder)
	}

	// Get GitHub URL
	gitURL, err := p.GetURL()
	if err != nil {
		return fmt.Errorf("failed to parse GitHub URL: %s", err.Error())
	}

	// Define clone options
	cloneOptions := &git.CloneOptions{
		URL:  gitURL,
		Tags: git.AllTags,
	}

	// Print initial information
	switch {
	case p.Version != "":
		ui.ColorizedOutput(ui.ColorProgress, "Cloning Policy %s/%s@%s\n", p.Organization, p.Repository, p.Version)
	default:
		ui.ColorizedOutput(ui.ColorProgress, "Cloning Policy %s/%s\n", p.Organization, p.Repository)
	}

	// Set output to stdout
	cloneOptions.Progress = os.Stdout

	// Clone the repository
	repoPath := filepath.Join(policyOrgFolder, p.Repository)
	_, err = git.PlainCloneContext(ctx, repoPath, false, cloneOptions)
	switch err {
	case nil:
	case git.ErrRepositoryAlreadyExists:
		_, err = git.PlainOpen(repoPath)
		if err != nil {
			return fmt.Errorf("failed to open repository: %s", err.Error())
		}
	default:
		return fmt.Errorf("failed to clone repository: %s", err.Error())
	}

	return nil
}

func (m *ManagerImpl) Load(ctx context.Context, p *config.Policy, execReq *ExecuteRequest) (*PolicyWrapper, error) {
	switch p.Type {
	case config.Hub:
		remotePolicy, err := m.ParsePolicySource(p)
		if err != nil {
			return nil, err
		}
		return m.loadRemotePolicy(ctx, remotePolicy, execReq)
	case config.Remote:
		remotePolicy, err := m.ParsePolicySource(p)
		if err != nil {
			return nil, err
		}
		return m.loadRemotePolicy(ctx, remotePolicy, execReq)
	case config.Local:
		return m.loadLocalPolicy(p)
	case config.Inline:
		return m.loadInlinePolicy(p)
	default:
		return nil, fmt.Errorf(`policy type value of "%s" is not valid`, p.Type)
	}
}

func (m *ManagerImpl) loadLocalPolicy(cfg *config.Policy) (*PolicyWrapper, error) {

	osFs := file.NewOsFs()
	var policyFolder = filepath.Dir(cfg.Source)
	var policyFilePath = cfg.Source

	// check if abs path has provided if no, get the default policy file
	if !strings.HasSuffix(cfg.Source, ".hcl") && !strings.HasSuffix(cfg.Source, ".json") {
		policyFolder := cfg.Source
		if info, err := osFs.Stat(policyFolder); err != nil || !info.IsDir() {
			return nil, fmt.Errorf("could not find policy '%s' in the folder '%s'. Try to download the policy first", cfg.Name, cfg.Source)
		}
		m.logger.Debug("internal repo folder set", "path", policyFolder)

		// Make sure policy file exists
		for _, extensionName := range defaultSupportedPolicyExtensions {
			currPolicyFile := filepath.Join(policyFolder, fmt.Sprintf("%s.%s", defaultPolicyFileName, extensionName))
			if _, err := osFs.Stat(currPolicyFile); err == nil {
				policyFilePath = currPolicyFile
				break
			}
		}
	}

	return readPolicy(policyFilePath, policyFolder)
}

func (m *ManagerImpl) loadInlinePolicy(cfg *config.Policy) (*PolicyWrapper, error) {
	parser := config.NewParser()

	policiesRaw, diags := parser.LoadFromSource("policy.hcl", []byte(cfg.Source), config.SourceHCL)
	if diags != nil && diags.HasErrors() {
		return nil, fmt.Errorf("failed to load policy file: %#v", diags.Error())
	}
	return decodePolicy(policiesRaw, diags, "")
}

func (m *ManagerImpl) loadRemotePolicy(ctx context.Context, remotePolicy *RemotePolicy, execReq *ExecuteRequest) (*PolicyWrapper, error) {
	if err := m.DownloadPolicy(ctx, remotePolicy); err != nil {
		return nil, err
	}
	var policyFilePath, policyFolder string

	// Check if given policy exists in our policy folder
	orgPolicyStr := filepath.Join(remotePolicy.Organization, remotePolicy.Repository)
	repoFolder := filepath.Join(m.policyDirectory, orgPolicyStr)

	osFs := file.NewOsFs()
	if info, err := osFs.Stat(repoFolder); err != nil || !info.IsDir() {
		return nil, fmt.Errorf("could not find policy '%s' locally. Try to download the policy first", orgPolicyStr)
	}
	m.logger.Debug("found repo folder", "path", repoFolder)

	if !execReq.SkipVersioning {
		// Checkout policy repository tag
		if err := m.checkoutPolicyVersion(repoFolder, remotePolicy.Version); err != nil {
			return nil, fmt.Errorf("failed to checkout repository tag: %s", err.Error())
		}
	}

	// If repository path was specified, also check if that exists
	policyFolder = repoFolder
	// if remotePolicy.Repository != "" {
	//	policyFolder = filepath.Join(repoFolder, remotePolicy.Repository)
	//	if info, err := osFs.Stat(policyFolder); err != nil || !info.IsDir() {
	//		return nil, fmt.Errorf("could not find policy '%s' in the folder '%s'. Try to download the policy first", orgPolicyStr, p.RepositoryPath)
	//	}
	//	m.logger.Debug("internal repo folder set", "path", policyFolder)
	//}

	// Make sure policy file exists
	for _, extensionName := range defaultSupportedPolicyExtensions {
		currPolicyFile := filepath.Join(policyFolder, fmt.Sprintf("%s.%s", defaultPolicyFileName, extensionName))
		if _, err := osFs.Stat(currPolicyFile); err == nil {
			policyFilePath = currPolicyFile
			break
		}
	}
	if policyFilePath == "" {
		return nil, fmt.Errorf("failed to find policy file; policy.%#v not found in %s", defaultSupportedPolicyExtensions, policyFolder)
	}
	m.logger.Debug("policy file found", "path", policyFilePath)

	return readPolicy(policyFilePath, policyFolder)
}

func (m *ManagerImpl) checkoutPolicyVersion(repoFolder string, checkoutVersion string) error {
	// Open git repo folder
	r, err := git.PlainOpen(repoFolder)
	if err != nil {
		return fmt.Errorf("failed to open policy repository folder: %s", err.Error())
	}

	// Fetch first to make sure we're up-to-date
	if err := r.Fetch(&git.FetchOptions{
		Tags:  git.AllTags,
		Force: true,
	}); err != nil && err != git.NoErrAlreadyUpToDate {
		return fmt.Errorf("failed to fetch latest changes: %s", err.Error())
	}
	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get work tree: %s", err.Error())
	}

	// Make sure we have the correct version checked out before we proceed.
	// NOTE: This is not "Thread-Safe" e.g. two threads or processes could interfere here and the output
	// could be unpredictable. A better solution would be to create a local lock file that prevents other
	// threads to execute at the same time.
	if checkoutVersion == "" {
		// Create a new map that stores the version->tag reference
		versionTagMap := make(map[*version.Version]string)

		// List all lightweight tags
		tagRefs, err := r.Tags()
		if err != nil {
			return fmt.Errorf("failed to list annotated repository tags: %s", err.Error())
		}
		_ = tagRefs.ForEach(func(reference *plumbing.Reference) error {
			// Try to convert tag to a version
			v, err := version.NewSemver(reference.Name().Short())
			if err != nil {
				// Ignore this tag if it is not a valid version
				return nil
			}

			// Add to our data structure
			versionTagMap[v] = reference.Name().Short()
			return nil
		})

		// List all annotated tags
		tags, err := r.TagObjects()
		if err != nil {
			return fmt.Errorf("failed to list lightweight repository tags: %s", err.Error())
		}
		_ = tags.ForEach(func(tag *object.Tag) error {
			// Try to convert tag to a version
			v, err := version.NewSemver(tag.Name)
			if err != nil {
				// Ignore this tag if it is not a valid version
				return nil
			}

			// Add to our data structures
			versionTagMap[v] = tag.Hash.String()
			return nil
		})

		// Sort versions
		var sortedVersions version.Collection
		for v := range versionTagMap {
			sortedVersions = append(sortedVersions, v)
		}
		sort.Sort(sortedVersions)
		if len(sortedVersions) != 0 {
			// TODO: Find the latest version for the used provider
			checkoutVersion = versionTagMap[sortedVersions[len(sortedVersions)-1]]
		}
	}

	// Get the tag reference
	var tagHash string
	ref, err := r.Tag(checkoutVersion)
	if err != nil {
		// It could be an annotated tag
		objRef, err := r.TagObject(plumbing.NewHash(checkoutVersion))
		if err != nil {
			return fmt.Errorf("failed to find provided tag (%s): %s", checkoutVersion, err.Error())
		}
		commit, err := objRef.Commit()
		if err != nil {
			return fmt.Errorf("failed to find annotated tag associated commit: %s", err.Error())
		}
		tagHash = commit.Hash.String()
	} else {
		tagHash = ref.Hash().String()
	}

	// Checkout given tag
	if err := w.Checkout(&git.CheckoutOptions{
		Hash:   plumbing.NewHash(tagHash),
		Create: false,
		Force:  true,
		Keep:   false,
	}); err != nil {
		return fmt.Errorf("failed to checkout tag (%s): %s", checkoutVersion, err.Error())
	}

	return nil
}

func (m *ManagerImpl) ParsePolicySource(policy *config.Policy) (*RemotePolicy, error) {
	switch policy.Type {
	case config.Hub:
		// transform hub policy config into RemotePolicy
		return &RemotePolicy{
			SourceControl: gitHubUrl,
			Organization:  cloudQueryOrg,
			Repository:    policy.Source,
			Version:       policy.Version,
		}, nil
	case config.Remote:
		// transform remote policy config into RemotePolicy
		parts, _ := url.Parse(policy.Source)
		sourceControl := fmt.Sprintf("%s://%s/", parts.Scheme, parts.Host)
		if parts.User != nil {
			sourceControl = fmt.Sprintf("%s://%s@%s/", parts.Scheme, parts.User, parts.Host)
		}
		pathParts := strings.Split(parts.Path, "/")

		if len(pathParts) != 3 {
			return nil, fmt.Errorf("cloud not parse policy source url")
		}

		organization := pathParts[1]
		repository := strings.TrimSuffix(pathParts[2], ".git")

		return &RemotePolicy{
			SourceControl: sourceControl,
			Organization:  organization,
			Repository:    repository,
			Version:       policy.Version,
		}, nil
	default:
		return nil, fmt.Errorf("unknown policy type, %s", policy.Type)
	}
}

func (m *ManagerImpl) ParsePolicyFromArgs(args []string) (*RemotePolicy, error) {
	// Make sure the mandatory args are given
	if len(args) < 1 {
		return nil, fmt.Errorf("invalid policy path. Repository name is required but got %#v", args)
	}
	policy := &RemotePolicy{SourceControl: gitHubUrl}

	// Parse and validate org/repository
	orgRepoSplit := strings.Split(args[0], pathDelimiter)

	// Parse org/repo
	switch len(orgRepoSplit) {
	case 2:
		policy.Organization = orgRepoSplit[0]
		policy.Repository = orgRepoSplit[1]
	case 1:
		policy.Repository = orgRepoSplit[0]
		policy.Organization = cloudQueryOrg
	default:
		return nil, fmt.Errorf("invalid policy path. Repository name malformed: %s", args[0])
	}

	// Parse version
	versionSplit := strings.Split(policy.Repository, versionDelimiter)
	if len(versionSplit) == 2 {
		policy.Version = versionSplit[1]
		policy.Repository = versionSplit[0]
	}

	return policy, nil
}

func readPolicy(policyPath, policyFolder string) (*PolicyWrapper, error) {
	parser := config.NewParser()
	policiesRaw, diags := parser.LoadHCLFile(policyPath)
	if diags != nil && diags.HasErrors() {
		return nil, fmt.Errorf("failed to load policy file: %#v", diags.Error())
	}
	return decodePolicy(policiesRaw, diags, policyFolder)
}

// readPolicy reads, normalizes and validates the policy file at policyPath, using policyFolder as base path.
func decodePolicy(policiesRaw hcl.Body, diags hcl.Diagnostics, policyFolder string) (*PolicyWrapper, error) {
	policies, diagsDecode := DecodePolicies(policiesRaw, diags, policyFolder)
	if diagsDecode != nil && diagsDecode.HasErrors() {
		return nil, fmt.Errorf("failed to parse policy file: %#v", diagsDecode.Error())
	}
	return policies, nil
}

func (r *RemotePolicy) GetURL() (string, error) {
	base, err := url.Parse(r.SourceControl)
	if err != nil {
		return "", err
	}
	org, err := url.Parse(r.Organization + "/")
	if err != nil {
		return "", err
	}
	repo, err := url.Parse(r.Repository + ".git")
	if err != nil {
		return "", err
	}
	return base.ResolveReference(org).ResolveReference(repo).String(), nil
}

func (p *Policy) getQueriesCount() int {
	count := 0
	if len(p.Policies) > 0 {
		for _, inner := range p.Policies {
			count += inner.getQueriesCount()
		}
	}
	count += len(p.Queries)
	return count
}

func (p *PolicyWrapper) getQueriesCount() int {
	count := 0
	for _, p := range p.Policies {
		count += p.getQueriesCount()
	}
	return count
}
