// Code generated by codegen; DO NOT EDIT.

package redis

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	pb "google.golang.org/genproto/googleapis/cloud/redis/v1"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "gcp_redis_instances",
		Resolver:  fetchInstances,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocationId"),
			},
			{
				Name:     "alternative_location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlternativeLocationId"),
			},
			{
				Name:     "redis_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RedisVersion"),
			},
			{
				Name:     "reserved_ip_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservedIpRange"),
			},
			{
				Name:     "secondary_ip_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryIpRange"),
			},
			{
				Name:     "host",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Host"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "current_location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentLocationId"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "state",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "redis_configs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RedisConfigs"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Tier"),
			},
			{
				Name:     "memory_size_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MemorySizeGb"),
			},
			{
				Name:     "authorized_network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizedNetwork"),
			},
			{
				Name:     "persistence_iam_identity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PersistenceIamIdentity"),
			},
			{
				Name:     "connect_mode",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ConnectMode"),
			},
			{
				Name:     "auth_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AuthEnabled"),
			},
			{
				Name:     "server_ca_certs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServerCaCerts"),
			},
			{
				Name:     "transit_encryption_mode",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TransitEncryptionMode"),
			},
			{
				Name:     "maintenance_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaintenancePolicy"),
			},
			{
				Name:     "maintenance_schedule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaintenanceSchedule"),
			},
			{
				Name:     "replica_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReplicaCount"),
			},
			{
				Name:     "nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Nodes"),
			},
			{
				Name:     "read_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReadEndpoint"),
			},
			{
				Name:     "read_endpoint_port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReadEndpointPort"),
			},
			{
				Name:     "read_replicas_mode",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReadReplicasMode"),
			},
		},
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListInstancesRequest{}
	it := c.Services.RedisCloudRedisClient.ListInstances(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
