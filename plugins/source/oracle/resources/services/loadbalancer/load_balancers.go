package loadbalancer

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/loadbalancer"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:      "oracle_loadbalancer_load_balancers",
		Resolver:  fetchLoadBalancers,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&loadbalancer.LoadBalancer{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchLoadBalancers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := loadbalancer.ListLoadBalancersRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].LoadbalancerLoadbalancerClient.ListLoadBalancers(ctx, request)

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
