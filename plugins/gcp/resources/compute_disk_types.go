package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	compute "google.golang.org/api/compute/v1"
)

func ComputeDiskTypes() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_disk_types",
		Description: "Represents a Disk Type resource.",
		Resolver:    fetchComputeDiskTypes,
		Multiplex:   client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "default_disk_size_gb",
				Description: "Server-defined default disk size in GB",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "deprecated_deleted",
				Description: "An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DELETED This is only informational and the status will not change unless the client explicitly changes it",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.Deleted"),
			},
			{
				Name:     "deprecated",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Deprecated"),
			},
			{
				Name:        "deprecated_obsolete",
				Description: "An optional RFC3339 timestamp on or after which the state of this resource is intended to change to OBSOLETE This is only informational and the status will not change unless the client explicitly changes it",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.Obsolete"),
			},
			{
				Name:        "deprecated_replacement",
				Description: "The URL of the suggested replacement for a deprecated resource The suggested replacement resource must be the same kind of resource as the deprecated resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.Replacement"),
			},
			{
				Name:        "deprecated_state",
				Description: "The deprecation state of this resource This can be ACTIVE, DEPRECATED, OBSOLETE, or DELETED Operations which communicate the end of life date for an image, can use ACTIVE Operations which create a new resource using a DEPRECATED resource will return successfully, but with a warning indicating the deprecated resource and recommending its replacement Operations which use OBSOLETE or DELETED resources will be rejected and result in an error",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Deprecated.State"),
			},
			{
				Name:        "description",
				Description: "An optional description of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#diskType for disk types",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "URL of the region where the disk type resides Only applicable for regional resources You must specify this field as part of the HTTP request URL It is not settable as a field in the request body",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "valid_disk_size",
				Description: "An optional textual description of the valid disk size, such as \"10GB-10TB\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "zone",
				Description: "URL of the zone where the disk type resides You must specify this field as part of the HTTP request URL It is not settable as a field in the request body",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeDiskTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.DiskTypes.AggregatedList(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var diskTypes []*compute.DiskType
		for _, items := range output.Items {
			diskTypes = append(diskTypes, items.DiskTypes...)
		}
		res <- diskTypes

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
