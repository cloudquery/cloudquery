package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func Compartments() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_compartments",
		Resolver:  fetchCompartments,
		Multiplex: client.TenancyMultiplex,
		Transform: client.TransformWithStruct(&identity.Compartment{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
