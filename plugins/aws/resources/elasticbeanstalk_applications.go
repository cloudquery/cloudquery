package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ElasticbeanstalkApplications() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticbeanstalk_applications",
		Description:  "Describes the properties of an application.",
		Resolver:     fetchElasticbeanstalkApplications,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn", "date_created"}},
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the application.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationArn"),
			},
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApplicationName"),
			},
			{
				Name:        "configuration_templates",
				Description: "The names of the configuration templates associated with this application.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "date_created",
				Description: "The date when the application was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "date_updated",
				Description: "The date when the application was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "User-defined description of the application.",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_lifecycle_config_service_role",
				Description: "The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a VersionLifecycleConfig for the application in one of the supporting calls (CreateApplication or UpdateApplicationResourceLifecycle)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.ServiceRole"),
			},
			{
				Name:        "max_age_rule_enabled",
				Description: "Specify true to apply the rule, or false to disable it.  This member is required.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.Enabled"),
			},
			{
				Name:        "max_age_rule_delete_source_from_s3",
				Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.DeleteSourceFromS3"),
			},
			{
				Name:        "max_age_rule_max_age_in_days",
				Description: "Specify the number of days to retain an application versions.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.MaxAgeInDays"),
			},
			{
				Name:        "max_count_rule_enabled",
				Description: "Specify true to apply the rule, or false to disable it.  This member is required.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.Enabled"),
			},
			{
				Name:        "max_count_rule_delete_source_from_s3",
				Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.DeleteSourceFromS3"),
			},
			{
				Name:        "max_count_rule_max_count",
				Description: "Specify the maximum number of application versions to retain.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.MaxCount"),
			},
			{
				Name:        "versions",
				Description: "The names of the versions for this application.",
				Type:        schema.TypeStringArray,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElasticbeanstalkApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config elasticbeanstalk.DescribeApplicationsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	output, err := svc.DescribeApplications(ctx, &config, func(options *elasticbeanstalk.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- output.Applications
	return nil
}
