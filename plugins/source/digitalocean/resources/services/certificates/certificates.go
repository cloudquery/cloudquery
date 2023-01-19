package certificates

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_certificates",
		Resolver:  fetchCertificatesCertificates,
		Transform: transformers.TransformWithStruct(&godo.Certificate{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
