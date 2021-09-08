package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Account() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_accounts",
		Description:  "Account represents a DigitalOcean Account",
		Resolver:     fetchAccounts,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uuid"}},
		Columns: []schema.Column{
			{
				Name:        "droplet_limit",
				Description: "The total number of Droplets current user or team may have active at one time.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "floating_ip_limit",
				Description: "The total number of Floating IPs the current user or team may have.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("FloatingIPLimit"),
			},
			{
				Name:        "volume_limit",
				Description: "The total number of volumes the current user or team may have.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "email",
				Description: "The email address used by the current user to register for DigitalOcean.",
				Type:        schema.TypeString,
			},
			{
				Name:        "uuid",
				Description: "The unique universal identifier for the current user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UUID"),
			},
			{
				Name:        "email_verified",
				Description: "If true, the user has verified their account via email. False otherwise.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "status",
				Description: "This value is one of \"active\", \"warning\" or \"locked\".",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_message",
				Description: "A human-readable message giving more details about the status of the account.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client)
	account, _, err := svc.DoClient.Account.Get(ctx)
	if err != nil {
		return err
	}
	res <- *account
	return nil
}
