package virtualnetwork

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func ReservedPublicIPs() *schema.Table {
	return &schema.Table{
		Name: "oracle_virtualnetwork_reserved_public_ips",
		Description: `Reserved public IPs.
See https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/PublicIp/ListPublicIps for more details.`,
		Resolver: fetchPublicIPs(
			core.ListPublicIpsScopeRegion,
			nil,
			core.ListPublicIpsLifetimeReserved,
		),
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.PublicIpPoolSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func EphemeralPublicIPs() *schema.Table {
	return &schema.Table{
		Name: "oracle_virtualnetwork_ephemeral_public_ips",
		Description: `Ephemeral public IPs assigned to a regional entity such as a NAT gateway.
See https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/PublicIp/ListPublicIps for more details.`,
		Resolver: fetchPublicIPs(
			core.ListPublicIpsScopeRegion,
			nil,
			core.ListPublicIpsLifetimeEphemeral,
		),
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.PublicIpPoolSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func AssignedPublicIPs() *schema.Table {
	return &schema.Table{
		Name: "oracle_virtualnetwork_assigned_public_ips",
		Description: `Ephemeral public IPs assigned to private IPs.
See https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/PublicIp/ListPublicIps for more details.`,
		Resolver: fetchPublicIPs(
			core.ListPublicIpsScopeAvailabilityDomain,
			func(c *client.Client) *string { return &c.AvailabilityDomain },
			core.ListPublicIpsLifetimeEphemeral,
		),
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.PublicIpPoolSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

type availabilityDomainGetter func(*client.Client) *string

func fetchPublicIPs(
	scope core.ListPublicIpsScopeEnum,
	availabilityDomainGetter availabilityDomainGetter,
	lifetime core.ListPublicIpsLifetimeEnum,
) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		cqClient := meta.(*client.Client)

		var availabilityDomain *string
		if availabilityDomainGetter != nil {
			availabilityDomain = availabilityDomainGetter(cqClient)
		}

		var page *string
		limit := 100
		for {
			request := core.ListPublicIpsRequest{
				CompartmentId:      common.String(cqClient.CompartmentOcid),
				AvailabilityDomain: availabilityDomain,
				Scope:              scope,
				Lifetime:           lifetime,
				Page:               page,
				Limit:              &limit,
			}

			response, err := cqClient.OracleClients[cqClient.Region].CoreVirtualnetworkClient.ListPublicIps(ctx, request)

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
}
