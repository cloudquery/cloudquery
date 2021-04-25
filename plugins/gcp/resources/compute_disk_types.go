package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeDiskTypes() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_disk_types",
		Resolver:     fetchComputeDiskTypes,
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
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "default_disk_size_gb",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "deprecated_deleted",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Deleted"),
			},
			{
				Name:     "deprecated",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Deprecated"),
			},
			{
				Name:     "deprecated_obsolete",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Obsolete"),
			},
			{
				Name:     "deprecated_replacement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.Replacement"),
			},
			{
				Name:     "deprecated_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Deprecated.State"),
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "valid_disk_size",
				Type: schema.TypeString,
			},
			{
				Name: "zone",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeDiskTypes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
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
