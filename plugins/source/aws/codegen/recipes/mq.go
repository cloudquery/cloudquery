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
			SubService:          "brokers",
			Struct:              &mq.DescribeBrokerOutput{},
			Description:         "https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers.html",
			SkipFields:          []string{"BrokerArn"},
			PreResourceResolver: "getMqBroker",
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
			SubService:  "broker_configurations",
			Struct:      &types.Configuration{},
			Description: "https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "broker_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
			Relations: []string{
				"BrokerConfigurationRevisions()",
			},
		},
		{
			SubService:          "broker_configuration_revisions",
			Struct:              &mq.DescribeConfigurationRevisionOutput{},
			Description:         "https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id-revisions.html",
			SkipFields:          []string{"Data"},
			PreResourceResolver: "getMqBrokerConfigurationRevision",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "broker_configuration_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
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
						Resolver: `schema.ParentColumnResolver("arn")`,
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
