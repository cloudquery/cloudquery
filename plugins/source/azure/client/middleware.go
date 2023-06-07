package client

import (
	"context"
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func LowercaseIDResolver(_ context.Context, _ schema.ClientMeta, resource *schema.Resource) error {
	resource.SetItem(lowercaseID(resource.GetItem()))
	return nil
}

func ChainRowResolvers(next ...schema.RowResolver) schema.RowResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
		for i := range next {
			if next[i] == nil {
				continue
			}
			if err := next[i](ctx, meta, resource); err != nil {
				return err
			}
		}
		return nil
	}
}

func lowercaseID(obj any) any {
	value := reflect.Indirect(reflect.ValueOf(obj))
	if value.Kind() != reflect.Struct {
		return obj
	}

	vt := value.Type()
	for i := 0; i < value.NumField(); i++ {
		if tag := strings.SplitN(vt.Field(i).Tag.Get("json"), ",", 2)[0]; tag != "id" {
			continue
		}

		f := value.Field(i)
		if f.Kind() == reflect.String || (f.Kind() == reflect.Ptr && f.Elem().Kind() == reflect.String) {
			if f.Kind() == reflect.String {
				f.SetString(strings.ToLower(f.String()))
			} else if f.Kind() == reflect.Ptr && f.Elem().Kind() == reflect.String {
				f.Elem().SetString(strings.ToLower(f.Elem().String()))
			}
		}
	}
	return obj
}
