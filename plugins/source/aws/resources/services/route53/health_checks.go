package route53

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func HealthChecks() *schema.Table {
	tableName := "aws_route53_health_checks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheck.html`,
		Resolver:    fetchRoute53HealthChecks,
		Transform:   transformers.TransformWithStruct(&Route53HealthCheckWrapper{}, transformers.WithUnwrapStructFields("HealthCheck")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveHealthCheckArn(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Description: `The tags associated with the health check.`,
			},
			{
				Name:     "cloud_watch_alarm_configuration_dimensions",
				Type:     schema.TypeJSON,
				Resolver: resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions,
			},
		},
	}
}
