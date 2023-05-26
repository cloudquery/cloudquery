package database

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func CloudExadataInfrastructures() *schema.Table {
	return &schema.Table{
		Name:      "oracle_database_cloud_exadata_infrastructures",
		Resolver:  fetchCloudExadataInfrastructures,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: transformers.TransformWithStruct(&database.CloudExadataInfrastructureSummary{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: []schema.Column{
			{
				Name:       "region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveOracleRegion,
				PrimaryKey: true,
			},
			{
				Name:       "compartment_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveCompartmentId,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}
