package objectstorage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "oracle_objectstorage_buckets",
		Resolver:  fetchBuckets,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&objectstorage.BucketSummary{}, transformers.WithPrimaryKeys("Namespace", "Name")),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchBuckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := objectstorage.ListBucketsRequest{
			NamespaceName: common.String(cqClient.ObjectStorageNamespace),
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].ObjectstorageObjectstorageClient.ListBuckets(ctx, request)

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
