package plugin

import (
	"context"
	"encoding/json"
	"fmt"
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
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
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
	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	for k, v := range googleAdsExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

type Client struct {
	plugin.UnimplementedDestination
	schduler   *scheduler.Scheduler
	syncClient *client.Client
	options    plugin.NewClientOptions
}

func newClient(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		options: options,
	}
	if options.NoConnection {
		return c, nil
	}
	spec := &client.Spec{}
	if err := json.Unmarshal(specBytes, spec); err != nil {
		return nil, err
	}
	spec.SetDefaults()
	syncClient, err := client.Configure(ctx, logger, *spec)
	if err != nil {
		return nil, err
	}
	c.syncClient = syncClient.(*client.Client)
	c.schduler = scheduler.NewScheduler(scheduler.WithLogger(logger), scheduler.WithConcurrency(uint64(spec.Concurrency)))
	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (*Client) Tables(ctx context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tables := getTables()
	tables, err := tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	tables := getTables()
	tables, err := tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}
	return c.schduler.Sync(ctx, c.syncClient, tables, res)
}

func getTables() schema.Tables {
	tables := []*schema.Table{
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
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
		for _, rel := range table.Relations {
			schema.AddCqIDs(rel)
		}
	}
	return tables
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"k8s",
		Version,
		newClient,
	)
}
