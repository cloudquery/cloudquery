package virtualnetwork

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func DRGs() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_drgs",
		Resolver:  fetchDRGs,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.Drg{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchDRGs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := core.ListDrgsRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].CoreVirtualnetworkClient.ListDrgs(ctx, request)

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
