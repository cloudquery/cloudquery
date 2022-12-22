package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AMPResources() []*Resource {
	return []*Resource{
		{
			Service:             "amp",
			SubService:          "workspaces",
			Struct:              new(types.WorkspaceDescription),
			Description:         "https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-WorkspaceDescription",
			Multiplex:           `client.ServiceAccountRegionMultiplexer("amp")`,
			PreResourceResolver: "describeWorkspace",
			PKColumns:           []string{"arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "alert_manager_definition",
					Type:     schema.TypeJSON,
					Resolver: `describeAlertManagerDefinition`,
				},
				codegen.ColumnDefinition{
					Name:     "logging_configuration",
					Type:     schema.TypeJSON,
					Resolver: `describeLoggingConfiguration`,
				},
			),
			Relations: []string{"RuleGroupsNamespaces()"},
		},
		{
			Service:             "amp",
			SubService:          "rule_groups_namespaces",
			Struct:              new(types.RuleGroupsNamespaceDescription),
			Description:         "https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-RuleGroupsNamespaceDescription",
			Multiplex:           "", // relation for workspace
			PreResourceResolver: "describeRuleGroupsNamespace",
			PKColumns:           []string{"arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "workspace_arn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("arn")`,
				},
			),
		},
	}
}
