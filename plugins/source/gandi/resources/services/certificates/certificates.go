package certificates

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
