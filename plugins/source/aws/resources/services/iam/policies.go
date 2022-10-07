// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Policies() *schema.Table {
	return &schema.Table{
		Name:      "aws_iam_policies",
		Resolver:  fetchIamPolicies,
		Multiplex: client.AccountMultiplex,
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamPolicyTags,
			},
			{
				Name:     "policy_version_list",
				Type:     schema.TypeJSON,
				Resolver: resolveIamPolicyVersionList,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "attachment_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AttachmentCount"),
			},
			{
				Name:     "create_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateDate"),
			},
			{
				Name:     "default_version_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultVersionId"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "is_attachable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsAttachable"),
			},
			{
				Name:     "path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Path"),
			},
			{
				Name:     "permissions_boundary_usage_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PermissionsBoundaryUsageCount"),
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
			},
			{
				Name:     "update_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdateDate"),
			},
		},
	}
}
