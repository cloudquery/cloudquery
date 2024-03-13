package spec

import (
	_ "embed"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/invopop/jsonschema"
	"golang.org/x/exp/maps"
)

type Spec struct {
	// Specify which subscriptions to sync data from.
	// Empty means all visible subscriptions.
	Subscriptions []string `json:"subscriptions" jsonschema:"minLength=1,uniqueItems=true,example=00000000-0000-0000-0000-000000000000"`

	// A list of subscription IDs that CloudQuery will skip syncing.
	// This is useful if CloudQuery is discovering the list of subscription IDs and there are some subscriptions that you want to not even attempt syncing.
	SkipSubscriptions []string `json:"skip_subscriptions" jsonschema:"minLength=1,uniqueItems=true,example=00000000-0000-0000-0000-000000000000"`

	// The name of the cloud environment to use.
	// See the [Azure CLI documentation](https://learn.microsoft.com/en-us/cli/azure/manage-clouds-azure-cli) for more information.
	CloudName string `json:"cloud_name" jsonschema:"minLength=1,example=AzureCloud"`

	// Enabling this setting will force all `id` column values to be lowercase.
	// This is useful to avoid case sensitivity and uniqueness issues around the `id` primary keys.
	NormalizeIDs bool `json:"normalize_ids"`

	// An OIDC token can be used to authenticate with Azure instead of `AZURE_CLIENT_SECRET`.
	// This is useful for Azure AD workload identity federation.
	// When using this option, the `AZURE_CLIENT_ID` and `AZURE_TENANT_ID` environment variables must be set.
	OIDCToken string `json:"oidc_token" jsonschema:"minLength=1,example=oidc_token"`

	// The best effort maximum number of Go routines to use.
	// Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency" jsonschema:"minimum=1,default=50000"`

	// During initialization the Azure source plugin discovers all resource groups
	// and enabled resource providers per subscription, to be used later on during the sync process.
	// The plugin runs the discovery process in parallel.
	// This setting controls the maximum number of concurrent requests to the Azure API during discovery.
	// Only accounts with many subscriptions should require modifying this setting,
	// to either lower it to avoid network errors, or to increase it to speed up the discovery process.
	DiscoveryConcurrency int `json:"discovery_concurrency" jsonschema:"minimum=1,default=400"`

	// Retry options to pass to the Azure Go SDK, see more details
	// [here](https://github.com/Azure/azure-sdk-for-go/blob/f951bf52fb68cbb978b7b95d41147693c1863366/sdk/azcore/policy/policy.go#L86).
	RetryOptions *RetryOptions `json:"retry_options"`
}

var (
	specCloudToConfig = map[string]cloud.Configuration{
		"AzurePublic":     cloud.AzurePublic,
		"AzureGovernment": cloud.AzureGovernment,
		"AzureChina":      cloud.AzureChina,
	}
	// note: this should also be updated if new keys are added to specCloudToConfig
	specCloudToConfigKeys = []any{"AzurePublic", "AzureGovernment", "AzureChina"}
)

func (s *Spec) CloudConfig() (cloud.Configuration, error) {
	if v, ok := specCloudToConfig[s.CloudName]; ok {
		return v, nil
	}

	return cloud.Configuration{}, fmt.Errorf("unknown Azure cloud name %q. Supported values are %q", s.CloudName, maps.Keys(specCloudToConfig))
}

func (s *Spec) SetDefaults() {
	if s.DiscoveryConcurrency <= 0 {
		s.DiscoveryConcurrency = 400
	}
	if s.Concurrency <= 0 {
		const defaultConcurrency = 50000
		s.Concurrency = defaultConcurrency
	}
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	cloudName := sc.Properties.Value("cloud_name")
	cloudName.Enum = append(cloudName.Enum, specCloudToConfigKeys...)
}

//go:embed schema.json
var JSONSchema string
