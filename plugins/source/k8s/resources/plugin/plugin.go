package plugin

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/admissionregistration"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/apps"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/certificates"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/coordination"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/core"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/crd"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/discovery"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/networking"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/nodes"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/policy"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/rbac"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/storage"
	"github.com/cloudquery/plugin-sdk/v2/caser"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"golang.org/x/exp/maps"
)

var Version = "Development"

var googleAdsExceptions = map[string]string{
	"admissionregistration": "Admission Registration",
	"crds":                  "Custom Resource Definitions (CRDs)",
	"csi":                   "Container Storage Interface (CSI)",
	"hpas":                  "Horizontal Pod Autoscalers (HPAs)",
	"k8s":                   "Kubernetes (K8s)",
	"pvcs":                  "Persistent Volume Claims (PVCs)",
	"pvs":                   "Persistent Volumes (PVs)",
	"rbac":                  "Role-Based Access Control (RBAC)",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := maps.Clone(source.DefaultTitleExceptions)
	for k, v := range googleAdsExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"k8s",
		Version,
		[]*schema.Table{
			discovery.EndpointSlices(),
			admissionregistration.MutatingWebhookConfigurations(),
			admissionregistration.ValidatingWebhookConfigurations(),
			apps.DaemonSets(),
			apps.Deployments(),
			apps.ReplicaSets(),
			apps.StatefulSets(),
			autoscaling.Hpas(),
			batch.Jobs(),
			batch.CronJobs(),
			certificates.SigningRequests(),
			coordination.Leases(),
			core.ComponentStatuses(),
			core.ConfigMaps(),
			core.Endpoints(),
			core.Events(),
			core.LimitRanges(),
			core.Namespaces(),
			core.Nodes(),
			core.Pvs(),
			core.Pvcs(),
			core.Pods(),
			core.ReplicationControllers(),
			core.ResourceQuotas(),
			core.Secrets(),
			core.Services(),
			core.ServiceAccounts(),
			crd.CRDs(),
			networking.Ingresses(),
			networking.NetworkPolicies(),
			networking.IngressClasses(),
			nodes.RuntimeClasses(),
			rbac.ClusterRoles(),
			rbac.ClusterRoleBindings(),
			rbac.Roles(),
			rbac.RoleBindings(),
			policy.PodDisruptionBudgets(),
			storage.CsiDrivers(),
			storage.CsiNodes(),
			storage.CsiStorageCapacities(),
			storage.StorageClasses(),
			storage.VolumeAttachments(),
		},
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
