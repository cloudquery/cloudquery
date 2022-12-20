package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
	"github.com/thoas/go-funk"
)

func ResolveTeamID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.TeamID)
}

func JSONTimeResolver(fieldName string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
		t := funk.Get(r.Item, fieldName, funk.WithAllowZero())
		switch v := t.(type) {
		case slack.JSONTime:
			return r.Set(col.Name, v.Time())
		case *slack.JSONTime:
			if v == nil {
				return nil
			}
			return r.Set(col.Name, v.Time())
		}
		panic("unknown type for ResolveJSONTime")
	}
}
