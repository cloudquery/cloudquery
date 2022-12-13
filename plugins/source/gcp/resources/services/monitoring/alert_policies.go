// Code generated by codegen; DO NOT EDIT.

package monitoring

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/monitoring/v3"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/monitoring/apiv3/v2"
)

func AlertPolicies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_monitoring_alert_policies",
		Resolver:  fetchAlertPolicies,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "documentation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Documentation"),
			},
			{
				Name:     "user_labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserLabels"),
			},
			{
				Name:     "conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Conditions"),
			},
			{
				Name:     "combiner",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("Combiner"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "validity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Validity"),
			},
			{
				Name:     "notification_channels",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NotificationChannels"),
			},
			{
				Name:     "creation_record",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreationRecord"),
			},
			{
				Name:     "mutation_record",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MutationRecord"),
			},
			{
				Name:     "alert_strategy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AlertStrategy"),
			},
		},
	}
}

func fetchAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
