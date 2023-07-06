package iot

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Jobs() *schema.Table {
	tableName := "aws_iot_jobs"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_Job.html`,
		Resolver:            fetchIotJobs,
		PreResourceResolver: getJobs,
		Transform:           transformers.TransformWithStruct(&types.Job{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: ResolveIotJobTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("JobArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIotJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListJobsInput{
		MaxResults: aws.Int32(250),
	}
	paginator := iot.NewListJobsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Jobs
	}
	return nil
}

func getJobs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot

	output, err := svc.DescribeJob(ctx, &iot.DescribeJobInput{
		JobId: resource.Item.(types.JobSummary).JobId,
	}, func(options *iot.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output.Job
	return nil
}

func ResolveIotJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*types.Job)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	return resolveIotTags(ctx, meta, svc, resource, c, i.JobArn)
}
