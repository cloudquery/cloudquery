package provider

import (
	"embed"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/resources/services/apps"
	"github.com/cloudquery/cq-provider-k8s/resources/services/batch"
	"github.com/cloudquery/cq-provider-k8s/resources/services/core"
	"github.com/cloudquery/cq-provider-k8s/resources/services/networking"
	"github.com/cloudquery/cq-provider-k8s/resources/services/rbac"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const ProviderName = "k8s"

var (
	//go:embed migrations/*/*.sql
	migrationFiles embed.FS

	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:         Version,
		Name:            ProviderName,
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		Config: func() provider.Config {
			return &client.Config{}
		},
		ResourceMap: map[string]*schema.Table{
			"apps.daemon_sets":            apps.DaemonSets(),
			"apps.deployments":            apps.Deployments(),
			"apps.replica_sets":           apps.ReplicaSets(),
			"apps.stateful_sets":          apps.StatefulSets(),
			"batch.cron_jobs":             batch.CronJobs(),
			"batch.jobs":                  batch.Jobs(),
			"core.endpoints":              core.Endpoints(),
			"core.limit_ranges":           core.LimitRanges(),
			"core.namespaces":             core.Namespaces(),
			"core.nodes":                  core.Nodes(),
			"core.pods":                   core.Pods(),
			"core.resource_quotas":        core.ResourceQuotas(),
			"core.service_accounts":       core.ServiceAccounts(),
			"core.services":               core.Services(),
			"networking.network_policies": networking.NetworkPolicies(),
			"rbac.role_bindings":          rbac.RoleBindings(),
			"rbac.roles":                  rbac.Roles(),
		},
		Migrations: migrationFiles,
	}
}
