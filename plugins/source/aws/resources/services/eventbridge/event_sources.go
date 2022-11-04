// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventSources() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_event_sources",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventSource.html`,
		Resolver:    fetchEventbridgeEventSources,
		Multiplex:   client.ServiceAccountRegionMultiplexer("events"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedBy"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "expiration_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ExpirationTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
		},
	}
}
