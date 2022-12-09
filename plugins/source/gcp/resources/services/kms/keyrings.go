// Code generated by codegen; DO NOT EDIT.

package kms

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Keyrings() *schema.Table {
	return &schema.Table{
		Name:      "gcp_kms_keyrings",
		Resolver:  fetchKeyrings,
		Multiplex: client.ProjectMultiplex("cloudkms.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
		},

		Relations: []*schema.Table{
			CryptoKeys(),
		},
	}
}
