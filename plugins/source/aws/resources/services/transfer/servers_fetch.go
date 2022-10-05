package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTransferServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Transfer
	input := transfer.ListServersInput{MaxResults: aws.Int32(1000)}
	for {
		result, err := svc.ListServers(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Servers

		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func getServer(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Transfer
	server := resource.Item.(types.ListedServer)

	desc, err := svc.DescribeServer(ctx, &transfer.DescribeServerInput{ServerId: server.ServerId})
	if err != nil {
		return err
	}
	resource.Item = desc.Server
	return nil
}

func resolveServersTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Transfer
	server := resource.Item.(*types.DescribedServer)
	input := transfer.ListTagsForResourceInput{Arn: server.Arn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &input)
		if err != nil {
			if cl.IsNotFoundError(err) {
				continue
			}
			return err
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}
