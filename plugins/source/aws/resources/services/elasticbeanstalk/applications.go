package elasticbeanstalk

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Applications() *schema.Table {
	tableName := "aws_elasticbeanstalk_applications"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationDescription.html`,
		Resolver:    fetchElasticbeanstalkApplications,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&types.ApplicationDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ApplicationArn"),
				PrimaryKey: true,
			},
			{
				Name:       "date_created",
				Type:       arrow.FixedWidthTypes.Timestamp_us,
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveElasticbeanstalkApplicationTags,
			},
		},
	}
}

func fetchElasticbeanstalkApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elasticbeanstalk.DescribeApplicationsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticbeanstalk
	output, err := svc.DescribeApplications(ctx, &config, func(options *elasticbeanstalk.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.Applications
	return nil
}

func resolveElasticbeanstalkApplicationTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ApplicationDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticbeanstalk
	tagsOutput, err := svc.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: p.ApplicationArn,
	}, func(o *elasticbeanstalk.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		// It takes a few minutes for an environment to be terminated
		// This ensures we don't error while trying to fetch related resources for a terminated environment
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tagsOutput.ResourceTags))
}
