package client

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

func ResolveParentColumn(field string) schema.ColumnResolver {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, funk.Get(r.Parent.Item, field))
	}
}

func ResolveAccountName(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.multiplexedAccount.Name)
}

type Nullable interface {
	Get() interface{}
	IsSet() bool
}

func GetNullable(i interface{}) (interface{}, error) {
	switch v := i.(type) {
	case datadog.NullableTime:
		return v.Get(), nil
	case datadog.NullableInt64:
		return v.Get(), nil
	}
	return nil, fmt.Errorf("unsupported datadog nullable type %T", i)
}

func NullableResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		data, err := GetNullable(funk.Get(r.Item, path, funk.WithAllowZero()))
		if err != nil {
			return err
		}
		return r.Set(c.Name, data)
	}
}
