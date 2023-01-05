// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ServiceAccountKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_account_keys",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey`,
		Resolver:    fetchServiceAccountKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "service_account_unique_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("unique_id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "key_algorithm",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("KeyAlgorithm"),
			},
			{
				Name:     "public_key_data",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("PublicKeyData"),
			},
			{
				Name:     "valid_after_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("ValidAfterTime"),
			},
			{
				Name:     "valid_before_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("ValidBeforeTime"),
			},
			{
				Name:     "key_origin",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("KeyOrigin"),
			},
			{
				Name:     "key_type",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("KeyType"),
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Disabled"),
			},
		},
	}
}
