// Code generated by codegen; DO NOT EDIT.

package apprunner

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:                "aws_apprunner_services",
		Description:         "https://docs.aws.amazon.com/apprunner/latest/api/API_Service.html",
		Resolver:            fetchApprunnerServices,
		PreResourceResolver: getService,
		Multiplex:           client.ServiceAccountRegionMultiplexer("apprunner"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "auto_scaling_configuration_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoScalingConfigurationSummary"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "instance_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceConfiguration"),
			},
			{
				Name:     "network_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkConfiguration"),
			},
			{
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceId"),
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceName"),
			},
			{
				Name:     "service_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceUrl"),
			},
			{
				Name:     "source_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceConfiguration"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "deleted_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletedAt"),
			},
			{
				Name:     "encryption_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionConfiguration"),
			},
			{
				Name:     "health_check_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheckConfiguration"),
			},
			{
				Name:     "observability_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ObservabilityConfiguration"),
			},
		},

		Relations: []*schema.Table{
			RestApiAuthorizers(),
		},
	}
}
