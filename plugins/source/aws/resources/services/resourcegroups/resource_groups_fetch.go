package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/resourcegroups/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchResourcegroupsResourceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config resourcegroups.ListGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Resourcegroups
	for {
		output, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.GroupIdentifiers
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func getResourceGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	group := resource.Item.(types.GroupIdentifier)
	svc := c.Services().Resourcegroups
	groupResponse, err := svc.GetGroup(ctx, &resourcegroups.GetGroupInput{
		Group: group.GroupArn,
	})
	if err != nil {
		return err
	}

	input := resourcegroups.GetGroupQueryInput{
		Group: groupResponse.Group.GroupArn,
	}
	output, err := svc.GetGroupQuery(ctx, &input)
	if err != nil {
		return err
	}
	resource.Item = models.ResourceGroupWrapper{
		Group:         groupResponse.Group,
		ResourceQuery: output.GroupQuery.ResourceQuery,
	}
	return nil
}

func resolveResourcegroupsResourceGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Resourcegroups
	group := resource.Item.(models.ResourceGroupWrapper)
	input := resourcegroups.GetTagsInput{
		Arn: group.GroupArn,
	}
	output, err := svc.GetTags(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.Tags)
}
