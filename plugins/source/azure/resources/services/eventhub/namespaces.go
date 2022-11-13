// Auto generated code - DO NOT EDIT.

package eventhub

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:        "azure_eventhub_namespaces",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub#EHNamespace`,
		Resolver:    fetchEventHubNamespaces,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "service_bus_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceBusEndpoint"),
			},
			{
				Name:     "cluster_arm_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArmID"),
			},
			{
				Name:     "metric_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricID"),
			},
			{
				Name:     "is_auto_inflate_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsAutoInflateEnabled"),
			},
			{
				Name:     "maximum_throughput_units",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaximumThroughputUnits"),
			},
			{
				Name:     "kafka_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("KafkaEnabled"),
			},
			{
				Name:     "zone_redundant",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ZoneRedundant"),
			},
			{
				Name:     "encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Encryption"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			networkRuleSets(),
		},
	}
}

func fetchEventHubNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().EventHub.Namespaces

	response, err := svc.List(ctx)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
