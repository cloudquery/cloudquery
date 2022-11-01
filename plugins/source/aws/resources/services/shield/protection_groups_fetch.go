package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchShieldProtectionGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.ListProtectionGroupsInput{}
	for {
		output, err := svc.ListProtectionGroups(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.ProtectionGroups

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveShieldProtectionGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ProtectionGroup)
	cli := meta.(*client.Client)
	svc := cli.Services().Shield
	config := shield.ListTagsForResourceInput{ResourceARN: r.ProtectionGroupArn}

	output, err := svc.ListTagsForResource(ctx, &config, func(o *shield.Options) {
		o.Region = cli.Region
	})
	if err != nil {
		if cli.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
