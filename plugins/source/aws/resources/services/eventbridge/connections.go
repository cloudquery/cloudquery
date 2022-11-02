// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Connections() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_connections",
		Description: "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Connection.html",
		Resolver:    fetchEventbridgeConnections,
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
				Resolver: schema.PathResolver("ConnectionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "authorization_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizationType"),
			},
			{
				Name:     "connection_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionState"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "last_authorized_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastAuthorizedTime"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReason"),
			},
		},
	}
}
