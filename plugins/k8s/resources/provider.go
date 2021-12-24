package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const ProviderName = "k8s"

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:   Version,
		Name:      ProviderName,
		Configure: client.Configure,
		Config: func() provider.Config {
			return &client.Config{}
		},
		ResourceMap: map[string]*schema.Table{
			"apps.daemon_sets":            AppsDaemonSets(),
			"apps.deployments":            AppsDeployments(),
			"apps.replica_sets":           AppsReplicaSets(),
			"apps.stateful_sets":          AppsStatefulSets(),
			"batch.cron_jobs":             BatchCronJobs(),
			"batch.jobs":                  BatchJobs(),
			"core.endpoints":              CoreEndpoints(),
			"core.limit_ranges":           CoreLimitRanges(),
			"core.namespaces":             CoreNamespaces(),
			"core.nodes":                  CoreNodes(),
			"core.pods":                   CorePods(),
			"core.resource_quotas":        CoreResourceQuotas(),
			"core.service_accounts":       CoreServiceAccounts(),
			"core.services":               CoreServices(),
			"networking.network_policies": NetworkingNetworkPolicies(),
			"rbac.role_bindings":          RbacRoleBindings(),
			"rbac.roles":                  RbacRoles(),
		},
	}
}
