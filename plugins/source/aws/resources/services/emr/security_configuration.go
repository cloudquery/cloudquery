package emr

import (
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SecurityConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_emr_security_configurations",
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeSecurityConfiguration.html`,
		Resolver:    fetchSecurityConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
		Transform:   transformers.TransformWithStruct(&types.SecurityConfigurationSummary{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "security_configuration",
				Type:     schema.TypeJSON,
				Resolver: resolveConfiguration,
			},
		},
	}
}
