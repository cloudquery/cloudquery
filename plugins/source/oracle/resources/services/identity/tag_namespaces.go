package identity

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func TagNamespaces() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_tag_namespaces",
		Resolver:  fetchTagNamespaces,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: transformers.TransformWithStruct(&identity.TagNamespaceSummary{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn,
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}
