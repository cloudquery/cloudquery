package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceShareInvitations() *schema.Table {
	tableName := "aws_ram_resource_share_invitations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareInvitation.html`,
		Resolver:    fetchRamResourceShareInvitations,
		Transform:   client.TransformWithStruct(&types.ResourceShareInvitation{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareInvitationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "receiver_combined",
				Type:     schema.TypeString,
				Resolver: resolveResourceShareInvitationReceiver,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
