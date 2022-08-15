package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugins/source/gcp/resources/services/bigquery"
	"github.com/cloudquery/plugins/source/gcp/resources/services/cloudbilling"
	"github.com/cloudquery/plugins/source/gcp/resources/services/cloudfunctions"
	"github.com/cloudquery/plugins/source/gcp/resources/services/cloudrun"
	"github.com/cloudquery/plugins/source/gcp/resources/services/compute"
	"github.com/cloudquery/plugins/source/gcp/resources/services/dns"
	"github.com/cloudquery/plugins/source/gcp/resources/services/domains"
	"github.com/cloudquery/plugins/source/gcp/resources/services/iam"
	"github.com/cloudquery/plugins/source/gcp/resources/services/kms"
	"github.com/cloudquery/plugins/source/gcp/resources/services/kubernetes"
	"github.com/cloudquery/plugins/source/gcp/resources/services/logging"
	"github.com/cloudquery/plugins/source/gcp/resources/services/memorystore"
	"github.com/cloudquery/plugins/source/gcp/resources/services/monitoring"
	"github.com/cloudquery/plugins/source/gcp/resources/services/resource_manager"
	"github.com/cloudquery/plugins/source/gcp/resources/services/security"
	"github.com/cloudquery/plugins/source/gcp/resources/services/sql"
	"github.com/cloudquery/plugins/source/gcp/resources/services/storage"
)

var (
	Version = "development"
)

const exampleConfig = `
# Optional. List of folders to get projects from. Required permission: resourcemanager.projects.list
# folder_ids:
# 	- "organizations/<ORG_ID>"
# 	- "folders/<FOLDER_ID>"
# Optional. Maximum level of folders to recurse into
# folders_max_depth: 5
# Optional. If not specified either using all projects accessible.
# project_ids:
# 	- "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"
# Optional. ServiceAccountKeyJSON passed as value instead of a file path, can be passed also via env: CQ_SERVICE_ACCOUNT_KEY_JSON
# service_account_key_json: <YOUR_JSON_SERVICE_ACCOUNT_KEY_DATA>
# Optional. GRPC Retry/backoff configuration, time units in seconds. Documented in https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md
# backoff_base_delay: 1
# backoff_multiplier: 1.6
# backoff_max_delay: 120
# backoff_jitter: 0.2
# backoff_min_connect_timeout = 0
# Optional. Max amount of retries for retrier, defaults to max 3 retries.
# max_retries: 3
`

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"gcp",
		Version,
		[]*schema.Table{
			bigquery.BigqueryDatasets(),
			cloudbilling.Accounts(),
			cloudbilling.Services(),
			cloudfunctions.CloudfunctionsFunction(),
			cloudrun.Services(),
			compute.ComputeAddresses(),
			compute.ComputeAutoscalers(),
			compute.ComputeBackendServices(),
			compute.ComputeDiskTypes(),
			compute.ComputeDisks(),
			compute.ComputeFirewalls(),
			compute.ComputeForwardingRules(),
			compute.ComputeImages(),
			compute.InstanceGroups(),
			compute.ComputeInstances(),
			compute.ComputeInterconnects(),
			compute.ComputeNetworks(),
			compute.ComputeProjects(),
			compute.ComputeSslCertificates(),
			compute.ComputeSslPolicies(),
			compute.ComputeSubnetworks(),
			compute.ComputeTargetHTTPProxies(),
			compute.ComputeTargetHTTPSProxies(),
			compute.ComputeTargetSslProxies(),
			compute.ComputeURLMaps(),
			compute.ComputeVpnGateways(),
			dns.DNSManagedZones(),
			dns.DNSPolicies(),
			domains.DomainsRegistration(),
			iam.IamRoles(),
			iam.IamServiceAccounts(),
			kms.KmsKeyrings(),
			kubernetes.Clusters(),
			logging.LoggingMetrics(),
			logging.LoggingSinks(),
			memorystore.RedisInstances(),
			monitoring.MonitoringAlertPolicies(),
			resource_manager.ResourceManagerFolders(),
			resource_manager.ResourceManagerProjects(),
			sql.SQLInstances(),
			security.Secrets(),
			storage.StorageBuckets(),
			storage.Metrics(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
		plugins.WithClassifyError(client.ClassifyError),
	)
}
