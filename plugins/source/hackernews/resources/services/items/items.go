package items

import (
	"reflect"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
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
			transformers.WithTypeTransformer(typeTransformer),
			transformers.WithResolverTransformer(resolverTransformer),
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

func typeTransformer(f reflect.StructField) (schema.ValueType, error) {
	if f.Name == "Time" {
		return schema.TypeTimestamp, nil
	}
	return transformers.DefaultTypeTransformer(f)
}

func resolverTransformer(f reflect.StructField, path string) schema.ColumnResolver {
	if f.Name == "Time" {
		return client.UnixTimeResolver(f.Name)
	}
	return transformers.DefaultResolverTransformer(f, path)
}
