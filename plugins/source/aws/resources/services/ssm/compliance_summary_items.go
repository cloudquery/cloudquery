// Code generated by codegen; DO NOT EDIT.

package ssm

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ComplianceSummaryItems() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_compliance_summary_items",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceSummaryItem.html`,
		Resolver:    fetchSsmComplianceSummaryItems,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
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
				Name: "compliance_type",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "compliant_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CompliantSummary"),
			},
			{
				Name:     "non_compliant_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NonCompliantSummary"),
			},
		},
	}
}
