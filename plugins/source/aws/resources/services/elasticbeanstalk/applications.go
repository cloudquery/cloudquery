package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ElasticbeanstalkApplications() *schema.Table {
	return &schema.Table{
		Name:          "aws_elasticbeanstalk_applications",
		Description:   "Describes the properties of an application.",
		Resolver:      fetchElasticbeanstalkApplications,
		Multiplex:     client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		IgnoreInTests: true,
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the application.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ApplicationArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:            "date_created",
				Description:     "The date when the application was created.",
				Type:            schema.TypeTimestamp,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:     "resource_lifecycle_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceLifecycleConfig"),
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
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchElasticbeanstalkApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elasticbeanstalk.DescribeApplicationsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	output, err := svc.DescribeApplications(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.Applications
	return nil
}
