package config

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConformancePackRuleCompliances() *schema.Table {
	tableName := "aws_config_conformance_pack_rule_compliances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_DescribeConformancePackCompliance.html`,
		Resolver:    fetchConfigConformancePackRuleCompliances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform:   client.TransformWithStruct(&models.ConformancePackComplianceWrapper{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "conformance_pack_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
