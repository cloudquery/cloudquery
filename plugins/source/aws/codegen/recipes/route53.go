package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Route53Resources() []*Resource {
	resources := []*Resource{

		{
			SubService: "delegation_sets",
			Struct:     &types.DelegationSet{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver:    `resolveDelegationSetArn()`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},

		{
			SubService: "domains",
			Struct:     &route53domains.GetDomainDetailOutput{},
			SkipFields: []string{"DomainName", "_"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:    "domain_name",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:        "tags",
					Description: "A list of tags",
					Type:        schema.TypeJSON,
					Resolver:    `resolveRoute53DomainTags`,
				},
			},
		},

		{
			SubService: "health_checks",
			Struct:     &types.HealthCheck{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveHealthCheckArn()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The tags associated with the health check.",
						Type:        schema.TypeJSON,
					},
					{
						Name:     "cloud_watch_alarm_configuration_dimensions",
						Type:     schema.TypeJSON,
						Resolver: `resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions`,
					},
				}...),
		},

		{
			SubService: "hosted_zones",
			Struct:     &types.HostedZone{},
			SkipFields: []string{"ARN"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveRoute53HostedZoneArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"HostedZoneQueryLoggingConfigs()",
				"HostedZoneResourceRecordSets()",
				"HostedZoneTrafficPolicyInstances()",
			},
		},
		{
			SubService: "hosted_zone_query_logging_configs",
			Struct:     &types.QueryLoggingConfig{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveRoute53HostedZoneQueryLoggingConfigsArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "hosted_zone_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "hosted_zone_resource_record_sets",
			Struct:     &types.ResourceRecordSet{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "hosted_zone_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "hosted_zone_traffic_policy_instances",
			Struct:     &types.TrafficPolicyInstance{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "arn",
						Description: "Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.",
						Type:        schema.TypeString,
						Resolver:    `resolveRoute53HostedZoneTrafficPolicyInstancesArn`,
					},
					{
						Name:     "hosted_zone_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},

		{
			SubService: "traffic_policies",
			Struct:     &types.TrafficPolicySummary{},
			SkipFields: []string{"ARN"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveTrafficPolicyArn()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				`TrafficPolicyVersions()`,
			},
		},
		{
			SubService: "traffic_policy_versions",
			Struct:     &types.TrafficPolicy{},
			SkipFields: []string{"Version", "Id"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "traffic_policy_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Id")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "version",
						Type:     schema.TypeInt,
						Resolver: `schema.PathResolver("Version")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "route53"
		r.Multiplex = "client.AccountMultiplex"
	}
	return resources
}
