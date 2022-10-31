package client

import (
	"context"

	"github.com/google/go-github/v48/github"
)

type GithubServices struct {
	Actions       ActionsService
	Billing       BillingService
	Issues        IssuesService
	Organizations OrganizationsService
	Repositories  RepositoriesService
	Teams         TeamsService
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_teams.go . TeamsService
type TeamsService interface {
	GetTeamMembershipBySlug(ctx context.Context, org, slug, user string) (*github.Membership, *github.Response, error)
	ListExternalGroups(ctx context.Context, org string, opts *github.ListExternalGroupsOptions) (*github.ExternalGroupList, *github.Response, error)
	ListTeamMembersByID(ctx context.Context, orgID, teamID int64, opts *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error)
	ListTeamReposByID(ctx context.Context, orgID, teamID int64, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
	ListTeams(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_billing.go . BillingService
type BillingService interface {
	GetActionsBillingOrg(ctx context.Context, org string) (*github.ActionBilling, *github.Response, error)
	GetPackagesBillingOrg(ctx context.Context, org string) (*github.PackageBilling, *github.Response, error)
	GetStorageBillingOrg(ctx context.Context, org string) (*github.StorageBilling, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_repositories.go . RepositoriesService
type RepositoriesService interface {
	GetContents(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error)
	ListByOrg(ctx context.Context, org string, opts *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_orgs.go . OrganizationsService
type OrganizationsService interface {
	Get(ctx context.Context, org string) (*github.Organization, *github.Response, error)
	GetOrgMembership(ctx context.Context, user, org string) (*github.Membership, *github.Response, error)
	ListHookDeliveries(ctx context.Context, org string, id int64, opts *github.ListCursorOptions) ([]*github.HookDelivery, *github.Response, error)
	ListHooks(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error)
	ListInstallations(ctx context.Context, org string, opts *github.ListOptions) (*github.OrganizationInstallations, *github.Response, error)
	ListMembers(ctx context.Context, org string, opts *github.ListMembersOptions) ([]*github.User, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_issues.go . IssuesService
type IssuesService interface {
	ListByOrg(ctx context.Context, org string, opts *github.IssueListOptions) ([]*github.Issue, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_actions.go . ActionsService
type ActionsService interface {
	ListWorkflows(ctx context.Context, owner, repo string, opts *github.ListOptions) (*github.Workflows, *github.Response, error)
}
