package database

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func AutonomousContainerDatabases() *schema.Table {
	return &schema.Table{
		Name:      "oracle_database_autonomous_container_databases",
		Resolver:  fetchAutonomousContainerDatabases,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&database.AutonomousContainerDatabaseSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
