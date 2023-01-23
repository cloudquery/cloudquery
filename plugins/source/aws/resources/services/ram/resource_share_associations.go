package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceShareAssociations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_associations",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareAssociation.html`,
		Resolver:    fetchRamResourceShareAssociations,
		Transform:   transformers.TransformWithStruct(&types.ResourceShareAssociation{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
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
				Name:     "associated_entity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociatedEntity"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "resource_share_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
