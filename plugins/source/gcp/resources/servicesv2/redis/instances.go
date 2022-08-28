// Code generated by codegen; DO NOT EDIT.

package redis

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
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
				Name:     "alternative_location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlternativeLocationId"),
			},
			{
				Name:     "auth_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AuthEnabled"),
			},
			{
				Name:     "authorized_network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizedNetwork"),
			},
			{
				Name:     "connect_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectMode"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "current_location_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentLocationId"),
			},
			{
				Name:     "customer_managed_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerManagedKey"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "host",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Host"),
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
				Name:     "memory_size_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MemorySizeGb"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Nodes"),
			},
			{
				Name:     "persistence_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PersistenceConfig"),
			},
			{
				Name:     "persistence_iam_identity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PersistenceIamIdentity"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
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
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReadReplicasMode"),
			},
			{
				Name:     "redis_configs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RedisConfigs"),
			},
			{
				Name:     "redis_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RedisVersion"),
			},
			{
				Name:     "replica_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReplicaCount"),
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
				Name:     "server_ca_certs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServerCaCerts"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "suspension_reasons",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SuspensionReasons"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
			{
				Name:     "transit_encryption_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitEncryptionMode"),
			},
		},
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Redis.Projects.Locations.Instances.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Instances

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
