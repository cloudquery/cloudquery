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

func Jobs() *schema.Table {
	tableName := "aws_iot_jobs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_Job.html`,
		Resolver:    fetchIotJobs,
		Transform:   transformers.TransformWithStruct(&types.Job{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotJobTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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

	for {
		response, err := svc.ListJobs(ctx, &input)
		if err != nil {
			return err
		}

		for _, s := range response.Jobs {
			job, err := svc.DescribeJob(ctx, &iot.DescribeJobInput{
				JobId: s.JobId,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- job.Job
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*types.Job)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.JobArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input)

		if err != nil {
			return err
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
