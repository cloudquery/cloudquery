package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func AlertPolicies() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_alert_policies",
		Resolver:     fetchAlertPolicies,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeUUID,
				Resolver: schema.UUIDResolver("UUID"),
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "compare",
				Type: schema.TypeString,
			},
			{
				Name: "value",
				Type: schema.TypeFloat,
			},
			{
				Name: "window",
				Type: schema.TypeString,
			},
			{
				Name: "entities",
				Type: schema.TypeStringArray,
			},
			{
				Name: "tags",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "alerts_email",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Alerts.Email"),
			},
			{
				Name: "enabled",
				Type: schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "digitalocean_alert_policy_alerts_slack",
				Resolver: fetchAlertPolicyAlertsSlacks,
				Columns: []schema.Column{
					{
						Name:        "alert_policy_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_alert_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("URL"),
					},
					{
						Name: "channel",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{
		PerPage: client.MaxItemsPerPage,
	}
	for {
		alertPolicies, resp, err := svc.DoClient.Monitoring.ListAlertPolicies(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- alertPolicies
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}

func fetchAlertPolicyAlertsSlacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	ap := parent.Item.(godo.AlertPolicy)
	res <- ap.Alerts.Slack
	return nil
}
