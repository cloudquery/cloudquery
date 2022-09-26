package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchResourcegroupsResourceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listResourceGroups, resourceGroupDetail)
}
func resolveResourcegroupsResourceGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().ResourceGroups
	group := resource.Item.(ResourceGroupWrapper)
	input := resourcegroups.GetTagsInput{
		Arn: group.GroupArn,
	}
	output, err := svc.GetTags(ctx, &input, func(options *resourcegroups.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.Tags)
}

// ResourceGroupWrapper fields are extracted from types.Group and types.ResourceQuery
// TODO: Use embedded structs once https://github.com/cloudquery/plugin-sdk/pull/75 is merged
type ResourceGroupWrapper struct {
	GroupArn    *string
	Name        *string
	Description *string
	Query       *string
	Type        types.QueryType
}

func listResourceGroups(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	var config resourcegroups.ListGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().ResourceGroups
	for {
		output, err := svc.ListGroups(ctx, &config, func(options *resourcegroups.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, item := range output.GroupIdentifiers {
			detailChan <- item.GroupArn
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resourceGroupDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	groupArn := listInfo.(*string)
	svc := c.Services().ResourceGroups
	groupResponse, err := svc.GetGroup(ctx, &resourcegroups.GetGroupInput{
		Group: groupArn,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}

	input := resourcegroups.GetGroupQueryInput{
		Group: groupResponse.Group.GroupArn,
	}
	output, err := svc.GetGroupQuery(ctx, &input)
	if err != nil {
		errorChan <- err
		return
	}
	resultsChan <- ResourceGroupWrapper{
		GroupArn:    groupResponse.Group.GroupArn,
		Name:        groupResponse.Group.Name,
		Description: groupResponse.Group.Description,
		Query:       output.GroupQuery.ResourceQuery.Query,
		Type:        output.GroupQuery.ResourceQuery.Type,
	}
}
