package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func TagNamespaces() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_tag_namespaces",
		Resolver:  fetchTagNamespaces,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&identity.TagNamespaceSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
