package wellarchitected

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ShareInvitations() *schema.Table {
	name := "aws_wellarchitected_share_invitations"
	return &schema.Table{
		Name:        name,
		Description: `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_ShareInvitation.html`,
		Transform: transformers.TransformWithStruct(new(types.ShareInvitationSummary),
			transformers.WithPrimaryKeys("ShareInvitationId"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:  fetchShareInvitations,
		Columns:   schema.ColumnList{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true)},
	}
}

func fetchShareInvitations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected

	for _, shareResourceType := range types.ShareResourceType("").Values() {
		p := wellarchitected.NewListShareInvitationsPaginator(service,
			&wellarchitected.ListShareInvitationsInput{
				MaxResults:        50,
				ShareResourceType: shareResourceType,
			},
		)
		for p.HasMorePages() {
			output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
				o.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- output.ShareInvitationSummaries
		}
	}

	return nil
}
