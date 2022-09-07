package hooks

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

func Hooks() *schema.Table {
	return &schema.Table{
		Name:        "github_hooks",
		Description: "Hook represents a GitHub (web and service) hook for a repository.",
		Resolver:    fetchHooks,
		Multiplex:   client.OrgMultiplex,
		Columns: []schema.Column{
			{
				Name:            "org",
				Description:     "The Github Organization of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveOrg,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "id",
				Type:            schema.TypeInt,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name: "created_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "updated_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "test_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TestURL"),
			},
			{
				Name:     "ping_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PingURL"),
			},
			{
				Name:          "last_response",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "config",
				Description:   "Only the following fields are used when creating a hook. Config is required.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name: "events",
				Type: schema.TypeStringArray,
			},
			{
				Name: "active",
				Type: schema.TypeBool,
			},
			{
				Name:        "deliveries",
				Description: "Webhook deliveries",
				Type:        schema.TypeJSON,
				Resolver:    resolveHookDeliveries,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchHooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{
		Page:    0,
		PerPage: 100,
	}
	for {
		hooks, resp, err := c.Github.Organizations.ListHooks(ctx, c.Org, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- hooks
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func resolveHookDeliveries(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	c := meta.(*client.Client)
	h := resource.Item.(*github.Hook)
	opts := &github.ListCursorOptions{
		PerPage: 100,
	}
	var deliveries []*github.HookDelivery
	for {
		hooks, resp, err := c.Github.Organizations.ListHookDeliveries(ctx, c.Org, *h.ID, opts)
		if err != nil {
			return errors.WithStack(err)
		}
		deliveries = append(deliveries, hooks...)
		if len(hooks) == 0 || resp.Cursor == "" {
			break
		}
		opts.Cursor = resp.Cursor
	}
	data, err := json.Marshal(deliveries)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(column.Name, data))
}
