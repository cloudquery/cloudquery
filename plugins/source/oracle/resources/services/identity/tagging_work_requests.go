package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func TaggingWorkRequests() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_tagging_work_requests",
		Resolver:  fetchTaggingWorkRequests,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&identity.TaggingWorkRequestSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
