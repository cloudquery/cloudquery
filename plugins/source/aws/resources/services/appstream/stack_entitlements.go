package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func StackEntitlements() *schema.Table {
	tableName := "aws_appstream_stack_entitlements"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Entitlement.html`,
		Resolver:    fetchAppstreamStackEntitlements,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.Entitlement{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
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
