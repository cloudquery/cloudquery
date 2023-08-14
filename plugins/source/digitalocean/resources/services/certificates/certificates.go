package certificates

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
