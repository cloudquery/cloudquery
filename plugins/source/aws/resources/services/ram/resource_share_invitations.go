package ram

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ResourceShareInvitations() *schema.Table {
	tableName := "aws_ram_resource_share_invitations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareInvitation.html`,
		Resolver:    fetchRamResourceShareInvitations,
		Transform:   transformers.TransformWithStruct(&types.ResourceShareInvitation{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ResourceShareInvitationArn"),
				PrimaryKey: true,
			},
			{
				Name:       "receiver_combined",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveResourceShareInvitationReceiver,
				PrimaryKey: true,
			},
		},
	}
}

func fetchRamResourceShareInvitations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input ram.GetResourceShareInvitationsInput = getResourceShareInvitationsInput()
	cl := meta.(*client.Client)
	svc := cl.Services().Ram
	paginator := ram.NewGetResourceShareInvitationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ram.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ResourceShareInvitations
	}
	return nil
}

func resolveResourceShareInvitationReceiver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	inv := resource.Item.(types.ResourceShareInvitation)
	if inv.ReceiverArn != nil {
		return resource.Set(c.Name, *inv.ReceiverArn)
	}
	if inv.ReceiverAccountId != nil {
		return resource.Set(c.Name, *inv.ReceiverAccountId)
	}
	return fmt.Errorf("aws:ram invitation receiver both account and arn is missing")
}
