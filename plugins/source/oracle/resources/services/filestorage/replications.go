package filestorage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/filestorage"
)

func Replications() *schema.Table {
	return &schema.Table{
		Name:      "oracle_filestorage_replications",
		Resolver:  fetchReplications,
		Multiplex: client.AvailibilityDomainCompartmentMultiplex,
		Transform: client.TransformWithStruct(&filestorage.ReplicationSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn, client.AvailabilityDomainColumn},
	}
}

func fetchReplications(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := filestorage.ListReplicationsRequest{
			CompartmentId:      common.String(cqClient.CompartmentOcid),
			AvailabilityDomain: common.String(cqClient.AvailabilityDomain),
			Page:               page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].FilestorageFilestorageClient.ListReplications(ctx, request)

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
