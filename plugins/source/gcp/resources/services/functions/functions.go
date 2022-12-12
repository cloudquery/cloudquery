// Code generated by codegen; DO NOT EDIT.

package functions

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/functions/apiv1"
)

func Functions() *schema.Table {
	return &schema.Table{
		Name:      "gcp_functions_functions",
		Resolver:  fetchFunctions,
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
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("Status"),
			},
			{
				Name:     "entry_point",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EntryPoint"),
			},
			{
				Name:     "runtime",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Runtime"),
			},
			{
				Name:     "timeout",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProtoDuration("Timeout"),
			},
			{
				Name:     "available_memory_mb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AvailableMemoryMb"),
			},
			{
				Name:     "service_account_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccountEmail"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "version_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VersionId"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "environment_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnvironmentVariables"),
			},
			{
				Name:     "build_environment_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BuildEnvironmentVariables"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "max_instances",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxInstances"),
			},
			{
				Name:     "min_instances",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinInstances"),
			},
			{
				Name:     "vpc_connector",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcConnector"),
			},
			{
				Name:     "vpc_connector_egress_settings",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("VpcConnectorEgressSettings"),
			},
			{
				Name:     "ingress_settings",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("IngressSettings"),
			},
			{
				Name:     "kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyName"),
			},
			{
				Name:     "build_worker_pool",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BuildWorkerPool"),
			},
			{
				Name:     "build_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BuildId"),
			},
			{
				Name:     "build_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BuildName"),
			},
			{
				Name:     "secret_environment_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecretEnvironmentVariables"),
			},
			{
				Name:     "secret_volumes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecretVolumes"),
			},
			{
				Name:     "source_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceToken"),
			},
			{
				Name:     "docker_repository",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DockerRepository"),
			},
			{
				Name:     "docker_registry",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("DockerRegistry"),
			},
		},
	}
}

func fetchFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListFunctionsRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := functions.NewCloudFunctionsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListFunctions(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
