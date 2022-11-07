// Code generated by codegen; DO NOT EDIT.

package ram

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceShareAssociatedPrincipals() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_associated_principals",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareAssociation.html`,
		Resolver:    fetchRamResourceShareAssociatedPrincipals,
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
				Name:     "associated_entity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociatedEntity"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "resource_share_arn",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceShareArn"),
				Description: `Resource Share ARN`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "association_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationType"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "external",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("External"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "resource_share_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareName"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
		},
	}
}
