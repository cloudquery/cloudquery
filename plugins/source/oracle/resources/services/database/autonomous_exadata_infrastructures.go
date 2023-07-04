package database

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func AutonomousExadataInfrastructures() *schema.Table {
	return &schema.Table{
		Name:      "oracle_database_autonomous_exadata_infrastructures",
		Resolver:  fetchAutonomousExadataInfrastructures,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&database.AutonomousExadataInfrastructureSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchAutonomousExadataInfrastructures(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := database.ListAutonomousExadataInfrastructuresRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].DatabaseDatabaseClient.ListAutonomousExadataInfrastructures(ctx, request)

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
