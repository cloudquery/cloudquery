package client

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
)

func ResolveAWSAccount(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("account_id", client.AccountID)
}

func ResolveAWSRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("region", client.Region)
}

func ResolveAWSNamespace(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("namespace", client.AutoscalingNamespace)
}

func ResolveWAFScope(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return errors.WithStack(r.Set(c.Name, meta.(*Client).WAFScope))
}

func ResolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return ResolveTagField("Tags")(ctx, meta, r, c)
}

func ResolveTagField(fieldName string) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		var val reflect.Value

		if reflect.TypeOf(r.Item).Kind() == reflect.Ptr {
			val = reflect.ValueOf(r.Item).Elem()
		} else {
			val = reflect.ValueOf(r.Item)
		}

		if val.Kind() != reflect.Struct {
			panic("need struct type")
		}
		f := val.FieldByName(fieldName)
		if f.IsNil() {
			return errors.WithStack(r.Set(c.Name, map[string]string{})) // can't have nil or the integration test will make a fuss
		} else if f.IsZero() {
			panic("no such field " + fieldName)
		}
		data := TagsToMap(f.Interface())
		return errors.WithStack(r.Set(c.Name, data))
	}
}

func ResolveTimestampField(path string, rfcs ...string) func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return func(ctx context.Context, cl schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		var val reflect.Value

		value := funk.Get(r.Item, path, funk.WithAllowZero())
		if value == nil {
			return errors.WithStack(r.Set(c.Name, nil))
		}

		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			val = reflect.ValueOf(value).Elem()
		} else {
			val = reflect.ValueOf(value)
		}

		switch val.Kind() {
		case reflect.Int32, reflect.Int64:
			return errors.WithStack(r.Set(c.Name, time.Unix(val.Int(), 0)))
		case reflect.String:
			// TODO: v2 SDK refactor: see if this is still used
			panic("schema.DateResolver not implemented in new SDK")
			// return schema.DateResolver(path, rfcs...)(ctx, cl, r, c)
		default:
			return errors.WithStack(r.Set(c.Name, nil))
		}
	}
}

/*
SliceJsonResolver resolves slice of objects into a map[string]interface{}.

	For example object: SliceJsonStruct{Nested: &SliceJsonStruct{
					Nested: &SliceJsonStruct{
						Value: []types1.Tag{{
							Key:   "k1",
							Value: "v1",
						}, {
							Key:   "k2",
							Value: "v2",
						}},
					},
				}}

can be converted to map[string]interface{}{"k1":"v1","k2":"v2"} by setting a resolver with next params:
SliceJsonResolver("Nested.Nested.Value", "Key", "Value")
*/
func SliceJsonResolver(path, keyPath, valuePath string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		var j map[string]interface{}
		field := funk.Get(r.Item, path, funk.WithAllowZero())
		s := reflect.ValueOf(field)
		if s.IsNil() {
			return r.Set(c.Name, j)
		}
		j = make(map[string]interface{})
		if reflect.TypeOf(field).Kind() != reflect.Slice {
			return errors.WithStack(fmt.Errorf("field: %s is not a slice", path))
		}
		for i := 0; i < s.Len(); i++ {
			key := funk.Get(s.Index(i).Interface(), keyPath, funk.WithAllowZero())
			value := funk.Get(s.Index(i).Interface(), valuePath, funk.WithAllowZero())
			k := reflect.ValueOf(key)
			if k.Kind() == reflect.Ptr {
				k = k.Elem()
			}
			if k.Kind() != reflect.String {
				return errors.WithStack(fmt.Errorf("key field: %s is not a string", path))
			}
			j[k.String()] = value
		}
		return r.Set(c.Name, j)
	}
}

// MarshaledJsonResolver  resolves a json string into a map[string]interface{}.
func MarshaledJsonResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		var j map[string]interface{}
		field := funk.Get(r.Item, path, funk.WithAllowZero())

		if field == nil {
			return r.Set(c.Name, nil)
		}

		var val reflect.Value
		val = reflect.ValueOf(field)
		if reflect.TypeOf(field).Kind() == reflect.Ptr {
			val = val.Elem()
		}

		var b []byte
		switch val.Kind() {
		case reflect.String:
			b = []byte(val.String())
		case reflect.Slice, reflect.Array:
			b = val.Bytes()
		}
		if len(b) == 0 {
			return r.Set(c.Name, nil)
		}

		j = make(map[string]interface{})
		err := json.Unmarshal(b, &j)
		if err != nil {
			return errors.WithStack(err)
		}

		return r.Set(c.Name, j)
	}
}
