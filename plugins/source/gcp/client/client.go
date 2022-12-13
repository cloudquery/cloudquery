package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/googleapis/gax-go/v2"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const maxProjectIdsToLog int = 100

type Client struct {
	projects []string
	orgs     []string

	ClientOptions []option.ClientOption
	CallOptions   []gax.CallOption
	// this is set by table client multiplexer
	ProjectId string
	OrgId     string
	// Logger
	logger zerolog.Logger
}

//revive:disable:modifies-value-receiver

// withProject allows multiplexer to create a new client with given projectId
func (c *Client) withProject(project string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("project_id", project).Logger()
	newClient.ProjectId = project
	return &newClient
}

// withOrg allows multiplexer to create a new client with given organizationId
func (c *Client) withOrg(org string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("org_id", org).Logger()
	newClient.OrgId = org
	return &newClient
}

func isValidJson(content []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(content, &v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ID() string {
	if c.OrgId != "" {
		return "org:" + c.OrgId
	}
	return c.ProjectId
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var err error
	c := Client{
		logger: logger,
	}
	var gcpSpec Spec
	if err := s.UnmarshalSpec(&gcpSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	gcpSpec.setDefaults()
	projects := gcpSpec.ProjectIDs
	if gcpSpec.BackoffRetries > 0 {
		c.CallOptions = append(c.CallOptions, gax.WithRetry(func() gax.Retryer {
			return &Retrier{
				backoff: gax.Backoff{
					Max: time.Duration(gcpSpec.BackoffDelay) * time.Second,
				},
				maxRetries: gcpSpec.BackoffRetries,
				codes:      []codes.Code{codes.ResourceExhausted},
			}
		}))
	}
	unaryInterceptor := grpc.WithUnaryInterceptor(logging.UnaryClientInterceptor(grpczerolog.InterceptorLogger(logger)))
	streamInterceptor := grpc.WithStreamInterceptor(logging.StreamClientInterceptor(grpczerolog.InterceptorLogger(logger)))

	serviceAccountKeyJSON := []byte(gcpSpec.ServiceAccountKeyJSON)
	// Add a fake request reason because it is not possible to pass nil options
	c.ClientOptions = append(c.ClientOptions,
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the same side with telemtry
		option.WithTelemetryDisabled(),
		option.WithGRPCDialOption(
			unaryInterceptor,
		),
		option.WithGRPCDialOption(
			streamInterceptor,
		))
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, fmt.Errorf("invalid json at service_account_key_json: %w", err)
		}
		c.ClientOptions = append(c.ClientOptions, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	if len(gcpSpec.ProjectFilter) > 0 && len(gcpSpec.FolderIDs) > 0 {
		return nil, fmt.Errorf("project_filter and folder_ids are mutually exclusive")
	}

	projectsClient, err := resourcemanager.NewProjectsClient(ctx, c.ClientOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create projects client: %w", err)
	}
	foldersClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create folders client: %w", err)
	}

	projectsAndParents := make(map[string]string) // project -> parent
	projectFolders := make(map[string][]string)   // project -> folders

	switch {
	case len(projects) == 0 && len(gcpSpec.FolderIDs) == 0 && len(gcpSpec.ProjectFilter) == 0:
		c.logger.Info().Msg("No project_ids, folder_ids, or project_filter specified - assuming all active projects")
		projectsAndParents, err = getProjectsV1(ctx, c.ClientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}
		projects = funk.Keys(projectsAndParents).([]string)

	case len(gcpSpec.FolderIDs) > 0:
		var folderIds []string

		for _, parentFolder := range gcpSpec.FolderIDs {
			c.logger.Info().Msg("Listing folders...")
			childFolders, err := listFolders(ctx, foldersClient, parentFolder, *gcpSpec.FolderRecursionDepth)
			if err != nil {
				return nil, fmt.Errorf("failed to list folders: %w", err)
			}
			folderIds = append(folderIds, childFolders...)
		}

		logFolderIds(&c.logger, folderIds)

		c.logger.Info().Msg("listing folder projects...")
		folderProjects, err := listProjectsInFolders(ctx, projectsClient, folderIds)
		projects = setUnion(projects, folderProjects)
		if err != nil {
			return nil, fmt.Errorf("failed to list projects: %w", err)
		}
		for _, p := range folderProjects {
			for _, f := range folderIds {
				projectFolders[p] = append(projectFolders[p], f)
			}
		}

	case len(gcpSpec.ProjectFilter) > 0:
		c.logger.Info().Msg("Listing projects with filter...")
		projectsAndParents, err = getProjectsV1WithFilter(ctx, gcpSpec.ProjectFilter, c.ClientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects with filter: %w", err)
		}

		projectsWithFilter := funk.Keys(projectsAndParents).([]string)
		projects = setUnion(projects, projectsWithFilter)
	}

	logProjectIds(&logger, projects)

	if len(projects) == 0 {
		return nil, fmt.Errorf("no active projects")
	}

	// get list of orgs
	projectVsOrg, projectVsFolder := siftProjectParents(projectsAndParents) // each project has one parent, either a folder or an org
	for p, f := range projectVsFolder {
		projectFolders[p] = append(projectFolders[p], f)
	}

	orgs := funk.Values(projectVsOrg).([]string)

	folderVsOrg := make(map[string]string)
	for p, f := range projectVsFolder {
		// do we have info about this folder's org already?
		if o := folderVsOrg[f]; o != "" {
			c.Logger().Debug().Str("project_id", p).Str("org_id", o).Msg("already have the org from previous data")
			orgs = append(orgs, o)
			continue
		}
		if o := projectVsOrg[p]; o != "" {
			c.Logger().Debug().Str("project_id", p).Str("org_id", o).Msg("already have the org")
			continue
		}

		// if our folder is in the same list as other folders, check them for org first
		var otherFolders []string
		for _, sameOrgFolderList := range projectFolders {
			if funk.Contains(sameOrgFolderList, f) {
				otherFolders = append(otherFolders, sameOrgFolderList...)
			}
		}
		// get the org for this folder
		found := false
		for _, otherFolder := range otherFolders {
			if o := folderVsOrg[otherFolder]; o != "" {
				folderVsOrg[f] = o
				c.Logger().Debug().Str("project_id", p).Str("org_id", o).Str("folder_id", otherFolder).Msg("already have the org from previous folder")
				orgs = append(orgs, o)
				found = true
				break
			}
		}
		if found {
			continue
		}

		c.Logger().Debug().Str("project_id", p).Msg("querying project org")
		o, folders, err := getOrganization(ctx, p, c.ClientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to get organization of project %s: %w", p, err)
		}
		if o == "" {
			c.Logger().Warn().Str("project_id", p).Msg("could not get projects organization")
			continue
		}
		orgs = append(orgs, o)
		for _, f := range folders {
			folderVsOrg[f] = o
		}
	}

	c.projects = projects
	c.orgs = funk.UniqString(orgs)

	if len(projects) == 1 {
		c.ProjectId = projects[0]
		c.OrgId = projectVsOrg[c.ProjectId]
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
func getProjectsV1(ctx context.Context, options ...option.ClientOption) (map[string]string, error) {
	projects := make(map[string]string)

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
			if project.Parent != nil {
				projects[project.ProjectId] = project.Parent.Type + "/" + project.Parent.Id
			} else {
				projects[project.ProjectId] = ""
			}
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

func getProjectsV1WithFilter(ctx context.Context, filter string, options ...option.ClientOption) (map[string]string, error) {
	projects := make(map[string]string)

	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	call := service.Projects.List().Filter(filter).Context(ctx)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			if project.LifecycleState != "ACTIVE" {
				continue
			}
			if project.Parent != nil {
				projects[project.ProjectId] = project.Parent.Type + "/" + project.Parent.Id
			} else {
				projects[project.ProjectId] = ""
			}
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	return projects, nil
}

// listFolders recursively lists the folders in the 'parent' folder. Includes the 'parent' folder itself.
// recursionDepth is the depth of folders to recurse - where 0 means not to recurse any folders.
func listFolders(ctx context.Context, folderClient *resourcemanager.FoldersClient, parent string, recursionDepth int) ([]string, error) {
	folders := []string{
		parent,
	}
	if recursionDepth <= 0 {
		return folders, nil
	}

	it := folderClient.ListFolders(ctx, &resourcemanagerpb.ListFoldersRequest{
		Parent: parent,
	})

	for {
		child, err := it.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if child.State == resourcemanagerpb.Folder_ACTIVE {
			childFolders, err := listFolders(ctx, folderClient, child.Name, recursionDepth-1)
			if err != nil {
				return nil, err
			}
			folders = append(folders, childFolders...)
		}
	}

	return folders, nil
}

func listProjectsInFolders(ctx context.Context, projectClient *resourcemanager.ProjectsClient, folders []string) ([]string, error) {
	var projects []string
	for _, folder := range folders {
		it := projectClient.ListProjects(ctx, &resourcemanagerpb.ListProjectsRequest{
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

			if project.State == resourcemanagerpb.Project_ACTIVE {
				projects = append(projects, project.ProjectId)
			}
		}
	}

	return projects, nil
}

func getOrganization(ctx context.Context, projectId string, options ...option.ClientOption) (string, []string, error) {
	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	ac, err := service.Projects.GetAncestry(projectId, &crmv1.GetAncestryRequest{}).Context(ctx).Do()
	if err != nil {
		return "", nil, err
	}

	org := ""
	folders := make([]string, 0, len(ac.Ancestor))
	for _, a := range ac.Ancestor {
		if a == nil || a.ResourceId == nil {
			continue
		}
		switch a.ResourceId.Type {
		case "organization":
			org = a.ResourceId.Id
		case "folder":
			folders = append(folders, a.ResourceId.Id)
		}
	}

	return org, folders, nil
}

func setUnion(a []string, b []string) []string {
	set := make(map[string]struct{}, len(a)+len(b)) // alloc max
	for _, s := range a {
		set[s] = struct{}{}
	}
	for _, s := range b {
		set[s] = struct{}{}
	}

	union := make([]string, 0, len(set))
	for s := range set {
		union = append(union, s)
	}
	return union
}

func siftProjectParents(projectsAndParents map[string]string) (map[string]string, map[string]string) {
	orgs := make(map[string]string)
	folders := make(map[string]string)
	for project, parent := range projectsAndParents {
		if strings.HasPrefix(parent, "organization/") {
			orgs[project] = strings.TrimPrefix(parent, "organization/")
			continue
		}

		if strings.HasPrefix(parent, "folder/") {
			folders[project] = strings.TrimPrefix(parent, "folder/")
		}
	}
	return orgs, folders
}
