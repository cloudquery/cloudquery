package spec

import (
	_ "embed"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// Spec defines GCP source plugin Spec
type Spec struct {
	// Specify projects to connect to.
	// If either `folder_ids` or `project_filter` is specified,
	// these projects will be synced in addition to the projects from the folder/filter.
	//
	// Empty or `null` value will use all projects available to the current authenticated account.
	ProjectIDs []string `json:"project_ids" jsonschema:"minLength=1"`

	// CloudQuery will sync from all the projects in the specified folders, recursively.
	// `folder_ids` must be of the format `folders/<folder_id>` or `organizations/<organization_id>`.
	// This feature requires the `resourcemanager.folders.list` permission.
	//
	// By default, CloudQuery will also sync from sub-folders recursively (up to depth `100`).
	// To reduce this, set `folder_recursion_depth` to a lower value (or to `0` to disable recursion completely).
	//
	// Mutually exclusive with `project_filter`.
	FolderIDs []string `json:"folder_ids" jsonschema:"pattern=^(folders|organizations)/(.)+$"`

	// The maximum depth to recurse into sub-folders.
	// `0` means no recursion (only the top-level projects in folders will be used for sync).
	FolderRecursionDepth *int `json:"folder_recursion_depth" jsonschema:"minimum=0,default=100"`

	// Specify organizations to use when syncing organization level resources (e.g.
	// [folders](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/resourcemanager/folders_fetch.go#L23)
	// or
	// [security findings](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/securitycenter/organization_findings.go#L43)).
	//
	// If `organization_filter` is specified, these organizations will be used in addition to the organizations from the filter.
	//
	// Empty or `null` value will use all organizations available to the current authenticated account).
	OrganizationIDs []string `json:"organization_ids" jsonschema:"minLength=1"`

	// A filter to determine the projects that are synced, mutually exclusive with `folder_ids`.
	//
	// For instance, to only sync projects where the name starts with `how-`, set `project_filter` to `name:how-*`.
	//
	// More examples:
	//
	// - `"name:how-* OR name:test-*"` matches projects starting with `how-` or `test-`
	// - `"NOT name:test-*"` matches all projects _not_ starting with `test-`
	//
	// For syntax and example queries refer to API References
	// [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#google.cloudresourcemanager.v1.Projects.ListProjects)
	// and
	// [here](https://cloud.google.com/sdk/gcloud/reference/topic/filters).
	ProjectFilter string `json:"project_filter"`

	// A filter to determine the organizations to use when syncing organization level resources (e.g.
	// [folders](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/resourcemanager/folders_fetch.go#L23)
	// or
	// [security findings](https://github.com/cloudquery/cloudquery/blob/0e384a84d1c9545b24c2eda9af00f111bab79c36/plugins/source/gcp/resources/services/securitycenter/organization_findings.go#L43)).
	//
	// For instance, to use only organizations from the `cloudquery.io` domain, set `organization_filter` to `domain:cloudquery.io`.
	//
	// For syntax and example queries refer to API Reference [here](https://cloud.google.com/resource-manager/reference/rest/v1/organizations/search#google.cloudresourcemanager.v1.SearchOrganizationsRequest).
	OrganizationFilter string `json:"organization_filter"`

	// GCP service account key content.
	//
	// Using service accounts is not recommended, but if it is used it is better to use
	// [environment or file variable substitution](/docs/advanced-topics/environment-variable-substitution).
	ServiceAccountKeyJSON string `json:"service_account_key_json"`

	// If specified APIs will be retried with exponential backoff if they are rate limited.
	// This is the max delay (in seconds) between retries.
	BackoffDelay int `json:"backoff_delay" jsonschema:"minimum=0,default=30"`

	// If specified APIs will be retried with exponential backoff if they are rate limited.
	// This is the max number of retries.
	BackoffRetries int `json:"backoff_retries" jsonschema:"minimum=0,default=0"`

	// If enabled CloudQuery will skip any resources that belong to a service that has been disabled or not been enabled.
	//
	// If you use this option on a large organization (with more than `500` projects)
	// you should also set the `backoff_retries` to a value greater than `0`, otherwise you may hit the API rate limits.
	//
	// In `>=v9.0.0` if an error is returned then CloudQuery will assume that all services are enabled
	// and will continue to attempt to sync all specified tables rather than just ending the sync.
	EnabledServicesOnly bool `json:"enabled_services_only"`

	// The best effort maximum number of Go routines to use.
	// Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency" jsonschema:"minimum=1,default=50000"`

	// The number of concurrent requests that CloudQuery will make to resolve enabled services.
	// This is only used when `enabled_services_only` is set to `true`.
	DiscoveryConcurrency int `json:"discovery_concurrency" jsonschema:"minimum=1,default=100"`

	// The scheduler to use when determining the priority of resources to sync.
	//
	// For more information about this, see [performance tuning](/docs/advanced-topics/performance-tuning).
	Scheduler scheduler.Strategy `json:"scheduler,omitempty"`

	// Service Account impersonation configuration.
	ServiceAccountImpersonation *CredentialsConfig `json:"service_account_impersonation"`
}

func (s *Spec) Validate() error {
	if len(s.ProjectFilter) > 0 && len(s.FolderIDs) > 0 {
		return fmt.Errorf("project_filter and folder_ids are mutually exclusive")
	}
	return nil
}

func (s *Spec) SetDefaults() {
	if s.BackoffRetries < 0 {
		const defaultBackoffRetries = 0
		s.BackoffRetries = defaultBackoffRetries
	}
	if s.BackoffDelay < 0 {
		const defaultBackoffDelay = 30
		s.BackoffDelay = defaultBackoffDelay
	}

	if s.FolderRecursionDepth == nil || *s.FolderRecursionDepth < 0 {
		// we do allow 0 as value
		var defaultRecursionDepth = 100
		s.FolderRecursionDepth = &defaultRecursionDepth
	}

	if s.DiscoveryConcurrency <= 0 {
		const defaultDiscoveryConcurrency = 100
		s.DiscoveryConcurrency = defaultDiscoveryConcurrency
	}

	if s.Concurrency <= 0 {
		const defaultConcurrency = 50000
		s.Concurrency = defaultConcurrency
	}

	s.ServiceAccountImpersonation.SetDefaults()
}

// JSONSchemaExtend is required to add `not` section for `project_filter` & `folder_ids` being mutually exclusive.
func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Not = &jsonschema.Schema{
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			one := uint64(1)
			properties := jsonschema.NewProperties()

			projectFilter := *sc.Properties.Value("project_filter")
			projectFilter.MinLength = &one
			properties.Set("project_filter", &projectFilter)

			folderIDs := *sc.Properties.Value("folder_ids").OneOf[0] // 0 is spec, 1 is null
			folderIDs.MinItems = &one
			items := *folderIDs.Items
			items.MinLength = &one
			items.Pattern = ""
			folderIDs.Items = &items

			properties.Set("folder_ids", &folderIDs)

			return properties
		}(),
		Required: []string{"project_filter", "folder_ids"},
	}
}

//go:embed schema.json
var JSONSchema string
