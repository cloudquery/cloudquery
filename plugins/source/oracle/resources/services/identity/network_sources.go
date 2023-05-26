package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func NetworkSources() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_network_sources",
		Resolver:  fetchNetworkSources,
		Multiplex: client.TenancyMultiplex,
		Transform: client.TransformWithStruct(&identity.NetworkSourcesSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn},
	}
}
