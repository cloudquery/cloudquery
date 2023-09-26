package spec

import (
	_ "embed"
	"log"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"google.golang.org/api/cloudresourcemanager/v1"
)

// Spec defines GCP source plugin Spec
type Spec struct {
	ProjectIDs                  []string           `json:"project_ids"`
	FolderIDs                   []string           `json:"folder_ids"`
	FolderRecursionDepth        *int               `json:"folder_recursion_depth"`
	OrganizationIDs             []string           `json:"organization_ids"`
	ProjectFilter               string             `json:"project_filter"`
	OrganizationFilter          string             `json:"organization_filter"`
	ServiceAccountKeyJSON       string             `json:"service_account_key_json"`
	BackoffDelay                int                `json:"backoff_delay"`
	BackoffRetries              int                `json:"backoff_retries"`
	DiscoveryConcurrency        int                `json:"discovery_concurrency"`
	EnabledServicesOnly         bool               `json:"enabled_services_only"`
	Concurrency                 int                `json:"concurrency"`
	Scheduler                   scheduler.Strategy `json:"scheduler,omitempty"`
	ServiceAccountImpersonation *CredentialsConfig `json:"service_account_impersonation"`
}

func (spec *Spec) SetDefaults() {
	if spec.BackoffRetries < 0 {
		const defaultBackoffRetries = 0
		spec.BackoffRetries = defaultBackoffRetries
	}
	if spec.BackoffDelay < 0 {
		const defaultBackoffDelay = 30
		spec.BackoffDelay = defaultBackoffDelay
	}

	if spec.FolderRecursionDepth == nil || *spec.FolderRecursionDepth < 0 {
		// we do allow 0 as value
		var defaultRecursionDepth = 100
		spec.FolderRecursionDepth = &defaultRecursionDepth
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

type CredentialsConfig struct {
	// TargetPrincipal is the email address of the service account to
	// impersonate. Required.
	TargetPrincipal string `json:"target_principal" jsonschema:"required,format=email"`
	// Scopes that the impersonated credential should have. Required.
	Scopes []string `json:"scopes" jsonschema:"pattern=^https://www.googleapis.com/auth/(.)+$,default=https://www.googleapis.com/auth/cloud-platform"`
	// Delegates are the service account email addresses in a delegation chain.
	// Each service account must be granted roles/iam.serviceAccountTokenCreator
	// on the next service account in the chain. Optional.
	Delegates []string `json:"delegates" jsonschema:"format=email"`
	// Subject is the subject field of a JWT (sub). This field should only be set if you
	// wish to impersonate as a user. This feature is useful when using domain
	// wide delegation. Optional.
	Subject string `json:"subject" jsonschema:"minLength=1"`
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

var jsonSchema string

func init() {
	data, err := jsonschema.Generate(new(Spec))
	if err != nil {
		log.Fatal(err)
	}
	jsonSchema = string(data)
}

func JSONSchema() string {
	return jsonSchema
}
