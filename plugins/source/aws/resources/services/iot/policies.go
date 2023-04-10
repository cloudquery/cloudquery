package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Policies() *schema.Table {
	tableName := "aws_iot_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_Policy.html`,
		Resolver:            fetchIotPolicies,
		PreResourceResolver: getPolicy,
		Transform:           transformers.TransformWithStruct(&types.Policy{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotPolicyTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIotPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListPoliciesInput{
		PageSize: aws.Int32(250),
	}
	paginator := iot.NewListPoliciesPaginator(svc, &input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Policies
	}
	return nil
}

func getPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot

	output, err := svc.GetPolicy(ctx, &iot.GetPolicyInput{
		PolicyName: resource.Item.(types.Policy).PolicyName,
	})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

func ResolveIotPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListTagsForResourceInput{
		ResourceArn: resource.Item.(*iot.GetPolicyOutput).PolicyArn,
	}
	tags := make(map[string]string)

	paginator := iot.NewListTagsForResourcePaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		client.TagsIntoMap(page.Tags, tags)
	}
	return resource.Set(c.Name, tags)
}
