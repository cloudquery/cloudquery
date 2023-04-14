package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_users",
		Resolver:  fetchUsers,
		Multiplex: client.TenancyMultiplex,
		Transform: transformers.TransformWithStruct(&identity.User{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
