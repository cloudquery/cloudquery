package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StackUserAssociations() *schema.Table {
	tableName := "aws_appstream_stack_user_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UserStackAssociation.html`,
		Resolver:    fetchAppstreamStackUserAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   client.TransformWithStruct(&types.UserStackAssociation{}),
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
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "authentication_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthenticationType"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
