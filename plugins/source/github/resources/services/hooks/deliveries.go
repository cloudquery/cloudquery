package hooks

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func deliveries() *schema.Table {
	return &schema.Table{
		Name:                "github_hook_deliveries",
		Resolver:            fetchDeliveries,
		PreResourceResolver: hooksGet,
		Transform: transformers.TransformWithStruct(&github.HookDelivery{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{
			client.OrgColumn,
			{
				Name:            "hook_id",
				Type:            schema.TypeInt,
				Resolver:        client.ResolveParentColumn("ID"),
				Description:     `Hook ID`,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
