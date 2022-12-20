package vercel

import (
	"context"
	"fmt"
)

const (
	projectsURL    = "/v9/projects"
	projectEnvsURL = "/v9/projects/%s/env"
)

type Project struct {
	AccountID                       string     `json:"accountId"`
	Analytics                       any        `json:"analytics,omitempty"`
	AutoExposeSystemEnvs            bool       `json:"autoExposeSystemEnvs,omitempty"`
	BuildCommand                    *string    `json:"buildCommand,omitempty"`
	CommandForIgnoringBuildStep     *string    `json:"commandForIgnoringBuildStep,omitempty"`
	CreatedAt                       MilliTime  `json:"createdAt"`
	DevCommand                      *string    `json:"devCommand,omitempty"`
	DirectoryListing                bool       `json:"directoryListing"`
	Env                             []any      `json:"env"`
	Framework                       *string    `json:"framework,omitempty"`
	GitForkProtection               *bool      `json:"gitForkProtection,omitempty"`
	ID                              string     `json:"id"`
	InstallCommand                  *string    `json:"installCommand,omitempty"`
	Name                            string     `json:"name"`
	NodeVersion                     string     `json:"nodeVersion"`
	OutputDirectory                 *string    `json:"outputDirectory,omitempty"`
	PasswordProtection              any        `json:"passwordProtection"`
	PublicSource                    *bool      `json:"publicSource,omitempty"`
	RootDirectory                   *string    `json:"rootDirectory,omitempty"`
	ServerlessFunctionRegion        *string    `json:"serverlessFunctionRegion,omitempty"`
	SourceFilesOutsideRootDirectory *bool      `json:"sourceFilesOutsideRootDirectory,omitempty"`
	UpdatedAt                       MilliTime  `json:"updatedAt"`
	Live                            *bool      `json:"live,omitempty"`
	Link                            any        `json:"link"`
	LatestDeployments               []any      `json:"latestDeployments"`
	Targets                         any        `json:"targets"`
	TransferStartedAt               *MilliTime `json:"transferStartedAt"`
	TransferCompletedAt             *MilliTime `json:"transferCompletedAt"`
	TransferredFromAccountID        *string    `json:"transferredFromAccountId"`
}

type ProjectEnv struct {
	ID                string     `json:"id"`
	Target            *string    `json:"target,omitempty"`
	Type              *string    `json:"type,omitempty"`
	Key               *string    `json:"key,omitempty"`
	Value             *string    `json:"value,omitempty"`
	ConfigurationId   *string    `json:"configurationId,omitempty"`
	GitBranch         *string    `json:"gitBranch,omitempty"`
	EdgeConfigId      *string    `json:"edgeConfigId,omitempty"`
	EdgeConfigTokenId *string    `json:"edgeConfigTokenId,omitempty"`
	Decrypted         *bool      `json:"decrypted,omitempty"`
	System            *bool      `json:"system,omitempty"`
	CreatedAt         MilliTime  `json:"createdAt"`
	CreatedBy         string     `json:"createdBy"`
	UpdatedAt         *MilliTime `json:"updatedAt,omitempty"`
	UpdatedBy         *string    `json:"updatedBy,omitempty"`
}

func (v *Client) ListProjects(ctx context.Context, pag *Paginator) ([]Project, *Paginator, error) {
	var list struct {
		Projects   []Project `json:"projects"`
		Pagination Paginator `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, projectsURL, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.Projects, &list.Pagination, nil
}

func (v *Client) ListProjectEnvs(ctx context.Context, projectId string, pag *Paginator) ([]ProjectEnv, *Paginator, error) {
	u := fmt.Sprintf(projectEnvsURL, projectId)

	var list struct {
		Envs       []ProjectEnv `json:"envs"`
		Pagination Paginator    `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, u, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.Envs, &list.Pagination, nil
}
