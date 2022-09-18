// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventBuses() *schema.Table {
	return &schema.Table{
		Name:      "aws_eventbridge_event_buses",
		Resolver:  fetchEventbridgeEventBuses,
		Multiplex: client.ServiceAccountRegionMultiplexer("events"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEventbridgeEventBusTags,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Policy"),
			},
		},

		Relations: []*schema.Table{
			EventBusRules(),
		},
	}
}
