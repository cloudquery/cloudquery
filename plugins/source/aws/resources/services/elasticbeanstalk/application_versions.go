package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ApplicationVersions() *schema.Table {
	tableName := "aws_elasticbeanstalk_application_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationVersionDescription.html`,
		Resolver:    fetchElasticbeanstalkApplicationVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&types.ApplicationVersionDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationVersionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchElasticbeanstalkApplicationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elasticbeanstalk.DescribeApplicationVersionsInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticbeanstalk
	// No paginator available
	for {
		output, err := svc.DescribeApplicationVersions(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.ApplicationVersions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
