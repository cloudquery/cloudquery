package client

import (
	"context"

	"github.com/google/go-github/v45/github"
)

type GithubServices struct {
	Teams         TeamsService
	Billing       BillingService
	Repositories  RepositoriesService
	Organizations OrganizationsService
	Issues        IssuesService
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_teams.go . TeamsService
type TeamsService interface {
	ListTeamReposByID(ctx context.Context, orgID, teamID int64, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
	ListTeamMembersByID(ctx context.Context, orgID, teamID int64, opts *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error)
	ListTeams(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Team, *github.Response, error)
	ListExternalGroups(ctx context.Context, org string, opts *github.ListExternalGroupsOptions) (*github.ExternalGroupList, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_billing.go . BillingService
type BillingService interface {
	GetStorageBillingOrg(ctx context.Context, org string) (*github.StorageBilling, *github.Response, error)
	GetPackagesBillingOrg(ctx context.Context, org string) (*github.PackageBilling, *github.Response, error)
	GetActionsBillingOrg(ctx context.Context, org string) (*github.ActionBilling, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_repositories.go . RepositoriesService
type RepositoriesService interface {
	ListByOrg(ctx context.Context, org string, opts *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_orgs.go . OrganizationsService
type OrganizationsService interface {
	Get(ctx context.Context, org string) (*github.Organization, *github.Response, error)
	ListInstallations(ctx context.Context, org string, opts *github.ListOptions) (*github.OrganizationInstallations, *github.Response, error)
	ListHooks(ctx context.Context, org string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error)
	ListHookDeliveries(ctx context.Context, org string, id int64, opts *github.ListCursorOptions) ([]*github.HookDelivery, *github.Response, error)
	ListMembers(ctx context.Context, org string, opts *github.ListMembersOptions) ([]*github.User, *github.Response, error)
	GetOrgMembership(ctx context.Context, user, org string) (*github.Membership, *github.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_issues.go . IssuesService
type IssuesService interface {
	ListByOrg(ctx context.Context, org string, opts *github.IssueListOptions) ([]*github.Issue, *github.Response, error)
}
