package identity

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func CostTrackingTags() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_cost_tracking_tags",
		Resolver:  fetchCostTrackingTags,
		Multiplex: client.TenancyMultiplex,
		Transform: transformers.TransformWithStruct(&identity.Tag{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}
