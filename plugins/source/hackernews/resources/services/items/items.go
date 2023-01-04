package items

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/hermanschaaf/hackernews"
)

func Items() *schema.Table {
	return &schema.Table{
		Name:          "hackernews_items",
		Description:   `https://github.com/HackerNews/API#items`,
		Resolver:      fetchItems,
		IsIncremental: true,
		Transform: transformers.TransformWithStruct(
			&hackernews.Item{},
			transformers.WithSkipFields("ID"),
		),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey:     true,
					IncrementalKey: true,
				},
			},
		},
	}
}
