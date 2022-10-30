package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IOTResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "billing_groups",
			Struct:     &iot.DescribeBillingGroupOutput{},
			SkipFields: []string{"BillingGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "things_in_group",
						Type:     schema.TypeStringArray,
						Resolver: `resolveIotBillingGroupThingsInGroup`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveIotBillingGroupTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("BillingGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "ca_certificates",
			Struct:      &types.CACertificateDescription{},
			Description: "https://docs.aws.amazon.com/iot/latest/apireference/API_CACertificateDescription.html",
			SkipFields:  []string{"CertificateArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "certificates",
						Type:     schema.TypeStringArray,
						Resolver: `ResolveIotCaCertificateCertificates`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("CertificateArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "certificates",
			Struct:      &types.CertificateDescription{},
			Description: "https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html",
			SkipFields:  []string{"CertificateArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "policies",
						Type:     schema.TypeStringArray,
						Resolver: `ResolveIotCertificatePolicies`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("CertificateArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "jobs",
			Struct:      &types.Job{},
			Description: "https://docs.aws.amazon.com/iot/latest/apireference/API_Job.html",
			SkipFields:  []string{"JobArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `ResolveIotJobTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("JobArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "policies",
			Struct:      &types.Policy{},
			Description: "https://docs.aws.amazon.com/iot/latest/apireference/API_Policy.html",
			SkipFields:  []string{"PolicyArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `ResolveIotPolicyTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("PolicyArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "security_profiles",
			Struct:     &iot.DescribeSecurityProfileOutput{},
			SkipFields: []string{"SecurityProfileArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "targets",
						Type:     schema.TypeStringArray,
						Resolver: `ResolveIotSecurityProfileTargets`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `ResolveIotSecurityProfileTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("SecurityProfileArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "streams",
			Struct:      &types.StreamInfo{},
			Description: "https://docs.aws.amazon.com/iot/latest/apireference/API_StreamInfo.html",
			SkipFields:  []string{"StreamArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("StreamArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "thing_groups",
			Struct:     &iot.DescribeThingGroupOutput{},
			SkipFields: []string{"ThingGroupArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "things_in_group",
						Type:     schema.TypeStringArray,
						Resolver: `ResolveIotThingGroupThingsInGroup`,
					},
					{
						Name:          "policies",
						Type:          schema.TypeStringArray,
						Resolver:      `ResolveIotThingGroupPolicies`,
						IgnoreInTests: true,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `ResolveIotThingGroupTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ThingGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "thing_types",
			Struct:      &types.ThingTypeDefinition{},
			Description: "https://docs.aws.amazon.com/iot/latest/apireference/API_ThingTypeDefinition.html",
			SkipFields:  []string{"ThingTypeArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `ResolveIotThingTypeTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ThingTypeArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "things",
			Struct:      &types.ThingAttribute{},
			Description: "https://docs.aws.amazon.com/iot/latest/apireference/API_ThingAttribute.html",
			SkipFields:  []string{"ThingArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:          "principals",
						Type:          schema.TypeStringArray,
						Resolver:      `ResolveIotThingPrincipals`,
						IgnoreInTests: true,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ThingArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "topic_rules",
			Struct:     &iot.GetTopicRuleOutput{},
			SkipFields: []string{"RuleArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `ResolveIotTopicRuleTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("RuleArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "iot"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("iot")`
	}
	return resources
}
