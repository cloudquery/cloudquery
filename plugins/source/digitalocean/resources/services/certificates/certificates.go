package certificates

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_certificates",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Certificates",
		Resolver:    fetchCertificatesCertificates,
		Transform:   transformers.TransformWithStruct(&godo.Certificate{}),
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
