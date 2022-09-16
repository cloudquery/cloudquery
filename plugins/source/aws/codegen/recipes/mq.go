package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func MQResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "brokers",
			Struct:     &mq.DescribeBrokerOutput{},
			SkipFields: []string{"BrokerArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("BrokerArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"BrokerConfigurations()",
				"BrokerUsers()",
			},
		},
		{
			SubService: "broker_configurations",
			Struct:     &types.Configuration{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "broker_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
			Relations: []string{
				"BrokerConfigurationRevisions()",
			},
		},
		{
			SubService: "broker_configuration_revisions",
			Struct:     &mq.DescribeConfigurationRevisionOutput{},
			SkipFields: []string{"Data"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "broker_configuration_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "data",
						Type:     schema.TypeJSON,
						Resolver: `resolveBrokerConfigurationRevisionsData`,
					},
				}...),
		},
		{
			SubService: "broker_users",
			Struct:     &mq.DescribeUserOutput{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "broker_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "mq"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("mq")`
	}
	return resources
}
