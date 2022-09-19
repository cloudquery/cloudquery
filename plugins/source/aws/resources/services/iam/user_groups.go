// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_iam_user_groups",
		Resolver:  fetchIamUserGroups,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("user_id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "create_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateDate"),
			},
			{
				Name:     "group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupId"),
			},
			{
				Name:     "group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupName"),
			},
			{
				Name:     "path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Path"),
			},
		},
	}
}
