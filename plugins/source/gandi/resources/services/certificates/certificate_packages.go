package certificates

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/go-gandi/go-gandi/certificate"
)

func CertificatePackages() *schema.Table {
	return &schema.Table{
		Name:      "gandi_certificate_packages",
		Resolver:  fetchCertificatePackages,
		Transform: transformers.TransformWithStruct(&certificate.Package{}),
		Columns: []schema.Column{
			{
				Name:        "sharing_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
