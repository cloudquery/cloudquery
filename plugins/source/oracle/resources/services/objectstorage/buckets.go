package objectstorage

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
