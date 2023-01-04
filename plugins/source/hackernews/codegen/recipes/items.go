package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/hermanschaaf/hackernews"
)

func ItemResources() []*Resource {
	return []*Resource{
		{
			Service:       "items", // this will be the directory name under resources/services
			TableName:     "items", // will become hackernews_items
			DataStruct:    &hackernews.Item{},
			Description:   "https://github.com/HackerNews/API#items",
			IsIncremental: true,
			SkipFields:    []string{"ID"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name: "id",
					Type: schema.TypeInt,
					Options: schema.ColumnCreationOptions{
						PrimaryKey:     true,
						IncrementalKey: true,
					},
					Resolver: `schema.PathResolver("ID")`,
				},
			},
		},
	}
}
