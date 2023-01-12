package virtualnetwork

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func fetchIpSecConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := core.ListIPSecConnectionsRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].CoreVirtualnetworkClient.ListIPSecConnections(ctx, request)

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
