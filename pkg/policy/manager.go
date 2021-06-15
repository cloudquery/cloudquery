package policy

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/go-version"

	"github.com/cloudquery/cloudquery/pkg/ui"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

const (
	pathDelimiter    = "/"
	versionDelimiter = "@"

	cloudQueryOrg = "cloudquery"
	gitHubUrl     = "https://github.com/"

	defaultLocalSubPath   = ".cq/policy/"
	defaultPolicyFileName = "policy"
)

var defaultSupportedPolicyExtensions = []string{"hcl", "json"}

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	// Pointer to the client config
	config *config.Config

	// Instance of a database connection pool
	pool *pgxpool.Pool
}

var (
	managerInstance *ManagerImpl
	once            sync.Once
)

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
}

// Manager is the interface that describes the interaction with the policy hub.
// Implemented by ManagerImpl.
type Manager interface {
	// ParsePolicyHubPath parses a given policy hub path and returns a Policy object.
	ParsePolicyHubPath(args []string, subPolicyPath string) (*Policy, error)

	// DownloadPolicy downloads the given policy.
	DownloadPolicy(ctx context.Context, p *Policy) error

	// RunPolicy runs the given policy.
	RunPolicy(ctx context.Context, p *Policy) (*ExecutionResult, error)
}

// NewManager returns the manager instance.
func NewManager(c *config.Config, pool *pgxpool.Pool) Manager {
	// Singleton instantiation
	once.Do(func() {
		managerInstance = &ManagerImpl{
			config: c,
			pool:   pool,
		}
	})
	return managerInstance
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
	policyOrgFolder := filepath.Join(m.config.CloudQuery.PolicyDirectory, defaultLocalSubPath, p.Organization)
	if err := osFs.MkdirAll(policyOrgFolder, 0744); err != nil {
		return fmt.Errorf("failed to create organization policy directory: %s", policyOrgFolder)
	}

	// Get GitHub URL
	gitURL, err := p.getGitHubURL()
	if err != nil {
		return fmt.Errorf("failed to parse GitHub URL: %s", err.Error())
	}

	// Define clone options (start with main branch)
	cloneOptions := &git.CloneOptions{
		URL:           gitURL,
		ReferenceName: plumbing.NewBranchReferenceName("main"),
		Depth:         1,
		SingleBranch:  true,
		Tags:          git.AllTags,
	}

	// Output progress information if necessary
	if ui.IsTerminal() {
		// Print initial information
		switch {
		case p.Version != "":
			ui.ColorizedOutput(ui.ColorProgress, fmt.Sprintf("Cloning Policy %s/%s@%s\n", p.Organization, p.Repository, p.Version))
		default:
			ui.ColorizedOutput(ui.ColorProgress, fmt.Sprintf("Cloning Policy %s/%s\n", p.Organization, p.Repository))
		}

		// Set output to stdout
		cloneOptions.Progress = os.Stdout
	}

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
	case git.ErrBranchNotFound:
		cloneOptions.ReferenceName = plumbing.NewBranchReferenceName("master")
		_, err = git.PlainCloneContext(ctx, repoPath, false, cloneOptions)
		if err != nil && err != git.ErrRepositoryAlreadyExists {
			return fmt.Errorf("failed to clone repository: %s", err.Error())
		}
	default:
		return fmt.Errorf("failed to clone repository: %s", err.Error())
	}

	return nil
}

