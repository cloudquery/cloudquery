package policy

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
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

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	// policyDirectory points to the local policy directory
	policyDirectory string

	// Instance of a database connection pool
	pool *pgxpool.Pool

	// Logger instance
	logger hclog.Logger
}

// Policy represents a single policy.
type Policy struct {
	// Organization is the policy org.
	Organization string

	// Repository is the policy repository.
	Repository string

	// RepositoryPath is the policy repository internal path.
	RepositoryPath string

	// Version is the git repository tag that should be used.
	Version string

	// SubPath is the policy sub-path.
	SubPath string

	// LocalPath is the policy local path.
	LocalPath string
}

// Manager is the interface that describes the interaction with the policy hub.
// Implemented by ManagerImpl.
type Manager interface {
	// ParsePolicyHubPath parses a given policy hub path and returns a Policy object.
	ParsePolicyHubPath(args []string, subPolicyPath string) (*Policy, error)

	// DownloadPolicy downloads the given policy.
	DownloadPolicy(ctx context.Context, p *Policy) error

	// RunPolicy runs the given policy.
	RunPolicy(ctx context.Context, execRequest *ExecuteRequest) (*ExecutionResult, error)
}

// NewManager returns a new manager instance.
func NewManager(policyDir string, pool *pgxpool.Pool, logger hclog.Logger) Manager {
	return &ManagerImpl{
		policyDirectory: policyDir,
		pool:            pool,
		logger:          logger,
	}
}

// ParsePolicyHubPath parses and validates the given arguments into the Policy struct.
// Given args should follow the following semantic structure:
// [(organization/)repository-name(@tag)] ([repository-path])
func (m *ManagerImpl) ParsePolicyHubPath(args []string, subPolicyPath string) (*Policy, error) {
	// Make sure the mandatory args are given
	if len(args) < 1 {
		return nil, fmt.Errorf("invalid policy path. Repository name is required but got %#v", args)
	}
	policy := &Policy{
		SubPath: subPolicyPath,
	}

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

	// Parse repository path if given
	if len(args) == 2 {
		policy.RepositoryPath = args[1]
	}
	return policy, nil
}

// DownloadPolicy downloads the given policy from GitHub and stores it in the local policy directory.
func (m *ManagerImpl) DownloadPolicy(ctx context.Context, p *Policy) error {
	// Make sure that the local policy organization folder exists
	osFs := file.NewOsFs()
	policyOrgFolder := filepath.Join(m.policyDirectory, p.Organization)
	if err := osFs.MkdirAll(policyOrgFolder, 0744); err != nil {
		return fmt.Errorf("failed to create organization policy directory: %s", policyOrgFolder)
	}

	// Get GitHub URL
	gitURL, err := p.getGitHubURL()
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

// RunPolicy runs the given policy.
func (m *ManagerImpl) RunPolicy(ctx context.Context, execReq *ExecuteRequest) (*ExecutionResult, error) {
	p := execReq.Policy
	var policyFilePath, policyFolder string
	osFs := file.NewOsFs()

	if p.LocalPath != "" {
		// Run local path policy
		// Make sure policy file exists
		for _, extensionName := range defaultSupportedPolicyExtensions {
			currPolicyFile := filepath.Join(p.LocalPath, fmt.Sprintf("%s.%s", defaultPolicyFileName, extensionName))
			policyFolder = p.LocalPath

			if strings.HasSuffix(p.LocalPath, ".hcl") || strings.HasSuffix(p.LocalPath, ".json") {
				currPolicyFile = p.LocalPath
				pathParts := strings.Split(p.LocalPath, "/")
				policyFolder = strings.Join(pathParts[:len(pathParts)-1], "/")
			}

			if _, err := osFs.Stat(currPolicyFile); err == nil {
				policyFilePath = currPolicyFile
				break
			}
		}
		if policyFilePath == "" {
			return nil, fmt.Errorf("failed to find policy file; not found in %s", p.LocalPath)
		}
	} else {
		// Check if given policy exists in our policy folder
		orgPolicyStr := filepath.Join(p.Organization, p.Repository)
		repoFolder := filepath.Join(m.policyDirectory, orgPolicyStr)
		if info, err := osFs.Stat(repoFolder); err != nil || !info.IsDir() {
			return nil, fmt.Errorf("could not find policy '%s' locally. Try to download the policy first", orgPolicyStr)
		}
		m.logger.Debug("found repo folder", "path", repoFolder)

		if !execReq.SkipVersioning {
			// Checkout policy repository tag
			if err := p.checkoutPolicyVersion(repoFolder); err != nil {
				return nil, fmt.Errorf("failed to checkout repository tag: %s", err.Error())
			}
		}

		// If repository path was specified, also check if that exists
		policyFolder := repoFolder
		if p.RepositoryPath != "" {
			policyFolder = filepath.Join(repoFolder, p.RepositoryPath)
			if info, err := osFs.Stat(policyFolder); err != nil || !info.IsDir() {
				return nil, fmt.Errorf("could not find policy '%s' in the folder '%s'. Try to download the policy first", orgPolicyStr, p.RepositoryPath)
			}
			m.logger.Debug("internal repo folder set", "path", policyFolder)
		}

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
	}

	policies, err := m.readPolicy(policyFilePath, policyFolder)
	if err != nil {
		return nil, err
	}
	m.logger.Debug("parsed policy file", "policies", policies)
	if policies == nil {
		return nil, nil
	}
	// Acquire connection from the connection pool
	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection from the connection pool: %s", err.Error())
	}
	defer conn.Release()

	// Traverse all policies recursively
	m.logger.Debug("starting to traverse policies")
	policyMap := make(map[string]*config.Policy)
	policyMap = m.traversePolicies(policies.Policies, "", policyMap)
	m.logger.Debug("finished traversing policies", "policyMap", policyMap)

	// Execute policies dependent on policy sub path
	executor := NewExecutor(conn, m.logger)
	var results *ExecutionResult
	switch p.SubPath {
	case "":
		m.logger.Debug("no policy sub path defined; executing all policies")
		for _, p := range policies.Policies {
			results, err = executor.ExecutePolicy(ctx, execReq, p)
			if err != nil {
				return nil, fmt.Errorf("failed to run policies: %s", err.Error())
			}
		}
	default:
		m.logger.Debug("policy sub path defined; only executing sub policy/query", "subpath", p.SubPath)
		results, err = m.runSubPolicyOrQuery(ctx, executor, policyMap, p.SubPath, execReq)
		if err != nil {
			return nil, fmt.Errorf("failed to run sub policy/query: %s", err.Error())
		}
	}
	return results, nil
}

