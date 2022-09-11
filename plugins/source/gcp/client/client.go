package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

type Client struct {
	// plugin   *plugins.SourcePlugin
	projects []string
	// All gcp services initialized by client
	Services *Services
	// this is set by table client multiplexer
	ProjectId string
	// Logger
	logger zerolog.Logger
}

const (
	serviceAccountEnvKey = "CQ_SERVICE_ACCOUNT_KEY_JSON"
)

//revive:disable:modifies-value-receiver

// withProject allows multiplexer to create a new client with given subscriptionId
func (c Client) withProject(project string) *Client {
	c.logger = c.logger.With().Str("project_id", project).Logger()
	c.ProjectId = project
	return &c
}

func isValidJson(content []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(content, &v)
	if err != nil {
		var syntaxError *json.SyntaxError
		if errors.As(err, &syntaxError) {
			return fmt.Errorf("the environment variable %s should contain valid JSON object. %w", serviceAccountEnvKey, err)
		}
		return err
	}
	return nil
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	c := Client{
		logger: logger,
		// plugin: p,
	}
	// providerConfig := config.(*Config)
	var gcpSpec Spec
	if err := s.UnmarshalSpec(&gcpSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	projects := gcpSpec.ProjectIDs

	serviceAccountKeyJSON := []byte(gcpSpec.ServiceAccountKeyJSON)
	if len(serviceAccountKeyJSON) == 0 {
		serviceAccountKeyJSON = []byte(os.Getenv(serviceAccountEnvKey))
	}

	// Add a fake request reason because it is not possible to pass nil options
	options := []option.ClientOption{option.WithRequestReason("cloudquery resource fetch")}
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, fmt.Errorf("invalid service account key JSON: %w", err)
		}
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	var err error
	c.Services, err = initServices(context.Background(), options)
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		c.logger.Info().Msg("No project_ids specified, assuming all active projects")
		var err error
		projects, err = getProjectsV1(ctx, options...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}
	}

	c.logger.Debug().Strs("projects", projects).Msg("Found projects")

	c.projects = projects
	if len(projects) == 1 {
		c.ProjectId = projects[0]
	}

	return &c, nil
}

// getProjectsV1 requires the `resourcemanager.projects.get` permission to list projects
func getProjectsV1(ctx context.Context, options ...option.ClientOption) ([]string, error) {
	var (
		projects []string
	)
	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	call := service.Projects.List().Filter("lifecycleState=ACTIVE").Context(ctx)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			projects = append(projects, project.ProjectId)
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("no active projects")
	}

	return projects, nil
}
