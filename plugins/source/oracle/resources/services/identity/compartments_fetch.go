package identity

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func fetchCompartments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := identity.ListCompartmentsRequest{
			CompartmentId:          common.String(cqClient.CompartmentOcid),
			CompartmentIdInSubtree: common.Bool(true),
			Page:                   page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].IdentityIdentityClient.ListCompartments(ctx, request)

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
