package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
)

//go:generate mockgen -package=mocks -destination=./mocks/mock_vercel_services.go . VercelServices
type VercelServices interface {
	ListDomains(ctx context.Context, pag *vercel.Paginator) ([]vercel.Domain, *vercel.Paginator, error)
	ListDomainRecords(ctx context.Context, domainName string, pag *vercel.Paginator) ([]vercel.DomainRecord, *vercel.Paginator, error)

	ListTeams(ctx context.Context, pag *vercel.Paginator) ([]vercel.Team, *vercel.Paginator, error)
	ListTeamMembers(ctx context.Context, teamID string, pag *vercel.Paginator) ([]vercel.TeamMember, *vercel.Paginator, error)

	ListProjects(ctx context.Context, pag *vercel.Paginator) ([]vercel.Project, *vercel.Paginator, error)
	ListProjectEnvs(ctx context.Context, projectId string, pag *vercel.Paginator) ([]vercel.ProjectEnv, *vercel.Paginator, error)

	ListDeployments(ctx context.Context, pag *vercel.Paginator) ([]vercel.Deployment, *vercel.Paginator, error)
	ListDeploymentChecks(ctx context.Context, deploymentId string, pag *vercel.Paginator) ([]vercel.DeploymentCheck, *vercel.Paginator, error)
}
