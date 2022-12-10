package client

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/googleapis/gax-go/v2"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/api/serviceusage/v1"
	"google.golang.org/grpc/codes"
)

const maxProjectIdsToLog int = 100

type Client struct {
	projects      []string
	ClientOptions []option.ClientOption
	// this is set by table client multiplexer

	CallOptions     []gax.CallOption
	ProjectId       string
	EnabledServices map[string]map[GcpService]bool
	// Logger
	logger zerolog.Logger
}

type GcpService string

const (
	BigQueryService             GcpService = "bigquery.googleapis.com"
	CloudBillingService         GcpService = "cloudbilling.googleapis.com"
	CloudFunctionsService       GcpService = "cloudfunctions.googleapis.com"
	CloudKmsService             GcpService = "cloudkms.googleapis.com"
	CloudResourceManagerService GcpService = "cloudresourcemanager.googleapis.com"
	ComputeService              GcpService = "compute.googleapis.com"
	DnsService                  GcpService = "dns.googleapis.com"
	DomainsService              GcpService = "domains.googleapis.com"
	IamService                  GcpService = "iam.googleapis.com"
	KubernetesService           GcpService = "container.googleapis.com"
	LoggingService              GcpService = "logging.googleapis.com"
	MonitoringService           GcpService = "monitoring.googleapis.com"
	SqlAdminService             GcpService = "sqladmin.googleapis.com"
	StorageService              GcpService = "storage-api.googleapis.com"
	ContainerAnalysisService    GcpService = "containeranalysis.googleapis.com"
	RediService                 GcpService = "redis.googleapis.com"
)

//revive:disable:modifies-value-receiver

// withProject allows multiplexer to create a new client with given subscriptionId
func (c *Client) withProject(project string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("project_id", project).Logger()
	newClient.ProjectId = project
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
	return c.ProjectId
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var err error

	c := Client{
		logger:          logger,
		EnabledServices: map[string]map[GcpService]bool{},
		// plugin: p,
	}
	// providerConfig := config.(*Config)
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

	serviceAccountKeyJSON := []byte(gcpSpec.ServiceAccountKeyJSON)
	// Add a fake request reason because it is not possible to pass nil options
	options := []option.ClientOption{option.WithRequestReason("cloudquery resource fetch")}
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, fmt.Errorf("invalid json at service_account_key_json: %w", err)
		}
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	if len(gcpSpec.ProjectFilter) > 0 && len(gcpSpec.FolderIDs) > 0 {
		return nil, fmt.Errorf("project_filter and folder_ids are mutually exclusive")
	}

	projectsClient, err := resourcemanager.NewProjectsClient(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create projects client: %w", err)
	}
	foldersClient, err := resourcemanager.NewFoldersClient(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create folders client: %w", err)
	}

	switch {
	case len(projects) == 0 && len(gcpSpec.FolderIDs) == 0 && len(gcpSpec.ProjectFilter) == 0:
		c.logger.Info().Msg("No project_ids, folder_ids, or project_filter specified - assuming all active projects")
		projects, err = getProjectsV1(ctx, options...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}

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

	case len(gcpSpec.ProjectFilter) > 0:
		c.logger.Info().Msg("Listing projects with filter...")
		projectsWithFilter, err := getProjectsV1WithFilter(ctx, gcpSpec.ProjectFilter, options...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects with filter: %w", err)
		}

		projects = setUnion(projects, projectsWithFilter)
	}

	logProjectIds(&logger, projects)

	if len(projects) == 0 {
		return nil, fmt.Errorf("no active projects")
	}

	c.projects = projects
	if len(projects) == 1 {
		c.ProjectId = projects[0]
	}
	if gcpSpec.EnabledServicesOnly {
		if err := c.configureEnabledServices(); err != nil {
			// TODO: log why we failed to grab enabled services
			return nil, err
		}
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

func getProjectsV1WithFilter(ctx context.Context, filter string, options ...option.ClientOption) ([]string, error) {
	var (
		projects []string
	)
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
			if project.LifecycleState == "ACTIVE" {
				projects = append(projects, project.ProjectId)
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

func (c *Client) configureEnabledServices() error {
	var esLock sync.Mutex
	g, ctx := errgroup.WithContext(context.Background())
	maxGoroutines := 10
	goroutinesSem := semaphore.NewWeighted(int64(maxGoroutines))
	for _, p := range c.projects {
		project := p
		if err := goroutinesSem.Acquire(ctx, 1); err != nil {
			return err
		}
		g.Go(func() error {
			defer goroutinesSem.Release(1)
			cl := c.withProject(project)
			svc, err := cl.fetchEnabledServices(ctx)
			esLock.Lock()
			c.EnabledServices[project] = svc
			esLock.Unlock()
			return err
		})
	}
	return g.Wait()
}

func (c *Client) fetchEnabledServices(ctx context.Context) (map[GcpService]bool, error) {
	enabled := make(map[GcpService]bool)
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

		item := resp.GetConfig()
		serviceName := GcpService(item.Name)

		enabled[serviceName] = true
	}
	return enabled, nil
}
