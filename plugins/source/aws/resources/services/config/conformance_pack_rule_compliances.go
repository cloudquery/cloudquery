package config

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConformancePackRuleCompliances() *schema.Table {
	return &schema.Table{
		Name:      "aws_config_conformance_pack_rule_compliances",
		Resolver:  fetchConfigConformancePackRuleCompliances,
		Multiplex: client.ServiceAccountRegionMultiplexer("config"),
		Transform: transformers.TransformWithStruct(&models.ConformancePackComplianceWrapper{}),
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
		},
	}
}
