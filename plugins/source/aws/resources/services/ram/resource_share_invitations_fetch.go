package ram

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRamResourceShareInvitations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input ram.GetResourceShareInvitationsInput = getResourceShareInvitationsInput()
	c := meta.(*client.Client)
	svc := c.Services().Ram
	for {
		response, err := svc.GetResourceShareInvitations(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ResourceShareInvitations

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
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
