package redis

import (
	"context"
	"encoding/json"
	"net"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RedisServices() *schema.Table {
	return &schema.Table{
		Name:         "azure_redis_services",
		Description:  "Azure Redis service",
		Resolver:     fetchRedisServices,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioning_state",
				Description: "Redis instance provisioning status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "hostname",
				Description: "Redis host name.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HostName"),
			},
			{
				Name:        "port",
				Description: "Redis non-SSL port.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "ssl_port",
				Description: "Redis SSL port.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "linked_server_ids",
				Description: "List of the linked servers associated with the cache.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveServiceLinkedServers,
			},
			{
				Name:        "instances",
				Description: "List of the Redis instances associated with the cache.",
				Type:        schema.TypeJSON,
				Resolver: resolveServiceJSONField(func(s redis.ResourceType) interface{} {
					if s.Instances == nil {
						return nil
					}
					items := make([]interface{}, len(*s.Instances))
					for i, v := range *s.Instances {
						items[i] = map[string]interface{}{
							"sslPort":    v.SslPort,
							"nonSslPort": v.NonSslPort,
							"zone":       v.Zone,
							"shardId":    v.ShardID,
							"isMaster":   v.IsMaster,
							"isPrimary":  v.IsPrimary,
						}
					}
					return items
				}),
			},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connection associated with the specified redis cache.",
				Type:        schema.TypeJSON,
				Resolver: resolveServiceJSONField(func(s redis.ResourceType) interface{} {
					if s.PrivateEndpointConnections == nil {
						return nil
					}
					items := make([]interface{}, len(*s.PrivateEndpointConnections))
					for i, v := range *s.PrivateEndpointConnections {
						m := map[string]interface{}{
							"id":                                v.ID,
							"name":                              v.Name,
							"type":                              v.Type,
							"privateLinkServiceConnectionState": v.PrivateLinkServiceConnectionState,
							"provisioningState":                 v.ProvisioningState,
						}
						if v.PrivateEndpoint != nil {
							m["privateEndpoint"] = map[string]interface{}{"id": v.PrivateEndpoint.ID}
						}
						items[i] = m
					}
					return items
				}),
				IgnoreInTests: true,
			},
			{
				Name:        "sku_name",
				Description: "The type of Redis cache to deploy.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_family",
				Description: "The SKU family to use.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "sku_capacity",
				Description: "The size of the Redis cache to deploy.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:          "subnet_id",
				Description:   "The full resource ID of a subnet in a virtual network to deploy the Redis cache in.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("SubnetID"),
				IgnoreInTests: true,
			},
			{
				Name:          "static_ip",
				Description:   "Static IP address.",
				Type:          schema.TypeInet,
				Resolver:      resolveServiceStaticIP,
				IgnoreInTests: true,
			},
			{
				Name:        "configuration",
				Description: "All Redis Settings.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("RedisConfiguration"),
			},
			{
				Name:        "version",
				Description: "Redis version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RedisVersion"),
			},
			{
				Name:        "enable_non_ssl_port",
				Description: "Specifies whether the non-ssl Redis server port (6379) is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EnableNonSslPort"),
			},
			{
				Name:          "replicas_per_master",
				Description:   "The number of replicas to be created per primary.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:          "replicas_per_primary",
				Description:   "The number of replicas to be created per primary.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:          "tenant_settings",
				Description:   "A dictionary of tenant settings.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "shard_count",
				Description:   "The number of shards to be created on a Premium Cluster Cache.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:        "minimum_tls_version",
				Description: "Requires clients to use a specified TLS version (or higher) to connect.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MinimumTLSVersion"),
			},
			{
				Name:        "public_network_access",
				Description: "Whether or not public endpoint access is allowed for this cache.",
				Type:        schema.TypeBool,
				Resolver:    resolveServicePublicNetworkAccess,
			},
		},
	}
}

func fetchRedisServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Redis
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

func resolveServiceLinkedServers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(redis.ResourceType)
	if s.LinkedServers == nil {
		return nil
	}
	ids := make([]string, 0, len(*s.LinkedServers))
	for _, v := range *s.LinkedServers {
		if v.ID != nil {
			ids = append(ids, *v.ID)
		}
	}
	return resource.Set(c.Name, ids)
}

func resolveServiceJSONField(getter func(s redis.ResourceType) interface{}) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		s := resource.Item.(redis.ResourceType)
		b, err := json.Marshal(getter(s))
		if err != nil {
			return err
		}
		return resource.Set(c.Name, b)
	}
}

func resolveServiceStaticIP(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(redis.ResourceType)
	if s.StaticIP == nil {
		return nil
	}
	ip := net.ParseIP(*s.StaticIP)
	if ip == nil {
		return nil
	}
	if v4 := ip.To4(); v4 != nil {
		ip = v4
	}
	return resource.Set(c.Name, ip)
}

func resolveServicePublicNetworkAccess(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(redis.ResourceType)
	return resource.Set(c.Name, s.PublicNetworkAccess == redis.PublicNetworkAccessEnabled)
}
