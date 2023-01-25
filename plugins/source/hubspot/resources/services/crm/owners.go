package crm

import (
	"github.com/clarkmcc/go-hubspot/generated/v3/owners"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Owners() *schema.Table {
	return &schema.Table{
		Name:        "hubspot_crm_owners",
		Resolver:    fetchOwners,
		Description: "https://developers.hubspot.com/docs/api/crm/owners",
		Transform:   transformers.TransformWithStruct(owners.PublicOwner{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
