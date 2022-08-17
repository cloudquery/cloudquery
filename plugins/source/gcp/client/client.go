package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
)

type Client struct {
	plugin   *plugins.SourcePlugin
	projects []string
	backoff  BackoffSettings
	// All gcp services initialized by client
	Services *Services
	// this is set by table client multiplexer
	ProjectId string
	// Logger
	logger zerolog.Logger
}

const (
	defaultProjectIdName = "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"
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

func Configure(ctx context.Context, p *plugins.SourcePlugin, s specs.Source) (schema.ClientMeta, error) {
	c := Client{
		plugin: p,
	}
	// providerConfig := config.(*Config)
	var gcpSpec Spec
	if err := s.UnmarshalSpec(&gcpSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	projects := gcpSpec.ProjectIDs
	if gcpSpec.FolderMaxDepth == 0 {
		gcpSpec.FolderMaxDepth = 5
	}

	serviceAccountKeyJSON := []byte(gcpSpec.ServiceAccountKeyJSON)
	if len(serviceAccountKeyJSON) == 0 {
		serviceAccountKeyJSON = []byte(os.Getenv(serviceAccountEnvKey))
	}

	// Add a fake request reason because it is not possible to pass nil options
	options := append([]option.ClientOption{option.WithRequestReason("cloudquery resource fetch")}, gcpSpec.ClientOptions()...)
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, fmt.Errorf("invalid service account key JSON: %w", err)
		}
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	if len(gcpSpec.FolderIDs) > 0 {
		c.logger.Warn().Msg("ProjectFilter config option is deprecated and will not work with the folder_ids feature")
	}

	var err error
	c.Services, err = initServices(context.Background(), options)
	if err != nil {
		return nil, err
	}

	if len(gcpSpec.FolderIDs) > 0 {
		c.logger.Debug().Strs("folder_ids", gcpSpec.FolderIDs).Msg("Listing folders")

		var folderList []string
		for _, f := range gcpSpec.FolderIDs {
			folderAndChildren, err := listFolders(ctx, c.Services.ResourceManager.Folders, f, int(gcpSpec.FolderMaxDepth)-1)
			if err != nil {
				return nil, fmt.Errorf("failed to list folders: %w", err)
			}
			folderList = append(folderList, folderAndChildren...)
		}
		c.logger.Debug().Strs("folder_ids", folderList).Msg("Found folders")

		proj, err := getProjects(ctx, c.Services.ResourceManager, folderList)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}
		appendWithoutDupes(&projects, proj)
	}
	if len(projects) == 0 {
		c.logger.Info().Msg("No project_ids specified, assuming all active projects")
		var err error
		projects, err = getProjectsV1(options...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}
	}

	c.logger.Debug().Strs("projects", projects).Msg("Found projects")

	c.projects = projects
	c.backoff = gcpSpec.Backoff()
	if len(projects) == 1 {
		c.ProjectId = projects[0]
	}

	return &c, nil
}

// getProjects requires the `resourcemanager.projects.list` permission, and at least one folder
func getProjects(ctx context.Context, service *cloudresourcemanager.Service, folders []string) ([]string, error) {
	if len(folders) == 0 {
		return nil, fmt.Errorf("no folders specified")
	}

	var (
		projects []string
		inactive int
	)

	for _, folder := range folders {
		call := service.Projects.List().Context(ctx).Parent(folder)

		for {
			output, err := call.Do()
			if err != nil {
				return nil, fmt.Errorf("failed to list projects in folder %s: %w", folder, err)
			}
			for _, project := range output.Projects {
				if project.State == "ACTIVE" {
					projects = append(projects, project.ProjectId)
				} else {
					inactive++
				}
			}
			if output.NextPageToken == "" {
				break
			}
			call.PageToken(output.NextPageToken)
		}
	}

	if len(projects) == 0 {
		if inactive > 0 {
			return nil, fmt.Errorf("project listing failed: no active projects")
		}
		return nil, fmt.Errorf("project listing failed")
	}

	return projects, nil
}

// getProjectsV1 requires the `resourcemanager.projects.get` permission to list projects
func getProjectsV1(options ...option.ClientOption) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	var (
		projects []string
		inactive int
	)

	call := service.Projects.List().Context(ctx)

	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			if project.LifecycleState == "ACTIVE" {
				projects = append(projects, project.ProjectId)
			} else {
				inactive++
			}
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	if len(projects) == 0 {
		if inactive > 0 {
			return nil, fmt.Errorf("project listing failed: no active projects")
		}
		return nil, fmt.Errorf("project listing failed")
	}

	return projects, nil
}

func listFolders(ctx context.Context, service *cloudresourcemanager.FoldersService, parent string, maxDepth int) ([]string, error) {
	folders := []string{
		parent,
	}
	if maxDepth <= 0 {
		return folders, nil
	}

	call := service.List().Context(ctx).Parent(parent)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("failed to list folders: %w", err)
		}
		for _, folder := range output.Folders {
			if folder.State != "ACTIVE" {
				continue
			}
			fList, err := listFolders(ctx, service, folder.Name, maxDepth-1)
			if err != nil {
				return nil, fmt.Errorf("failed to list folders: %w", err)
			}
			folders = append(folders, fList...)
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	return folders, nil
}

func appendWithoutDupes(dst *[]string, src []string) {
	dstMap := make(map[string]struct{}, len(*dst))
	for i := range *dst {
		dstMap[(*dst)[i]] = struct{}{}
	}
	for i := range src {
		if _, ok := dstMap[src[i]]; ok {
			continue
		}
		dstMap[src[i]] = struct{}{}
		*dst = append(*dst, src[i])
	}
}
