package client

import "fmt"

type ResourceDiscovery struct {
	IncludeListId []string `json:"include_list_id"`
	ExcludeListId []string `json:"exclude_list_id"`
	IncludeFilter []string `json:"include_filter"`
}

// If entire object is nil then all projects will be included in the sync
// If only exclude statements are defined then all of that resource will found and exclusions will be applied
type HierarchyDiscovery struct {
	// if an organization is listed as excluded, all projects and folders under that organization will be excluded
	// if an organization is listed as included, all project and folder underneath it will be included, unless otherwise excluded
	// If no organizations are specified, then organizations will not be included or excluded on the basis of organizations this also means that no orgs will be used to multiplex
	// Explicit include/exclude lists will only be applied after any filter statements have been executed
	Organizations ResourceDiscovery `json:"organizations"`

	// if a folder is listed as excluded, all projects and folders under that folder will be excluded
	// if a folder is listed as included, all project and folder underneath it will be included, unless otherwise excluded
	// if no folders are specified, then folders will not be included or excluded on the basis of organizations this also means that no folders will be used to multiplex
	// Explicit include/exclude lists will only be applied after any filter statements have been executed
	Folders ResourceDiscovery `json:"folders"`

	// Explicit include/exclude lists will only be applied after any filter statements have been executed
	Projects ResourceDiscovery `json:"projects"`
}

// Spec defines GCP source plugin Spec
type Spec struct {
	ProjectIDs                  []string           `json:"project_ids"`
	ServiceAccountKeyJSON       string             `json:"service_account_key_json"`
	FolderFilter                string             `json:"folder_filter"`
	FolderIDs                   []string           `json:"folder_ids"`
	FolderRecursionDepth        *int               `json:"folder_recursion_depth"`
	ProjectFilter               string             `json:"project_filter"`
	BackoffDelay                int                `json:"backoff_delay"`
	BackoffRetries              int                `json:"backoff_retries"`
	DiscoveryConcurrency        *int               `json:"discovery_concurrency"`
	EnabledServicesOnly         bool               `json:"enabled_services_only"`
	OrganizationIDs             []string           `json:"organization_ids"`
	OrganizationFilter          string             `json:"organization_filter"`
	ServiceAccountImpersonation *CredentialsConfig `json:"service_account_impersonation"`

	// Experimental
	HierarchyDiscovery HierarchyDiscovery `json:"hierarchy_discovery"`
}

type CredentialsConfig struct {
	// TargetPrincipal is the email address of the service account to
	// impersonate. Required.
	TargetPrincipal string `json:"target_principal"`
	// Scopes that the impersonated credential should have. Required.
	Scopes []string `json:"scopes"`
	// Delegates are the service account email addresses in a delegation chain.
	// Each service account must be granted roles/iam.serviceAccountTokenCreator
	// on the next service account in the chain. Optional.
	Delegates []string `json:"delegates"`
	// Subject is the sub field of a JWT. This field should only be set if you
	// wish to impersonate as a user. This feature is useful when using domain
	// wide delegation. Optional.
	Subject string `json:"subject"`
}

func (spec *Spec) validate() error {
	if !spec.HierarchyDiscovery.isNull() {
		if len(spec.ProjectIDs) > 0 {
			return fmt.Errorf("cannot specify both project_ids and projects")
		}
		if spec.FolderFilter != "" {
			return fmt.Errorf("cannot specify both folder_filter and projects")
		}
		if spec.ProjectFilter != "" {
			return fmt.Errorf("cannot specify both project_filter and projects")
		}

		if len(spec.OrganizationIDs) > 0 {
			return fmt.Errorf("cannot specify both organization_ids and projects")
		}
		if spec.OrganizationFilter != "" {
			return fmt.Errorf("cannot specify both organization_filter and projects")
		}
	}
	return nil
}
func (spec *Spec) setDefaults() {
	var defaultRecursionDepth = 100
	if spec.FolderRecursionDepth == nil {
		spec.FolderRecursionDepth = &defaultRecursionDepth
	}

	var defaultDiscoveryConcurrency = 100
	if spec.DiscoveryConcurrency == nil {
		spec.DiscoveryConcurrency = &defaultDiscoveryConcurrency
	}
	if spec.ServiceAccountImpersonation != nil {
		if len(spec.ServiceAccountImpersonation.Scopes) == 0 {
			spec.ServiceAccountImpersonation.Scopes = []string{"https://www.googleapis.com/auth/cloud-platform"}
		}
	}
}

func (rd ResourceDiscovery) isIncludeNull() bool {
	if len(rd.IncludeListId) > 0 {
		return false
	}
	if len(rd.IncludeFilter) > 0 {
		return false
	}
	return true
}

func (rd ResourceDiscovery) isExcludeNull() bool {
	return len(rd.ExcludeListId) <= 0
}

func (rd ResourceDiscovery) isNull() bool {
	return rd.isIncludeNull() && rd.isExcludeNull()
}

func (hd HierarchyDiscovery) isNull() bool {
	return hd.Organizations.isNull() && hd.Projects.isNull() && hd.Folders.isNull()
}
