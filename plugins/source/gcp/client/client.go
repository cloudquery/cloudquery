package client

import (
	"context"
	"encoding/json"
	"fmt"

	resourcemanagerv3 "cloud.google.com/go/resourcemanager/apiv3"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	resourcemanagerv3pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
)

const maxProjectIdsToLog int = 100

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
		return err
	}
	return nil
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var err error

	c := Client{
		logger: logger,
		// plugin: p,
	}
	// providerConfig := config.(*Config)
	var gcpSpec Spec
	if err := s.UnmarshalSpec(&gcpSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	gcpSpec.setDefaults()

	projects := gcpSpec.ProjectIDs

	serviceAccountKeyJSON := []byte(gcpSpec.ServiceAccountKeyJSON)

	// Add a fake request reason because it is not possible to pass nil options
	options := []option.ClientOption{option.WithRequestReason("cloudquery resource fetch")}
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, fmt.Errorf("invalid json at service_account_key_json: %w", err)
		}
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	c.Services, err = initServices(context.Background(), options)
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 && len(gcpSpec.FolderIDs) == 0 {
		c.logger.Info().Msg("No project_ids or folder_ids specified, assuming all active projects")
		projects, err = getProjectsV1(ctx, options...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}
	} else {
		folderIds := []string{}

		for _, parentFolder := range gcpSpec.FolderIDs {
			c.logger.Info().Msg("Listing folders..")
			childFolders, err := listFolders(ctx, c.Services.ResourcemanagerFoldersClient, parentFolder, *gcpSpec.FolderRecursionDepth)
			if err != nil {
				return nil, fmt.Errorf("failed to list folders: %w", err)
			}
			folderIds = append(folderIds, childFolders...)
		}

		logFolderIds(&c.logger, folderIds)

		c.logger.Info().Msg("listing folder projects..")
		folderProjects, err := listProjectsInFolders(ctx, c.Services.ResourcemanagerProjectsClient, folderIds)
		projects = append(projects, folderProjects...)
		if err != nil {
			return nil, fmt.Errorf("failed to list projects: %w", err)
		}
	}

	logProjectIds(&logger, projects)

	c.projects = projects
	if len(projects) == 1 {
		c.ProjectId = projects[0]
	}

	return &c, nil
}

func logFolderIds(logger *zerolog.Logger, folderIds []string) {
	// If there are too many folders, just log the first maxProjectIdsToLog.
	if len(folderIds) > maxProjectIdsToLog {
		logger.Info().Interface("folder_ids", folderIds[:maxProjectIdsToLog]).Msgf("Found %d folders. First %d: ", len(folderIds), maxProjectIdsToLog)
		logger.Debug().Interface("folder_ids", folderIds).Msg("All folders: ")
	} else {
		logger.Info().Interface("folder_ids", folderIds).Msgf("Found %d projects in folders", len(folderIds))
	}
}

func logProjectIds(logger *zerolog.Logger, projectIds []string) {
	// If there are too many folders, just log the first maxProjectIdsToLog.
	if len(projectIds) > maxProjectIdsToLog {
		logger.Info().Interface("projects", projectIds[:maxProjectIdsToLog]).Msgf("Found %d projects. First %d: ", len(projectIds), maxProjectIdsToLog)
		logger.Debug().Interface("projects", projectIds).Msg("All projects: ")
	} else {
		logger.Info().Interface("projects", projectIds).Msgf("Found %d projects in folders", len(projectIds))
	}
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

// listFolders recursively lists the folders in the 'parent' folder. Includes the 'parent' folder itself.
// recursionDepth is the depth of folders to recurse - where 0 means not to recurse any folders.
func listFolders(ctx context.Context, folderClient *resourcemanagerv3.FoldersClient, parent string, recursionDepth int) ([]string, error) {
	folders := []string{
		parent,
	}
	if recursionDepth <= 0 {
		return folders, nil
	}

	it := folderClient.ListFolders(ctx, &resourcemanagerv3pb.ListFoldersRequest{
		Parent: parent,
	})

	for {
		folder, err := it.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if folder.State == resourcemanagerv3pb.Folder_ACTIVE {
			folders = append(folders, folder.Name)
		}
	}

	return folders, nil
}

func listProjectsInFolders(ctx context.Context, projectClient *resourcemanagerv3.ProjectsClient, folders []string) ([]string, error) {
	projects := []string{}
	for _, folder := range folders {
		it := projectClient.ListProjects(ctx, &resourcemanagerv3pb.ListProjectsRequest{
			Parent: folder,
		})

		for {
			project, err := it.Next()

			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}

			if project.State == resourcemanagerv3pb.Project_ACTIVE {
				projects = append(projects, project.ProjectId)
			}
		}
	}

	return projects, nil
}
