// Code generated by codegen; DO NOT EDIT.

package ssm

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InstanceComplianceItems() *schema.Table {
	return &schema.Table{
		Name:      "aws_ssm_instance_compliance_items",
		Resolver:  fetchSsmInstanceComplianceItems,
		Multiplex: client.ServiceAccountRegionMultiplexer("ssm"),
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
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "instance_arn",
				Type:     schema.TypeString,
				Resolver: resolveInstanceComplianceItemInstanceARN,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "compliance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ComplianceType"),
			},
			{
				Name:     "details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Details"),
			},
			{
				Name:     "execution_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExecutionSummary"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceId"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "severity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Severity"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
		},
	}
}
