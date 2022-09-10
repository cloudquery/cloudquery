package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApplicationVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_application_versions",
		Description: "Describes the properties of an application version.",
		Resolver:    fetchElasticbeanstalkApplicationVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
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
				Name:        "application_name",
				Description: "The name of the application to which the application version belongs.",
				Type:        schema.TypeString,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the application version.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ApplicationVersionArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:          "build_arn",
				Description:   "Reference to the artifact from the AWS CodeBuild build.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "date_created",
				Description: "The creation date of the application version.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "date_updated",
				Description: "The last modified date of the application version.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The description of the application version.",
				Type:        schema.TypeString,
			},
			{
				Name: "source_build_information",
				Type: schema.TypeJSON,
			},
			{
				Name:     "source_bundle",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceBundle"),
			},
			{
				Name:        "status",
				Description: "The processing status of the application version",
				Type:        schema.TypeString,
			},
			{
				Name:        "version_label",
				Description: "A unique identifier for the application version.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticbeanstalkApplicationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elasticbeanstalk.DescribeApplicationVersionsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk

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
