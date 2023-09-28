package spec

import (
	_ "embed"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/invopop/jsonschema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

// Spec defines GCP source plugin Spec
type Spec struct {
	ProjectIDs                  []string           `json:"project_ids" jsonschema:"minLength=1"`
	FolderIDs                   []string           `json:"folder_ids" jsonschema:"pattern=^(folders|organizations)/(.)+$"`
	FolderRecursionDepth        *int               `json:"folder_recursion_depth" jsonschema:"minimum=0,default=100"`
	OrganizationIDs             []string           `json:"organization_ids" jsonschema:"minLength=1"`
	ProjectFilter               string             `json:"project_filter"`
	OrganizationFilter          string             `json:"organization_filter"`
	ServiceAccountKeyJSON       string             `json:"service_account_key_json"`
	BackoffDelay                int                `json:"backoff_delay" jsonschema:"minimum=0,default=30"`
	BackoffRetries              int                `json:"backoff_retries" jsonschema:"minimum=0,default=0"`
	DiscoveryConcurrency        int                `json:"discovery_concurrency" jsonschema:"minimum=1,default=100"`
	EnabledServicesOnly         bool               `json:"enabled_services_only"`
	Concurrency                 int                `json:"concurrency" jsonschema:"minimum=1,default=50000"`
	Scheduler                   scheduler.Strategy `json:"scheduler,omitempty" jsonschema:"type=string,default=dfs,enum=dfs,enum=round-robin,enum=shuffle"`
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

// JSONSchemaExtend is required to add `not` section for `project_filter` & `folder_ids` being mutually exclusive.
// We use value receiver because of https://github.com/invopop/jsonschema/issues/102
func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Not = &jsonschema.Schema{Required: []string{"project_filter", "folder_ids"}}
}

//go:embed schema.json
var JSONSchema string

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
