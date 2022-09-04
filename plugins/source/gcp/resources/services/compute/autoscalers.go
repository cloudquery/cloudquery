// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func Autoscalers() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_autoscalers",
		Resolver:  fetchAutoscalers,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "autoscaling_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoscalingPolicy"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "recommended_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RecommendedSize"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "scaling_schedule_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ScalingScheduleStatus"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StatusDetails"),
			},
			{
				Name:     "target",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Target"),
			},
			{
				Name:     "zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Zone"),
			},
		},
	}
}

func fetchAutoscalers(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListAutoscalersRequest{}
	it := c.Services.ComputeAutoscalersClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.Value.Autoscalers

	}
	return nil
}
