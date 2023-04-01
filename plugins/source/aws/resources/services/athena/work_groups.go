package athena

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkGroups() *schema.Table {
	tableName := "aws_athena_work_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_WorkGroup.html`,
		Resolver:            fetchAthenaWorkGroups,
		PreResourceResolver: getWorkGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:           transformers.TransformWithStruct(&types.WorkGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveAthenaWorkGroupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveAthenaWorkGroupTags,
			},
		},

		Relations: []*schema.Table{
			workGroupPreparedStatements(),
			workGroupQueryExecutions(),
			workGroupNamedQueries(),
		},
	}
}

func fetchAthenaWorkGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListWorkGroupsInput{}
	paginator := athena.NewListWorkGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.WorkGroups
	}

	return nil
}

func getWorkGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena

	wg := resource.Item.(types.WorkGroupSummary)
	dc, err := svc.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
		WorkGroup: wg.Name,
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.WorkGroup
	return nil
}

func resolveAthenaWorkGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.WorkGroup)
	return resource.Set(c.Name, createWorkGroupArn(cl, *dc.Name))
}

func resolveAthenaWorkGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	wg := resource.Item.(types.WorkGroup)
	arnStr := createWorkGroupArn(cl, *wg.Name)
	params := athena.ListTagsForResourceInput{ResourceARN: &arnStr}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}

func createWorkGroupArn(cl *client.Client, groupName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Athena),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("workgroup/%s", groupName),
	}.String()
}
