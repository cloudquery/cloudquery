package virtualnetwork

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func PrivateIPs() *schema.Table {
	return &schema.Table{
		Name:      "oracle_virtualnetwork_private_ips",
		Resolver:  fetchPrivateIPs,
		Multiplex: client.TenancyMultiplex,
		Transform: client.TransformWithStruct(new(core.PrivateIp)),
	}
}

func fetchPrivateIPs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	limit := 100
	for {
		request := core.ListPrivateIpsRequest{
			Page:  page,
			Limit: &limit,
		}

		response, err := cqClient.OracleClients[cqClient.Region].CoreVirtualnetworkClient.ListPrivateIps(ctx, request)

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
