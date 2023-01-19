package account

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/fastly/go-fastly/v7/fastly"
)

func AccountEvents() *schema.Table {
	return &schema.Table{
		Name:        "fastly_account_events",
		Description: `https://developer.fastly.com/reference/api/account/events/`,
		Resolver:    fetchAccountEvents,
		Transform:   transformers.TransformWithStruct(&fastly.Event{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IP"),
			},
		},
	}
}
