package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v45/github"
)

//go:generate cq-gen --resource  --config hooks.hcl --output .
func Hooks() *schema.Table {
	return &schema.Table{
		Name:        "github_hooks",
		Description: "Hook represents a GitHub (web and service) hook for a repository.",
		Resolver:    fetchHooks,
		Multiplex:   client.OrgMultiplex,
		IgnoreError: client.IgnoreError,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"org", "id"}},
		Columns: []schema.Column{
			{
				Name:        "org",
				Description: "The Github Organization of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
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
				Name:     "id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ID"),
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
		},
		Relations: []*schema.Table{
			{
				Name:        "github_hook_deliveries",
				Description: "HookDelivery represents the data that is received from GitHub's Webhook Delivery API  GitHub API docs: - https://docs.github.com/en/rest/webhooks/repo-deliveries#list-deliveries-for-a-repository-webhook - https://docs.github.com/en/rest/webhooks/repo-deliveries#get-a-delivery-for-a-repository-webhook",
				Resolver:    fetchHookDeliveries,
				Columns: []schema.Column{
					{
						Name:        "hook_cq_id",
						Description: "Unique CloudQuery ID of github_hooks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name:     "guid",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("GUID"),
					},
					{
						Name:     "delivered_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("DeliveredAt.Time"),
					},
					{
						Name: "redelivery",
						Type: schema.TypeBool,
					},
					{
						Name: "duration",
						Type: schema.TypeFloat,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "status_code",
						Type: schema.TypeBigInt,
					},
					{
						Name: "event",
						Type: schema.TypeString,
					},
					{
						Name: "action",
						Type: schema.TypeString,
					},
					{
						Name:     "installation_id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("InstallationID"),
					},
					{
						Name:     "repository_id",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("RepositoryID"),
					},
					{
						Name:     "request_headers",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Request.Headers"),
					},
					{
						Name:     "request_raw_payload",
						Type:     schema.TypeByteArray,
						Resolver: resolveHookDeliveriesRequestRawPayload,
					},
					{
						Name:     "response_headers",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Response.Headers"),
					},
					{
						Name:     "response_raw_payload",
						Type:     schema.TypeByteArray,
						Resolver: resolveHookDeliveriesResponseRawPayload,
					},
				},
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
			return diag.WrapError(err)
		}
		res <- hooks
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
func fetchHookDeliveries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	h := parent.Item.(*github.Hook)
	opts := &github.ListCursorOptions{
		PerPage: 100,
	}
	for {
		hooks, resp, err := c.Github.Organizations.ListHookDeliveries(ctx, c.Org, *h.ID, opts)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- hooks
		if len(hooks) == 0 || resp.Cursor == "" {
			return nil
		}
		opts.Cursor = resp.Cursor
	}
}
func resolveHookDeliveriesRequestRawPayload(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	hd := resource.Item.(*github.HookDelivery)
	data, err := hd.Request.RawPayload.MarshalJSON()
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}

func resolveHookDeliveriesResponseRawPayload(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	hd := resource.Item.(*github.HookDelivery)
	data, err := hd.Response.RawPayload.MarshalJSON()
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
