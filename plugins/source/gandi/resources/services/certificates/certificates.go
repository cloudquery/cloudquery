package certificates

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/go-gandi/go-gandi/certificate"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:      "gandi_certificates",
		Resolver:  fetchCertificates,
		Transform: transformers.TransformWithStruct(&certificate.CertificateType{}),
		Columns: []schema.Column{
			{
				Name:        "sharing_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
			},
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
