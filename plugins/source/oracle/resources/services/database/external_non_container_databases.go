package database

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func ExternalNonContainerDatabases() *schema.Table {
	return &schema.Table{
		Name:      "oracle_database_external_non_container_databases",
		Resolver:  fetchExternalNonContainerDatabases,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&database.ExternalNonContainerDatabaseSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
