package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func CostTrackingTags() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_cost_tracking_tags",
		Resolver:  fetchCostTrackingTags,
		Multiplex: client.TenancyMultiplex,
		Transform: client.TransformWithStruct(&identity.Tag{}),
		Columns:   schema.ColumnList{client.RegionColumn},
	}
}
