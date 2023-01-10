package monitoring

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
)

func AlertPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_monitoring_alert_policies",
		Description: `https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies#AlertPolicy`,
		Resolver:    fetchAlertPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("monitoring.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.AlertPolicy{}, client.Options()...),
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
		},
	}
}

func fetchAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListAlertPoliciesRequest{
		Name: "projects/" + c.ProjectId,
	}
	gcpClient, err := monitoring.NewAlertPolicyClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListAlertPolicies(ctx, req, c.CallOptions...)
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
