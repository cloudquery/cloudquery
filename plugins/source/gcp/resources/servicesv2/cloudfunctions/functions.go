// Code generated by codegen; DO NOT EDIT.

package cloudfunctions

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Functions() *schema.Table {
	return &schema.Table{
		Name:      "gcp_cloudfunctions_functions",
		Resolver:  fetchFunctions,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "available_memory_mb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AvailableMemoryMb"),
			},
			{
				Name:     "build_environment_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BuildEnvironmentVariables"),
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
				Name:     "build_worker_pool",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BuildWorkerPool"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "docker_registry",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DockerRegistry"),
			},
			{
				Name:     "docker_repository",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DockerRepository"),
			},
			{
				Name:     "entry_point",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EntryPoint"),
			},
			{
				Name:     "environment_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnvironmentVariables"),
			},
			{
				Name:     "event_trigger",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EventTrigger"),
			},
			{
				Name:     "https_trigger",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HttpsTrigger"),
			},
			{
				Name:     "ingress_settings",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IngressSettings"),
			},
			{
				Name:     "kms_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyName"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
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
				Name: "name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "runtime",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Runtime"),
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
				Name:     "service_account_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccountEmail"),
			},
			{
				Name:     "source_archive_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceArchiveUrl"),
			},
			{
				Name:     "source_repository",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceRepository"),
			},
			{
				Name:     "source_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceToken"),
			},
			{
				Name:     "source_upload_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceUploadUrl"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "timeout",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Timeout"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UpdateTime"),
			},
			{
				Name:     "version_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VersionId"),
			},
			{
				Name:     "vpc_connector",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcConnector"),
			},
			{
				Name:     "vpc_connector_egress_settings",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcConnectorEgressSettings"),
			},
		},
	}
}

func fetchFunctions(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Cloudfunctions.Projects.Locations.Functions.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Functions

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
