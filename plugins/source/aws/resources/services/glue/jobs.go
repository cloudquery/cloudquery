package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Jobs() *schema.Table {
	tableName := "aws_glue_jobs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Job.html`,
		Resolver:    fetchGlueJobs,
		Transform:   transformers.TransformWithStruct(&types.Job{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueJobArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueJobTags,
			},
		},

		Relations: []*schema.Table{
			jobRuns(),
		},
	}
}

func fetchGlueJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetJobsInput{}
	for {
		result, err := svc.GetJobs(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Jobs
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func resolveGlueJobArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, jobARN(cl, aws.ToString(resource.Item.(types.Job).Name)))
}
func resolveGlueJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(jobARN(cl, aws.ToString(resource.Item.(types.Job).Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
