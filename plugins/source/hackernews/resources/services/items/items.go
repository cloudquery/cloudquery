package items

import (
	"reflect"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:           "id",
				Type:           arrow.PrimitiveTypes.Int64,
				Resolver:       schema.PathResolver("ID"),
				PrimaryKey:     true,
				IncrementalKey: true,
			},
		},
	}
}

func typeTransformer(f reflect.StructField) (arrow.DataType, error) {
	if f.Name == "Time" {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}
	return transformers.DefaultTypeTransformer(f)
}

func resolverTransformer(f reflect.StructField, path string) schema.ColumnResolver {
	if f.Name == "Time" {
		return client.UnixTimeResolver(f.Name)
	}
	return transformers.DefaultResolverTransformer(f, path)
}
