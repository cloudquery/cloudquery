package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_users",
		Resolver:  fetchUsers,
		Multiplex: client.TenancyMultiplex,
		Transform: client.TransformWithStruct(&identity.User{}),
		Columns:   schema.ColumnList{client.RegionColumn},
	}
}
