package database

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func AutonomousDatabases() *schema.Table {
	return &schema.Table{
		Name:      "oracle_database_autonomous_databases",
		Resolver:  fetchAutonomousDatabases,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&database.AutonomousDatabaseSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchAutonomousDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := database.ListAutonomousDatabasesRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].DatabaseDatabaseClient.ListAutonomousDatabases(ctx, request)

		if err != nil {
			return err
		}

		res <- response.Items

		if response.OpcNextPage == nil {
			break
		}

		page = response.OpcNextPage
	}

	return nil
}