// RunPolicy runs the given policy.
func (m *ManagerImpl) RunPolicy(ctx context.Context, p *Policy) (*ExecutionResult, error) {
	// Check if given policy exists in our policy folder
	osFs := file.NewOsFs()
	orgPolicyStr := filepath.Join(p.Organization, p.Repository)
	repoFolder := filepath.Join(m.config.CloudQuery.PolicyDirectory, defaultLocalSubPath, orgPolicyStr)
	if info, err := osFs.Stat(repoFolder); err != nil || !info.IsDir() {
		return nil, fmt.Errorf("could not find policy '%s' locally. Try to download the policy first", orgPolicyStr)
	}

	// If repository path was specified, also check if that exists
	policyFolder := repoFolder
	if p.RepositoryPath != "" {
		policyFolder = filepath.Join(repoFolder, p.RepositoryPath)
		if info, err := osFs.Stat(policyFolder); err != nil || !info.IsDir() {
			return nil, fmt.Errorf("could not find policy '%s' in the folder '%s'. Try to download the policy first", orgPolicyStr, p.RepositoryPath)
		}
	}

	// Checkout policy repository tag
	if err := p.checkoutPolicyVersion(repoFolder); err != nil {
		return nil, fmt.Errorf("failed to checkout repository tag: %s", err.Error())
	}

	// Make sure policy file exists
	var policyFilePath string
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

	// Read policy file
	parser := config.NewParser(nil)
	policiesRaw, diags := parser.LoadHCLFile(policyFilePath)
	if diags != nil && diags.HasErrors() {
		return nil, fmt.Errorf("failed to load policy file: %#v", diags.Errs())
	}
	policies, diagsDecode := parser.DecodePolicies(policiesRaw, diags)
	if diagsDecode != nil && diagsDecode.HasErrors() {
		return nil, fmt.Errorf("failed to parse policy file: %#v", diagsDecode.Errs())
	}

	// Acquire connection from the connection pool
	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection from the connection pool: %s", err.Error())
	}
	defer conn.Release()

	// Prepare execution
	executor := NewExecutor(conn)
	execResults := &ExecutionResult{
		Passed:  true,
		Results: make(map[string]*QueryResult),
	}

	// Traverse all policies recursively
	policyMap := make(map[string]*config.Policy)
	policyMap = traversePolicies(policies.Policies, "", policyMap)

	// No sub path provided. Execute everything.
	if p.SubPath == "" {
		for _, policy := range policyMap {
			// Execute policy
			results, err := executor.ExecutePolicy(ctx, policy)
			if err != nil {
				return nil, err
			}

			// Collect results
			collectExecutionResults(execResults, results...)
		}
		return execResults, nil
	}

	// If we are here, the user only wants to execute a sub policy/view/query so we have to
	// find the corresponding element.

	// If the given path points directly to a sub policy
	if executePolicy, ok := policyMap[p.SubPath]; ok {
		// Execute policy
		results, err := executor.ExecutePolicy(ctx, executePolicy)
		if err != nil {
			return nil, err
		}

		// Make sure we also execute sub policies from this policy
		for _, subPolicy := range executePolicy.Policies {
			subPolicyResults, err := executor.ExecutePolicy(ctx, subPolicy)
			if err != nil {
				return nil, err
			}
			results = append(results, subPolicyResults...)
		}

		// Collect results
		collectExecutionResults(execResults, results...)
		return execResults, nil
	}

	// Must be a query so get the policy path and the last element
	pathSplit := strings.Split(p.SubPath, pathDelimiter)
	if len(pathSplit) <= 1 {
		// Sub path is malformed
		return nil, fmt.Errorf("malformed sub path: %s", p.SubPath)
	}

	// Get the policy path and last element
	policyPath := pathSplit[len(pathSplit)-2]
	elementName := pathSplit[len(pathSplit)-1]

	// Get query parent policy
	parentPolicy, ok := policyMap[policyPath]
	if !ok {
		return nil, fmt.Errorf("cannot find sub query parent policy %s in %s", policyPath, p.SubPath)
	}

	for _, query := range parentPolicy.Queries {
		if query.Name == elementName {
			// Execute query
			res, err := executor.ExecuteQuery(ctx, query)
			if err != nil {
				return nil, err
			}

			// Collect results
			collectExecutionResults(execResults, res)
			return execResults, nil
		}
	}

	return nil, fmt.Errorf("cannot find sub query %s in %s", elementName, p.SubPath)
}

// traversePolicies is a recursive function that traverses p until all policies are resolved.
// All traversed policies gets stored into m where the policy level is used as a key.
func traversePolicies(p []*config.Policy, levelPath string, m map[string]*config.Policy) map[string]*config.Policy {
	for id, policy := range p {
		// Add current level to level path
		levelPath = strings.ToLower(filepath.Join(levelPath, policy.Name))

		// Add policy to map
		m[levelPath] = p[id]

		// Check if this policy has sub policies
		if len(policy.Policies) > 0 {
			newM := traversePolicies(policy.Policies, levelPath, m)

			// Merge maps
			for k, v := range newM {
				m[k] = v
			}
		}
	}
	return m
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

func (p *Policy) checkoutPolicyVersion(repoFolder string) error {
	// Open git repo folder
	r, err := git.PlainOpen(repoFolder)
	if err != nil {
		return fmt.Errorf("failed to open policy repository folder: %s", err.Error())
	}

	// Make sure we have the correct version checked out before we proceed.
	// NOTE: This is not "Thread-Safe" e.g. two threads or processes could interfere here and the output
	// could be unpredictable. A better solution would be to create a local lock file that prevents other
	// threads to execute at the same time.
	checkoutVersion := p.Version
	if checkoutVersion == "" {
		// Create a new map that stores the version->tag reference
		versionTagMap := make(map[*version.Version]string)

		// List all annotated tags
		tagRefs, err := r.Tags()
		if err != nil {
			return fmt.Errorf("failed to list annotated repository tags: %s", err.Error())
		}
		_ = tagRefs.ForEach(func(reference *plumbing.Reference) error {
			// Try to convert tag to a version
			v, err := version.NewSemver(reference.Name().String())
			if err != nil {
				// Ignore this tag if it is not a valid version
				return nil
			}

			// Add to our data structure
			versionTagMap[v] = reference.Name().String()
			return nil
		})

		// List all lightweight tags
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
			versionTagMap[v] = tag.Name
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
			checkoutVersion = versionTagMap[sortedVersions[0]]
		}
	}

	// Get the tag reference
	ref, err := r.Tag(checkoutVersion)
	if err != nil {
		return fmt.Errorf("failed to find provided tag (%s): %s", checkoutVersion, err.Error())
	}

	// Get working tree
	workTree, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get work tree: %s", err.Error())
	}

	// Checkout given tag
	if err := workTree.Checkout(&git.CheckoutOptions{
		Hash:   ref.Hash(),
		Create: false,
		Force:  true,
		Keep:   false,
	}); err != nil {
		return fmt.Errorf("failed to checkout tag (%s): %s", checkoutVersion, err.Error())
	}

	return nil
}

// collectExecutionResults collects all query results and adds them to the
// execution results struct.
func collectExecutionResults(execResult *ExecutionResult, results ...*QueryResult) {
	for _, res := range results {
		if !res.Passed {
			execResult.Passed = false
		}
		execResult.Results[res.Name] = res
	}
}
