package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/googleapis/gax-go/v2"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxIdsToLog int = 100

type Client struct {
	projects  []string
	orgs      []*crmv1.Organization
	folderIds []string

	ClientOptions []option.ClientOption
	CallOptions   []gax.CallOption

	EnabledServices map[string]map[string]any
	// this is set by table client project multiplexer
	ProjectId string
	// this is set by table client Org multiplexer
	OrgId string
	Org   *crmv1.Organization
	// this is set by table client Folder multiplexer
	FolderId string
	// this is set by table client Location multiplexer
	Location string
	// Logger
	logger zerolog.Logger

	Backend state.Client
}

func (c *Client) WithBackend(backend state.Client) *Client {
	newClient := *c
	newClient.Backend = backend
	return &newClient
}

//revive:disable:modifies-value-receiver

// withProject allows multiplexer to create a new client with given projectId
func (c *Client) withProject(project string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("project_id", project).Logger()
	newClient.ProjectId = project
	return &newClient
}

func (c *Client) withLocation(location string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("location", location).Logger()
	newClient.Location = location
	return &newClient
}

// withOrg allows multiplexer to create a new client with given organization
func (c *Client) withOrg(org *crmv1.Organization) *Client {
	orgId := strings.TrimPrefix(org.Name, "organizations/")
	newClient := *c
	newClient.logger = c.logger.With().Str("org_id", orgId).Logger()
	newClient.OrgId = orgId
	newClient.Org = org
	return &newClient
}

// withFolder allows multiplexer to create a new client with given folderId
func (c *Client) withFolder(folder string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("folder_id", folder).Logger()
	newClient.FolderId = folder
	return &newClient
}

