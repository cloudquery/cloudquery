package policy

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

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

	defaultLocalSubPath = ".cq/policy/"
)

// ManagerImpl is the manager implementation struct.
type ManagerImpl struct {
	// Pointer to the client config
	config *config.Config
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
}

// NewManager returns the manager instance.
func NewManager(c *config.Config) Manager {
	// Singleton instantiation
	once.Do(func() {
		managerInstance = &ManagerImpl{
			config: c,
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
	r, err := git.PlainCloneContext(ctx, repoPath, false, cloneOptions)
	switch err {
	case nil:
	case git.ErrRepositoryAlreadyExists:
		r, err = git.PlainOpen(repoPath)
		if err != nil {
			return fmt.Errorf("failed to open repository: %s", err.Error())
		}
	case git.ErrBranchNotFound:
		cloneOptions.ReferenceName = plumbing.NewBranchReferenceName("master")
		r, err = git.PlainCloneContext(ctx, repoPath, false, cloneOptions)
		if err != nil && err != git.ErrRepositoryAlreadyExists {
			return fmt.Errorf("failed to clone repository: %s", err.Error())
		}
	default:
		return fmt.Errorf("failed to clone repository: %s", err.Error())
	}

	// Switch to version tag if provided
	if p.Version != "" {
		ref, err := r.Tag(p.Version)
		if err != nil {
			return fmt.Errorf("failed to find provided tag (%s): %s", p.Version, err.Error())
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
			return fmt.Errorf("failed to checkout tag (%s): %s", p.Version, err.Error())
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
