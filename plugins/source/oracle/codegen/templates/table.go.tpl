package {{.Service}}

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/oracle/oci-go-sdk/v65/{{.Package}}"
)

func {{.SubService | ToCamel}}() *schema.Table {
	return &schema.Table{
		Name:      "oracle_{{.Service}}_{{.SubService}}",
		Resolver:  fetch{{.SubService | ToCamel}},
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: transformers.TransformWithStruct(&{{.Package}}.{{.StructName}}{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: []schema.Column{
			{
				Name:     "region",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveOracleRegion,
				PrimaryKey: true,
			},
			{
				Name:     "compartment_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveCompartmentId,
				PrimaryKey: true,
			},
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}
