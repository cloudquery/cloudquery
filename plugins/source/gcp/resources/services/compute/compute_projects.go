package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeProjects() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_projects",
		Description: "Represents a Project resource which is used to organize resources in a Google Cloud Platform environment",
		Resolver:    fetchComputeProjects,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "common_instance_metadata_fingerprint",
				Description: "Specifies a fingerprint for this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CommonInstanceMetadata.Fingerprint"),
			},
			{
				Name:        "common_instance_metadata_items",
				Description: "Array of key/value pairs The total size of all keys and values must be less than 512 KB",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeProjectCommonInstanceMetadataItems,
			},
			{
				Name:        "common_instance_metadata_kind",
				Description: "Type of the resource Always compute#metadata for metadata",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CommonInstanceMetadata.Kind"),
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("CreationTimestamp"),
			},
			{
				Name:        "default_network_tier",
				Description: "This signifies the default network tier used for configuring resources of the project and can only take the following values: PREMIUM, STANDARD Initially the default network tier is PREMIUM",
				Type:        schema.TypeString,
			},
			{
				Name:        "default_service_account",
				Description: "Default service account used by VMs running in this project",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional textual description of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:          "enabled_features",
				Description:   "Restricted features enabled for use on this project",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "compute_project_id",
				Description: "The unique identifier for the resource This identifier is defined by the server This is not the project ID, and is just a unique ID used by Compute Engine to identify resources",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#project for projects",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The project ID For example: my-example-project",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "usage_export_location_bucket_name",
				Description: "The name of an existing bucket in Cloud Storage where the usage report object is stored",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UsageExportLocation.BucketName"),
			},
			{
				Name:        "usage_export_location_report_name_prefix",
				Description: "An optional prefix for the name of the usage report object stored in bucketName",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UsageExportLocation.ReportNamePrefix"),
			},
			{
				Name:        "xpn_project_status",
				Description: "The role this project has in a shared VPC configuration Currently, only projects with the host role, which is specified by the value HOST, are differentiated",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_project_quotas",
				Description: "A quotas entry",
				Resolver:    fetchComputeProjectQuotas,
				Columns: []schema.Column{
					{
						Name:        "project_cq_id",
						Description: "Unique ID of gcp_compute_projects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "project_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("project_id"),
					},
					{
						Name:        "limit",
						Description: "Quota limit for this metric",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "metric",
						Description: "Name of the quota metric",
						Type:        schema.TypeString,
					},
					{
						Name:        "owner",
						Description: "Owning resource This is the resource on which this quota is applied",
						Type:        schema.TypeString,
					},
					{
						Name:        "usage",
						Description: "Current usage of this metric",
						Type:        schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	output, err := c.Services.Compute.Projects.Get(c.ProjectId).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	res <- output
	return nil
}
func resolveComputeProjectCommonInstanceMetadataItems(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*compute.Project)
	m := make(map[string]interface{})
	for _, i := range p.CommonInstanceMetadata.Items {
		m[i.Key] = i.Value
	}
	return errors.WithStack(resource.Set(c.Name, m))
}
func fetchComputeProjectQuotas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*compute.Project)
	res <- p.Quotas
	return nil
}