// readPolicy reads, normalizes and validates the policy file at policyPath, using policyFolder as base path.
func (m *ManagerImpl) readPolicy(policyPath, policyFolder string) (*config.PolicyWrapper, error) {
	parser := config.NewParser()
	policiesRaw, diags := parser.LoadHCLFile(policyPath)
	if diags != nil && diags.HasErrors() {
		return nil, fmt.Errorf("failed to load policy file: %#v", diags.Error())
	}
	policies, diagsDecode := parser.DecodePolicies(policiesRaw, diags, policyFolder)
	if diagsDecode != nil && diagsDecode.HasErrors() {
		return nil, fmt.Errorf("failed to parse policy file: %#v", diagsDecode.Error())
	}
	return policies, nil
}

// runSubPolicyOrQuery
func (m *ManagerImpl) runSubPolicyOrQuery(
	ctx context.Context, exec *Executor,
	policyMap map[string]*config.Policy,
	subPath string,
	execReq *ExecuteRequest) (*ExecutionResult, error) {
	// If the given path points directly to a sub policy
	if policy, ok := policyMap[subPath]; ok {
		m.logger.Debug("running sub policy only", "policy", policy)

		// Collect all sub policies from this policy
		subPolicyMap := make(map[string]*config.Policy)
		for k, v := range policyMap {
			// Add policy if prefix is the same.
			// This will also add the actual policy since k == p.SubPath.
			if strings.HasPrefix(k, subPath) {
				subPolicyMap[k] = v
			}
		}
		if policy, ok := policyMap[subPath]; ok {
			m.logger.Debug("running sub policy only", "policy", policy)
			return exec.ExecutePolicy(ctx, execReq, policy)
		}
	}

	// Must be a query so get the policy path and the last element
	pathSplit := strings.Split(subPath, pathDelimiter)
	if len(pathSplit) <= 1 {
		// Sub path is malformed
		return nil, fmt.Errorf("malformed sub path: %s", subPath)
	}

	// Get the policy path and last element
	policyPath := pathSplit[len(pathSplit)-2]
	elementName := pathSplit[len(pathSplit)-1]

	// Get query parent policy
	parentPolicy, ok := policyMap[policyPath]
	if !ok {
		return nil, fmt.Errorf("cannot find sub query parent policy %s in %s", policyPath, subPath)
	}

	for _, query := range parentPolicy.Queries {
		if query.Name == elementName {
			// Execute query
			res, err := exec.ExecuteQuery(ctx, query)
			if err != nil {
				return nil, err
			}

			// Collect results
			execResults := &ExecutionResult{
				Passed:  true,
				Results: make(map[string]*QueryResult),
			}
			collectExecutionResults(execResults, policyPath, res)
			if execReq.UpdateCallback != nil {
				execReq.UpdateCallback(query.Name, res.Type, execResults.Passed)
			}
			return execResults, nil
		}
	}
	return nil, fmt.Errorf("failed to find sub policy/query: %s", subPath)
}

// traversePolicies is a recursive function that traverses p until all policies are resolved.
// All traversed policies gets stored into m where the policy level is used as a key.
func (m *ManagerImpl) traversePolicies(p []*config.Policy, levelPath string, policyMap map[string]*config.Policy) map[string]*config.Policy {
	for id, policy := range p {
		// Add current level to level path
		subLevelPath := policyPathJoin(levelPath, policy.Name)

		// Add policy to map
		policyMap[subLevelPath] = p[id]
		m.logger.Debug("added policy to policy map", "key", subLevelPath, "policy", p[id])

		// Check if this policy has sub policies
		if len(policy.Policies) > 0 {
			newM := m.traversePolicies(policy.Policies, subLevelPath, policyMap)

			// Merge maps
			for k, v := range newM {
				policyMap[k] = v
			}
		}
	}
	return policyMap
}

func (p *Policy) checkoutPolicyVersion(repoFolder string) error {
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
	checkoutVersion := p.Version
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

func (p *Policy) getGitHubURL() (string, error) {
	base, err := url.Parse(gitHubUrl)
	if err != nil {
		return "", err
	}
	org, err := url.Parse(p.Organization + "/")
	if err != nil {
		return "", err
	}
	repo, err := url.Parse(p.Repository + ".git")
	if err != nil {
		return "", err
	}
	return base.ResolveReference(org).ResolveReference(repo).String(), nil
}

// policyPathJoin joins policy path names with "/"
func policyPathJoin(paths ...string) string {
	return strings.Join(paths, "/")
}
