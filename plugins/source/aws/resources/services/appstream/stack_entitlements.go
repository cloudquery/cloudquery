package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func StackEntitlements() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_stack_entitlements",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Entitlement.html`,
		Resolver:    fetchAppstreamStackEntitlements,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
		Transform:   transformers.TransformWithStruct(&types.Entitlement{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "stack_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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
