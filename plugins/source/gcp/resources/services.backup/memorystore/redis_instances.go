package memorystore

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/redis/v1"
)

//go:generate cq-gen --resource redis_instances --config gen.hcl --output .
func RedisInstances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_memorystore_redis_instances",
		Description: "A Memorystore for Redis instance.",
		Resolver:    fetchMemorystoreRedisInstances,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "id",
				Description: "Memorystore for Redis instance ID",
				Type:        schema.TypeString,
				Resolver:    ResolveMemorystoreRedisInstanceID,
			},
			{
				Name:        "alternative_location_id",
				Description: "If specified, at least one node will be provisioned in this zone in addition to the zone specified in `location_id`",
				Type:        schema.TypeString,
			},
			{
				Name:        "auth_enabled",
				Description: "Indicates whether OSS Redis AUTH is enabled for the instance",
				Type:        schema.TypeBool,
			},
			{
				Name:        "authorized_network",
				Description: "The full name of the Google Compute Engine network (https://cloud.google.com/vpc/docs/vpc) to which the instance is connected",
				Type:        schema.TypeString,
			},
			{
				Name:        "connect_mode",
				Description: "The network connect mode of the Redis instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_time",
				Description: "The time the instance was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "current_location_id",
				Description: "The current zone where the Redis primary node is located",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_managed_key",
				Description: "The KMS key reference that the customer provides when trying to create the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "display_name",
				Description: "An arbitrary and optional user-provided name for the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "host",
				Description: "Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "Resource labels to represent user provided metadata",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location_id",
				Description: "The zone where the instance will be provisioned",
				Type:        schema.TypeString,
			},
			{
				Name:          "maintenance_policy",
				Description:   "The maintenance policy for the instance",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "maintenance_schedule",
				Description:   "Date and time of upcoming maintenance events which have been scheduled",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "memory_size_gb",
				Description: "Redis memory size in GiB",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "Unique name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "nodes",
				Description: "Redis instance nodes properties",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "persistence_config",
				Description: "Persistence configuration parameters",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "persistence_iam_identity",
				Description: "Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage",
				Type:        schema.TypeString,
			},
			{
				Name:        "port",
				Description: "The port number of the exposed Redis endpoints",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "read_endpoint",
				Description: "Hostname or IP address of the exposed readonly Redis endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "read_endpoint_port",
				Description: "The port number of the exposed readonly redis endpoint",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "read_replicas_mode",
				Description: "Read replicas mode for the instance",
				Type:        schema.TypeString,
			},
			{
				Name:          "redis_configs",
				Description:   "Redis configuration parameters",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "redis_version",
				Description: "The version of Redis software",
				Type:        schema.TypeString,
			},
			{
				Name:        "replica_count",
				Description: "The number of replica nodes",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "reserved_ip_range",
				Description: "IP range for node placement",
				Type:        schema.TypeString,
			},
			{
				Name:        "secondary_ip_range",
				Description: "Additional IP range for node placement",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_message",
				Description: "Additional information about the current status of the instance, if available",
				Type:        schema.TypeString,
			},
			{
				Name:          "suspension_reasons",
				Description:   "Reasons that causes instance in `SUSPENDED` state",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "tier",
				Description: "The service tier of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "transit_encryption_mode",
				Description: "The TLS mode of the Redis instance",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_memorystore_redis_instance_server_ca_certs",
				Description: "List of server CA certificates for the instance",
				Resolver:    fetchMemorystoreRedisInstanceRedisInstanceServerCaCerts,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"sha1_fingerprint"}},
				Columns: []schema.Column{
					{
						Name:        "redis_instance_cq_id",
						Description: "Unique CloudQuery ID of gcp_memorystore_redis_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cert",
						Description: "PEM representation",
						Type:        schema.TypeString,
					},
					{
						Name:        "create_time",
						Description: "The time when the certificate was created in RFC 3339 format",
						Type:        schema.TypeString,
					},
					{
						Name:        "expire_time",
						Description: "The time when the certificate expires in RFC 3339 format",
						Type:        schema.TypeString,
					},
					{
						Name:        "serial_number",
						Description: "Serial number, as extracted from the certificate",
						Type:        schema.TypeString,
					},
					{
						Name:        "sha1_fingerprint",
						Description: "Sha1 Fingerprint of the certificate",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchMemorystoreRedisInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func ResolveMemorystoreRedisInstanceID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	// name is in the form projects/{project_id}/locations/{location_id}/instances/{instance_id}
	parts := strings.Split(resource.Item.(*redis.Instance).Name, "/")
	if len(parts) != 6 {
		return errors.WithStack(
			fmt.Errorf(
				"name of Redis instance (%q) not in the form `projects/{project_id}/locations/{location_id}/instances/{instance_id}`",
				resource.Item.(*redis.Instance).Name,
			),
		)
	}
	return errors.WithStack(resource.Set("id", parts[5]))
}
func fetchMemorystoreRedisInstanceRedisInstanceServerCaCerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(*redis.Instance)
	res <- instance.ServerCaCerts
	return nil
}
