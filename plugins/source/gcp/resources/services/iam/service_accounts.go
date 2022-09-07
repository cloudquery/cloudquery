// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:      "gcp_iam_service_accounts",
		Resolver:  fetchServiceAccounts,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Disabled"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "oauth_2_client_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Oauth2ClientId"),
			},
			{
				Name:     "unique_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UniqueId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "server_response",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServerResponse"),
			},
		},
	}
}
