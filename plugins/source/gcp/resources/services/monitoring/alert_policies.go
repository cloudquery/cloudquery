// Code generated by codegen; DO NOT EDIT.

package monitoring

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	pb "google.golang.org/genproto/googleapis/monitoring/v3"
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
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Combiner"),
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

func fetchAlertPolicies(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListAlertPoliciesRequest{}
	it := c.Services.MonitoringAlertPolicyClient.ListAlertPolicies(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
