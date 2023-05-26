package virtualnetwork

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func Drgs() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_drgs",
		Resolver:  fetchDrgs,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: transformers.TransformWithStruct(&core.Drg{},
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
