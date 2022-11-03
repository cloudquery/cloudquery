// Code generated by codegen; DO NOT EDIT.

package kms

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func KeyGrants() *schema.Table {
	return &schema.Table{
		Name:        "aws_kms_key_grants",
		Description: "https://docs.aws.amazon.com/kms/latest/APIReference/API_GrantListEntry.html",
		Resolver:    fetchKmsKeyGrants,
		Multiplex:   client.ServiceAccountRegionMultiplexer("kms"),
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
				Name:     "key_arn",
				Type:     schema.TypeString,
				Resolver: resolveKeyGrantsKeyArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "grant_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GrantId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "constraints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Constraints"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "grantee_principal",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GranteePrincipal"),
			},
			{
				Name:     "issuing_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IssuingAccount"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "operations",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Operations"),
			},
			{
				Name:     "retiring_principal",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RetiringPrincipal"),
			},
		},
	}
}
