package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource application_versions --config gen.hcl --output .
func ApplicationVersions() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticbeanstalk_application_versions",
		Description:  "Describes the properties of an application version.",
		Resolver:     fetchElasticbeanstalkApplicationVersions,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "application_name",
				Description: "The name of the application to which the application version belongs.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the application version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationVersionArn"),
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
				Name:          "source_location",
				Description:   "The location of the source code, as a formatted string, depending on the value of SourceRepository  * For CodeCommit, the format is the repository name and commit ID, separated by a forward slash",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("SourceBuildInformation.SourceLocation"),
				IgnoreInTests: true,
			},
			{
				Name:        "source_repository",
				Description: "Location where the repository is stored.  * CodeCommit  * S3  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceBuildInformation.SourceRepository"),
			},
			{
				Name:        "source_type",
				Description: "The type of repository.  * Git  * Zip  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceBuildInformation.SourceType"),
			},
			{
				Name:        "source_bundle_s3_bucket",
				Description: "The Amazon S3 bucket where the data is located.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceBundle.S3Bucket"),
			},
			{
				Name:        "source_bundle_s3_key",
				Description: "The Amazon S3 key where the data is located.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceBundle.S3Key"),
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
		output, err := svc.DescribeApplicationVersions(ctx, &config, func(options *elasticbeanstalk.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		res <- output.ApplicationVersions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
