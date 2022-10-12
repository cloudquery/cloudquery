// Code generated by codegen; DO NOT EDIT.

package route53

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TrafficPolicyVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_traffic_policy_versions",
		Description: "https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html",
		Resolver:    fetchRoute53TrafficPolicyVersions,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "traffic_policy_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "document",
				Type:     schema.TypeJSON,
				Resolver: client.MarshaledJsonResolver("Document"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "comment",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Comment"),
			},
		},
	}
}
