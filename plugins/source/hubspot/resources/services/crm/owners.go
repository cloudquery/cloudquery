package crm

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/clarkmcc/go-hubspot/generated/v3/owners"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Owners() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_owners",
		Resolver:    fetchOwners,
		Description: "https://developers.hubspot.com/docs/api/crm/owners",
		Transform:   transformers.TransformWithStruct(owners.PublicOwner{}),
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
