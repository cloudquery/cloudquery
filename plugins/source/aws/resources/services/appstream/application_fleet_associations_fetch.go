package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAppstreamApplicationFleetAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	parentApplication := parent.Item.(types.Application)

	var input appstream.DescribeApplicationFleetAssociationsInput
	input.ApplicationArn = parentApplication.Arn

	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeApplicationFleetAssociations(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ApplicationFleetAssociations
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
