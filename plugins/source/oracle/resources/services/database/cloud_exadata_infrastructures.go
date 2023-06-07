package database

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func CloudExadataInfrastructures() *schema.Table {
	return &schema.Table{
		Name:      "oracle_database_cloud_exadata_infrastructures",
		Resolver:  fetchCloudExadataInfrastructures,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&database.CloudExadataInfrastructureSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
