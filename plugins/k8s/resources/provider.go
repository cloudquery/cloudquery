package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const ProviderName = "k8s"

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      ProviderName,
		Configure: client.Configure,
		Config: func() provider.Config {
			return &client.Config{}
		},
		ResourceMap: map[string]*schema.Table{

			"core.namespaces":    CoreNamespaces(),
			"core.nodes":         CoreNodes(),
			"core.pods":          CorePods(),
			"core.services":      CoreServices(),
			"apps.stateful_sets": AppsStatefulSets(),
			"apps.replica_sets":  AppsReplicaSets(),
			"rbac.roles":         RbacRoles(),
			"rbac.role_bindings": RbacRoleBindings(),
			"apps.daemon_sets":   AppsDaemonSets(),
		},
	}
}
