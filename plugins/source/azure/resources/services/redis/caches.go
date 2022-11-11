// Auto generated code - DO NOT EDIT.

package redis

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Caches() *schema.Table {
	return &schema.Table{
		Name:        "azure_redis_caches",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis#ResourceType`,
		Resolver:    fetchRedisCaches,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostName"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "ssl_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SslPort"),
			},
			{
				Name:     "access_keys",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessKeys"),
			},
			{
				Name:     "linked_servers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedServers"),
			},
			{
				Name:     "instances",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Instances"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetID"),
			},
			{
				Name:     "static_ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StaticIP"),
			},
			{
				Name:     "redis_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RedisConfiguration"),
			},
			{
				Name:     "redis_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RedisVersion"),
			},
			{
				Name:     "enable_non_ssl_port",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableNonSslPort"),
			},
			{
				Name:     "replicas_per_master",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReplicasPerMaster"),
			},
			{
				Name:     "replicas_per_primary",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReplicasPerPrimary"),
			},
			{
				Name:     "tenant_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TenantSettings"),
			},
			{
				Name:     "shard_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ShardCount"),
			},
			{
				Name:     "minimum_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimumTLSVersion"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicNetworkAccess"),
			},
			{
				Name:     "zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Zones"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
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
	}
}

func fetchRedisCaches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Redis.Caches

	response, err := svc.ListBySubscription(ctx)

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
