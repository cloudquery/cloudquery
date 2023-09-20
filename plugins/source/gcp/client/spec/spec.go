package spec

import (
	_ "embed"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"google.golang.org/api/cloudresourcemanager/v1"
)

// Spec defines GCP source plugin Spec
type Spec struct {
	ProjectIDs                  []string           `json:"project_ids"`
	FolderIDs                   []string           `json:"folder_ids"`
	FolderRecursionDepth        int                `json:"folder_recursion_depth"`
	ProjectFilter               string             `json:"project_filter"`
	ServiceAccountKeyJSON       string             `json:"service_account_key_json"`
	BackoffDelay                int                `json:"backoff_delay"`
	BackoffRetries              int                `json:"backoff_retries"`
	DiscoveryConcurrency        int                `json:"discovery_concurrency"`
	EnabledServicesOnly         bool               `json:"enabled_services_only"`
	OrganizationIDs             []string           `json:"organization_ids"`
	OrganizationFilter          string             `json:"organization_filter"`
	ServiceAccountImpersonation *CredentialsConfig `json:"service_account_impersonation"`
	Concurrency                 int                `json:"concurrency"`
	Scheduler                   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (spec *Spec) SetDefaults() {
	if spec.BackoffRetries <= 0 {
		const defaultBackoffRetries = 3
		spec.BackoffRetries = defaultBackoffRetries
	}
	if spec.BackoffDelay <= 0 {
		const defaultBackoffDelay = 1
		spec.BackoffDelay = defaultBackoffDelay
	}

	if spec.FolderRecursionDepth < 0 {
		const defaultRecursionDepth = 100
		spec.FolderRecursionDepth = defaultRecursionDepth
	}

	if spec.DiscoveryConcurrency <= 0 {
		const defaultDiscoveryConcurrency = 100
		spec.DiscoveryConcurrency = defaultDiscoveryConcurrency
	}

	if spec.Concurrency <= 0 {
		const defaultConcurrency = 50000
		spec.Concurrency = defaultConcurrency
	}

	spec.ServiceAccountImpersonation.SetDefaults()
}

//go:embed schema.json
var JSONSchema string

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

func (c *CredentialsConfig) SetDefaults() {
	if c == nil {
		return
	}
	if len(c.Scopes) == 0 {
		// `https://www.googleapis.com/auth/cloud-platform`
		// We use this as some APIs don't utilize the read-only alternative `https://www.googleapis.com/auth/cloud-platform.read-only`
		// See https://developers.google.com/identity/protocols/oauth2/scopes for more details.
		c.Scopes = []string{cloudresourcemanager.CloudPlatformScope}
	}
}
