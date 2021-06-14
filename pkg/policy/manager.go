package policy

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

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
	RunPolicy(ctx context.Context, p *Policy) error
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
func (m *ManagerImpl) RunPolicy(ctx context.Context, p *Policy) error {
	// Check if given policy exists in our policy folder
	osFs := file.NewOsFs()
	orgPolicyStr := filepath.Join(p.Organization, p.Repository)
	repoFolder := filepath.Join(m.config.CloudQuery.PolicyDirectory, defaultLocalSubPath, orgPolicyStr)
	if info, err := osFs.Stat(repoFolder); err != nil || !info.IsDir() {
		return fmt.Errorf("could not find policy '%s' locally. Try to download the policy first", orgPolicyStr)
	}

	// If repository path was specified, also check if that exists
	policyFolder := repoFolder
	if p.RepositoryPath != "" {
		policyFolder = filepath.Join(repoFolder, p.RepositoryPath)
		if info, err := osFs.Stat(policyFolder); err != nil || !info.IsDir() {
			return fmt.Errorf("could not find policy '%s' in the folder '%s'. Try to download the policy first", orgPolicyStr, p.RepositoryPath)
		}
	}

	// Checkout policy repository tag
	if err := p.checkoutPolicyVersion(repoFolder); err != nil {
		return fmt.Errorf("failed to checkout repository tag: %s", err.Error())
	}

	// Make sure policy file exists
	var policyFilePath string
	for _, extensionName := range defaultSupportedPolicyExtensions {
		currPolicyFile := filepath.Join(repoFolder, fmt.Sprintf("%s.%s", defaultPolicyFileName, extensionName))
		if _, err := osFs.Stat(currPolicyFile); err == nil {
			policyFilePath = currPolicyFile
			break
		}
	}
	if policyFilePath == "" {
		return fmt.Errorf("failed to find policy file; policy.%#v not found", defaultSupportedPolicyExtensions)
	}

	// Read policy file
	parser := config.NewParser(nil)
	policiesRaw, diags := parser.LoadHCLFile(policyFilePath)
	if diags != nil && diags.HasErrors() {
		return fmt.Errorf("failed to load policy file: %#v", diags.Errs())
	}
	policies, diagsDecode := parser.DecodePolicies(policiesRaw, diags)
	if diagsDecode != nil && diagsDecode.HasErrors() {
		return fmt.Errorf("failed to parse policy file: %#v", diagsDecode.Errs())
	}

	// Acquire connection from the connection pool
	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection from the connection pool: %s", err.Error())
	}
	defer conn.Release()

	// Iterate all policies
	var subLevelPath string
	for _, policy := range policies.Policies {
		// Set sub level path
		subLevelPath = filepath.Join(subLevelPath, policy.Name)

		// Check if the user wants to only execute a specific sub-policy/view/query
		switch {
		case p.SubPath != "" && !strings.HasPrefix(p.SubPath, subLevelPath){
			// Sub path was defined but this is not the right leaf
			continue
		}
		}
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
