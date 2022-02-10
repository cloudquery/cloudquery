package client

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/api/option"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/cloudresourcemanager/v1"
)

const defaultProjectIdName = "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"

const serviceAccountEnvKey = "CQ_SERVICE_ACCOUNT_KEY_JSON"

type Client struct {
	projects []string
	logger   hclog.Logger
	// All gcp services initialized by client
	Services *Services
	// this is set by table client multiplexer
	ProjectId string
}

func NewGcpClient(log hclog.Logger, projects []string, services *Services) *Client {
	return &Client{
		logger:   log,
		projects: projects,
		Services: services,
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// withProject allows multiplexer to create a new client with given subscriptionId
func (c Client) withProject(project string) *Client {
	return &Client{
		projects:  c.projects,
		Services:  c.Services,
		logger:    c.logger.With("project_id", project),
		ProjectId: project,
	}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	projects := providerConfig.ProjectIDs
	serviceAccountKeyJSON := []byte(providerConfig.ServiceAccountKeyJSON)

	if len(serviceAccountKeyJSON) == 0 {
		serviceAccountKeyJSON = []byte(os.Getenv(serviceAccountEnvKey))
	}

	var err error
	if len(providerConfig.ProjectIDs) == 0 {
		projects, err = getProjects(serviceAccountKeyJSON, logger, providerConfig.ProjectFilter)
		if err != nil {
			return nil, err
		}
		logger.Info("No project_ids specified in config.yml assuming all active projects", "count", len(projects))
	}
	if err := validateProjects(projects); err != nil {
		return nil, err
	}
	services, err := initServices(context.Background(), serviceAccountKeyJSON)
	if err != nil {
		return nil, err
	}
	client := NewGcpClient(logger, projects, services)
	return client, nil
}

func validateProjects(projects []string) error {
	for _, project := range projects {
		if project == defaultProjectIdName {
			return fmt.Errorf("please specify a valid project_id in config.yml instead of <CHANGE_THIS_TO_YOUR_PROJECT_ID>")
		}
	}
	return nil
}

func getProjects(serviceAccountKeyJSON []byte, logger hclog.Logger, filter string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	// Add a fake request reason because it is not possible to pass nil options
	options := []option.ClientOption{option.WithRequestReason("cloudquery resource fetch")}
	if len(serviceAccountKeyJSON) != 0 {
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	service, err := cloudresourcemanager.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}

	call := service.Projects.List()
	if filter != "" {
		call.Filter(filter)
	}

	projects := make([]string, 0)
	inactiveProjects := 0
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			if project.LifecycleState == "ACTIVE" {
				projects = append(projects, project.ProjectId)
			} else {
				logger.Info("Project state is not active. Project will be ignored", "project_id", project.ProjectId)
				inactiveProjects++
			}
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	if len(projects) == 0 {
		if inactiveProjects > 0 {
			return nil, fmt.Errorf("project listing failed: no active projects")
		}
		return nil, fmt.Errorf("project listing failed")
	}

	return projects, nil
}
