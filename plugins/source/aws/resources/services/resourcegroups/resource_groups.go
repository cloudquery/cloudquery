package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource resource_groups --config gen.hcl --output .
func ResourceGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_resourcegroups_resource_groups",
		Resolver:     fetchResourcegroupsResourceGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("resource-groups"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveResourcegroupsResourceGroupTags,
			},
			{
				Name:        "arn",
				Description: "The ARN of the resource group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Group.GroupArn"),
			},
			{
				Name:        "group",
				Description: "The name of the resource group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Group.Name"),
			},
			{
				Name:        "group_description",
				Description: "The description of the resource group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Group.Description"),
			},
			{
				Name:        "resource_query",
				Description: "The query that defines a group or a search",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceQuery.Query"),
			},
			{
				Name:        "resource_query_type",
				Description: "The type of the query",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceQuery.Type"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchResourcegroupsResourceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listResourceGroups, resourceGroupDetail))
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, output.Tags))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type ResourceGroupWrapper struct {
	*types.Group
	*types.ResourceQuery
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
			return diag.WrapError(err)
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
		errorChan <- diag.WrapError(err)
		return
	}

	input := resourcegroups.GetGroupQueryInput{
		Group: groupResponse.Group.GroupArn,
	}
	output, err := svc.GetGroupQuery(ctx, &input)
	if err != nil {
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- ResourceGroupWrapper{
		groupResponse.Group,
		output.GroupQuery.ResourceQuery,
	}
}
