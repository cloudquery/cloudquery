package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
)

const defaultProjectIdName = "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"

const serviceAccountEnvKey = "CQ_SERVICE_ACCOUNT_KEY_JSON"

type Client struct {
	projects []string
	logger   hclog.Logger
	backoff  BackoffSettings

	// All gcp services initialized by client
	Services *Services
	// this is set by table client multiplexer
	ProjectId string
}

func NewGcpClient(log hclog.Logger, bo BackoffSettings, projects []string, services *Services) *Client {
	return &Client{
		projects: projects,
		logger:   log,
		backoff:  bo,
		Services: services,
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// withProject allows multiplexer to create a new client with given subscriptionId
func (c Client) withProject(project string) *Client {
	return &Client{
		backoff:   c.backoff,
		projects:  c.projects,
		Services:  c.Services,
		logger:    c.logger.With("project_id", project),
		ProjectId: project,
	}
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

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	var diags diag.Diagnostics
	providerConfig := config.(*Config)
	projects := providerConfig.ProjectIDs
	if providerConfig.FolderMaxDepth == 0 {
		providerConfig.FolderMaxDepth = 5
	}

	serviceAccountKeyJSON := []byte(providerConfig.ServiceAccountKeyJSON)
	if len(serviceAccountKeyJSON) == 0 {
		serviceAccountKeyJSON = []byte(os.Getenv(serviceAccountEnvKey))
	}

	// Add a fake request reason because it is not possible to pass nil options
	options := append([]option.ClientOption{option.WithRequestReason("cloudquery resource fetch")}, providerConfig.ClientOptions()...)
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, diag.FromError(err, diag.USER)
		}
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	if providerConfig.ProjectFilter != "" && len(providerConfig.FolderIDs) > 0 {
		logger.Warn("ProjectFilter config option is deprecated and will not work with the folder_ids feature")
	}

	services, err := initServices(context.Background(), options)
	if err != nil {
		return nil, diags.Add(classifyError(err, diag.INTERNAL, projects))
	}

	if len(providerConfig.FolderIDs) > 0 {
		logger.Debug("Listing folders", "folder_ids", providerConfig.FolderIDs)

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		var folderList []string
		for _, f := range providerConfig.FolderIDs {
			folderAndChildren, err := listFolders(ctx, logger, services.ResourceManager.Folders, f, int(providerConfig.FolderMaxDepth)-1)
			if err != nil {
				return nil, diags.Add(classifyError(fmt.Errorf("folder listing failed: %w", err), diag.INTERNAL, projects))
			}
			folderList = append(folderList, folderAndChildren...)
		}
		logger.Debug("Found folders", "folder_ids", folderList)

		proj, err := getProjects(logger, services.ResourceManager, folderList)
		if err != nil {
			return nil, diags.Add(classifyError(fmt.Errorf("get projects failed: %w", err), diag.INTERNAL, projects))
		}
		appendWithoutDupes(&projects, proj)
	}
	if len(projects) == 0 {
		logger.Info("No project_ids specified, assuming all active projects")
		var err error
		projects, err = getProjectsV1(logger, providerConfig.ProjectFilter, options...)
		if err != nil {
			return nil, diags.Add(classifyError(fmt.Errorf("get projects(v1) failed: %w", err), diag.INTERNAL, projects))
		}
	}

	logger.Debug("Found projects", "projects", projects)

	diags = diags.Add(validateProjects(projects))
	if diags.HasErrors() {
		return nil, diags
	}

	client := NewGcpClient(logger, providerConfig.Backoff(), projects, services)
	return client, diags
}

func validateProjects(projects []string) diag.Diagnostics {
	for _, project := range projects {
		if project == defaultProjectIdName {
			return diag.FromError(errors.New("please specify a valid project_id in config.hcl instead of <CHANGE_THIS_TO_YOUR_PROJECT_ID>"), diag.USER)
		}
	}
	return nil
}

// getProjects requires the `resourcemanager.projects.list` permission, and at least one folder
func getProjects(logger hclog.Logger, service *cloudresourcemanager.Service, folders []string) ([]string, error) {
	if len(folders) == 0 {
		return nil, fmt.Errorf("no folders specified")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var (
		projects []string
		inactive int
	)

	for _, folder := range folders {
		call := service.Projects.List().Context(ctx).Parent(folder)

		for {
			output, err := call.Do()
			if err != nil {
				return nil, err
			}
			for _, project := range output.Projects {
				if project.State == "ACTIVE" {
					projects = append(projects, project.ProjectId)
				} else {
					logger.Info("Project state is not active. Project will be ignored", "project_id", project.ProjectId, "project_state", project.State)
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
func getProjectsV1(logger hclog.Logger, filter string, options ...option.ClientOption) ([]string, error) {
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
	if filter != "" {
		call.Filter(filter)
	}

	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			if project.LifecycleState == "ACTIVE" {
				projects = append(projects, project.ProjectId)
			} else {
				logger.Info("Project state is not active. Project will be ignored", "project_id", project.ProjectId, "project_state", project.LifecycleState)
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

func listFolders(ctx context.Context, logger hclog.Logger, service *cloudresourcemanager.FoldersService, parent string, maxDepth int) ([]string, error) {
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
			return nil, err
		}
		for _, folder := range output.Folders {
			if folder.State != "ACTIVE" {
				logger.Info("Folder state is not active. Folder will be ignored", "folder_id", folder.Name, "folder_name", folder.DisplayName, "folder_state", folder.State)
				continue
			}
			fList, err := listFolders(ctx, logger, service, folder.Name, maxDepth-1)
			if err != nil {
				return nil, err
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
