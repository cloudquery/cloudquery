package client

const (
	defaultConcurrency = 50000
)

// Spec defines GCP source plugin Spec
type Spec struct {
	ProjectIDs                  []string           `json:"project_ids"`
	ServiceAccountKeyJSON       string             `json:"service_account_key_json"`
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
	Concurrency                 int                `json:"concurrency"`
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

func (spec *Spec) SetDefaults() {
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
	if spec.Concurrency == 0 {
		spec.Concurrency = defaultConcurrency
	}
}
