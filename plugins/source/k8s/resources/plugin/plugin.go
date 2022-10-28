package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/apps"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/core"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/networking"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/rbac"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"k8s",
		Version,
		[]*schema.Table{
			apps.DaemonSets(),
			apps.Deployments(),
			apps.ReplicaSets(),
			apps.StatefulSets(),
			batch.CronJobs(),
			batch.Jobs(),
			core.Endpoints(),
			core.LimitRanges(),
			core.Namespaces(),
			core.Nodes(),
			core.Pods(),
			core.ResourceQuotas(),
			core.Secrets(),
			core.ServiceAccounts(),
			core.Services(),
			networking.NetworkPolicies(),
			rbac.RoleBindings(),
			rbac.Roles(),
		},
		client.Configure,
	)
}
