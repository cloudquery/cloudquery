// Code generated by codegen; DO NOT EDIT.

package certificates

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_certificates",
		Resolver: fetchCertificatesCertificates,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
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
				Name:     "dns_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DNSNames"),
			},
			{
				Name:     "not_after",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NotAfter"),
			},
			{
				Name:     "sha_1_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SHA1Fingerprint"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
