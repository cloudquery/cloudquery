package elasticbeanstalk

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func configurationOptions() *schema.Table {
	tableName := "aws_elasticbeanstalk_configuration_options"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html`,
		Resolver:    fetchElasticbeanstalkConfigurationOptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&models.ConfigurationOptionDescriptionWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "environment_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchElasticbeanstalkConfigurationOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.EnvironmentDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticbeanstalk
	configOptionsIn := elasticbeanstalk.DescribeConfigurationOptionsInput{
		ApplicationName: p.ApplicationName,
		EnvironmentName: p.EnvironmentName,
	}
	output, err := svc.DescribeConfigurationOptions(ctx, &configOptionsIn, func(options *elasticbeanstalk.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		// It takes a few minutes for an environment to be terminated
		// This ensures we don't error while trying to fetch related resources for a terminated environment
		if client.IsInvalidParameterValueError(err) {
			return nil
		}
		return err
	}

	arnStr := arn.ARN{
		Partition: cl.Partition,
		Service:   "elasticbeanstalk",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("application/%s", aws.ToString(p.ApplicationName)),
	}.String()

	for _, option := range output.Options {
		res <- models.ConfigurationOptionDescriptionWrapper{
			ConfigurationOptionDescription: option, ApplicationArn: arnStr,
		}
	}

	return nil
}
