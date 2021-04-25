package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CloudfunctionsFunction() *schema.Table {
	return &schema.Table{
		Name:         "gcp_cloudfunctions_functions",
		Resolver:     fetchCloudfunctionsFunctions,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "available_memory_mb",
				Type: schema.TypeBigInt,
			},
			{
				Name: "build_environment_variables",
				Type: schema.TypeJSON,
			},
			{
				Name: "build_id",
				Type: schema.TypeString,
			},
			{
				Name: "build_worker_pool",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "entry_point",
				Type: schema.TypeString,
			},
			{
				Name: "environment_variables",
				Type: schema.TypeJSON,
			},
			{
				Name:     "event_trigger_event_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventTrigger.EventType"),
			},
			{
				Name:     "event_trigger_resource",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventTrigger.Resource"),
			},
			{
				Name:     "event_trigger_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventTrigger.Service"),
			},
			{
				Name:     "https_trigger_security_level",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HttpsTrigger.SecurityLevel"),
			},
			{
				Name:     "https_trigger_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HttpsTrigger.Url"),
			},
			{
				Name: "ingress_settings",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "max_instances",
				Type: schema.TypeBigInt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "network",
				Type: schema.TypeString,
			},
			{
				Name: "runtime",
				Type: schema.TypeString,
			},
			{
				Name: "service_account_email",
				Type: schema.TypeString,
			},
			{
				Name: "source_archive_url",
				Type: schema.TypeString,
			},
			{
				Name:     "source_repository_deployed_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceRepository.DeployedUrl"),
			},
			{
				Name:     "source_repository_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceRepository.Url"),
			},
			{
				Name: "source_token",
				Type: schema.TypeString,
			},
			{
				Name: "source_upload_url",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "timeout",
				Type: schema.TypeString,
			},
			{
				Name: "update_time",
				Type: schema.TypeString,
			},
			{
				Name: "version_id",
				Type: schema.TypeBigInt,
			},
			{
				Name: "vpc_connector",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_connector_egress_settings",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudfunctionsFunctions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.CloudFunctions.Projects.Locations.Functions.List("projects/" + c.ProjectId + "/locations/-").Context(ctx).PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}
		res <- output.Functions
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
