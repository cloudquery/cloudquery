// Code generated by codegen using template resource_manual.go.tpl; DO NOT EDIT.

package codegen

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/services"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_images",
		Resolver:  services.FetchImages,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "filename",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Filename"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "require_signed_ur_ls",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RequireSignedURLs"),
			},
			{
				Name:     "variants",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Variants"),
			},
			{
				Name:     "uploaded",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Uploaded"),
			},
		},
	}
}
