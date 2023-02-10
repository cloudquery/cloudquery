package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceShareInvitations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_invitations",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareInvitation.html`,
		Resolver:    fetchRamResourceShareInvitations,
		Transform:   transformers.TransformWithStruct(&types.ResourceShareInvitation{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			client.AccountPKColumn(false),
			client.RegionPKColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareInvitationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
