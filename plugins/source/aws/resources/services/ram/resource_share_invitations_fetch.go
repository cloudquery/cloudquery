package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
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
