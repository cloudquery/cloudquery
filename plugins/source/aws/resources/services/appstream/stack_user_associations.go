package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func StackUserAssociations() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_stack_user_associations",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UserStackAssociation.html`,
		Resolver:    fetchAppstreamStackUserAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
		Transform:   transformers.TransformWithStruct(&types.UserStackAssociation{}),
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
