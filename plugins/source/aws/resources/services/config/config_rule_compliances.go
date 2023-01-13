package config

import (
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConfigRuleCompliances() *schema.Table {
	return &schema.Table{
		Name:      "aws_config_config_rule_compliances",
		Resolver:  fetchConfigConfigRuleCompliances,
		Multiplex: client.ServiceAccountRegionMultiplexer("config"),
		Transform: transformers.TransformWithStruct(&types.ComplianceByConfigRule{}),
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