func isValidJson(content []byte) error {
	var v map[string]any
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
	if c.FolderId != "" {
		return "folder:" + c.FolderId
	}
	if c.Location != "" {
		return "project:" + c.ProjectId + ":location:" + c.Location
	}
	return "project:" + c.ProjectId
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, spec *Spec) (schema.ClientMeta, error) {
	var err error
	c := Client{
		logger:          logger,
		EnabledServices: map[string]map[string]any{},
	}

	projects := spec.ProjectIDs
	organizations := make([]*crmv1.Organization, 0)
	if spec.BackoffRetries > 0 {
		c.CallOptions = append(c.CallOptions, gax.WithRetry(func() gax.Retryer {
			return &Retrier{
				backoff: gax.Backoff{
					Max: time.Duration(spec.BackoffDelay) * time.Second,
				},
				maxRetries: spec.BackoffRetries,
				codes:      []codes.Code{codes.ResourceExhausted},
			}
		}))
	}
	unaryInterceptor := grpc.WithUnaryInterceptor(logging.UnaryClientInterceptor(grpczerolog.InterceptorLogger(logger)))
	streamInterceptor := grpc.WithStreamInterceptor(logging.StreamClientInterceptor(grpczerolog.InterceptorLogger(logger)))

	serviceAccountKeyJSON := []byte(spec.ServiceAccountKeyJSON)
	// Add a fake request reason because it is not possible to pass nil options
	c.ClientOptions = append(c.ClientOptions,
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the same side with telemetry
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
	if spec.ServiceAccountImpersonation != nil && spec.ServiceAccountImpersonation.TargetPrincipal != "" {
		// Base credentials sourced from ADC or provided client options.
		ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
			TargetPrincipal: spec.ServiceAccountImpersonation.TargetPrincipal,
			Scopes:          spec.ServiceAccountImpersonation.Scopes,
			// Optionally supply delegates.
			Delegates: spec.ServiceAccountImpersonation.Delegates,
			// Specify user to impersonate
			Subject: spec.ServiceAccountImpersonation.Subject,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to generate token source: %w", err)
		}
		c.ClientOptions = append(c.ClientOptions, option.WithTokenSource(ts))
	}
	if len(spec.ProjectFilter) > 0 && len(spec.FolderIDs) > 0 {
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

	switch {
	case len(projects) == 0 && len(spec.FolderIDs) == 0 && len(spec.ProjectFilter) == 0:
		c.logger.Info().Msg("No project_ids, folder_ids, or project_filter specified - assuming all active projects")
		projects, err = getProjectsV1(ctx, c.ClientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}

	case len(spec.FolderIDs) > 0:
		var folderIds []string

		for _, parentFolder := range spec.FolderIDs {
			c.logger.Info().Msg("Listing folders...")
			childFolders, err := listFolders(ctx, foldersClient, parentFolder, *spec.FolderRecursionDepth)
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

	case len(spec.ProjectFilter) > 0:
		c.logger.Info().Msg("Listing projects with filter...")
		projectsWithFilter, err := getProjectsV1WithFilter(ctx, spec.ProjectFilter, c.ClientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects with filter: %w", err)
		}

		projects = setUnion(projects, projectsWithFilter)
	}

	if len(spec.OrganizationIDs) == 0 && len(spec.OrganizationFilter) == 0 {
		c.logger.Info().Msg("No organization_ids or organization_filter specified - assuming all organizations")
		c.logger.Info().Msg("Listing organizations...")

		organizations, err = getOrganizations(ctx, "", c.ClientOptions...)
		if err != nil {
			c.logger.Err(err).Msg("failed to get organizations")
		}
	} else {
		if len(spec.OrganizationIDs) > 0 {
			for _, orgID := range spec.OrganizationIDs {
				c.logger.Info().Msgf("Getting spec organization %q...", orgID)
				org, err := getOrganization(ctx, orgID, c.ClientOptions...)
				if err != nil {
					return nil, fmt.Errorf("failed to get spec organization: %w", err)
				}
				organizations = append(organizations, org)
			}
		}
		if len(spec.OrganizationFilter) > 0 {
			c.logger.Info().Msg("Listing organizations with filter...")
			organizationsWithFilter, err := getOrganizations(ctx, spec.OrganizationFilter, c.ClientOptions...)
			if err != nil {
				return nil, fmt.Errorf("failed to get organizations with filter: %w", err)
			}
			for i := range organizationsWithFilter {
				found := false
				for _, org := range organizations {
					if organizationsWithFilter[i].Name == org.Name {
						found = true
						break
					}
				}
				if !found {
					organizations = append(organizations, organizationsWithFilter[i])
				}
			}
		}
	}

	logProjectIds(&logger, projects)
	logOrganizationIds(&logger, organizations)

	if len(projects) == 0 {
		return nil, fmt.Errorf("no active projects")
	}

	c.projects = projects
	c.folderIds = spec.FolderIDs
	c.orgs = organizations
	if err != nil {
		c.logger.Err(err).Msg("failed to get organizations")
	}
	c.logger.Info().Interface("orgs", c.orgs).Msg("Retrieved organizations")

	if len(projects) == 1 {
		c.ProjectId = projects[0]
	}
	if spec.EnabledServicesOnly {
		if err := c.configureEnabledServices(ctx, *spec.DiscoveryConcurrency); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func logFolderIds(logger *zerolog.Logger, folderIds []string) {
	// If there are too many folders, just log the first maxProjectIdsToLog.
	if len(folderIds) > maxIdsToLog {
		logger.Info().Interface("folder_ids", folderIds[:maxIdsToLog]).Msgf("Found %d folders. First %d: ", len(folderIds), maxIdsToLog)
		logger.Debug().Interface("folder_ids", folderIds).Msg("All folders: ")
	} else {
		logger.Info().Interface("folder_ids", folderIds).Msgf("Found %d projects in folders", len(folderIds))
	}
}

func logProjectIds(logger *zerolog.Logger, projectIds []string) {
	// If there are too many folders, just log the first maxIdsToLog.
	if len(projectIds) > maxIdsToLog {
		logger.Info().Interface("projects", projectIds[:maxIdsToLog]).Msgf("Found %d projects. First %d: ", len(projectIds), maxIdsToLog)
		logger.Debug().Interface("projects", projectIds).Msg("All projects: ")
	} else {
		logger.Info().Interface("projects", projectIds).Msgf("Found %d projects in folders", len(projectIds))
	}
}

func logOrganizationIds(logger *zerolog.Logger, organizations []*crmv1.Organization) {
	// If there are too many organizations, just log the first maxIdsToLog.
	organizationIds := make([]string, len(organizations))
	for i, org := range organizations {
		organizationIds[i] = org.Name
	}
	if len(organizationIds) > maxIdsToLog {
		logger.Info().Interface("organizations", organizationIds[:maxIdsToLog]).Msgf("Found %d organizations. First %d: ", len(organizationIds), maxIdsToLog)
		logger.Debug().Interface("organizations", organizationIds).Msg("All organizations: ")
	} else {
		logger.Info().Interface("organizations", organizationIds).Msgf("Found %d organizations in folders", len(organizationIds))
	}
}

// getProjectsV1 requires the `resourcemanager.projects.get` permission to list projects
func getProjectsV1(ctx context.Context, options ...option.ClientOption) ([]string, error) {
	var projects []string

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

func getProjectsV1WithFilter(ctx context.Context, filter string, options ...option.ClientOption) ([]string, error) {
	var projects []string

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
			projects = append(projects, project.ProjectId)
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

func getOrganizations(ctx context.Context, filter string, options ...option.ClientOption) ([]*crmv1.Organization, error) {
	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	var orgs []*crmv1.Organization
	if err := service.Organizations.Search(&crmv1.SearchOrganizationsRequest{Filter: filter}).Context(ctx).Pages(ctx, func(page *crmv1.SearchOrganizationsResponse) error {
		orgs = append(orgs, page.Organizations...)
		return nil
	}); err != nil {
		return nil, err
	}

	return orgs, nil
}

func getOrganization(ctx context.Context, id string, options ...option.ClientOption) (*crmv1.Organization, error) {
	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	return service.Organizations.Get("organizations/" + id).Context(ctx).Do()
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

func (c *Client) configureEnabledServices(ctx context.Context, concurrency int) error {
	var esLock sync.Mutex
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)
	for _, p := range c.projects {
		project := p
		g.Go(func() error {
			cl := c.withProject(project)
			svc, err := cl.fetchEnabledServices(ctx)
			if err != nil {
				switch status.Code(err) {
				case codes.ResourceExhausted:
					c.logger.Warn().Err(err).Msgf("failed to list enabled services because of rate limiting. Sync will continue without filtering out disabled services for this project: %s. Consider setting larger values for `backoff_retries` and `backoff_delay`", project)
				case codes.PermissionDenied:
					c.logger.Warn().Err(err).Msgf("failed to list enabled services because of insufficient permissions. Sync will continue without filtering out disabled services for this project: %s", project)
				default:
					c.logger.Err(err).Msg("failed to list enabled services")
					return err
				}
				return nil
			}
			// Only update the enabled services if we were able to list all services successfully
			esLock.Lock()
			c.EnabledServices[project] = svc
			esLock.Unlock()
			return err
		})
	}
	return g.Wait()
}

func (c *Client) fetchEnabledServices(ctx context.Context) (map[string]any, error) {
	enabled := make(map[string]any)
	req := &pb.ListServicesRequest{
		Parent:   "projects/" + c.ProjectId,
		PageSize: 200,
		Filter:   "state:ENABLED",
	}
	gcpClient, err := serviceusage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return nil, err
	}
	it := gcpClient.ListServices(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		enabled[resp.GetConfig().Name] = resp
	}
	return enabled, nil
}
